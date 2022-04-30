package database

import (
	"simple-dairy/tool"

	"github.com/jmoiron/sqlx"
)

type DataBaseMgr struct {
	accountDao *sqlx.DB
	contentDao *sqlx.DB
}

func (d *DataBaseMgr) Init() {
	var err error
	d.accountDao, err = sqlx.Open("mysql",tool.GSetting.AccessToken)
	if err != nil {
		panic(err)
	}
}
