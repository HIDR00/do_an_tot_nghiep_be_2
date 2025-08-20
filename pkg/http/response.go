package http

type BaseResponse struct {
	Success   bool        `json:"success"`
	ErrorCode *string     `json:"error_code"`
	Data      interface{} `json:"data"`
	TraceID   string      `json:"trace_id"`
}
