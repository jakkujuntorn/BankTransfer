package users

import (
	"banktransfer/models"
	"gorm.io/gorm"
)

type I_Rero_Users interface {
	CreateUser(data *models.User) error // create User อย่างดียว ไม่สร้าง Account
	GetUser_ByUsername(username string) (*models.UserResponse, error)
	UpdateUser(username string, newData *models.UserUpdate) error
	ChangePassword(newPassword *models.UserChangePassword) error
	// Transaction_Postgres(func(*repo_user) error) error
	Transaction_Postgres(func(*I_Rero_Users) error) error
}

type repo_user struct {
	db *gorm.DB
}

func New_Repo_User(db *gorm.DB) I_Rero_Users {
	return &repo_user{
		db: db,
	}
}

func (ru *repo_user) Transaction_Postgres(fn func(*I_Rero_Users) error) error {
	db := ru.db.Begin()

	newRepo := New_Repo_User(db)
	err := fn(&newRepo)
	if err != nil {
		db.Rollback()
		return err
	}

	db.Commit()
	return nil
}

// ChangePassword implements I_Rero_Users
func (*repo_user) ChangePassword(newPassword *models.UserChangePassword) error {
	panic("unimplemented")
}

// CreateUser implements I_Rero_Users
func (ru *repo_user) CreateUser(data *models.User) error {

	tx := ru.db.Table("users").Create(&data)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// GetUser_ById implements I_Rero_Users
func (ru *repo_user) GetUser_ByUsername(username string) (*models.UserResponse, error) {
	user := new(models.UserResponse)
	tx := ru.db.Where("username = ?", username).First(user)
	if tx.Error != nil {
		return &models.UserResponse{}, tx.Error
	}

	return user, nil
}

// UpdateUser implements I_Rero_Users
func (ru *repo_user) UpdateUser(username string, newData *models.UserUpdate) error {
	// find date  before
	dataFind := new(models.UserResponse)
	tx := ru.db.Where("username = ?", username).First(dataFind)
	if tx.Error != nil {
		return tx.Error
	}

	// update new data
	dataFind.FullName = newData.FullName
	dataFind.Email = newData.Email

	tx = ru.db.Where("username =?", username).Updates(dataFind)
	if tx.Error != nil {
		return tx.Error
	}

	return nil

}
