package users

import (
	"context"
	"fmt"
	
)

var ctx = context.Background()

type Controller_User struct {
	controller UserClient
}

type I_User interface {
	CreateUser(data *Create_UserRequest) (*StatusResponse, error)
	GetUser_ByUsername(data *Username) (*Get_UserResponse, error)
	UpdateUser(data *User_UpdateRequest) (*StatusResponse, error)
	ChangePassword(data *Change_PasswordRequest) (*StatusResponse, error)
}

func New_Controller_User(uc UserClient) I_User {
	return &Controller_User{controller: uc}
}

func (cu *Controller_User) CreateUser(data *Create_UserRequest) (*StatusResponse, error) {
	errorUser, err := cu.controller.CreateUser(ctx, data)
	if err != nil {
		fmt.Println("Client CreateUser error:", err.Error())
		return &StatusResponse{}, err
	}

	fmt.Println("Client CreateUser:", errorUser.GetStatus())
	return &StatusResponse{}, nil
}

func (cu *Controller_User) GetUser_ByUsername(data *Username) (*Get_UserResponse, error) {
	// username := Username{
	// 	Username: data.GetUsername(),
	// }
	userResponse, err := cu.controller.GetUser_ByUsername(ctx,data)
	if err != nil {
		fmt.Println(err.Error())
		return &Get_UserResponse{}, err
	}
	fmt.Println("GetUser_ByUsername : ",userResponse)
	return &Get_UserResponse{}, nil
}

func (cu *Controller_User) UpdateUser(data *User_UpdateRequest) (*StatusResponse, error) {
	status, err := cu.controller.UpdateUser(ctx, data)
	
	if err != nil {
		fmt.Println("Error Update Client : ",err.Error())
		return &StatusResponse{}, err
	}

	// complete  proto ถึงจะมีค่ามาให้ใช้
	fmt.Println("proto  : ",status.GetStatus())
	return status, nil
}

func (cu *Controller_User) ChangePassword(data *Change_PasswordRequest) (*StatusResponse, error) {
	status, err := cu.controller.ChangePassword(ctx, data)
	if err != nil {
		fmt.Println("Change password error : ", err.Error())
		return &StatusResponse{}, err
	}
	fmt.Println("status : ",status.GetStatus())
	return &StatusResponse{}, nil
}
