package main

// import (
// 	accountProto "banktransfer/account"
// 	accountRepo "banktransfer/account/repo"
// 	accountService "banktransfer/account/service"

// 	tansferProto "banktransfer/transfers"
// 	tansferRepo "banktransfer/transfers/repo"
// 	tansferService "banktransfer/transfers/service"

// 	usersProto "banktransfer/users"
// 	usersRepo "banktransfer/users/repo"
// 	usersService "banktransfer/users/service"

// 	// "banktransfer/models"
// 	"banktransfer/util"

// 	"fmt"
// 	"os"
// 	_ "time"

// 	"github.com/rs/zerolog"
// 	"github.com/rs/zerolog/log"

// 	"banktransfer/db"

// 	"net"

// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/reflection"
// 	// "google.golang.org/grpc/credentials/insecure"
// )

// func main2() {

// 	fmt.Printf("")
// 	// ***************  Load config  **************
// 	config, err := util.LoadConfig(".")
// 	if err != nil {
// 		log.Fatal().Err(err).Msg("cannot load config")
// 	}

// 	if config.Environment == "development" {
// 		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
// 	}

// 	// Start Server  gRPC ****************
// 	option := []grpc.ServerOption{}
// 	// ฝั่ง server จะใช้ NewServer
// 	s := grpc.NewServer(option...)

// 	//สร้าง net.Listen ก่อน ทำแค่ฝั่ง server
// 	listener, err := net.Listen("tcp", ":50051")
// 	if err != nil {
// 		// log.Fatal(err.Error())
// 		fmt.Println(err.Error())
// 	}

// 	//***************** DB ************************
// 	gormDB := db.Postgres_init()
// 	fmt.Println("Started Postgres .....")
// 	_ = gormDB
// 	// ********************************************

// 	// ************* Account ***************
// 	// สร้าแค่ repo   ส่วน service ไปสร้างใน gRCPserver
// 	account_Repo := accountRepo.New_Repo_Account(gormDB)

// 	// ********************** User *******************
// 	user_Repo := usersRepo.New_Repo_User(gormDB)
// 	// user_Repo.GetUser_ByUsername("john5")

// 	// **********  Transfer *************
// 	tansfer_Repo := tansferRepo.New_Repo_Transfer(gormDB)

// 	// usersProto.RegisterUserServer(s,usersService.New_Service_User(user_Repo))
// 	//******************** Run gRPC ****************
// 	// ทำ run service ต่างๆให้ครบก่อน
// 	// แล้วค้อยส่งไป register proto ต่างๆ***

// 	// reflection.Register(s)

// 	// fmt.Println("gRPC Start Server .........")

// 	// err = s.Serve(listener)
// 	// if err != nil {
// 	// 	// log.Fatal(err)
// 	// 	fmt.Println(err.Error())
// 	// }

// 	gRPC_Server(s, listener, account_Repo, user_Repo, tansfer_Repo)

// 	// Create user
// 	// dt := models.User{
// 	// 	Username:          "russy",
// 	// 	HashedPassword:    "1150",
// 	// 	FullName:          "russy five",
// 	// 	Email:             "russy@gmail.com",
// 	// 	PasswordChangedAt: time.Now().Local().UTC(),
// 	// 	CreatedAt:         time.Now().Local().UTC(),
// 	// 	IsEmailVerified:   false,
// 	// }
// 	// repo_User := users.New_Repo_User(gormDB)
// 	// errREpo_User:=repo_User.CreateUser(&dt)
// 	// if errREpo_User != nil {
// 	// 	fmt.Println(errREpo_User)
// 	// }

// 	// Create Account **************************
// 	// 	da:= models.CreateAccountParams{
// 	// 		Owner: "russy",
// 	// 		Balance: 3000,
// 	// 		Currency: "USD",
// 	// 		CreatedAt: time.Now().Local(),
// 	// 	}
// 	// repo_account := account.New_Repo_Account(gormDB)

// 	// _ = repo_account

// 	// 	errRepo_Account:= repo_account.CreateAccount(&da)
// 	// if errRepo_Account != nil {
// 	// 	fmt.Println(errRepo_Account)
// 	// }

// 	//  update account **************************
// 	// ua:=models.UpdateAccountParams{
// 	// 	ID: 4,
// 	// 	Balance: -100,
// 	// }
// 	// errUpdate_Blance:=repo_account.UpdateAccount_Blance(&ua)
// 	// fmt.Println(errUpdate_Blance)

// 	// get account *****************************
// 	// data, err := repo_account.GetAccount(7, "russy")
// 	// if err != nil {
// 	// 	fmt.Println(err.Error())
// 	// }
// 	// fmt.Println(data)

// 	// list account *************************
// 	// data2, err := repo_account.ListAccount("john5")
// 	// if err != nil {
// 	// 	fmt.Println(err.Error())
// 	// }
// 	// fmt.Println(data2)

// 	// Tansfer *******************************
// 	// repo_Transfer := repo_transfers.New_Repo_User(gormDB)
// 	// _ = repo_Transfer

