package database

import (
	"simple-dairy/data"
	"time"

	"github.com/jmoiron/sqlx"
)

type DataBaseMgr struct {
	accountDao *sqlx.DB
	contentDao *sqlx.DB
}

var GDataMgr = NewDataBaseMgr()

func NewDataBaseMgr() *DataBaseMgr {
	d := &DataBaseMgr{}
	d.Init()
	return d
}

func (d *DataBaseMgr) Init() {
	var err error
	// sqlx 自带连接池。不用画蛇添足。
	d.accountDao, err = sqlx.Open("mysql", data.GSetting.AccountSQLConnection)
	d.contentDao, err = sqlx.Open("mysql", data.GSetting.ContentSQLConnection)
	if err != nil {
		panic(err)
	}
}

type DBUser struct {
	ID         uint64 `db:"id"`
	PWD        string `db:"pwd"`
	Permission uint8  `db:"permission"`
}

type DBContent struct {
	ID   uint64    `db:"id"`
	Name string    `db:"name"`
	Tag  string    `db:"tag"` // 标签，使用逗号分隔
	Time time.Time `db:"time"`
	Data string    `db:"data"`
}

type DBContentTag struct {
	Tag       string `db:"tag"`
	ContentID uint64 `db:"content_id"`
}

