// ------------------------------------------------------------
// package test
// @file      : user_test.go
// @author    : WeiTao
// @contact   : 15537588047@163.com
// @time      : 2024/10/4 13:23
// ------------------------------------------------------------
package test

import (
	"context"
	"github.com/golang/mock/gomock"
	"learn_go/test/mock"
	"testing"
)

func TestGetUserByMobile(t *testing.T) {
	// mock准备工作
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	userData := mock.NewMockUserData(ctl)
	userData.EXPECT().GetUserByMobile(gomock.Any(), "15023076751").Return(
		User{
			Mobile:   "15023076751",
			Password: "123456",
			NickName: "TestUser",
		},
		nil,
	)
	// 实际调用过程
	userServer := UserServer{
		Db: userData,
	}
	user, err := userServer.GetUserByMobile(context.Background(), "15023076751")
	// 判断过程
	if err != nil {
		t.Error("GetUserByMobile error:", err)
	}
	if user.Mobile != "15023076751" || user.Password != "123456" || user.NickName != "TestUserModified" {
		t.Error("GetUserByMobile result is not expected.")
	}
}
