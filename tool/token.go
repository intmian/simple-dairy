package tool

import (
	"strconv"
	"time"
	"github.com/intmian/mian_go_lib/tool/cipher"
	"github.com/pochard/commons/randstr"
)

const MAX_TOKEN_IP = 3 // 最多允许同一个IP登录的token数量

type TokenDetail struct {
	id         string
	ip         []string // 已登录的ip
	createTime int64
}

type tokenMgr struct {
	tokenMap map[string]*TokenDetail
}

func (t *tokenMgr) isTokenValid(token string, id string, ip string) bool {
	tokenDetail, ok := t.tokenMap[token]
	if !ok {
		return false
	}
	if tokenDetail.id != id {
		return false
	}
	isExistIP := false
	for _, v := range tokenDetail.ip {
		if v == ip {
			isExistIP = true
		}
	}
	if !isExistIP {
		tokenDetail.ip = append(tokenDetail.ip, ip)
	}
	if len(tokenDetail.ip) > MAX_TOKEN_IP {
		return false
	}

	// 有效期仅三天
	if (tokenDetail.createTime + 3*24*60*60) < time.Now().Unix() {
		return false
	}
	return true
}

// 创建token
func (t *tokenMgr) createToken(id string, ip string) string {
	token := strconv.FormatInt(time.Now().Unix(), 10)
	token += randstr.RandomAlphanumeric(10)
	// sha256
	token = cipher.s
	tokenDetail := &TokenDetail{
		id:         id,
		ip:         []string{ip},
		createTime: time.Now().Unix(),
	}
	t.tokenMap[token] = tokenDetail
	return token
}
