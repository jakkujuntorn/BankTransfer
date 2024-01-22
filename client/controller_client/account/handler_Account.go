package account

import (
	"fmt"
	"context"
	grpc "google.golang.org/grpc"
)

type Handler_Account struct {
	controller I_Account_Service
}

func New_Handler_Account(ac I_Account_Service) AccountClient {
	return &Handler_Account{
		controller: ac,
	}
}

// CreateAccount implements AccountClient
func (ha *Handler_Account) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	fmt.Println("CreateAccount Controller Handler")
	
	status,err := ha.controller.CreateAccount(in)
	if err != nil {
		return &StatusResponse{}, err
	}

	return status, nil
}

// GetAccount implements AccountClient
func (ha *Handler_Account) GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*GetAccountResponse, error) {
	fmt.Println("GetAccount Controller Handler")

	account, err := ha.controller.GetAccount(ctx, in)
	if err != nil {
		return &GetAccountResponse{}, err
	}
	return account, nil
}

// ListAccount implements AccountClient
func (ha *Handler_Account) ListAccount(ctx context.Context, in *ListAccountRequest, opts ...grpc.CallOption) (*GetListAccount_Response, error) {
	listAccount, err := ha.controller.ListAccount(ctx, in)
	if err != nil {
		return &GetListAccount_Response{}, err
	}
	return listAccount, nil
}

// DeleteAccount implements AccountClient
func (*Handler_Account) DeleteAccount(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	panic("unimplemented")
}
