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
	g_Server := grpc.NewServer(option...)

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
	account_Repo := accountRepo.New_Repo_Account(gormDB)

	// ********************** User *******************
	user_Repo := usersRepo.New_Repo_User(gormDB)

	// **********  Transfer *************
	tansfer_Repo := tansferRepo.New_Repo_Transfer(gormDB)


	//******************** Run gRPC ****************
	gRPC_Server(g_Server, listener, account_Repo, user_Repo, tansfer_Repo)

}

// paramiter อื่น เป็น interface ของ Proto Account, Taransfer, User
func gRPC_Server(g_Server *grpc.Server,
	listener net.Listener,
	account_Repo accountRepo.I_Repo_Account,
	user_Repo usersRepo.I_Rero_Users,
	transfer_Repo tansferRepo.I_Repo_Transfers,
) error {

	//******************* Account ********************************
	accountProto.RegisterAccountServer(g_Server, accountService.New_Service_Account(account_Repo))
	// accountProto.RegisterAccountServer(s,ac_Service )

	// ******************** User *********************************
	usersProto.RegisterUserServer(g_Server, usersService.New_Service_User(user_Repo))

	// ******************* Transfer ******************************
	tansferProto.RegisterTransfersServer(g_Server, tansferService.New_Service_Tranfers(transfer_Repo))

	// register EVAN *******************
	reflection.Register(g_Server)

	fmt.Println("gRPC Start Server .........")

	err := g_Server.Serve(listener)
	if err != nil {
		// log.Fatal(err)
		fmt.Println(err.Error())
	}

	return nil
}
