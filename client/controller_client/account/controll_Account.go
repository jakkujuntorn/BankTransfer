package account

import (
	"banktransferclient/models"
	"fmt"
	// _"banktransferclient/service_proto/accountProto"
	"context"
	// "google.golang.org/grpc"
)

var ctx = context.Background()

type Controller_Account struct {
	controller AccountClient
}

type I_Account interface {
	// CreateAccount(*models.CreateAccountParams) error
	CreateAccount(*CreateAccountRequest) error
	GetAccount(id int, owner string) (*models.Account, error)
	ListAccount(owner string) ([]models.Account, error)
	DeleteAccount(id int) error
}

func New_Controller_Account(ac AccountClient) I_Account {
	fmt.Println("")
	return &Controller_Account{
		controller: ac,
	}
}

func (ca *Controller_Account) CreateAccount(data *CreateAccountRequest) error {

	responseData, err := ca.controller.CreateAccount(ctx, data)
	_ = responseData
	if err != nil {
		return err
	}

	return nil
}

func (ca *Controller_Account) GetAccount(id int, owner string) (*models.Account, error) {
	data := GetAccountRequest{
		Id:    &ID{Id: int32(id)},
		Owner: &Owner{Owner: owner},
	}

	responseData, err := ca.controller.GetAccount(ctx, &data)
	_ = responseData
	if err != nil {
		return &models.Account{}, err
	}
	fmt.Println(responseData)
	// fmt.Println(responseData.CreatedDate.AsTime())
	return &models.Account{}, nil
}

func (ca *Controller_Account) ListAccount(owner string) ([]models.Account, error) {
	data := ListAccountRequest{
		Owner: &Owner{Owner: owner},
	}
	responseData, err := ca.controller.ListAccount(ctx, &data)
	
	if err != nil {
		return []models.Account{}, err
	}


	// ปั้นข้อมูลใหม่ **************
	for _,v :=range responseData.ListAccount{
		fmt.Println(v)
		fmt.Println("**************")
	}


	return []models.Account{}, nil
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
