package main

import (
	controllerAccount "banktransferclient/controller_client/account"
	controllerTransfer "banktransferclient/controller_client/transfers"
	controllerUsers "banktransferclient/controller_client/users"

	"log"
	"fmt"
	"time"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func main() {
	// สร้าง insecure
	creds := insecure.NewCredentials()

	// gRPC Server *******************
	// ฝั่ง client จะใช้ Dial
	cc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	//********************** Account ******************************
	//  gRPC ************************
	newAccountClient := controllerAccount.NewAccountClient(cc)
	
	// Service Account *****************************
	serviceAccount := controllerAccount.New_Service_Account(newAccountClient)

	// Handler Account ***************************
	handlerAccount := controllerAccount.New_Handler_Account(serviceAccount)


	// *************************** User ******************************
	// gRPC ********************
	newUserClient := controllerUsers.NewUserClient(cc)

	// Service USer *********************
	serviceUser := controllerUsers.New_Service_User(newUserClient)

	// Handler user *******************
	handlerUser:=controllerUsers.New_Handler_Account(serviceUser)

	

	// ************************* Transfer *****************************************
	// gRPC ************************
	newTransferClient := controllerTransfer.NewTransfersClient(cc)
	
	// Service Transfer ***************************
	newTransferService:=controllerTransfer.New_Service_Transfer(newTransferClient)
	
	// Handler Transfer ***************************
	handlerTransfer := controllerTransfer.New_Handler_Transfer(newTransferService)
	

	

	// ******************  Time condition for taransfer  **************************
	// ถ้าส่งมาแบบนี้ "2023-12-27 14:46:04.495189+07" mepy'w'
	setTime1 := time.Date(2023, 01, 01, 8, 59, 00, 11111111, time.UTC)
	setTime2 := time.Date(2023, 12, 31, 8, 59, 00, 11111111, time.UTC)

	// fmt.Println(setTime1) // 2023-01-01 08:59:00.011111111 +0000 UTC
	// fmt.Println(setTime2)
	startTime, _ := time.Parse(time.RFC3339, setTime1.Format(time.RFC3339))
	endTime, _ := time.Parse(time.RFC3339, setTime2.Format(time.RFC3339))

	_ = startTime
	_ = endTime

	// fmt.Println("*******Time ***********")
	// fmt.Println(startTime) //2023-01-01 08:59:00 +0000 UTC
	// fmt.Println(startTime.Local().String()) //2023-01-01 15:59:00 +0700 +07
	// fmt.Println(startTime.Unix()) // 1672563540

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

	
	// controllTransfer.GetTransfer_ById(&controllerTransfer.GetTransfer_ByIdRequest{
	// 	AccountID: 1,
	// 	Owner:     "john5",
	// 	StartTime: &timestamppb.Timestamp{
	// 		Seconds: startTime.Unix(),
	// 		Nanos:   int32(time.Now().Nanosecond()),
	// 	},
	// 	EndTime: &timestamppb.Timestamp{
	// 		Seconds: endTime.Unix(),
	// 		Nanos:   int32(time.Now().Nanosecond()),
	// 	},
	// })

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



	// **********************  gRPC Service ******************************
	run_gRPCGateWay(handlerAccount,handlerUser,handlerTransfer)

}

func run_gRPCGateWay(
	handlerAccount controllerAccount.AccountClient,
	habdlerUSer controllerUsers.UserClient,
	handlerTransfer controllerTransfer.TransfersClient,
	) {

	// mux ************************
	mux := runtime.NewServeMux()

	// endPoint:=flag.String("grpc-server-endpoint","localhost:8081")
	//************************ gate ways **************************
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	
	// สร้าง  GateWay or EndPoint **************************
	err := controllerAccount.RegisterAccountHandlerFromEndpoint(context.Background(), mux, "localhost:8080", opts)
	if err != nil {
		fmt.Println(err.Error())
	}

	// ************************  ต้องใช้ func นี้ด้วย **************************
	// Account  *************************************
	// func พวกนี้จะรับ Interface Client ที่  Proto Gen มา
	// ต้องสร้าง struct ที่ confrom  ตาม interface Client ที่ Proto Gen ขึ้นมา
	err = controllerAccount.RegisterAccountHandlerClient(context.Background(), mux, handlerAccount)
	
	// User  **************************************
	// func พวกนี้จะรับ Interface Client ที่  Proto Gen มา
	// ต้องสร้าง struct ที่ confrom  ตาม interface Client ที่ Proto Gen ขึ้นมา
	err = controllerUsers.RegisterUserHandlerClient(context.Background(), mux,habdlerUSer)

	// Transfer  ***********************************
	// func พวกนี้จะรับ Interface Client ที่  Proto Gen มา
	// ต้องสร้าง struct ที่ confrom  ตาม interface Client ที่ Proto Gen ขึ้นมา
	err = controllerTransfer.RegisterTransfersHandlerClient(context.Background(), mux,handlerTransfer)
	
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Start gRPC Gateway.............")
	// ****************** http **********************************
	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}
