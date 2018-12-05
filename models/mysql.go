package message

import (
	log "github.com/golang/glog"

	"git.zapa.cloud/evn-gateway/DelayQueue/helper"
	"github.com/astaxie/beego/orm"
)

type MySQLMessage struct {
	Id         int64  `json:"-" orm:"auto"`
	TimeStamp  int64  `json:"timestamp"`
	Data       string `json:"Message" orm:"type(text)"`
	RetryCount int    `json:"retry_count"`
	Delay      int    `json:"delay"`
}

const MYSQL = "mysql"

func init() {
	if helper.GetDbType() == MYSQL {
		orm.RegisterModel(new(MySQLMessage))
	}
}


func Add(m *MySQLMessage) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func ByTime(now int64) (ml []*MySQLMessage, err error) {
	o := orm.NewOrm()
	cond := orm.NewCondition().And("time_stamp__lt", now)

	qs := o.QueryTable("MySQLMessage")
	qs = qs.SetCond(cond).Limit(LIMIT_RANGE_MSG)
	_, err = qs.All(&ml)
	if err != nil && err != orm.ErrNoRows {
		return nil, err
	}

	return ml, nil
}

func All() (ml []*MySQLMessage, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("MySQLMessage").All(&ml)
	if err != nil && err != orm.ErrNoRows {
		return nil, err
	}

	return ml, nil
}

func List(offset int, limit int) (ml []*MySQLMessage, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("MySQLMessage").Offset(offset).Limit(limit).All(&ml)
	if err != nil && err != orm.ErrNoRows {
		return nil, err
	}

	return ml, nil
}

func SearchBy(data string) ([]*MySQLMessage, error) {
	o := orm.NewOrm()
	var msgList []*MySQLMessage

	num, err := o.Raw("select * from `Message` where data like ?", "%"+data+"%").QueryRows(&msgList)
	if err != nil {
		return nil, err
	}
	log.Info("Message nums: ", num)
	return msgList, nil
}

func Force(id int64) (err error) {
	o := orm.NewOrm()
	msg := MySQLMessage{Id: id}
	if err = o.Read(&msg, "id"); err == nil {
		msg.TimeStamp = helper.Now() - 2

		var num int64
		if num, err = o.Update(&msg, "TimeStamp"); err == nil {
			log.Info("Number of records updated in database:", num)
		}
	}
	return
}

func Delete(m *MySQLMessage) error {
	o := orm.NewOrm()
	num, err := o.Delete(m)
	if err == nil {
		log.Info("Number of records deleted in database:", num)
	}
	return err
}
