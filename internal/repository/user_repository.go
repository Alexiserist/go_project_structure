package repository

import (
	"go_project_structure/auth"
	"go_project_structure/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]models.User, error)
	CreateUser(user models.User) (models.User, error)
	DeleteUser(user models.User) (error)
	FindOneByKey(id uint) (models.User,error)
}

type userRepository struct{
	db *gorm.DB
	authService auth.AuthService
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepository{
		authService: auth.NewAuthRepository(),
		db : db,
	}
}

func (r *userRepository) FindAll() ([]models.User, error){
	var user []models.User
	if err:= r.db.Find(&user).Error; err != nil {
		return nil,err
	}
	return user, nil
}

func (r *userRepository) FindOneByKey(id uint) (models.User,error){
	var user models.User
	err := r.db.First(&user, id).Error;
	if err != nil {
		return user, err
	}
	return user,nil
}


func (r *userRepository) CreateUser(user models.User) (models.User,error) {
	encoded, err := r.authService.EncodingPassword(user.Password);
	if err != nil {
		return user,err
	}
	user.Password = encoded;
	if err := r.db.Save(&user).Error; err != nil {
		return user, err
	}
	return user,nil
}


func (r *userRepository) DeleteUser(user models.User) (error) {
	if err := r.db.Delete(&user,user).Error; err != nil {
		return err
	}
	return nil
}