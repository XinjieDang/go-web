package utils

type HttpUtil struct {
	httpCode int64
	httpMsg  string
}

func (h *HttpUtil) error() *HttpUtil {
	h.httpCode = 500
	h.httpMsg = "系统异常"
	return h
}
