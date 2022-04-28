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
	tokenMap    map[string]*TokenDetail
	id2tokenMap map[string]string
}

func (t *tokenMgr) IsTokenValid(token string, id string, ip string) bool {
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

	if _, ok := t.id2tokenMap[id]; !ok {
		t.id2tokenMap[id] = token
	}
	return true
}

func makeToken(id string, ip string) string {
	token := strconv.FormatInt(time.Now().Unix(), 10)
	token += randstr.RandomAlphanumeric(10)
	tokenSha256 := cipher.Sha2562String(token)
	return tokenSha256
}

// 创建token
func (t *tokenMgr) CreateToken(id string, ip string) string {
	token := makeToken(id, ip)
	loopNum := 0
	for {
		token = makeToken(id, ip)
		tokenValue, ok := t.tokenMap[token]
		if !ok {
			break
		}
		// 如果token过期，则重新生成
		if tokenValue.createTime+3*24*60*60 >= time.Now().Unix() {
			break
		}
		loopNum++
		// 无法生成新的token，则返回空
		if loopNum > 100 {
			return ""
		}
	}
	tokenDetail := &TokenDetail{
		id:         id,
		ip:         []string{ip},
		createTime: time.Now().Unix(),
	}
	t.tokenMap[token] = tokenDetail
	t.id2tokenMap[id] = token
	return token
}

func (t *tokenMgr) ClearOutTimeToken() {
	for k, v := range t.id2tokenMap {
		if t.tokenMap[v].createTime+3*24*60*60 < time.Now().Unix() {
			delete(t.id2tokenMap, k)
		}
	}
	for k, v := range t.tokenMap {
		if v.createTime+3*24*60*60 < time.Now().Unix() {
			delete(t.tokenMap, k)
		}
	}
}

func (t *tokenMgr) ClearToken(token string) {
	id := t.tokenMap[token].id
	delete(t.id2tokenMap, id)
	delete(t.tokenMap, token)
}

func MakeTokenMgr() *tokenMgr {
	return &tokenMgr{
		tokenMap: make(map[string]*TokenDetail),
	}
}
