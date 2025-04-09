package repository

import (
	"github.com/michaelrodriguess/user-service/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUserRepository(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetAllAdminsUser() ([]model.GetsUsersResponse, error) {
	var users []model.User

	err := r.db.Where("role = 'admin'").Find(&users).Error
	if err != nil {
		return nil, err
	}

	var adminsResponse []model.GetsUsersResponse

	i := 0
	for i < len(users) {
		adminsResponse = append(adminsResponse, model.GetsUsersResponse{
			ID:          users[i].ID,
			Name:        users[i].Name,
			Email:       users[i].Email,
			Role:        users[i].Role,
			User_Status: users[i].User_Status,
			CreatedAt:   users[i].CreatedAt,
			UpdatedAt:   users[i].UpdatedAt,
		})
		i++
	}

	return adminsResponse, nil
}

func (r *UserRepository) GetAllUsers() ([]model.GetsUsersResponse, error) {
	var users []model.User

	err := r.db.Where("status = true").Find(&users).Error
	if err != nil {
		return nil, err
	}

	var usersResponse []model.GetsUsersResponse

	i := 0
	for i < len(users) {
		usersResponse = append(usersResponse, model.GetsUsersResponse{
			ID:          users[i].ID,
			Name:        users[i].Name,
			Email:       users[i].Email,
			Role:        users[i].Role,
			User_Status: users[i].User_Status,
			CreatedAt:   users[i].CreatedAt,
			UpdatedAt:   users[i].UpdatedAt,
		})
		i++
	}

	return usersResponse, nil
}
