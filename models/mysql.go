package message

import (
	log "github.com/golang/glog"

	"git.zapa.cloud/evn-gateway/DelayQueue/helper"
	"github.com/astaxie/beego/orm"
	"encoding/json"
)

type MySqlMessage struct {
	Id         int64  `json:"-" orm:"auto"`
	TimeStamp  int64  `json:"timestamp"`
	Data       string `json:"Message" orm:"type(text)"`
	RetryCount int    `json:"retry_count"`
	Delay      int    `json:"delay"`
}

func init() {
	orm.RegisterModel(new(MySqlMessage))
}

func (m *MySqlMessage) Save() (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func (m *MySqlMessage) Force() (err error) {
	o := orm.NewOrm()
	if err = o.Read(m, "id"); err == nil {
		m.TimeStamp = helper.Now() - 2

		var num int64
		if num, err = o.Update(m, "TimeStamp"); err == nil {
			log.Info("Number of records updated in database:", num)
		}
	}
	return
}

func (m *MySqlMessage) json() string {
	b, _ := json.Marshal(m)
	return string(b)
}

func (m *MySqlMessage) byte() []byte {
	b, _ := json.Marshal(m)
	return b
}

func (m *MySqlMessage) Delete() error {
	o := orm.NewOrm()
	num, err := o.Delete(m)
	if err == nil {
		log.Info("Number of records deleted in database:", num)
	}
	return err
}

func All() (ml []*MySqlMessage, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("MySqlMessage").All(&ml)
	if err != nil && err != orm.ErrNoRows {
		return nil, err
	}

	return ml, nil
}

func List(offset int, limit int) (ml []*MySqlMessage, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("MySqlMessage").Offset(offset).Limit(limit).All(&ml)
	if err != nil && err != orm.ErrNoRows {
		return nil, err
	}

	return ml, nil
}

func SearchBy(data string) ([]*MySqlMessage, error) {
	o := orm.NewOrm()
	var msgList []*MySqlMessage

	num, err := o.Raw("select * from `Message` where data like ?", "%"+data+"%").QueryRows(&msgList)
	if err != nil {
		return nil, err
	}
	log.Info("Message nums: ", num)
	return msgList, nil
}
