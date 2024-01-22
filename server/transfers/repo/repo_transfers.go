package repo_transfers

import (
	"banktransfer/models"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"

	transferProto "banktransfer/transfers"
)

type I_Repo_Transfers interface {
	// func ที่ต้องรับ struct เพราะต้องใช้งาน DB ไม่ใช้ใช้งาน interface
	transaction_Postgres(func(*repo_Transfers) error) error

	// โอน เดิม
	Create_Transfer(dataTransfer *models.CreateTransferParams) error

	// โอน Proto 
	// ใช้ ไม่ได้ เพราะ struct ของ proto มัน error ถ้าใช้ตรงๆ ******
	Create_Transfer_Proto(dataTransfer *transferProto.Create_TransferRequest) error
	//ฝาก
	Create_Deposit(dataDeposit *models.Create_Deposit_and_Withdraw) error
	// ถอน
	Create_Withdraw(dataWithdraw *models.Create_Deposit_and_Withdraw) error

	// เช็คการทำ transfers ทั้งหมด ฝาก, โอน, ถอน
	// ปรับ entries แล้วไปดึงค่าจาก entries แทน *****************
	GetTransfer_ById(idAccount int64, owner string, startTime, endTime string) ([]models.ListTransfers_ForOwner, error)

	// เช็คการโอนทั้งหมด (Transfer เท่านั้น) ของเราทุกบัญชี (ตาม owner) และเลือกช่วงเวลา
	// เช็คใน transfers DB (จะรู้ว่าดราโอนไปให้ใครหรือ ใครโอนมาให้เรา)
	// ปรับ entries แล้วไปดึงค่าจาก entries แทน ******************
	GetTransfer_ByOwner(ownerAccount, startTime, endTime string) ([]models.ListTransfers_ForOwner, error)

	// ไม่ได้ใช้แล้ว
	ListStatement(accountId int64) ([]models.Transfer, error)

	// หบัวจาก ฝาก,ถอน, โอน ต้องมาอัพเดท
	UpdateAccount_Blance([]models.UpdateAccountParams) error

	// **************************** Check Data  Core ****************
	checkDataFor_Transfer(dataCheck *models.CreateTransferParams) (destinationAccountData *models.Account, err error)
	// ไม่ได้ใช้ *******************
	checkDataFor_Transfer_Proto(dataCheck *transferProto.Create_TransferRequest) (destinationAccountData *models.Account, err error)
	checkDataFor_Deposit(dataCheck *models.Create_Deposit_and_Withdraw) error
	checkDataFor_Withdraw(dataCheck *models.Create_Deposit_and_Withdraw) error

	// Check data inside ****************

	// เช็ค account คนที่จะโอน หรือ เช็ค account คนจะฝากและถอน
	checkAccountOwner(owner string, id int64) error
	// เช็ค account คนที่จะรับโอน
	checkAccountTo(id int64) error
	// เช็ค สกุลเงิน
	checkCurrency(fromAccount, toAccount string) error
	// เช็ค จำนนเงินคนที่จะโอน
	checkAccountBalance(amount, balbnce int64) error
}

type repo_Transfers struct {
	db *gorm.DB
}

// ตรงนี้ return repo_Transfers  เพราะ ต้องใช้ตรงทำ transaction_Postgres
// ถ้า return interface จะ ทำ transaction_Postgres ไม่ได้ *************
func New_Repo_Transfer(db *gorm.DB) *repo_Transfers {
	return &repo_Transfers{
		db: db,
	}
}

