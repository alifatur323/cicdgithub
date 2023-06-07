package user

import (
	"testing"
)

type testUserParam struct {
	Id       uint   `json:"Id"`
	Username string `json:"Username"`
	Email    string `json:"Email"`
}

type testUseCaseUser struct {
	useCaseUser UseCaseUser
	Data        testUserParam `json:"data"`
}

func (uc testUseCaseUser) TestGetUserById(t *testing.T) {
	//result := uc.useCaseUser.GetUserById(1)
	//expected := json.Marshal(UseCase{Id: 1, Username: "jhon", Email: "email"})
	//expected := json.Marshal(uc.Data)
	//assert.JSONEq(t, result, expected)
}
