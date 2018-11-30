package helper

import (
	"time"

	config "github.com/micro/go-config"
	"github.com/micro/go-config/source/file"
	log "github.com/golang/glog"
)



func Now() int64 {
	loc, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}
	return time.Now().In(loc).Unix()
}


func NowPlus(extraSecond int32) int64 {
	return Now() + int64(extraSecond)
}


func Config(v interface{}, params ...string)  {
	config.Load(file.NewSource(
		file.WithPath("conf/config.json"),
	))
	err := config.Get(params...).Scan(v)
	if err != nil {
		log.Error("configuration error: ", err)
		panic(err)
	}
}

func GetQueueType() string {
	config.Load(file.NewSource(
		file.WithPath("conf/config.json"),
	))
	queueType := config.Get("queueType").String("")
	return queueType
}


func GetDbType() string {
	config.Load(file.NewSource(
		file.WithPath("conf/config.json"),
	))
	dbType := config.Get("dbType").String("")
	return dbType
}