// ฝากเงิน *******************
func (ra *repo_Transfers) Create_Deposit(dataDeposit *models.Create_Deposit_and_Withdraw) error {
	// check data for Deposit เช็คแค่ว่ามีบัญชีที่จะโอนหรือไม่ *******
	erCheckData := ra.checkDataFor_Deposit(dataDeposit)
	if erCheckData != nil {
		return erCheckData
	}

	// ปั้นข้อมูลใหม่ สำหรับ Update balance
	new_dataDeposit := []models.UpdateAccountParams{
		{
			ID:      dataDeposit.AccountID,
			Balance: dataDeposit.Amount,
		},
	}

	//************* ทำ transaction *********************
	return ra.transaction_Postgres(func(p *repo_Transfers) error {
		//create entries ********************
		dataEntries := models.Entry{
			// FromAccountID *********************
			AccountID: dataDeposit.AccountID,
			Amount:    dataDeposit.Amount,
			// CreatedAt: time.Now(), // มันใสค่าให้เองถึงเราไม่ใสค่าให้ ***
			Entries_type:        "deposit",
			Owner:               dataDeposit.Owner,
			Destination_account: dataDeposit.AccountID,
		}

		tx := p.db.Table("entries").Create(&dataEntries)
		if tx.Error != nil {
			fmt.Println("entries naja")
			return tx.Error
		}

		// update accounts ****************
		errUpdate := ra.UpdateAccount_Blance(new_dataDeposit)
		if errUpdate != nil {
			return errUpdate
		}

		// fmt.Println("Deposit Complete")
		return nil
	})
}

// ถอนเงิน *****************
func (ra *repo_Transfers) Create_Withdraw(dataWithdraw *models.Create_Deposit_and_Withdraw) error {
	// check บัญชีที่จะถอน และจำนวนเงิน
	erCheckData := ra.checkDataFor_Withdraw(dataWithdraw)
	if erCheckData != nil {
		return erCheckData
	}

	// ปั้นข้อมูลใหม่ สำหรับ Update balance *************
	new_dataDeposit := []models.UpdateAccountParams{
		{
			ID:      dataWithdraw.AccountID,
			Balance: -dataWithdraw.Amount,
		},
	}

	//************* ทำ transaction *********************
	return ra.transaction_Postgres(func(p *repo_Transfers) error {
		//create entries ********************
		dataEntries := models.Entry{
			AccountID:    dataWithdraw.AccountID,
			Amount:       -dataWithdraw.Amount,
			CreatedAt:    time.Now(),
			Entries_type: "withdraw",
			Owner:        dataWithdraw.Owner,
		}

		// create entries *********************
		tx := p.db.Table("entries").Create(&dataEntries)
		if tx.Error != nil {
			fmt.Println(tx.Error.Error())
			return tx.Error
		}

		// update accounts ****************
		errUpdate := ra.UpdateAccount_Blance(new_dataDeposit)
		if errUpdate != nil {
			return errUpdate
		}
		// fmt.Println("Withdraw Complete")
		return nil
	})

}

func (ra *repo_Transfers) checkDataFor_Deposit(dataCheck *models.Create_Deposit_and_Withdraw) error {
	// เช็คบัญชีว่ามีอนู่รึป่าว
	// tx := ra.checkAccountOwner(dataCheck.Owner, dataCheck.AccountID)
	// if tx != nil {
	// 	return tx
	// }
	// return nil

	return ra.checkAccountOwner(dataCheck.Owner, dataCheck.AccountID)
}
func (ra *repo_Transfers) checkDataFor_Withdraw(dataCheck *models.Create_Deposit_and_Withdraw) error {

	// เช็คบัญชีที่จะถอนวาสมีอยู๋ริงไหม ******************
	check_Account := new(models.CreateAccountParams)
	tx := ra.db.Table("accounts").Where("owner =?", dataCheck.Owner).Where("id =?", dataCheck.AccountID).First(&check_Account)
	if tx.Error != nil {
		return tx.Error
	}

	// เช็คจำนวนเงินที่จะถอนว่าพอรึป่าว *****************
	if check_Account.Balance < dataCheck.Amount {
		return errors.New("money not enough")
	}

	return nil
}

