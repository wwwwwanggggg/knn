package controller

type Controller struct {
	Hello
	User
	Admin
}

func New() *Controller {
	Controller := &Controller{}
	return Controller
}
