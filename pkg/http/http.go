package http

type CommonResponse struct {
	RetCode RetCode     `json:"ret_code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
