package transfers

// vdo ไป create   อย่างอย่าง

import (
	"banktransfer/models"
	transferProto "banktransfer/transfers"
	transfer_Repo "banktransfer/transfers/repo"
	"errors"

	context "context"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type service_Transfer struct {
	transferProto.UnimplementedTransfersServer
	transferRepo transfer_Repo.I_Repo_Transfers
}

func New_Service_Tranfers(transferRepo transfer_Repo.I_Repo_Transfers) transferProto.TransfersServer {
	return &service_Transfer{transferRepo: transferRepo}
}

// Create_Transfer implements transfers.TransfersServer
func (st *service_Transfer) Create_Transfer(ctx context.Context, dataTransfer *transferProto.Create_TransferRequest) (*transferProto.StatusResponse, error) {
	fmt.Println("Create_Transfer ***************")

	// err := st.transferRepo.Create_Transfer_Proto(dataTransfer)
	// if err != nil {
	// 	return &transferProto.StatusResponse{}, err
	// }

	// ปั้นข้อมูลใหม่ ****************
	dataTo_Db := models.CreateTransferParams{
		Owner:         dataTransfer.GetOwner(),
		FromAccountID: int64(dataTransfer.GetFromAccountID()),
		ToAccountID:   int64(dataTransfer.GetToAccountID()),
		Amount:        int64(dataTransfer.GetAmount()),
	}

	err := st.transferRepo.Create_Transfer(&dataTo_Db)
	if err != nil {
		return &transferProto.StatusResponse{}, err
	}

	return &transferProto.StatusResponse{Status: "transfer complete"}, nil
}

// Create_Deposit implements transfers.TransfersServer
func (st *service_Transfer) Create_Deposit(ctx context.Context, cd *transferProto.Create_DepositRequest) (*transferProto.StatusResponse, error) {
	fmt.Println("Create_Deposit ****************")

	// ปั้นข้อมูลใหม่ ****************************
	dataDB := models.Create_Deposit_and_Withdraw{
		Owner:     cd.GetOwner(),
		AccountID: int64(cd.GetAccountID()),
		Amount:    int64(cd.GetAmount()),
	}

	// to DB **********************************
	err := st.transferRepo.Create_Deposit(&dataDB)
	if err != nil {
		return nil, err
	}

	return &transferProto.StatusResponse{Status: "deposit complete"}, nil
}

// Create_Withdraw implements transfers.TransfersServer
func (st *service_Transfer) Create_Withdraw(ctx context.Context, cw *transferProto.Create_WithdrawRequest) (*transferProto.StatusResponse, error) {
	// fmt.Println("Create_Withdraw ***************")

	// ปั้นข้อมูลใหม่ ******************************
	dataDB := models.Create_Deposit_and_Withdraw{
		Owner:     cw.GetOwner(),
		AccountID: int64(cw.GetAccountID()),
		Amount:    int64(cw.GetAmount()),
	}

	// to DB *******************************
	err := st.transferRepo.Create_Withdraw(&dataDB)
	if err != nil {
		return nil, err
	}

	return &transferProto.StatusResponse{Status: "withdraw complete"}, nil
}

// GetTransfer_ById implements transfers.TransfersServer
func (st *service_Transfer) GetTransfer_ById(ctx context.Context, gb *transferProto.GetTransfer_ByIdRequest) (*transferProto.GetTransfer_Response, error) {

	// fmt.Println("GetTransfer_ById ***************")

	// to Repo ***************************
	dataTrandfer_ByID, err := st.transferRepo.GetTransfer_ById(
		int64(gb.GetAccountID()),
		gb.GetOwner(),
		gb.GetStartTime()[0:10],
		gb.GetEndTime()[0:10],	
	)
	if err != nil {
		fmt.Println("Error Get Transfer : ", err.Error())
	}
	// ต้องเช็คด้วยถ้าไม่เจอ **************************************
	if len(dataTrandfer_ByID) == 0 {
		return &transferProto.GetTransfer_Response{}, errors.New("recore not found")
	}

	//  ปั้นข้อมูลใหม่ สำหรับ response  **********************
	dataTransferResponse := transferProto.GetTransfer_Response{}

	for _, v := range dataTrandfer_ByID {
		dataTrnsfer := transferProto.DataTransfer_Response{
			AccountId: int32(v.AccountID),
			Amount:    int32(v.Amount),
			CreatedAt: &timestamppb.Timestamp{
				Seconds: v.CreatedAt.Unix(),
				Nanos:   int32(v.CreatedAt.Nanosecond()),
			},
			EntriesType:        v.EntriesType,
			Owner:              v.Owner,
			DestinationAccount: int32(v.Destination_account),
		}
		dataTransferResponse.DataTransfer = append(dataTransferResponse.DataTransfer, &dataTrnsfer)
	}

	return &dataTransferResponse, nil
}

// GetTransfer_ByOwner implements transfers.TransfersServer
func (st *service_Transfer) GetTransfer_ByOwner(ctx context.Context, gb *transferProto.GetTransfer_ByOwnerRequest) (*transferProto.GetTransfer_Response, error) {
	fmt.Println("GetTransfer_ByOwner ***************")
	
	// to Repo ***************************
	dataTrandfer_ByID, err := st.transferRepo.GetTransfer_ByOwner(
		gb.GetOwner(),
		gb.GetStartTime()[0:10],
		gb.GetEndTime()[0:10],
	)

	if err != nil {
		fmt.Println("Error Get Transfer : ", err.Error())
	}
	// ต้องเช็คด้วยถ้าไม่เจอ **************************************
	if len(dataTrandfer_ByID) == 0 {
		return &transferProto.GetTransfer_Response{}, errors.New("recore not found")
	}

	//  ปั้นข้อมูลใหม่ สำหรับ response  **********************
	dataTransferResponse := transferProto.GetTransfer_Response{}
	for _, v := range dataTrandfer_ByID {
		dataTrnsfer := transferProto.DataTransfer_Response{
			AccountId: int32(v.AccountID),
			Amount:    int32(v.Amount),
			CreatedAt: &timestamppb.Timestamp{
				Seconds: v.CreatedAt.Unix(),
				Nanos:   int32(v.CreatedAt.Nanosecond()),
			},
			EntriesType:        v.EntriesType,
			Owner:              v.Owner,
			DestinationAccount: int32(v.Destination_account),
		}
		dataTransferResponse.DataTransfer = append(dataTransferResponse.DataTransfer, &dataTrnsfer)
	}

	return &dataTransferResponse, nil
}

// ListStatement implements transfers.TransfersServer
func (*service_Transfer) ListStatement(context.Context, *transferProto.ListStatementRequest) (*transferProto.ListStatementResponse, error) {
	fmt.Println("ListStatement ***************")
	return &transferProto.ListStatementResponse{}, nil
}

// mustEmbedUnimplementedTransfersServer implements transfers.TransfersServer
// func (service_Transfer) mustEmbedUnimplementedTransfersServer() {
// 	panic("unimplemented")
// }
