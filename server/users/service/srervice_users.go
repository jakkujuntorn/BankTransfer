package users

import (
	userProto "banktransfer/users"
	userRepo "banktransfer/users/repo"

	"banktransfer/models"

	// "banktransfer/transfers/repo"
	context "context"
	"fmt"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type service_User struct {
	userProto.UnimplementedUserServer
	repo userRepo.I_Rero_Users
}

func New_Service_User(repo userRepo.I_Rero_Users) userProto.UserServer {
	return &service_User{repo: repo}
}

// CreateUser implements users.UserServer
func (su *service_User) CreateUser(context.Context, *userProto.Create_UserRequest) (*userProto.StatusResponse, error) {
	fmt.Println("CreateUser *******************")
	return &userProto.StatusResponse{
		Status: "error proto server",
	}, nil
}

// GetUser_ByUsername implements users.UserServer
func (su *service_User) GetUser_ByUsername(ctx context.Context, username *userProto.Username) (*userProto.Get_UserResponse, error) {

	userResponse, err := su.repo.GetUser_ByUsername(username.GetUsername())
	if err != nil {
		return &userProto.Get_UserResponse{}, err
	}

	return &userProto.Get_UserResponse{
		Username: userResponse.Username,
		Fullname: userResponse.FullName,
		Email:    userResponse.Email,
		PasswordChangedAt: &timestamppb.Timestamp{
			Seconds: userResponse.CreatedAt.Unix(),
			Nanos:   int32(userResponse.CreatedAt.Nanosecond()),
		},
		CreatedAt: &timestamppb.Timestamp{
			Seconds: userResponse.CreatedAt.Unix(),
			Nanos:   int32(userResponse.CreatedAt.Nanosecond()),
		},
	}, nil

}

// ChangePassword implements users.UserServer
func (su *service_User) ChangePassword(ctx context.Context, data *userProto.Change_PasswordRequest) (*userProto.StatusResponse, error) {
	fmt.Println("ChangePassword *******************")

	newPassword := models.UserChangePassword{Password: data.GetPassword()}

	err := su.repo.ChangePassword(data.GetUsername(), &newPassword)

	if err != nil {
		return &userProto.StatusResponse{}, err
	}

	return &userProto.StatusResponse{Status: "complete"}, nil
}

// UpdateUser implements users.UserServer
func (su *service_User) UpdateUser(ctx context.Context, data *userProto.User_UpdateRequest) (*userProto.StatusResponse, error) {
	fmt.Println("UpdateUser *******************")
	newData := models.UserUpdate{
		FullName: data.GetFullname(),
		Email:    data.GetEmail(),
	}

	err := su.repo.UpdateUser(data.GetUsername(), &newData)
	if err != nil {
		// fmt.Println("ERror Update :",err)
		// return &userProto.StatusResponse{Status: err.Error()}, err
		return &userProto.StatusResponse{}, err
	}
	return &userProto.StatusResponse{Status: "complete"}, nil
}

// mustEmbedUnimplementedUserServer implements users.UserServer
// func (*service_User) mustEmbedUnimplementedUserServer() {
// 	panic("unimplemented")
// }
