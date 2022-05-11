package repositories

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"testing"
	"user-context/diamond/domain/account/entity"
	"user-context/utils/mock"
)

func TestAccountAdapter_CheckIsExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mock.NewMockAccountRepository(ctrl)
	first := "account_00001"
	second := "account_0000a"

	//期望值
	repository.EXPECT().CheckIsExist(entity.Account{ID: first}).Return(true)
	repository.EXPECT().CheckIsExist(entity.Account{ID: second}).Return(false)

	//输出结果
	if repository.CheckIsExist(entity.Account{ID: first}) {
		fmt.Println("first exist")
	}
	if repository.CheckIsExist(entity.Account{ID: second}) {
		fmt.Println("second no exist")
	}
}

func TestAccountAdapter_Insert(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mock.NewMockAccountRepository(ctrl)

	//期望值
	repository.EXPECT().Insert(entity.Account{ID: "", NickName: "bba", PassWord: "abc_abc_abc", Mobile: "152000111000", Email: "", ThirdPartyID: ""}).Return(nil)
	repository.EXPECT().Insert(entity.Account{}).Return(errors.New("struct exception"))

	//输出结果
	if err := repository.Insert(entity.Account{ID: "", NickName: "bba", PassWord: "abc_abc_abc", Mobile: "152000111000", Email: "", ThirdPartyID: ""}); err != nil {
		fmt.Println("err: ", err)
		return
	} else {
		fmt.Println("insert success")
	}
	if err := repository.Insert(entity.Account{}); err != nil {
		fmt.Println("err: ", err)
		return
	} else {
		fmt.Println("insert success")
	}
}
