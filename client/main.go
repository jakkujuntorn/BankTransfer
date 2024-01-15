package main

import (
	controllerAccount "banktransferclient/controller_client/account"
	controllerTransfer "banktransferclient/controller_client/transfers"
	controllerUsers "banktransferclient/controller_client/users"

	"log"

	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	// สร้าง insecure
	creds := insecure.NewCredentials()

	// ฝั่ง client จะใช้ Dial
	cc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	//********************** Account ******************************
	// inteface ยังไม่เหมือน proto ****************
	// จาก gRPC ************************
	newAccountClient := controllerAccount.NewAccountClient(cc)
	// my *****************************
	controlerAccount := controllerAccount.New_Controller_Account(newAccountClient)
	_ = controlerAccount
	// Func Account ********************
	// err = controlerAccount.CreateAccount(&controllerAccount.CreateAccountRequest{})

	// GetAccount เทสแล้ว ############################
	// _,err = controlerAccount.GetAccount(2,"john5") 

	// ListAccount เทสแล้ว ############################
	// _,err = controlerAccount.ListAccount("john5") 

	// err = controlerAccount.DeleteAccount(1)

	
	// ************************* Transfer *****************************************
	// gRPC ************************
	newTransferClient := controllerTransfer.NewTransfersClient(cc)
	// my **************
	controllTransfer := controllerTransfer.New_Controller_Transfer(newTransferClient)
	_ = controllTransfer

	// func Transfer
	// controllTransfer.Create_Transfer(&controllerTransfer.Create_TransferRequest{})

	// controllTransfer.Create_Deposit(&controllerTransfer.Create_DepositRequest{
	// 	Owner: "russy",
	// 	AccountID: 7,
	// 	Amount: 700,
	// }) //  ลองใหม่แก้ logic

	// controllTransfer.Create_Withdraw(&controllerTransfer.Create_WithdrawRequest{
	// 	Owner: "russy",
	// 	AccountID: 7,
	// 	Amount: 700,
	// }) // ลองใหม่แก้ logic

	// ******************  Time condition for taransfer  **************************
	// ถ้าส่งมาแบบนี้ "2023-12-27 14:46:04.495189+07" mepy'w'
	setTime1 := time.Date(2023, 01, 01, 8, 59, 00, 11111111, time.UTC)
	setTime2 := time.Date(2023, 12, 31, 8, 59, 00, 11111111, time.UTC)
	
	startTime, _ := time.Parse(time.RFC3339, setTime1.Format(time.RFC3339))
	endTime, _ := time.Parse(time.RFC3339, setTime2.Format(time.RFC3339))

	// fmt.Println("******************")
	// fmt.Println(startTime)
	// fmt.Println(startTime.Local().String())
	// fmt.Println(startTime.Unix())

	// อันนี้ใส time แบบ ธรมดา ############################
	// startTime, err := time.Parse("2006-01-02", "2023-01-01")
	// if err != nil {
	// 	fmt.Println("Error Time : ", err.Error())
	// }
	// _ = startTime

	// endTime, err := time.Parse("2006-01-02", "2023-12-30")
	// if err != nil {
	// 	fmt.Println("Error Time : ", err.Error())
	// }
	// _ = endTime

	// GetTransfer_ById เทส แล้ว ออก client ############################
	controllTransfer.GetTransfer_ById(&controllerTransfer.GetTransfer_ByIdRequest{
		AccountID: 1,
		Owner:     "john5",
		StartTime: &timestamppb.Timestamp{
			Seconds: startTime.Unix(),
			Nanos:   int32(time.Now().Nanosecond()),
		},
		EndTime: &timestamppb.Timestamp{
			Seconds: endTime.Unix(),
			Nanos:   int32(time.Now().Nanosecond()),
		},
	}) 

// GetTransfer_ByOwner  เทสแล้ว ออก client ############################
	// controllTransfer.GetTransfer_ByOwner(&controllerTransfer.GetTransfer_ByOwnerRequest{
	// 	Owner:     "russy",
	// 	StartTime: &timestamppb.Timestamp{
	// 		Seconds: startTime.Unix(),
	// 		Nanos:   int32(time.Now().Nanosecond()),
	// 	},
	// 	EndTime: &timestamppb.Timestamp{
	// 		Seconds: endTime.Unix(),
	// 		Nanos:   int32(time.Now().Nanosecond()),
	// 	},
	// })
	
	// ****************************ไม่ใช้แล้ว *********************************
	// อันนี้ไม่ต้องมีก็ได้ เพรา getTransfer  Byid มันน่าจะครบแล้ว ************************
	// controllTransfer.ListStatement(&controllerTransfer.ListStatementRequest{})

	// *************************** User ******************************
	// gRPC ********************
	newUserClient := controllerUsers.NewUserClient(cc)

	// my *********************************
	controllUser := controllerUsers.New_Controller_User(newUserClient)
	_ = controllUser

	//func User *********************

	// controllUser.CreateUser(&controllerUsers.Create_UserRequest{})

	// controllUser.GetUser_ByUsername(&controllerUsers.Username{Username: "john5"}) // ************************

	// UpdateUser เทสแล้ว ############################
	// controllUser.UpdateUser(&controllerUsers.User_UpdateRequest{
	// 	Username: "russy",
	// 	Fullname: "russy jack",
	// 	Email: "rj@gmail.com",
	// }) 

	// ChangePassword  เทสแล้ว ############################
	// controllUser.ChangePassword(&controllerUsers.Change_PasswordRequest{
	// 	Username: "russy",
	// 	Password: "8888",
	// }) 

	if err != nil {
		fmt.Println("Client Error : ", err.Error())
	}
}
