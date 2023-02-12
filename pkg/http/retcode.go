package http

type RetCode int

const (
	RetCodeOk RetCode = 0

	// 1xxx为请求相关错误
	RetCodeInvalidRequest RetCode = 1000
	RetCodeBadRequest     RetCode = 1001

	// 2xxx为业务内部处理相关错误
	RetCodeInternalErr RetCode = 2000
)