// 	// ctf := models.CreateTransferParams{
// 	// 	Owner:         "russy",
// 	// 	FromAccountID: 8, // 7-3, 8-1
// 	// 	ToAccountID:   1,
// 	// 	Amount:        500,
// 	// }
// 	// _ = ctf
// 	// errRepoTransfer := repo_Transfer.Create_Transfer(&ctf)
// 	// if errRepoTransfer != nil {
// 	// 	fmt.Println(errRepoTransfer)
// 	// }

// 	// ฝาก ****************************************
// 	// deposit := models.Create_Deposit_and_Withdraw{
// 	// 	Owner:     "john5",
// 	// 	AccountID: 1,
// 	// 	Amount:    100,
// 	// }
// 	// err = repo_Transfer.Create_Deposit(&deposit)
// 	// if err != nil {
// 	// 	fmt.Println(err.Error())
// 	// }

// 	//ถอน *****************************************
// 	// withdraw := models.Create_Deposit_and_Withdraw{
// 	// 		Owner:     "john5",
// 	// 		AccountID: 1,
// 	// 		Amount:    100,
// 	// 	}
// 	// 	err = repo_Transfer.Create_Withdraw(&withdraw)
// 	// 	if err != nil {
// 	// 		fmt.Println(err.Error())
// 	// 	}

// 	// Get transfer by id ****************************************
// 	// id, errGet := repo_Transfer.GetTransfer_ById(3,"2023-01-01" , "2023-10-20") // 3,7
// 	// if errGet != nil {
// 	// 	fmt.Println(err.Error())
// 	// }
// 	// fmt.Println(id)

// 	// fmt.Println("*********************************")
// 	// // Get transfer by owner *************************************
// 	// owner, err := repo_Transfer.GetTransfer_ByOwner("russy","2023-01-01" , "2023-10-20") // "john50", "russy"
// 	// if err != nil {
// 	// 	fmt.Println(err.Error())
// 	// }
// 	// fmt.Println(len(owner))

// 	// ListTransfers *******************
// 	// 	statement_Repo := entries.New_Repo_Entries(gormDB)

// 	// 	statamentData, err := statement_Repo.GetStaement(1, "", "")

// 	// 	if len(statamentData) == 0 {
// 	// 		fmt.Println("record not found")
// 	// 	}
// 	// 	if err != nil {
// 	// 		fmt.Println(err.Error())
// 	// 	}

// 	// 	for _, v := range statamentData {
// 	// 		fmt.Println(v)
// 	// 	}

// 	// }

// 	// func main2() {
// 	// 	// option เอาไว้ทำอะไร ทำตาม vdo
// 	// 	option := []grpc.ServerOption{}
// 	// 	// ฝั่ง server จะใช้ NewServer
// 	// 	s := grpc.NewServer(option...)

// 	// 	//สร้าง net.Listen ก่อน ทำแค่ฝั่ง server
// 	// 	listener, err := net.Listen("tcp", ":50051")
// 	// 	if err != nil {
// 	// 		log.Fatal(err)
// 	// 	}

// 	// 	// มาจาก gRPC ที่ build  มา
// 	// 	// เอามาบริการงาน gRPC ที่ทำไว้
// 	// 	// รับ พารามิเตอร์ 2 ตัว ไปดูว่ามาจากอะไรบ้าง *******
// 	// 	// ตรงนี้สำคัญต้องใส เพื่อมาบริการงาน grpc มันจะลิ้งเข้าไปหา fun ที่ทำงานจากตรงนี้ ******
// 	// 	// ชื่อมันอิงมาจาก proto
// 	// 	// parameter  s = grpc, services.NewCalculatorServer() = func ที่ส้รางเองในไฟล์ service ที่ conform ตาม interface CalculatorServer ในไฟล์ proto ที่  build
// 	// 	services.RegisterCalculatorServer(s, services.NewServer_Service())

// 	// 	// ทำให้ evans เห็น service ทั้งหมด *****************
// 	// 	// ต้องการ reflection.GRPCServer
// 	// 	reflection.Register(s)

// 	// 	fmt.Println("gRPC Start Server .........")

// 	// 	err = s.Serve(listener)
// 	// 	if err != nil {
// 	// 		log.Fatal(err)
// 	// 	}

// 	// 	// stop gRPC
// 	// 	// s.Stop()
// 	// }

// }

// // paramiter อื่น เป็น interface ของ Proto Account, Taransfer, User
// func gRPC_Server2(s *grpc.Server,
// 	listener net.Listener,

// 	account_Repo accountRepo.I_Repo_Account,
// 	user_Repo usersRepo.I_Rero_Users,
// 	transfer_Repo tansferRepo.I_Repo_Transfers,
// ) error {

// 	//******************* Account ********************************
// 	accountProto.RegisterAccountServer(s, accountService.New_Service_Account(account_Repo))
// 	// accountProto.RegisterAccountServer(s,ac_Service )

// 	// ******************** User *********************************
// 	usersProto.RegisterUserServer(s, usersService.New_Service_User(user_Repo))

// 	// ******************* Transfer ******************************
// 	tansferProto.RegisterTransfersServer(s, tansferService.New_Service_Tranfers(transfer_Repo))

// 	// register EVAN *******************
// 	reflection.Register(s)

// 	fmt.Println("gRPC Start Server .........")

// 	err := s.Serve(listener)
// 	if err != nil {
// 		// log.Fatal(err)
// 		fmt.Println(err.Error())
// 	}

// 	return nil
// }
