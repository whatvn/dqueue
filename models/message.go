package message

import (
	log "github.com/golang/glog"

	"git.zapa.cloud/evn-gateway/DelayQueue/helper"
	"github.com/astaxie/beego/orm"
)

type Message struct {
	Id         int64  `json:"-" orm:"auto"`
	TimeStamp  int64  `json:"time_stamp"`
	Data       string `json:"message" orm:"type(text)"`
	RetryCount int    `json:"retry_count"`
	Delay      int    `json:"delay"`
}

func init() {
	orm.RegisterModel(new(Message))
}

func AddMessage(m *Message) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func GetMessagesByTimeStamp(now int64) (ml []*Message, err error) {
	o := orm.NewOrm()
	cond := orm.NewCondition().And("time_stamp__lt", now)

	qs := o.QueryTable("message")
	qs = qs.SetCond(cond).Limit(LIMIT_RANGE_MSG)
	_, err = qs.All(&ml)
	if err != nil && err != orm.ErrNoRows {
		return nil, err
	}

	return ml, nil
}

func GetAllMessages() (ml []*Message, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("message").All(&ml)
	if err != nil && err != orm.ErrNoRows {
		return nil, err
	}

	return ml, nil
}

func GetMessageList(offset int, limit int) (ml []*Message, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("message").Offset(offset).Limit(limit).All(&ml)
	if err != nil && err != orm.ErrNoRows {
		return nil, err
	}

	return ml, nil
}

func GetMessageListByData(data string) ([]*Message, error) {
	o := orm.NewOrm()
	var msgList []*Message

	num, err := o.Raw("select * from `message` where data like ?", "%"+data+"%").QueryRows(&msgList)
	if err != nil {
		return nil, err
	}
	log.Info("message nums: ", num)
	return msgList, nil
}

func UpdateTimestampMessage(data string) (err error) {
	o := orm.NewOrm()
	msg := Message{Data: data}
	if err = o.Read(&msg, "Data"); err == nil {
		msg.TimeStamp = helper.Now() - 2

		var num int64
		if num, err = o.Update(&msg, "TimeStamp"); err == nil {
			log.Info("Number of records updated in database:", num)
		}
	}
	return
}

func DeleteMessage(m *Message) error {
	o := orm.NewOrm()
	num, err := o.Delete(m)
	if err == nil {
		log.Info("Number of records deleted in database:", num)
	}
	return err
}
