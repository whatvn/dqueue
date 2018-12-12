package database

import (
	"github.com/astaxie/beego/orm"
	"time"
	"os"
	"strconv"
	"github.com/whatvn/dqueue/helper"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/golang/glog"
	"fmt"
)

type mysql struct {}

func (database *mysql) Init() error {
	var mysql DbConfig
	helper.Config(&mysql, "hosts", "mysql")
	return initialize(&DbConfig{
		User:     mysql.User,
		Password: mysql.Password,
		Address:  mysql.Address,
		Database: mysql.Database,
	})
}

func newMySqlDatabase() Database {
	return &mysql{}
}

type DbConfig struct {
	User     string
	Password string
	Address  string
	Database string
}

var (
	maxIdle  = 100
	maxConn  = 200
	lifeTime = 300
)


func initialize(conf *DbConfig) error {
	var connectString = fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&collation=utf8mb4_general_ci&loc=%s",
		conf.User, conf.Password, conf.Address, conf.Database, "Asia%2fBangkok")

	orm.RegisterDriver("mysql", orm.DRMySQL)
	log.Info("connect database with connect string: ", connectString)

	if err := orm.RegisterDataBase("default", "mysql",
		connectString, maxIdle, maxConn); err != nil {
		log.Info("error when RegisterDataBase: ", err)
		return err
	}
	//
	db, err := orm.GetDB("default")
	if err != nil {
		log.Fatal("get default DB error:" + err.Error())
		return err
	}
	db.SetConnMaxLifetime(time.Duration(lifeTime) * time.Second)
	orm.Debug, orm.DebugLog = false, orm.NewLog(os.Stdout)
	if dbMigrate := os.Getenv("DB_MIGRATE"); len(dbMigrate) == 0 {
		log.Info("init database without migrate, config: ", connectString)
	} else {
		if isMigrate, _ := strconv.ParseBool(dbMigrate); isMigrate {
			if err := orm.RunSyncdb("default", false, false); err != nil {
				log.Info("error when migrate db", err)
				return err
			}
			log.Info("init database migrate mode, config: ", connectString)
		}
	}
	return nil
}
