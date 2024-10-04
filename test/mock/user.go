// Package mock  ------------------------------------------------------------
// @file      : user.go
// @author    : WeiTao
// @contact   : 15537588047@163.com
// @time      : 2024/10/4 13:16
// ------------------------------------------------------------
package mock

import "context"

type User struct {
	Mobile   string
	Password string
	NickName string
}

type UserServer struct {
	Db UserData
}

func (us *UserServer) GetUserByMobile(ctx context.Context, mobile string) (User, error) {
	user, err := us.Db.GetUserByMobile(ctx, mobile)
	if err != nil {
		return User{}, err
	}
	if user.NickName == "TestUser" {
		user.NickName = "TestUserModified"
	}
	return user, nil
}

type UserData interface {
	GetUserByMobile(ctx context.Context, mobile string) (User, error)
}
