package handler

type Handler struct{}

//NewHandler возвращает указатель на структуру *Handler
func NewHandler() *Handler {
	var h = Handler{}
	return &h
}
