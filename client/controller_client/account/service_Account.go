package account

import (
	auth "banktransferclient/util/auth"
	"context"
	"fmt"
)

var ctx = context.Background()

type Controller_Account struct {
	controller AccountClient
}

type I_Account_Service interface {
	CreateAccount(*CreateAccountRequest) error
	GetAccount(ctx context.Context, in *GetAccountRequest) (*GetAccountResponse, error)
	ListAccount(ctx context.Context, in *ListAccountRequest) (*GetListAccount_Response, error)
	DeleteAccount(id int) error
}

func New_Service_Account(ac AccountClient) I_Account_Service {
	fmt.Println("")
	return &Controller_Account{
		controller: ac,
	}
}

func (ca *Controller_Account) CreateAccount(data *CreateAccountRequest) error {

	//  ปั้น token **********************
	jwtToken, err := auth.New_JWT("ratthakorn")
	if err != nil {
		return err
	}

	//  VerifyToken ********************
	payload, err := jwtToken.AuthorizeUser(ctx)
	if err != nil {
		return err
	}

	// ทำ message มายาก เลยตั้งปั้นข้อมูลยาก ******
	data.Owner = &Owner{Owner: payload.Audience}

	// to Server ***************************
	responseData, err := ca.controller.CreateAccount(ctx, data)
	_ = responseData
	if err != nil {
		return err
	}

	return nil
}

func (ca *Controller_Account) GetAccount(ctx context.Context, in *GetAccountRequest) (*GetAccountResponse, error) {

	//  ปั้น token **********************
	jwtToken, err := auth.New_JWT("ratthakorn")
	if err != nil {
		return &GetAccountResponse{}, err
	}

	//  VerifyToken ********************
	payload, err := jwtToken.AuthorizeUser(ctx)
	if err != nil {
		return &GetAccountResponse{}, err
	}

	in.Owner = payload.Audience

	// to Server *******************************
	responseData, err := ca.controller.GetAccount(ctx, in)
	_ = responseData
	if err != nil {
		return &GetAccountResponse{}, err
	}

	return responseData, nil
}

func (ca *Controller_Account) ListAccount(ctx context.Context, in *ListAccountRequest) (*GetListAccount_Response, error) {

	//  ปั้น token **********************
	jwtToken, err := auth.New_JWT("ratthakorn")
	if err != nil {
		return &GetListAccount_Response{}, err
	}

	//  VerifyToken ********************
	payload, err := jwtToken.AuthorizeUser(ctx)
	if err != nil {
		return &GetListAccount_Response{}, err
	}

	in.Owner = &Owner{Owner: payload.Audience}

	// to Server *************************
	responseData, err := ca.controller.ListAccount(ctx, in)

	if err != nil {
		return &GetListAccount_Response{}, err
	}

	return responseData, nil
}

func (ca *Controller_Account) DeleteAccount(id int) error {
	data := DeleteAccountRequest{}
	responseData, err := ca.controller.DeleteAccount(ctx, &data)
	_ = responseData
	if err != nil {
		return err
	}
	return nil
}
