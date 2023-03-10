/*
 * @Date: 2023-02-03 16:47:48
 * @Author: Bruce
 * @Description: UserServer's methods
 */

package handler

import (
	"context"
	"crypto/sha512"
	"fmt"
	"strings"
	"time"

	"github.com/anaskhan96/go-password-encoder"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"Bruce_shop/srvs/user_srv/global"
	"Bruce_shop/srvs/user_srv/model"
	"Bruce_shop/srvs/user_srv/proto"
)

type UserServer struct {
	proto.UnimplementedUserServer
}

// Model2Response Convert the object of user Model into the object of proto
func Model2Response(user model.User) proto.UserInfoResponse {
	// 在grpc的message中字段有默认值，不能随便赋值nil进去，容易出错
	// 要搞清楚，哪些字段是有默认值
	userInfoRsp := proto.UserInfoResponse{
		Id:       user.ID,
		Password: user.Password,
		NickName: user.NickName,
		Gender:   user.Gender,
		Mobile:   user.Mobile,
		Role:     int32(user.Role),
	}
	if user.Birthday != nil {
		userInfoRsp.BirthDay = uint64(user.Birthday.Unix())
	}
	return userInfoRsp
}

// Paginate Get the data according to page/pageSize parameters
func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

// GetUserList Query the whole userInfo
func (s *UserServer) GetUserList(_ context.Context, req *proto.PageInfo) (*proto.UserListResponse, error) {
	var users []model.User
	fmt.Println("用户列表")
	result := global.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	resp := &proto.UserListResponse{}
	resp.Total = int32(result.RowsAffected)

	global.DB.Scopes(Paginate(int(req.Pn), int(req.PSize))).Find(&users)
	for _, user := range users {
		userInfoRsp := Model2Response(user)
		// resp.Data = append(resp.Data, &(Model2Response(user)))   注意这样写是不行的，必须要声明一个变量来取地址
		resp.Data = append(resp.Data, &userInfoRsp)
	}
	return resp, nil
}

// GetUserByMobile Query the userInformation according to the user's mobile parameter
func (s *UserServer) GetUserByMobile(_ context.Context, req *proto.MobileRequest) (*proto.UserInfoResponse, error) {
	var user model.User
	result := global.DB.Where(&model.User{
		Mobile: req.Mobile,
	}).First(&user)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	userInfoResp := Model2Response(user)
	return &userInfoResp, nil
}

// GetUserById Query the userInformation according to the user's id parameter
func (s *UserServer) GetUserById(_ context.Context, req *proto.IdRequest) (*proto.UserInfoResponse, error) {
	var user model.User
	result := global.DB.First(&user, req.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	userInfoResp := Model2Response(user)
	return &userInfoResp, nil
}

// CreateUser Create a userInfo
func (s *UserServer) CreateUser(_ context.Context, req *proto.CreateUserInfo) (*proto.UserInfoResponse, error) {
	var user model.User
	result := global.DB.Where(&model.User{Mobile: req.Mobile}).First(&user)
	if result.RowsAffected == 1 {
		return nil, status.Errorf(codes.AlreadyExists, "用户已存在")
	}
	user.Mobile = req.Mobile
	user.NickName = req.NickName

	// 密码加密
	options := &password.Options{
		SaltLen:      16,
		Iterations:   100,
		KeyLen:       32,
		HashFunction: sha512.New,
	}
	salt, encodedPwd := password.Encode(req.Password, options)
	// 最终生成的密码,使用$进行分割,$算法$盐值$密码
	user.Password = fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)

	// 保存
	result = global.DB.Create(&user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}

	userInfoResp := Model2Response(user)
	return &userInfoResp, nil
}

// UpdateUser Update the userInfo according to the user's id ,empty引用的库"github.com/golang/protobuf/ptypes/empty"
func (s *UserServer) UpdateUser(_ context.Context, req *proto.UpdateUserInfo) (*empty.Empty, error) {
	var user model.User
	result := global.DB.First(&user, req.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	// Go语言将int类型转换成time类型
	birthDay := time.Unix(int64(req.BirthDay), 0)
	user.NickName = req.NickName
	user.Birthday = &birthDay
	user.Gender = req.Gender
	result = global.DB.Save(&user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	return &empty.Empty{}, nil
}

// CheckPassword check user's password whether is correct according to password parameter
func (s *UserServer) CheckPassword(_ context.Context, req *proto.PasswordCheckInfo) (*proto.CheckResponse, error) {
	options := &password.Options{
		SaltLen:      16,
		Iterations:   100,
		KeyLen:       32,
		HashFunction: sha512.New,
	}
	passwordInfo := strings.Split(req.EncryptoPassword, "$")
	check := password.Verify(req.Password, passwordInfo[2], passwordInfo[3], options)

	return &proto.CheckResponse{
		Success: check,
	}, nil
}
