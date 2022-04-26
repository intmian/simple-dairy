package tool

import "github.com/intmian/mian_go_lib/tool/misc"

type SettingData struct {
	PwdToken string `json:"pwd_token"`
	KeyToken string `json:"key_token"`  // 文本key 通过将原生key奇数位取出后，经过两次以上256进制转换后的key
}

type Setting struct {
	*SettingData
	*misc.TJsonTool
}

func NewSetting() *Setting {
	s := Setting{}
	s.TJsonTool = misc.NewTJsonTool("\\data\\setting.json",s.SettingData)
	return &s
}

var GSetting = NewSetting()

