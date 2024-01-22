package account // การตั้งชื่อต้องต้ังแบบนี้ด้วย ***

import (
	accountProto "banktransfer/account"
	account_Repo "banktransfer/account/repo"
	_ "banktransfer/models"
	"context"
	"errors"
	"fmt"
	"strconv"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type service_Account struct {
	accountRepo account_Repo.I_Repo_Account
	accountProto.UnimplementedAccountServer
}

// func New_Service_Account(ac account_Repo.I_Repo_Account) accountProto.AccountServer {
func New_Service_Account(ac account_Repo.I_Repo_Account) accountProto.AccountServer {
	return &service_Account{accountRepo: ac}
}

func (sa *service_Account) CreateAccount(ctx context.Context, ac *accountProto.CreateAccountRequest) (*accountProto.StatusResponse, error) {
	fmt.Println("Create Account Server .....................")
	return &accountProto.StatusResponse{}, nil
}

func (sa *service_Account) GetAccount(ctx context.Context, data *accountProto.GetAccountRequest) (*accountProto.GetAccountResponse, error) {
	// fmt.Println("Get Account server service")
	
	id, err := strconv.Atoi(data.GetId())
	if err != nil {
		return &accountProto.GetAccountResponse{}, err
	}

	// to Repo ********************
	accountDB, err := sa.accountRepo.GetAccount(id, data.GetOwner())
	if err != nil {
		return &accountProto.GetAccountResponse{}, err
	}

	// proto buf message ******************
	accountResponse := &accountProto.GetAccountResponse{
		Id:       &accountProto.ID{Id: int32(accountDB.ID)},
		Owner:    &accountProto.Owner{Owner: accountDB.Owner},
		Balance:  &accountProto.Balance{Balance: int32(accountDB.Balance)},
		Currency: &accountProto.Currency{Currency: accountDB.Currency},
		// ดูเรื่องเวลา ยังไม่ถูกใจ **********
		// CreatedDate: &timestamppb.Timestamp{
		// 	Seconds: accountDB.CreatedAt.Unix(),
		// 	Nanos:   int32(accountDB.CreatedAt.Nanosecond()),
		// },
		CreatedDate: timestamppb.New(accountDB.CreatedAt),
	}

	return accountResponse, nil
}

//ListAccount  (return Array) ****************
func (sa *service_Account) ListAccount(ctx context.Context, data *accountProto.ListAccountRequest) (*accountProto.GetListAccount_Response, error) {
	// fmt.Println("Owner :", data.GetOwner().Owner)
	accountListDB, err := sa.accountRepo.ListAccount(data.GetOwner().Owner)
	if err != nil {
		return nil, err
	}

	// ถ้าข้อมูลไม่มีจะเป็น array ว่าง ต้องเอามาเช็คเอง
	if len(accountListDB) == 0 {
		return nil, errors.New(" list account record not found")
	}

	// ปั้นข้อมุลใหม่ *************
	account_List := accountProto.GetListAccount_Response{}

	for _, v := range accountListDB {
		account := accountProto.GetAccountResponse{
			Id:       &accountProto.ID{Id: int32(v.ID)},
			Owner:    &accountProto.Owner{Owner: v.Owner},
			Balance:  &accountProto.Balance{Balance: int32(v.Balance)},
			Currency: &accountProto.Currency{Currency: v.Currency},
			// CreatedDate: v.CreatedAt,
			// CreatedDate: &timestamppb.Timestamp{
			// 	Seconds: v.CreatedAt.Unix(),
			// 	Nanos:   int32(v.CreatedAt.Nanosecond()),
			// },
			CreatedDate:timestamppb.New(v.CreatedAt),
		}
		account_List.ListAccount = append(account_List.ListAccount, &account)
	}

	return &account_List, nil
}

func (sa *service_Account) DeleteAccount(context.Context, *accountProto.DeleteAccountRequest) (*accountProto.StatusResponse, error) {
	fmt.Println("Delete Account Server .....................")
	return &accountProto.StatusResponse{}, nil
}

// func (sa *service_Account) mustEmbedUnimplementedAccountServer() {}
