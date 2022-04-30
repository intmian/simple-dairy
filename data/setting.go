package data

import "github.com/intmian/mian_go_lib/tool/misc"

type SettingData struct {
	AccountSQLConnection string `json:"account_sql_connection"`
	ContentSQLConnection string `json:"content_sql_connection"`
}

type Setting struct {
	*SettingData
	*misc.TJsonTool
}

func NewSetting() *Setting {
	s := Setting{}
	s.TJsonTool = misc.NewTJsonTool("\\data\\setting.json", s.SettingData)
	return &s
}

var GSetting = NewSetting()
