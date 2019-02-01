// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import models "github.com/ecojuntak/gorb/models"

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: u
func (_m *UserRepository) Create(u models.User) (models.User, error) {
	ret := _m.Called(u)

	var r0 models.User
	if rf, ok := ret.Get(0).(func(models.User) models.User); ok {
		r0 = rf(u)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.User) error); ok {
		r1 = rf(u)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *UserRepository) Delete(id int) (bool, error) {
	ret := _m.Called(id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int) bool); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, u
func (_m *UserRepository) Update(id int, u models.User) (models.User, error) {
	ret := _m.Called(id, u)

	var r0 models.User
	if rf, ok := ret.Get(0).(func(int, models.User) models.User); ok {
		r0 = rf(id, u)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, models.User) error); ok {
		r1 = rf(id, u)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// User provides a mock function with given fields: id
func (_m *UserRepository) User(id int) (models.User, error) {
	ret := _m.Called(id)

	var r0 models.User
	if rf, ok := ret.Get(0).(func(int) models.User); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Users provides a mock function with given fields:
func (_m *UserRepository) Users() ([]models.User, error) {
	ret := _m.Called()

	var r0 []models.User
	if rf, ok := ret.Get(0).(func() []models.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}