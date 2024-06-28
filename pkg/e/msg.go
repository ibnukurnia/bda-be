package e

var msgFlags = map[int]string{
	SUCCESS:     "ok",
	ERROR:       "internal server error",
	BAD_REQUEST: "bad request",
}

func GetMsg(code int) string {
	msg, ok := msgFlags[code]
	if ok {
		return msg
	}

	return msgFlags[ERROR]
}
