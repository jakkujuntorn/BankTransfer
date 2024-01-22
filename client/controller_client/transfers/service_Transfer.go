package transfers

import (
	auth "banktransferclient/util/auth"
	 "banktransferclient/util"
	"context"
	"errors"
	"fmt"
	"time"
)

var ctx = context.Background()

type Controller_Transfer struct {
	controller TransfersClient
}

type I_Transfer interface {
	Create_Transfer(ctx context.Context, in *Create_TransferRequest) (*StatusResponse, error)
	Create_Deposit(ctx context.Context, in *Create_DepositRequest) (*StatusResponse, error)
	Create_Withdraw(ctx context.Context, in *Create_WithdrawRequest) (*StatusResponse, error)
	GetTransfer_ById(ctx context.Context, in *GetTransfer_ByIdRequest) (*GetTransfer_Response, error)
	GetTransfer_ByOwner(ctx context.Context, in *GetTransfer_ByOwnerRequest) (*GetTransfer_Response, error)
	ListStatement(ctx context.Context, in *ListStatementRequest) (*ListStatementResponse, error)
}

func New_Service_Transfer(tc TransfersClient) I_Transfer {
	return &Controller_Transfer{controller: tc}
}

func (ct *Controller_Transfer) Create_Transfer(ctx context.Context, in *Create_TransferRequest) (*StatusResponse, error) {

	//  ปั้น token **********************
	jwtToken, err := auth.New_JWT("ratthakorn")
	if err != nil {
		return &StatusResponse{}, err
	}

	//  VerifyToken ********************
	payload, err := jwtToken.AuthorizeUser(ctx)
	if err != nil {
		return &StatusResponse{}, err
	}

	in.Owner = &payload.Audience

	// เช็ค money value **************************
	err =util.Validate_MoneyValue(in.GetAmount())
	if err != nil {
		return &StatusResponse{}, err
	}

	// to Server ******************************
	data, err := ct.controller.Create_Transfer(ctx, in)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (ct *Controller_Transfer) Create_Deposit(ctx context.Context, in *Create_DepositRequest) (*StatusResponse, error) {
	// fmt.Println("Create_Deposit service")
	//  ปั้น token **********************
	jwtToken, err := auth.New_JWT("ratthakorn")
	if err != nil {
		return &StatusResponse{}, err
	}

	//  VerifyToken ********************
	payload, err := jwtToken.AuthorizeUser(ctx)
	if err != nil {
		return &StatusResponse{}, err
	}

	in.Owner = &payload.Audience

	// เช็ค money value **************************
	err =util.Validate_MoneyValue(in.GetAmount())
	if err != nil {
		return &StatusResponse{}, err
	}

	// to Server ***************************
	status, err := ct.controller.Create_Deposit(ctx, in)
	if err != nil {
		fmt.Println("Error Create_Deposit client :", err.Error())
		return &StatusResponse{}, err
	}

	return status, nil
}

func (ct *Controller_Transfer) Create_Withdraw(ctx context.Context, in *Create_WithdrawRequest) (*StatusResponse, error) {
	//  ปั้น token **********************
	jwtToken, err := auth.New_JWT("ratthakorn")
	if err != nil {
		return &StatusResponse{}, err
	}

	//  VerifyToken ********************
	payload, err := jwtToken.AuthorizeUser(ctx)
	if err != nil {
		return &StatusResponse{}, err
	}

	in.Owner = &payload.Audience

	// เช็ค money value **************************
	err =util.Validate_MoneyValue(in.GetAmount())
	if err != nil {
		return &StatusResponse{}, err
	}

	// to Server ***************************
	status, err := ct.controller.Create_Withdraw(ctx, in)
	if err != nil {
		return &StatusResponse{}, err
	}

	return status, nil
}

func (ct *Controller_Transfer) GetTransfer_ById(ctx context.Context, in *GetTransfer_ByIdRequest) (*GetTransfer_Response, error) {

	//  ปั้น token **********************
	jwtToken, err := auth.New_JWT("ratthakorn")
	if err != nil {
		return &GetTransfer_Response{}, err
	}

	//  VerifyToken ********************
	payload, err := jwtToken.AuthorizeUser(ctx)
	if err != nil {
		return &GetTransfer_Response{}, err
	}

	in.Owner = &payload.Audience

	// แปลง string จาก front end เป็น time date ****************************
	// setTime1 := time.Date(2023, 01, 01, 8, 59, 00, 11111111, time.UTC)
	// startTime, _ := time.Parse(time.RFC3339, setTime1.Format(time.RFC3339))

	// ใช้อันนี้ เหมาะสุด *****************************
	startTime, err := time.Parse("2006-01-02", in.GetStartTime()[0:10])
	if err != nil {
		return &GetTransfer_Response{}, err
	}
	endtTime, err := time.Parse("2006-01-02", in.GetEndTime()[0:10])
	if err != nil {
		return &GetTransfer_Response{}, err
	}

	// เช็ค Time query ***************************************
	// fmt.Println("Start Time :",startTime.Unix()) // เทียบได้
	// fmt.Println("End Time :",endtTime.Unix())// เทียบได้
	if startTime.Unix() > endtTime.Unix() {
		fmt.Println("time format is wrong")
		return &GetTransfer_Response{}, errors.New("wrong time entered")
	}

	// to Server *************************
	dataTransfer, err := ct.controller.GetTransfer_ById(ctx, in)
	if err != nil {
		// fmt.Println(err.Error()) // record not found ***********
		return &GetTransfer_Response{}, err
	}

	// fmt.Println(data.GetDataTransfer()[1])
	// loop protobuf ทำแบบปกติได้
	//************* ข้อมูล tansfer ถ้าติด ลบ คือโอนออก  ถ้าค่า + คือมีคนโอนเข้ามา ********
	// for i := 0; i < len(dataTransfer.GetDataTransfer()); i++ {
	// 	fmt.Println(dataTransfer.GetDataTransfer()[i])
	// 	fmt.Println("****************************")
	// }

	return dataTransfer, nil
}

func (ct *Controller_Transfer) GetTransfer_ByOwner(ctx context.Context, in *GetTransfer_ByOwnerRequest) (*GetTransfer_Response, error) {

	//  ปั้น token **********************
	jwtToken, err := auth.New_JWT("ratthakorn")
	if err != nil {
		return &GetTransfer_Response{}, err
	}

	//  VerifyToken ********************
	payload, err := jwtToken.AuthorizeUser(ctx)
	if err != nil {
		return &GetTransfer_Response{}, err
	}

	in.Owner = &payload.Audience

	// ใช้อันนี้ เหมาะสุด *****************************
	startTime, err := time.Parse("2006-01-02", in.GetStartTime()[0:10])
	if err != nil {
		return &GetTransfer_Response{}, err
	}
	endtTime, err := time.Parse("2006-01-02", in.GetEndTime()[0:10])
	if err != nil {
		return &GetTransfer_Response{}, err
	}

	// เช็ค Time query ***************************************
	if startTime.Unix() > endtTime.Unix() {
		fmt.Println("time format is wrong")
		return &GetTransfer_Response{}, errors.New("wrong time entered")
	}

	// to Server ************************************
	dataTransfer, err := ct.controller.GetTransfer_ByOwner(ctx, in)
	if err != nil {
		fmt.Println("Error GetTransfer_ByOwner client: ", err.Error())
		return &GetTransfer_Response{}, err
	}

	return dataTransfer, nil
}

//ไม่ใช้แล้ว *********************************************
func (ct *Controller_Transfer) ListStatement(ctx context.Context, in *ListStatementRequest) (*ListStatementResponse, error) {
	data, err := ct.controller.ListStatement(ctx, in)
	if err != nil {
		return data, err
	}
	return &ListStatementResponse{}, nil
}
