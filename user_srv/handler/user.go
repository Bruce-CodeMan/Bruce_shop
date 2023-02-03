/*
 * @Date: 2023-02-03 16:47:48
 * @Author: Bruce
 * @Description:
 */

package handler

import (
	"context"

	"gorm.io/gorm"

	"Bruce_shop/user_srv/global"
	"Bruce_shop/user_srv/model"
	"Bruce_shop/user_srv/proto"
)

type UserServer struct{}

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

// 获取用户列表
func (s *UserServer) GetUserList(ctx context.Context, req *proto.PageInfo) (*proto.UserListResponse, error) {
	var users []model.User
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
