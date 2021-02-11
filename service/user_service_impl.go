package service
import (
	"github.com/hariszaki17/go-api-clean/repository"
	"github.com/hariszaki17/go-api-clean/model"
	"github.com/hariszaki17/go-api-clean/entity"
)

// NewUserService expose global
func NewUserService(userRepository *repository.UserRepository) UserService {
	return &userServiceImpl{
		UserRepository: *userRepository,
	}
}

type userServiceImpl struct  {
	UserRepository repository.UserRepository
}

func (service *userServiceImpl) Create(request model.CreateUserRequest) (response model.CreateUserResponse)  {
	user := entity.User{
		Username:	request.Username,
		Password:	service.UserRepository.Encrypt(request.Password),
		Role:		request.Role,
	}

	service.UserRepository.Insert(user)

	response = model.CreateUserResponse{
		Username:	user.Username,
		Role:		user.Role,
	}
	return response
}

func (service *userServiceImpl) List() (responses []model.GetUserResponse) {
	users :=  service.UserRepository.FindAll()
	for _, user := range users {
		responses = append(responses, model.GetUserResponse{
			ID:			user.ID,
			Username:	user.Username,
			Role:		user.Role,
		})
	}
	return responses
}

func (service *userServiceImpl) DeleteAll() (response model.DeleteAllUserResponse) {
	service.UserRepository.DeleteAll()
	response = model.DeleteAllUserResponse{
		Message: "Successfully deleted all record",
	}
	return response
}

func (service *userServiceImpl) validatePassword(username, password string) {
	user := entity.User{
		Username:	username,
		Password:	password,
	}
	service.UserRepository.ValidatePassword(user.Username, user.Password)
}

func (service *userServiceImpl) Login(request model.LoginUserRequest) (response model.LoginUserResponse) {
	user := entity.User{
		Username:	request.Username,
		Password:	request.Password,
	}
	service.validatePassword(user.Username, user.Password)
	return model.LoginUserResponse{
		Message:	"Berhasil Login",
	}
}