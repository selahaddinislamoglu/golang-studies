package handler

type HandlerInterface interface {
	GetUserActivity(username string) ([]string, error)
}
