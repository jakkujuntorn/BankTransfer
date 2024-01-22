package users

import (
	auth "banktransferclient/util/auth"
	"context"
	"fmt"
	"time"
)

var ctx = context.Background()

type Controller_User struct {
	controller UserClient
}

type I_User_Service interface {
	LoginUser(ctx context.Context, in *User_LoginRequest) (*Get_UserResponse, error)
	CreateUser(ctx context.Context, data *Create_UserRequest) (*StatusResponse, error)
	GetUser_ByUsername(ctx context.Context, data *Username) (*Get_UserResponse, error)
	UpdateUser(ctx context.Context, data *User_UpdateRequest) (*StatusResponse, error)
	ChangePassword(ctx context.Context, data *Change_PasswordRequest) (*StatusResponse, error)
}

func New_Service_User(uc UserClient) I_User_Service {
	return &Controller_User{controller: uc}
}


func (cu *Controller_User) LoginUser(ctx context.Context, in *User_LoginRequest) (*Get_UserResponse, error) {

	data, err := cu.controller.LoginUser(ctx, in)
	if err != nil {
		return &Get_UserResponse{}, err
	}

	//  ปั้น token **********************
	jwtToken, err := auth.New_JWT("ratthakorn")
	if err != nil {
		return &Get_UserResponse{}, err
	}
	// crete token ********************
	token, err := jwtToken.CreateToken(data.GetUsername(), 60*time.Minute)
	data.Token = token

	return data, nil
}


func (cu *Controller_User) CreateUser(ctx context.Context, data *Create_UserRequest) (*StatusResponse, error) {
	
	// to Server ***************************
	status, err := cu.controller.CreateUser(ctx, data)
	if err != nil {
		fmt.Println("Client CreateUser error:", err.Error())
		return &StatusResponse{}, err
	}

	// fmt.Println("Client CreateUser:", errorUser.GetStatus())
	return status, nil
}

func (cu *Controller_User) GetUser_ByUsername(ctx context.Context, data *Username) (*Get_UserResponse, error) {
	
	//  ปั้น token **********************
	jwtToken, err := auth.New_JWT("ratthakorn")
	if err != nil {
		return &Get_UserResponse{}, err
	}
	
	//  VerifyToken ********************
	payload,err:= jwtToken.AuthorizeUser(ctx)
	if err != nil {
		return &Get_UserResponse{}, err
	}

	data.Username= payload.Audience

	// to Server ***************************
	userResponse, err := cu.controller.GetUser_ByUsername(ctx, data)
	if err != nil {
		return &Get_UserResponse{}, err
	}

	return userResponse, nil
}


// เทสแล้ว ******************************
func (cu *Controller_User) UpdateUser(ctx context.Context, data *User_UpdateRequest) (*StatusResponse, error) {
	
	//  ปั้น token **********************
	jwtToken, err := auth.New_JWT("ratthakorn")
	if err != nil {
		return &StatusResponse{}, err
	}
	
	//  VerifyToken ********************
	payload,err:= jwtToken.AuthorizeUser(ctx)
	if err != nil {
		return &StatusResponse{}, err
	}

	data.Username = payload.Audience

	// to Server ***************************
	status, err := cu.controller.UpdateUser(ctx, data)

	if err != nil {
		fmt.Println("Error Update Client : ", err.Error())
		return &StatusResponse{}, err
	}

	return status, nil
}

func (cu *Controller_User) ChangePassword(ctx context.Context, data *Change_PasswordRequest) (*StatusResponse, error) {

	//  ปั้น token **********************
	jwtToken, err := auth.New_JWT("ratthakorn")
	if err != nil {
		return &StatusResponse{}, err
	}
	
	//  VerifyToken ********************
	payload,err:= jwtToken.AuthorizeUser(ctx)
	if err != nil {
		return &StatusResponse{}, err
	}

	data.Username= &payload.Audience

	// to Server *****************************
	status, err := cu.controller.ChangePassword(ctx, data)
	if err != nil {
		fmt.Println("Change password error : ", err.Error())
		return &StatusResponse{}, err
	}

	return status, nil

}
