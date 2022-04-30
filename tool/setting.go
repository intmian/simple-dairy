package tool

import "github.com/intmian/mian_go_lib/tool/misc"

type SettingData struct {
	AccessToken          string `json:"access_token"`
	AdminToken           string `json:"admin_token"`
	ReadWriteToken       string `json:"read_write_token"`
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
