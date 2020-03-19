package usecase

import (
	"fmt"
	. "github.com/petergtz/pegomock"
	"github.com/soluty/x/clean/adapter/mock"
	"github.com/soluty/x/clean/entity"
	"testing"
)

func TestInteractor_UserSelect(t *testing.T) {
	//mockCtrl := gomock.NewController(t)
	//defer mockCtrl.Finish()
	//
	//authToken := "anyToken"
	//rick := testData.User("rick")
	//i := mock.NewMockedInteractor(mockCtrl)
	//i.UserRW.EXPECT().GetByName(rick.Name).Return(&rick, nil).Times(1)
	//i.AuthHandler.EXPECT().GenUserToken(rick.Name).Return(authToken, nil).Times(1)
	//
	//retUser, token, err := i.GetUCHandler().UserGet(rick.Name)
	//assert.NoError(t, err)
	//assert.Equal(t, rick, *retUser)
	//assert.Equal(t, authToken, token)

	RegisterMockTestingT(t)

	rick := &entity.User{
		Name: "rick",
	}

	repo := mock.NewMockUserRepo()
	When(repo.GetById(1)).ThenReturn(rick, nil)
	a, e := repo.GetById(1)
	fmt.Println(a, e)
}
