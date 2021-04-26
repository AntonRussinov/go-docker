package model

//Response модель ответа
type Response struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ErrorResponse struct {
	Error      string `json:"error"`
	HTTPStatus int    `json:"httpStatus"`
}
