package transfers

import (
	context "context"
	"fmt"

	grpc "google.golang.org/grpc"
)

type Handler_Transfers struct {
	controller I_Transfer
}

func New_Handler_Transfer(tr I_Transfer) TransfersClient {
	fmt.Println("")
	return &Handler_Transfers{
		controller: tr,
	}
}

// Create_Transfer implements TransfersClient
// โอนเงิน
func (ht *Handler_Transfers) Create_Transfer(ctx context.Context, in *Create_TransferRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	status, err := ht.controller.Create_Transfer(ctx, in)
	if err != nil {
		return &StatusResponse{}, err
	}
	return status, nil
}

// Create_Deposit implements TransfersClient
// ฝากเงิน
func (ht *Handler_Transfers) Create_Deposit(ctx context.Context, in *Create_DepositRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	status, err := ht.controller.Create_Deposit(ctx, in)
	if err != nil {
		return &StatusResponse{}, err
	}
	return status, nil
}

// Create_Withdraw implements TransfersClient
// ถอนเงิน
func (ht *Handler_Transfers) Create_Withdraw(ctx context.Context, in *Create_WithdrawRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	status, err := ht.controller.Create_Withdraw(ctx, in)
	if err != nil {
		return &StatusResponse{}, err
	}
	return status, nil
}

// GetTransfer_ById implements TransfersClient
func (ht *Handler_Transfers) GetTransfer_ById(ctx context.Context, in *GetTransfer_ByIdRequest, opts ...grpc.CallOption) (*GetTransfer_Response, error) {
	// หรือใช้ optinal แล้วมาปั้นใหม่ที่ service ***
	dataTransfer, err := ht.controller.GetTransfer_ById(ctx, in)
	if err != nil {
		return &GetTransfer_Response{}, err
	}
	return dataTransfer, nil
}

// GetTransfer_ByOwner implements TransfersClient
func (ht *Handler_Transfers) GetTransfer_ByOwner(ctx context.Context, in *GetTransfer_ByOwnerRequest, opts ...grpc.CallOption) (*GetTransfer_Response, error) {
	dataTransfer, err := ht.controller.GetTransfer_ByOwner(ctx, in)
	if err != nil {
		return &GetTransfer_Response{}, err
	}
	return dataTransfer, nil
}

// ListStatement implements TransfersClient
func (*Handler_Transfers) ListStatement(ctx context.Context, in *ListStatementRequest, opts ...grpc.CallOption) (*ListStatementResponse, error) {
	panic("unimplemented")
}
