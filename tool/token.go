package tool

import "strconv"

func MakeUserKey(ip string,port int,param uint32) string {
	return ip + ":" + strconv.Itoa(port) + ":" + strconv.Itoa(int(param))
}

type tokenMgr struct {

}