func (ra *repo_Transfers) UpdateAccount_Blance(accountUpdate []models.UpdateAccountParams) error {

	// logic อันนี้ได้ แต่ดีรึป่าวไม่รู้ *************************
	// accountParams.Balance = accountParams.Balance + accountUpdate.Balance
	// tx = ra.db.Table("accounts").Where("id =?", accountUpdate.ID).Updates(&accountParams)
	// if tx.Error != nil {
	// 	return tx.Error
	// }

	// หรือใช้ query update  ไปเลย ใช้  blance = blance+ ค่าใหม่
	myQuery := `UPDATE  accounts SET  balance = balance+? WHERE  id = ?`
	for _, v := range accountUpdate {
		tx := ra.db.Exec(myQuery, v.Balance, v.ID)
		if tx.Error != nil {
			return tx.Error
		}
	}

	return nil
}

func (rt *repo_Transfers) checkDataFor_Transfer(dataCheck *models.CreateTransferParams) (destinationAccountData *models.Account, err error) {
	// ในแต่ละ func มันจะ  query ทุก func ซึ่งมันเยอะไป
	//  query ที่เดียวแล้วส่งค่าเข้าไปเช็คแทน

	// query From Account ********************
	fromAccount := new(models.Account)
	rt.db.Table("accounts").Where("id=?", dataCheck.FromAccountID).Where("owner=?", dataCheck.Owner).First(fromAccount)

	// query To Account **********************
	toAccount := new(models.Account)
	rt.db.Table("accounts").Where("id=?", dataCheck.ToAccountID).First(toAccount)
	// เอาข้อมูลที่ได้ส่งเข้าไปใน func ต่างๆ เพื่อเช็ค error

	var transferErr error
	// เช็ค error ต่างๆ ******************
	// เช็ค owner, accoutTo, currency, blance
	if transferErr = rt.checkAccountOwner(fromAccount.Owner, fromAccount.ID); transferErr == nil {
		if transferErr = rt.checkAccountTo(fromAccount.ID); transferErr == nil {
			if transferErr = rt.checkCurrency(fromAccount.Currency, toAccount.Currency); transferErr == nil {
				if transferErr = rt.checkAccountBalance(dataCheck.Amount, fromAccount.Balance); transferErr == nil {
					return toAccount, nil
				}
			}
		}
	}

	if transferErr != nil {
		errorTese := fmt.Sprintf(transferErr.Error())
		return &models.Account{}, errors.New(errorTese)
	}

	// ต้องเช็ค  owner  และ account ว่ามีอยู๋จริงไหม  เช็คจาก accounts
	// เอา id กับ owner มาเช็ค
	// อันนี้ใช้รวมกัน การเช็ค checkDataFor_Deposit_and_Withdraw *************************
	// from_Account := new(models.CreateAccountParams)
	// tx := rt.db.Table("accounts").Where("owner =?", dataCheck.Owner).Where("id =?", dataCheck.FromAccountID).First(&from_Account)
	// if tx.Error != nil {
	// 	// ไม่เจอบัญชี สำหรับโอนเงิน
	// 	errorTese := fmt.Sprintf(tx.Error.Error() + " for:" + "From account")
	// 	return errors.New(errorTese)
	// 	// return tx.Error
	// }

	// เช็ค จำนวนเงินที่จะโอน ของ  from_Account *************
	// if from_Account.Balance < dataCheck.Amount {
	// 	errorTese := fmt.Sprintf("money not enough")
	// 	return errors.New(errorTese)
	// }

	// เช็ค  account ปลายทางด้วยว่ามีจริงไหม เช็คจาก accounts
	// เอา id มาเช็ค
	// to_Account := new(models.CreateAccountParams)
	// tx = rt.db.Table("accounts").Where("id =?", dataCheck.ToAccountID).First(&to_Account)
	// if tx.Error != nil {
	// 	// ไม่เจอ บัญชีปลายทางสำหรับโอน
	// 	errorTese := fmt.Sprintf(tx.Error.Error() + " for:" + "To account")
	// 	return errors.New(errorTese)
	// 	// return tx.Error
	// }

	// เช็ค current ว่า ปะเภทเดียวกันรึป่าว ********
	// if from_Account.Currency != to_Account.Currency {
	// 	return errors.New("cruuency not match")
	// }

	return &models.Account{}, nil
}

