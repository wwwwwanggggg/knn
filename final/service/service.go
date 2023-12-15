package service

type Service struct {
	Hello
	User
	Admin
}

func New() *Service {
	service := &Service{}
	return service
}
