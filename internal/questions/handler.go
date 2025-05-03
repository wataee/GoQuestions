package questions


type Handler struct {
	service QuestionsService
}

func NewHandler(service QuestionsService) *Handler {
	return &Handler{service: service}
}