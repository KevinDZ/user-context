package repositories

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"testing"
	"user-context/rhombic/domain/account/entity"
	"user-context/utils/mock"
)

func TestAccountAdapter_CheckIsExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mock.NewMockAccountRepository(ctrl)
	first := "account_00001"
	second := "account_0000a"

	//期望值
	repository.EXPECT().CheckIsExist(entity.Entity{ID: first}).Return(true)
	repository.EXPECT().CheckIsExist(entity.Entity{ID: second}).Return(false)

	//输出结果
	if repository.CheckIsExist(entity.Entity{ID: first}) {
		fmt.Println("first exist")
	}
	if repository.CheckIsExist(entity.Entity{ID: second}) {
		fmt.Println("second no exist")
	}
}

func TestAccountAdapter_Insert(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mock.NewMockAccountRepository(ctrl)

	//期望值
	repository.EXPECT().Insert(entity.Entity{ID: "", Name: "bba", PassWord: "abc_abc_abc", Phone: "152000111000", Email: "", ThirdPartyID: "", Event: ""}).Return(nil)
	repository.EXPECT().Insert(entity.Entity{}).Return(errors.New("struct exception"))

	//输出结果
	if err := repository.Insert(entity.Entity{ID: "", Name: "bba", PassWord: "abc_abc_abc", Phone: "152000111000", Email: "", ThirdPartyID: "", Event: ""}); err != nil {
		fmt.Println("err: ", err)
		return
	} else {
		fmt.Println("insert success")
	}
	if err := repository.Insert(entity.Entity{}); err != nil {
		fmt.Println("err: ", err)
		return
	} else {
		fmt.Println("insert success")
	}
}
