package user

import (
	"errors"
	"server/pkg/utils"
	"sync"
)

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
		Email:    input.Email,
		Name:     input.Name,
		Password: hashed,
		Role:     "user",
	}

	if err := s.repo.Create(user); err != nil {
		return err
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		utils.SendEmail(user.Email, "Регистрация", "Добро пожаловать в GOSinvest!")
	}()

	go func() {
		defer wg.Done()
		utils.SendSMS(user.Email, "Вы зарегестрировались")
	}()

	wg.Wait()
	return nil
}

func (s *service) Login(input LoginDTO) (string, error) {
	user, err := s.repo.FindByEmail(input.Email)
	if err != nil || !utils.CheckPassword(input.Password, user.Password) {
		return "", errors.New("Invalit credentials")
	}
	return utils.GenerateToken(user.ID, user.Role)
}
