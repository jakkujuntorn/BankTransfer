package main

import (
	"banktransfer/account"
	"banktransfer/entries"
	repo_transfers "banktransfer/transfers/repo"
	_ "banktransfer/users"

	"banktransfer/models"
	"banktransfer/util"

	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	_ "time"

	"banktransfer/db"
)

func main() {

	fmt.Printf("")
	// ***************  Load config  **************
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	//***************** DB ************************
	gormDB := db.Postgres_init()
	fmt.Println("Started Postgres .....")
	_ = gormDB

	// Create user
	// dt := models.User{
	// 	Username:          "russy",
	// 	HashedPassword:    "1150",
	// 	FullName:          "russy five",
	// 	Email:             "russy@gmail.com",
	// 	PasswordChangedAt: time.Now().Local().UTC(),
	// 	CreatedAt:         time.Now().Local().UTC(),
	// 	IsEmailVerified:   false,
	// }
	// repo_User := users.New_Repo_User(gormDB)
	// errREpo_User:=repo_User.CreateUser(&dt)
	// if errREpo_User != nil {
	// 	fmt.Println(errREpo_User)
	// }

	// Create Account **************************
	// 	da:= models.CreateAccountParams{
	// 		Owner: "russy",
	// 		Balance: 3000,
	// 		Currency: "USD",
	// 		CreatedAt: time.Now().Local(),
	// 	}
	repo_account := account.New_Repo_Account(gormDB)

	_ = repo_account

	// 	errRepo_Account:= repo_account.CreateAccount(&da)
	// if errRepo_Account != nil {
	// 	fmt.Println(errRepo_Account)
	// }

	//  update account **************************
	// ua:=models.UpdateAccountParams{
	// 	ID: 4,
	// 	Balance: -100,
	// }
	// errUpdate_Blance:=repo_account.UpdateAccount_Blance(&ua)
	// fmt.Println(errUpdate_Blance)

	// get account *****************************
	// data, err := repo_account.GetAccount(7, "russy")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(data)

	// list account *************************
	// data2, err := repo_account.ListAccount("john5")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(data2)

	// Tansfer *******************************
	repo_Transfer := repo_transfers.New_Repo_User(gormDB)
	_ = repo_Transfer

	ctf := models.CreateTransferParams{
		Owner:         "russy",
		FromAccountID: 8, // 7-3, 8-1
		ToAccountID:   1,
		Amount:        500,
	}
	_ = ctf
	// errRepoTransfer := repo_Transfer.Create_Transfer(&ctf)
	// if errRepoTransfer != nil {
	// 	fmt.Println(errRepoTransfer)
	// }

	// ฝาก ****************************************
	// deposit := models.Create_Deposit_and_Withdraw{
	// 	Owner:     "john5",
	// 	AccountID: 1,
	// 	Amount:    100,
	// }
	// err = repo_Transfer.Create_Deposit(&deposit)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	//ถอน *****************************************
	// withdraw := models.Create_Deposit_and_Withdraw{
	// 		Owner:     "john5",
	// 		AccountID: 1,
	// 		Amount:    100,
	// 	}
	// 	err = repo_Transfer.Create_Withdraw(&withdraw)
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 	}

	// Get transfer by id ****************************************
	// id, errGet := repo_Transfer.GetTransfer_ById(3,"2023-01-01" , "2023-10-20") // 3,7
	// if errGet != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(id)

	// fmt.Println("*********************************")
	// // Get transfer by owner *************************************
	// owner, err := repo_Transfer.GetTransfer_ByOwner("russy","2023-01-01" , "2023-10-20") // "john50", "russy"
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(len(owner))

	// ListTransfers *******************
	statement_Repo := entries.New_Repo_Entries(gormDB)

	statamentData, err := statement_Repo.GetStaement(1, "", "")
	
	if len(statamentData) == 0 {
		fmt.Println("record not found")
	}
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, v := range statamentData {
		fmt.Println(v)
	}

}
