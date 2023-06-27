package modelBase

type Resp struct {
	Data   interface{} `json:"data,omitempty"`
	Error  *Error      `json:"error,omitempty"`
	Errors []Error     `json:"errors,omitempty"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
