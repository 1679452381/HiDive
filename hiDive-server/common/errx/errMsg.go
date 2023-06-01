package errx

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[SERVER_COMMON_ERROR] = "服务开小小差了,请稍后再试"
	message[REQUEST_PARAM_ERROR] = "参数错误"
	message[TOKEN_EXPIRE_ERROR] = "token超时"
	message[TOKEN_GENERATE_ERROR] = "生成token"
	message[DB_ERROR] = "数据库繁忙,请稍后再试"
}

func MapErrMsg(code uint32) string {
	if msg, ok := message[code]; ok {
		return msg
	} else {
		return "服务开小小差了,请稍后再试"
	}
}

func IsCodeErr(code uint32) bool {
	if _, ok := message[code]; ok {
		return true
	} else {
		return false
	}
}
