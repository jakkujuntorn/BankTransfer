package users

import (
	"banktransfer/models"
	userProto "banktransfer/users"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type I_Rero_Users interface {
	Login(*userProto.User_LoginRequest) (userLoginResponse *models.UserLogin, err error)
	CreateUser(data *models.User) error // create User อย่างดียว ไม่สร้าง Account
	GetUser_ByUsername(username string) (*models.UserResponse, error)
	UpdateUser(username string, newData *models.UserUpdate) error
	ChangePassword(data *userProto.Change_PasswordRequest) error
	
	// ไม่ได้ใช้ ******************
	// Transaction_Postgres(func(*repo_user) error) error
	Transaction_Postgres(func(*I_Rero_Users) error) error
}

type repo_user struct {
	db *gorm.DB
}

// User ในการเปิด บัญชี
func New_Repo_User(db *gorm.DB) I_Rero_Users {
	fmt.Println("")
	return &repo_user{
		db: db,
	}
}

// Login ****************
func (ru *repo_user) Login(user *userProto.User_LoginRequest) (userLoginResponse *models.UserLogin, err error) {
	userlogin := new(models.UserLogin)

	tx := ru.db.Table("users").Where("username = ?", user.GetUsername()).First(&userlogin)
	if tx.Error != nil {
		return &models.UserLogin{}, tx.Error
	}

	return userlogin, nil
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
func (ru *repo_user) ChangePassword(data *userProto.Change_PasswordRequest) error {
	// find date  before
	dataFind := new(models.User)
	tx := ru.db.Table("users").Where("username = ?", data.GetUsername()).First(dataFind)
	if tx.Error != nil {
		return tx.Error
	}

	
	// ปั้นข้อมูลใหม่ ******************
	dataFind.HashedPassword = data.GetNewpassword()
	dataFind.PasswordChangedAt = time.Now()

	tx = ru.db.Table("users").Where("username =?", data.GetUsername()).Updates(dataFind)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
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
	fmt.Println("GetUser_ByUsername Server REpo **************")
	user := new(models.UserResponse)

	tx := ru.db.Table("users").Where("username = ?", username).First(user)
	if tx.Error != nil {
		return &models.UserResponse{}, tx.Error
	}

	return user, nil
}

// UpdateUser implements I_Rero_Users
func (ru *repo_user) UpdateUser(username string, newData *models.UserUpdate) error {
	// find date  before
	dataFind := new(models.UserResponse)
	tx := ru.db.Table("users").Where("username = ?", username).First(dataFind)
	if tx.Error != nil {
		return tx.Error
	}

	// update new data
	dataFind.FullName = newData.FullName
	dataFind.Email = newData.Email

	tx = ru.db.Table("users").Where("username =?", username).Updates(dataFind)
	if tx.Error != nil {
		return tx.Error
	}

	return nil

}
