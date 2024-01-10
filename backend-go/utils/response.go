package utils

type ErrorCode struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Message     string      `json:"message"`
	Code        int         `json:"code"`
	Data        interface{} `json:"data"`
	Description *string     `json:"description"`
}

func ResponseOK(data interface{}, description ...string) *Response {
	var des *string
	if len(description) > 0 {
		des = &description[0]
	} else {
		des = nil
	}
	return &Response{
		Message:     "ok",
		Code:        0,
		Data:        data,
		Description: des,
	}
}
