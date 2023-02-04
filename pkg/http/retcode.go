package http

type RetCode int

const (
	Ok RetCode = 0

	// 1xxx为请求相关错误
	InvalidRequest RetCode = 1000
	BadRequest     RetCode = 1001

	// 2xxx为业务内部处理相关错误
	InternalErr RetCode = 2000
)
