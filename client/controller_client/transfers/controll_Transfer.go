package transfers

import (
	_ "banktransferclient/models"
	"context"
	"errors"
	"fmt"
)

var ctx = context.Background()

type Controller_Transfer struct {
	controller TransfersClient
}

type I_Transfer interface {
	// ใช้ตัวแปรที่  gen มาได้ไหมจะได้ไม่ต้อง ปั้นข้อมูลใหม่
	Create_Transfer(in *Create_TransferRequest) (*StatusResponse, error)
	Create_Deposit(in *Create_DepositRequest) (*StatusResponse, error)
	Create_Withdraw(in *Create_WithdrawRequest) (*StatusResponse, error)
	GetTransfer_ById(in *GetTransfer_ByIdRequest) (*GetTransfer_Response, error)
	GetTransfer_ByOwner(in *GetTransfer_ByOwnerRequest) (*GetTransfer_Response, error)
	// ListStatement(in *ListStatementRequest) (*ListStatementResponse, error)
}

func New_Controller_Transfer(tc TransfersClient) I_Transfer {
	return &Controller_Transfer{controller: tc}
}

func (ct *Controller_Transfer) Create_Transfer(in *Create_TransferRequest) (*StatusResponse, error) {
	data, err := ct.controller.Create_Transfer(ctx, in)
	if err != nil {
		return data, err
	}

	fmt.Println(data.GetStatus())
	return &StatusResponse{}, nil
}

func (ct *Controller_Transfer) Create_Deposit(in *Create_DepositRequest) (*StatusResponse, error) {
	data, err := ct.controller.Create_Deposit(ctx, in)
	if err != nil {
		fmt.Println("Error Create_Deposit client :", err.Error())
		return data, err
	}

	fmt.Println("Status : ", data.GetStatus())
	return &StatusResponse{}, nil
}

func (ct *Controller_Transfer) Create_Withdraw(in *Create_WithdrawRequest) (*StatusResponse, error) {
	data, err := ct.controller.Create_Withdraw(ctx, in)
	if err != nil {
		fmt.Println("Error Create_Withdraw client :", err.Error())
		return data, err
	}
	fmt.Println("Status : ", data.GetStatus())
	return &StatusResponse{}, nil
}

func (ct *Controller_Transfer) GetTransfer_ById(in *GetTransfer_ByIdRequest) (*GetTransfer_Response, error) {

	// เช็ค Time query *********
	if in.GetStartTime().Seconds > in.GetEndTime().Seconds {
		fmt.Println("time format is wrong")
		return &GetTransfer_Response{}, errors.New("time format is wrong")
	}

	data, err := ct.controller.GetTransfer_ById(ctx, in)
	if err != nil {
		fmt.Println(err.Error())
		return data, err
	}

	// fmt.Println(data.GetDataTransfer()[1])
	// loop protobuf ทำแบบปกติได้
for i := 0; i < len(data.GetDataTransfer()); i++ {
	fmt.Println(data.GetDataTransfer()[i])
	fmt.Println("****************************")
}

//********************* ข้อมูล tansfer ถ้าติด ลบ คือโอนออก  ถ้าค่า + คือมีคนโอนเข้ามา *******************************
// ใส  condition ด้วย เพื่อแสดงค่าต่างกัน 

// protobuf for range ค่าไม่ออก
// 	for v := data.GetDataTransfer(){
// fmt.Println(v)
// fmt.Println("****************************")
// 	}

	return &GetTransfer_Response{}, nil
}

func (ct *Controller_Transfer) GetTransfer_ByOwner(in *GetTransfer_ByOwnerRequest) (*GetTransfer_Response, error) {

	// เช็ค Time query *********
	if in.GetStartTime().Seconds > in.GetEndTime().Seconds {
		fmt.Println("time format is wrong")
		return &GetTransfer_Response{}, errors.New("time format is wrong")
	}

	data, err := ct.controller.GetTransfer_ByOwner(ctx, in)
	if err != nil {
		fmt.Println("Error GetTransfer_ByOwner client: ", err.Error())
		return data, err
	}

	fmt.Println(len(data.GetDataTransfer()))
	fmt.Println(data.GetDataTransfer())
	return &GetTransfer_Response{}, nil
}


//ไม่ใช้แล้ว *********************************************
// func (ct *Controller_Transfer) ListStatement(in *ListStatementRequest) (*ListStatementResponse, error) {
// 	data, err := ct.controller.ListStatement(ctx, in)
// 	if err != nil {
// 		return data, err
// 	}
// 	return &ListStatementResponse{}, nil
}
