package main

import (
	accountProto "banktransfer/account"
	accountRepo "banktransfer/account/repo"
	accountService "banktransfer/account/service"

	tansferProto "banktransfer/transfers"
	tansferRepo "banktransfer/transfers/repo"
	tansferService "banktransfer/transfers/service"

	usersProto "banktransfer/users"
	usersRepo "banktransfer/users/repo"
	usersService "banktransfer/users/service"

	// "banktransfer/models"
	"banktransfer/util"

	"fmt"
	"os"
	_ "time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"banktransfer/db"

	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	// "google.golang.org/grpc/credentials/insecure"
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

	// Start Server  gRPC ****************
	option := []grpc.ServerOption{}
	// ฝั่ง server จะใช้ NewServer
	s := grpc.NewServer(option...)

	//สร้าง net.Listen ก่อน ทำแค่ฝั่ง server ***************
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		// log.Fatal(err.Error())
		fmt.Println(err.Error())
	}

	//***************** DB ************************
	gormDB := db.Postgres_init()
	fmt.Println("Started Postgres .....")
	_ = gormDB
	
	// ************* Account ***************
	// สร้าแค่ repo   ส่วน service ไปสร้างใน gRCPserver
	account_Repo := accountRepo.New_Repo_Account(gormDB)

	// ********************** User *******************
	user_Repo := usersRepo.New_Repo_User(gormDB)
	// user_Repo.GetUser_ByUsername("john5")

	// **********  Transfer *************
	tansfer_Repo := tansferRepo.New_Repo_Transfer(gormDB)


	//******************** Run gRPC ****************
	gRPC_Server(s, listener, account_Repo, user_Repo, tansfer_Repo)

}

// paramiter อื่น เป็น interface ของ Proto Account, Taransfer, User
func gRPC_Server(s *grpc.Server,
	listener net.Listener,
	account_Repo accountRepo.I_Repo_Account,
	user_Repo usersRepo.I_Rero_Users,
	transfer_Repo tansferRepo.I_Repo_Transfers,
) error {

	//******************* Account ********************************
	accountProto.RegisterAccountServer(s, accountService.New_Service_Account(account_Repo))
	// accountProto.RegisterAccountServer(s,ac_Service )

	// ******************** User *********************************
	usersProto.RegisterUserServer(s, usersService.New_Service_User(user_Repo))

	// ******************* Transfer ******************************
	tansferProto.RegisterTransfersServer(s, tansferService.New_Service_Tranfers(transfer_Repo))

	// register EVAN *******************
	reflection.Register(s)

	fmt.Println("gRPC Start Server .........")

	err := s.Serve(listener)
	if err != nil {
		// log.Fatal(err)
		fmt.Println(err.Error())
	}

	return nil
}