// proto ***************************
func (rt *repo_Transfers) checkDataFor_Transfer_Proto(dataCheck *transferProto.Create_TransferRequest) (destinationAccountData *models.Account, err error) {

	fmt.Println("Create_Transfer ***************")
	// ในแต่ละ func มันจะ  query ทุก func ซึ่งมันเยอะไป
	//  query ที่เดียวแล้วส่งค่าเข้าไปเช็คแทน

	// query From Account ********************
	fromAccount := new(models.Account)
	rt.db.Table("accounts").Where("id=?", dataCheck.GetFromAccountID()).Where("owner=?", dataCheck.GetOwner()).First(fromAccount)

	// query To Account **********************
	toAccount := new(models.Account)
	rt.db.Table("accounts").Where("id=?", dataCheck.GetToAccountID()).First(toAccount)
	// เอาข้อมูลที่ได้ส่งเข้าไปใน func ต่างๆ เพื่อเช็ค error

	var transferErr error
	// เช็ค error ต่างๆ ******************
	// เช็ค owner, accoutTo, currency, blance
	if transferErr = rt.checkAccountOwner(fromAccount.Owner, fromAccount.ID); transferErr == nil {
		if transferErr = rt.checkAccountTo(fromAccount.ID); transferErr == nil {
			if transferErr = rt.checkCurrency(fromAccount.Currency, toAccount.Currency); transferErr == nil {
				if transferErr = rt.checkAccountBalance(int64(dataCheck.GetAmount()), fromAccount.Balance); transferErr == nil {
					return toAccount, nil
				}
			}
		}
	}

	if transferErr != nil {
		errorTese := fmt.Sprintf(transferErr.Error())
		return &models.Account{}, errors.New(errorTese)
	}

	return &models.Account{}, nil
}

// CreateTransfer implements I_Repo_Transfers
func (rt *repo_Transfers) Create_Transfer(dataTransfer *models.CreateTransferParams) error {

	// Check Data for Transfers
	// ตรงนี้มีข้อมูล  ของ ผู้รับโอนด้วย ***************
	// อันเก่าใช้กับ models ปกติ
	destinationAccountData, errCheck := rt.checkDataFor_Transfer(dataTransfer)
	if errCheck != nil {
		return errCheck
	}

	//****************** entries ***************
	// from Account ต้อง - blance
	// to Account ต้อง + blance
	// ทำเป็น array  จะได้ create  ครั้งเดียว
	// dataEntries  อันเก่า ********************************
	dataEntries := []models.Entry{
		{
			// FromAccountID *********************
			AccountID:           dataTransfer.FromAccountID,
			Amount:              -dataTransfer.Amount,
			CreatedAt:           time.Now(),
			Entries_type:        "transfer",
			Owner:               dataTransfer.Owner,
			Destination_account: dataTransfer.ToAccountID,
		},
		{
			// ToAccountID  *******************
			AccountID:           dataTransfer.ToAccountID,
			Amount:              dataTransfer.Amount,
			CreatedAt:           time.Now(),
			Entries_type:        "transfer",
			Owner:               destinationAccountData.Owner,
			Destination_account: dataTransfer.FromAccountID,
		},
	}

	// *************** accounts *****************
	// dataAccount อันเก่า *****************************
	dataAccount := []models.UpdateAccountParams{
		{
			// FromAccountID  *******************
			ID:      dataTransfer.FromAccountID,
			Balance: -dataTransfer.Amount,
		},
		{
			// ToAccountID  *******************
			ID:      dataTransfer.ToAccountID,
			Balance: dataTransfer.Amount,
		},
	}

	// ใสหรือไม่ใสก็ได้มั้ง *******************
	// dataTransfer.CreatedAt = time.Now()

	// from Account ต้อง - blance
	// to Account ต้อง + blance
	// ทำเป็น array  จะได้ create  ครั้งเดียว
	//**************** ทำ TRanscation  ****************
	return rt.transaction_Postgres(func(p *repo_Transfers) error {
		// Create Trasfer to DB ***********************
		tx := p.db.Table("transfers").Create(&dataTransfer)
		if tx.Error != nil {
			// fmt.Println("TRnsfer naja")
			return tx.Error
		}

		// create Entries  to DB   **************************
		// อันนี้ อาจไม่ต้องใช้แล้ว
		//ปรับ DB Transfer ใหม่ มันรองรับ entries แล้ว ***************
		tx = p.db.Table("entries").Create(&dataEntries)
		if tx.Error != nil {
			// fmt.Println("entries naja")
			return tx.Error
		}

		// update Accounts to DB  ***************************
		errUpdate := p.UpdateAccount_Blance(dataAccount)
		if errUpdate != nil {
			// fmt.Println("accounts naja")
			return errUpdate
		}
		// fmt.Println("Transfer is complete")
		return nil
	})

}

