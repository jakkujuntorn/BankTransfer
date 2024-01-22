package account // การตั้งชื่อต้องต้ังแบบนี้ด้วย ***

import (
	"banktransfer/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type I_Repo_Account interface {
	CreateAccount(*models.CreateAccountParams) error //  1 user เปิดได้ไม่เกิน 3 account
	GetAccount(id int, owner string) (*models.Account, error)
	ListAccount(owner string) ([]models.Account, error)
	DeleteAccount(id int) error
}

type repo_Account struct {
	db *gorm.DB
}

func New_Repo_Account(db *gorm.DB) I_Repo_Account {
	return &repo_Account{db}
}

func (*repo_Account) DeleteAccount(id int) error {
	return nil
}

//
func (ra *repo_Account) CreateAccount(account *models.CreateAccountParams) error {
	fmt.Println("CreateAccount Server ..................")
	// *********** เงื่อนไขการ create account ************
	//เปิดได้แค่ 3 account
	// ห้ามเปิด บัญชีที่สกุลเงินซ้ำกัน


	var ownerCount int64

	tx := ra.db.Table("accounts").Where("owner = ?", account.Owner).Count(&ownerCount)
	if tx.Error != nil {
		return tx.Error
	}

	// fmt.Println(ownerCount)
	// เช็คจำนวน  account มีได้แค่ 3 account
	if ownerCount >= 3 {
		fmt.Println("limit of create account *******")
		return errors.New("limit of create account")
	}

	tx = ra.db.Table("accounts").Create(&account)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// GetAccount implements I_Repo_Account
// ดึง account อันเดียว ของ owner นั้น
func (ra *repo_Account) GetAccount(id int, owner string) (*models.Account, error) {
	// fmt.Println("GetAccount server REpo")
	
	dataAccount := new(models.Account)
	tx := ra.db.Table("accounts").Where("id =?", id).Where("owner=?", owner).First(dataAccount)
	if tx.Error != nil {
		return dataAccount, tx.Error
	}

	return dataAccount, nil
}

// ListAccount implements I_Repo_Account
// ดึง account ทั้งหมดของ  owner นั้นๆ
func (ra *repo_Account) ListAccount(owner string) ([]models.Account, error) {
	dataAccount := []models.Account{}
	tx := ra.db.Table("accounts").Where("owner =?", owner).Find(&dataAccount)
	if tx.Error != nil {
		return dataAccount, tx.Error
	}
	return dataAccount, nil
}
