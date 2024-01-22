package users

import (
	userProto "banktransfer/users"
	userRepo "banktransfer/users/repo"
	"banktransfer/util"

	"banktransfer/models"

	// "banktransfer/transfers/repo"
	context "context"
	"fmt"

	"errors"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type service_User struct {
	userProto.UnimplementedUserServer
	repo userRepo.I_Rero_Users
}

func New_Service_User(repo userRepo.I_Rero_Users) userProto.UserServer {
	return &service_User{repo: repo}
}

// Login ***************
func (su *service_User) LoginUser(ctx context.Context, user *userProto.User_LoginRequest) (*userProto.Get_UserResponse, error) {
	fmt.Println("PAssword Server:", user.GetPassword())
	data, err := su.repo.Login(user)
	if err != nil {
		return &userProto.Get_UserResponse{}, err
	}

	// เอา password มาเช็ค ** ********
	errCompairPassword := util.CompairPassword(user.GetPassword(),data.HashedPassword)
	if errCompairPassword != nil {
		return &userProto.Get_UserResponse{}, errors.New("password is wrong")
	}

	// if data.HashedPassword != user.GetPassword() {
	// 	return &userProto.Get_UserResponse{}, errors.New("password is wrong")
	// }

	ff := userProto.Get_UserResponse{
		Username:          data.Username,
		Email:             data.Email,
		Fullname:          data.FullName,
		PasswordChangedAt: timestamppb.New(data.PasswordChangedAt),
		CreatedAt:         timestamppb.New(data.CreatedAt),
	}

	return &ff, nil
}

// CreateUser implements users.UserServer
func (su *service_User) CreateUser(ctx context.Context, ur *userProto.Create_UserRequest) (*userProto.StatusResponse, error) {
	fmt.Println("CreateUser Server service *******************")
	fmt.Println("Email Server: ", ur.GetEmail())
	return &userProto.StatusResponse{
		Status: "error proto server",
	}, nil
}

// GetUser_ByUsername implements users.UserServer
func (su *service_User) GetUser_ByUsername(ctx context.Context, username *userProto.Username) (*userProto.Get_UserResponse, error) {
	fmt.Println("GetUser_ByUsername Server service ***********")
	// ค่า  username.GetUsername() จะไม่มีเพราะ  method get จะไม่มี body มา
	// fmt.Println("Username : ",username.GetUsername())

	// ต้องเอา username จาก  header *****
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

	// hash pasword **********************
	hashPassword, err := util.HashPassword(data.GetNewpassword())
	if err != nil {
		return &userProto.StatusResponse{}, err
	}

	// ปั้นข้อมูลใหม่ ****************
	data.Newpassword = hashPassword
	
	
	// repo *****************
	err = su.repo.ChangePassword(data)
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
	fmt.Println("UpdateUser Server:", newData)

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