// Proto ***************
func (rt *repo_Transfers) Create_Transfer_Proto(dataTransfer *transferProto.Create_TransferRequest) error {

	// Check Data for Transfers
	// ตรงนี้มีข้อมูล  ของ ผู้รับโอนด้วย ***************
	// อันเก่าใช้กับ models ปกติ
	destinationAccountData, errCheck := rt.checkDataFor_Transfer_Proto(dataTransfer)
	if errCheck != nil {
		return errCheck
	}

	//****************** entries ***************
	// from Account ต้อง - blance
	// to Account ต้อง + blance
	// ทำเป็น array  จะได้ create  ครั้งเดียว

	// dataEntries  อันใหม่  ********************************
	dataEntries := []models.Entry{
		{
			// FromAccountID *********************
			AccountID:           int64(dataTransfer.GetFromAccountID()),
			Amount:              -int64(dataTransfer.GetAmount()),
			CreatedAt:           time.Now(),
			Entries_type:        "transfer",
			Owner:               dataTransfer.GetOwner(),
			Destination_account: int64(dataTransfer.GetToAccountID()),
		},
		{
			// ToAccountID  *******************
			AccountID:           int64(dataTransfer.GetToAccountID()),
			Amount:              int64(dataTransfer.GetAmount()),
			CreatedAt:           time.Now(),
			Entries_type:        "transfer",
			Owner:               destinationAccountData.Owner,
			Destination_account: int64(dataTransfer.GetFromAccountID()),
		},
	}

	// *************** accounts *****************
	// dataAccount อันใหม่ *****************************
	dataAccount := []models.UpdateAccountParams{
		{
			// FromAccountID  *******************
			ID:      int64(dataTransfer.GetFromAccountID()),
			Balance: -int64(dataTransfer.GetAmount()),
		},
		{
			// ToAccountID  *******************
			ID:      int64(dataTransfer.GetToAccountID()),
			Balance: int64(dataTransfer.GetAmount()),
		},
	}

	// ใสหรือไม่ใสก็ได้มั้ง *******************
	// dataTransfer.CreatedAt = time.Now()

	// from Account ต้อง - blance
	// to Account ต้อง + blance
	// ทำเป็น array  จะได้ create  ครั้งเดียว
	//**************** ทำ TRanscation  ****************
	return rt.transaction_Postgres(func(p *repo_Transfers) error {
		// Create Trasfer to DB ***********************
		fmt.Println("transaction_Postgres ***************")
		tx := p.db.Table("transfers").Create(&dataTransfer)
		if tx.Error != nil {
			// fmt.Println("TRnsfer naja")
			return tx.Error
		}

		// create Entries  to DB   **************************
		// อันนี้ อาจไม่ต้องใช้แล้ว
		//ปรับ DB Transfer ใหม่ มันรองรับ entries แล้ว ***************
		tx = p.db.Table("entries").Create(&dataEntries)
		if tx.Error != nil {
			// fmt.Println("entries naja")
			return tx.Error
		}

		// update Accounts to DB  ***************************
		errUpdate := p.UpdateAccount_Blance(dataAccount)
		if errUpdate != nil {
			// fmt.Println("accounts naja")
			return errUpdate
		}
		// fmt.Println("Transfer is complete")
		return nil
	})

}

