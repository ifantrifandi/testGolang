package user

type Service interface {
	FindUser() []User
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindUser() []User{
	users := s.repository.Find()

	return users
}