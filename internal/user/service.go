package user

import "server/pkg/utils"

type Service interface {
	Register(input RegisterDTO) error
	Login(input LoginDTO) (string, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) Register(input RegisterDTO) error {
	hashed, err := utils.HashPassword(input.Password)

	if err != nil {
		return err
	}

	user := &User{
		Email: input.Email,
		Name: input.Name,
		Email: input.Email,
	}
}

func Login() {

}