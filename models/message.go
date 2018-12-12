package message

import (
	"github.com/whatvn/dqueue/protobuf"
	"github.com/whatvn/dqueue/helper"
	"github.com/astaxie/beego/orm"
	"github.com/whatvn/dqueue/queue"
	"github.com/tidwall/buntdb"
	log "github.com/golang/glog"
	"strconv"
)

type Message interface {
	Save() (int64, error)
	Force() error
	Delete() error
}



func NewMessage(queueRequest *delayQueue.QueueRequest) Message {

	msg := MySqlMessage{
		TimeStamp:  helper.NowPlus(queueRequest.Delay * queueRequest.RetryCount),
		Data:       queueRequest.Messsage,
		RetryCount: int(queueRequest.RetryCount),
		Delay:      int(queueRequest.Delay),
	}
	switch helper.GetDbType() {
	case MYSQL:
		return &msg
	case MEMORY:
		return &MemoryMessage{
			MySqlMessage: msg,
		}
	default:
		panic(NotImplementError)

	}
}

func NewMessageById(id int64) Message {
	msg := MySqlMessage{
		Id: id,
	}
	switch helper.GetDbType() {
	case MYSQL:
		return &msg
	case MEMORY:
		return &MemoryMessage{
			MySqlMessage: msg,
		}
	default:
		panic(NotImplementError)

	}
}

func Publish(q queue.Queue) error {
	now := helper.Now()
	switch helper.GetDbType() {
	case MYSQL:
		var ml = make([]*MySqlMessage, 0)
		o := orm.NewOrm()
		cond := orm.NewCondition().And("time_stamp__lt", now)

		qs := o.QueryTable("MySqlMessage")
		qs = qs.SetCond(cond).Limit(LIMIT_RANGE_MSG)
		_, err := qs.All(&ml)
		if err != nil && err != orm.ErrNoRows {
			return err
		}

		for _, m := range ml {
			err = q.PublishMessage(m.byte())
			if err == nil {
				m.Delete()
			}
		}
		return nil

	case MEMORY:
		deleteChan := make(chan *MemoryMessage, 1)
		err := GetMemDb().Engine.View(func(tx *buntdb.Tx) error {
			err := tx.AscendRange("ts", NewTsRange(0).ToJson(), NewTsRange(now).ToJson(), func(key, value string) bool {
				bs := []byte(value)

				err := q.PublishMessage(bs)
				log.Error(err)
				if err == nil {
					m := &MemoryMessage{}
					mId, _ := strconv.ParseInt(key, 10, 0)
					log.Info("push message id to delete chan ", mId)
					m.Id = mId
					deleteChan <- m
				}

				return true
			})
			close(deleteChan)
			return err
		})
		log.Error("AscendRange error: ", err)
		if err == nil {
			for m := range deleteChan {
				log.Info("m: ", m)
				if err = m.Delete(); err != nil {
					return err
				}
			}
			return nil
		}
		return err

	default:
		panic(NotImplementError)
	}
}
