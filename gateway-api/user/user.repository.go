package user

import (
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]User, error)
	FindByID(id uint) (User, error)
	Save(user User) (*User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) FindAll() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) FindByID(id uint) (User, error) {
	var user User
	err := r.db.First(&user, id).Error
	return user, err
}

func (r *userRepository) Save(user User) (*User, error) {
	err := r.db.Save(&user).Error
	savedUser := User{
        Id: user.Id,
        Name: user.Name,
        Email: user.Email,
        Is_Active: user.Is_Active,
        Password_changed_at: user.Password_changed_at,
        Updated: user.Updated,
        Created: user.Created,
    }
    return &savedUser, err
}
