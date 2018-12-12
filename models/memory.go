package message

import (
	"github.com/whatvn/dqueue/helper"

	"github.com/whatvn/dqueue/database"
	"sync/atomic"
	"encoding/json"
)

var (
	db *database.MemoryDb
	id int64 = 0
)

type MemoryMessage struct {
	MySqlMessage
}

func GetMemDb()  *database.MemoryDb {
	return db
}
func init()  {
	if helper.GetDbType() == MEMORY {
		db = database.NewMemDb()
	}
}

func (m *MemoryMessage) Save() (int64, error) {
	atomic.AddInt64(&id, 1)
	return id, db.Insert(id, m.json())
}

func (m *MemoryMessage) json() string  {
	b, _ := json.Marshal(m)
	return string(b)
}

func (m *MemoryMessage) byte() []byte  {
	b, _ := json.Marshal(m)
	return b
}

func (m *MemoryMessage) Delete() error {
	return db.Delete(m.Id)
}

func (m *MemoryMessage) Force() error {
	m.TimeStamp = helper.NowPlus(-100)
	return db.Insert(m.Id, m.json())
}

