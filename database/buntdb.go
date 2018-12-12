package database

import (
	"github.com/tidwall/buntdb"
	"fmt"
)

type MemoryDb struct {
	Engine *buntdb.DB
}

func NewMemDb() *MemoryDb {
	mdb := &MemoryDb{}
	err := mdb.Init()
	if err != nil {
		panic(err)
	}
	return mdb
}

func (db *MemoryDb) Init() error  {

	bdb, err := buntdb.Open(":memory:")
	if err != nil {
		return err
	}
	err = bdb.Update(func(tx *buntdb.Tx) error {
		return tx.CreateIndex("ts", "*", buntdb.IndexJSON("timestamp"))
	})
	db.Engine = bdb
	return err
}

func (db *MemoryDb) Insert(id int64, data string) error {
	ids := fmt.Sprintf("%d", id)
	return db.Engine.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(ids, string(data), nil)
		return err
	})
}

func (db *MemoryDb) Delete(id int64) error {
	fmt.Println("delete id", id)
	ids := fmt.Sprintf("%d", id)
	return db.Engine.Update(func(tx *buntdb.Tx) error {
		fmt.Println("delete id", id)
		_, err := tx.Delete(ids)
		fmt.Println("delete error: ", err)
		return err
	})
}
