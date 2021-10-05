package service

import (
	"ginPractical/dto"
	"ginPractical/entity"
	"ginPractical/repository"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyCredential(email string, password string) interface{}
	CreateUser(user dto.UserCreateDTO) entity.User
	FindByEmail(email string) entity.User
	isDuplicateEmail(email string) bool
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRep repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRep,
	}
}

func (service *authService) VerifyCredential(email string, password string) interface{} {
	res := service.userRepository.VerifyCredentials(email, password)

	if v, ok := res.(entity.User); ok {
		comparePassword := comparePassword(v.Password, []byte(password))
		if v.Email == email && comparePassword {
			return res
		}
		return false
	}
	return false
}

func (service *authService) CreateUser(user dto.UserCreateDTO) entity.User {
	userToCreate := entity.User{
		Name:     user.Name,
		Password: user.Password,
		Email:    user.Email,
	}

	return userToCreate

}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
