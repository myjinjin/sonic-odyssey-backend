// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	usecase "github.com/myjinjin/sonic-odyssey-backend/internal/usecase"
	mock "github.com/stretchr/testify/mock"
)

// UserUsecase is an autogenerated mock type for the UserUsecase type
type UserUsecase struct {
	mock.Mock
}

// GetUserByID provides a mock function with given fields: userID
func (_m *UserUsecase) GetUserByID(userID uint) (*usecase.GetUserByIDOutput, error) {
	ret := _m.Called(userID)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByID")
	}

	var r0 *usecase.GetUserByIDOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*usecase.GetUserByIDOutput, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(uint) *usecase.GetUserByIDOutput); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*usecase.GetUserByIDOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PatchUser provides a mock function with given fields: userID, input
func (_m *UserUsecase) PatchUser(userID uint, input *usecase.PatchUserInput) error {
	ret := _m.Called(userID, input)

	if len(ret) == 0 {
		panic("no return value specified for PatchUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, *usecase.PatchUserInput) error); ok {
		r0 = rf(userID, input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ResetPassword provides a mock function with given fields: password, flowID
func (_m *UserUsecase) ResetPassword(password string, flowID string) error {
	ret := _m.Called(password, flowID)

	if len(ret) == 0 {
		panic("no return value specified for ResetPassword")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(password, flowID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendPasswordRecoveryEmail provides a mock function with given fields: baseURL, email
func (_m *UserUsecase) SendPasswordRecoveryEmail(baseURL string, email string) error {
	ret := _m.Called(baseURL, email)

	if len(ret) == 0 {
		panic("no return value specified for SendPasswordRecoveryEmail")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(baseURL, email)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SignUp provides a mock function with given fields: _a0
func (_m *UserUsecase) SignUp(_a0 usecase.SignUpInput) (*usecase.SignUpOutput, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for SignUp")
	}

	var r0 *usecase.SignUpOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(usecase.SignUpInput) (*usecase.SignUpOutput, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(usecase.SignUpInput) *usecase.SignUpOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*usecase.SignUpOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(usecase.SignUpInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserUsecase creates a new instance of UserUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserUsecase {
	mock := &UserUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
