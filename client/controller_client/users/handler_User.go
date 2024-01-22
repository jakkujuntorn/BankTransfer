package users

import (
	context "context"
	"fmt"
	grpc "google.golang.org/grpc"
)

type Handler_Users struct {
	service I_User_Service
}

func New_Handler_Account(ac I_User_Service) UserClient {
	return &Handler_Users{
		service: ac,
	}
}

// เทสแล้ว ข้อมูลมา ****************************
// เหลือ เช็ค password กับ DB *******************
func (hu *Handler_Users)LoginUser(ctx context.Context, in *User_LoginRequest, opts ...grpc.CallOption) (*Get_UserResponse, error){
	fmt.Println("")
	dataLogin,err:=hu.service.LoginUser(ctx,in)
	if err != nil {
		return &Get_UserResponse{}, err
	}
	return dataLogin,nil
}

// ChangePassword implements UserClient
func (hu *Handler_Users) ChangePassword(ctx context.Context, in *Change_PasswordRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	status, err := hu.service.ChangePassword(ctx, in)
	if err != nil {
		return &StatusResponse{}, err
	}
	return status, nil
}

// CreateUser implements UserClient
func (hu *Handler_Users) CreateUser(ctx context.Context, in *Create_UserRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	
	status,err:=hu.service.CreateUser(ctx, in)
	if err != nil {
		return &StatusResponse{}, err
	}
	return status, nil
}

// GetUser_ByUsername implements UserClient
func (hu *Handler_Users) GetUser_ByUsername(ctx context.Context, in *Username, opts ...grpc.CallOption) (*Get_UserResponse, error) {
	dataUser, err := hu.service.GetUser_ByUsername(ctx, in)
	if err != nil {
		return &Get_UserResponse{}, err
	}
	return dataUser, nil
}

// UpdateUser implements UserClient
// เหลือ ดึง username จาก token, xั้นข้อมูล response
func (hu *Handler_Users) UpdateUser(ctx context.Context, in *User_UpdateRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	status, err := hu.service.UpdateUser(ctx, in)
	if err != nil {
		return &StatusResponse{}, err
	}
	return status, nil
}
