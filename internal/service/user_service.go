package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/michaelrodriguess/user-service/internal/model"
	"github.com/michaelrodriguess/user-service/internal/repository"
	"github.com/michaelrodriguess/user-service/pkg/client"
)

type UserService struct {
	repo       *repository.UserRepository
	authClient *client.AuthClient
}

func NewUserService(repo *repository.UserRepository, authClient *client.AuthClient) *UserService {
	return &UserService{
		repo:       repo,
		authClient: authClient,
	}
}

func (s *UserService) CreateUserService(req model.UserRequest) (*model.UserResponse, error) {
	existingUser, _ := s.repo.FindByEmail(req.Email)
	if existingUser != nil {
		return nil, errors.New("email already registered")
	}

	if req.Role == "" {
		req.Role = "user"
	}

	user := model.User{
		ID:          uuid.New().String(),
		Name:        req.Name,
		Email:       req.Email,
		Role:        req.Role,
		User_Status: true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err := s.repo.CreateUserRepository(&user)
	if err != nil {
		return nil, err
	}

	token, err := s.authClient.GenerateToken(user.Email, req.Password, user.Role)
	if err != nil {
		return nil, err
	}

	return &model.UserResponse{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		Role:        user.Role,
		User_Status: user.User_Status,
		AccessToken: token,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}, nil
}

func (s *UserService) GetAllAdminsUser() ([]model.GetsUsersResponse, error) {

	users, err := s.repo.GetAllAdminsUser()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) GetAllUsers() ([]model.GetsUsersResponse, error) {

	users, err := s.repo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}