// Transaction_Postgres implements I_Repo_Transfers
func (rt *repo_Transfers) transaction_Postgres(fn func(*repo_Transfers) error) error {
	fmt.Println("")
	dbPostgres := rt.db.Begin()

	newRepo := New_Repo_Transfer(dbPostgres)

	// func ต้องได้ struct ถึงจะเรียกใช้ recive Func ของมันได้
	// ภายใน func นี้ถ้ามีการเรียกใช้ query แล้วมี  error ออกมามันจะไม่ commit ให้
	// transcation ทั้งหมดเลยถูกยกเลิก
	err := fn(newRepo)

	// คีร์หลักคือตรงนี้ ถ้าใน การทำ transcation มี error มันจะมาหยุดตรงนี้ และไม่มีถึ
	// การทำงานของ commit
	if err != nil {
		fmt.Println("Error in Tracscation")
		dbPostgres.Rollback()
		return err
	}

	dbPostgres.Commit()
	return nil
}

// GetTransfer implements I_Repo_Transfers
// เช็คการโอนทั้งหมด ของบัญชีนั้นๆ (ตาม account) และเลือกช่วงเวลา
func (rt *repo_Transfers) GetTransfer_ById(idAccount int64, owner string, startTime, endTime string) (dataTransfers []models.ListTransfers_ForOwner, err error) {

	tx := rt.db.Table("entries").Where("account_id=?", idAccount).
		Where("owner = ?", owner).
		Where("created_at between ? and ?", startTime, endTime).Find(&dataTransfers)
	if tx.Error != nil {
		return dataTransfers, tx.Error
	}
	return dataTransfers, nil
}

// เช็คการโอนทั้งหมด ของเราทุกบัญชี (ตาม owner) และเลือกช่วงเวลา
func (rt *repo_Transfers) GetTransfer_ByOwner(ownerAccount, startTime, endTime string) (dataTransfers []models.ListTransfers_ForOwner, err error) {
	tx := rt.db.Table("entries").Where("owner = ?", ownerAccount).
		Where("created_at between ? and ?", startTime, endTime).
		Find(&dataTransfers)
	if tx.Error != nil {
		return dataTransfers, tx.Error
	}
	return dataTransfers, nil
}

// ListTransfers implements I_Repo_Transfers
func (rt *repo_Transfers) ListStatement(accountId int64) (statement []models.Transfer, err error) {
	tx := rt.db.Table("entries").Where("account_id = ?", accountId).Find(&statement)
	if tx.Error != nil {
		return statement, tx.Error
	}
	return statement, nil
}

// เช็ค account คนที่จะโอน หรือ เช็ค account คนจะฝากและถอน
func (rt *repo_Transfers) checkAccountOwner(owner string, id int64) error {
	from_Account := new(models.CreateAccountParams)
	tx := rt.db.Table("accounts").Where("owner =?", owner).Where("id =?", id).First(&from_Account)
	if tx.Error != nil {
		errorTese := fmt.Sprintf(tx.Error.Error() + " for:" + "From account")
		return errors.New(errorTese)
	}
	return nil
}

// เช็ค account คนที่จะรับโอน
func (rt *repo_Transfers) checkAccountTo(id int64) error {
	to_Account := new(models.CreateAccountParams)
	tx := rt.db.Table("accounts").Where("id =?", id).First(&to_Account)
	if tx.Error != nil {
		errorTese := fmt.Sprintf(tx.Error.Error() + " for:" + "To account")
		return errors.New(errorTese)
	}
	return nil
}

// เช็ค สกุลเงิน
func (rt *repo_Transfers) checkCurrency(fromAccount, toAccount string) error {
	// check currency
	if fromAccount != toAccount {
		return errors.New("cruuency not match")
	}
	return nil
}

// เช็ค จำนนเงินคนที่จะโอน
func (rt *repo_Transfers) checkAccountBalance(amount, balance int64) error {
	// เช็คเงินในการโอน
	if balance < amount {
		errorTese := fmt.Sprintf("money not enough")
		return errors.New(errorTese)
	}
	return nil
}
