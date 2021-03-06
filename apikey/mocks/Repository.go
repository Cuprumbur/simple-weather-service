// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	model "github.com/Cuprumbur/weather-service/model"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: id
func (_m *Repository) Delete(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAllApiKeys provides a mock function with given fields:
func (_m *Repository) FindAllApiKeys() ([]*model.ApiKey, error) {
	ret := _m.Called()

	var r0 []*model.ApiKey
	if rf, ok := ret.Get(0).(func() []*model.ApiKey); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.ApiKey)
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

// FindApiKey provides a mock function with given fields: id
func (_m *Repository) FindApiKey(id int) (*model.ApiKey, error) {
	ret := _m.Called(id)

	var r0 *model.ApiKey
	if rf, ok := ret.Get(0).(func(int) *model.ApiKey); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ApiKey)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindApiKeys provides a mock function with given fields: detectorID
func (_m *Repository) FindApiKeys(detectorID int) ([]*model.ApiKey, error) {
	ret := _m.Called(detectorID)

	var r0 []*model.ApiKey
	if rf, ok := ret.Get(0).(func(int) []*model.ApiKey); ok {
		r0 = rf(detectorID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.ApiKey)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(detectorID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: key
func (_m *Repository) Store(key model.ApiKey) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.ApiKey) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateScopes provides a mock function with given fields: id, scopes
func (_m *Repository) UpdateScopes(id int, scopes []string) error {
	ret := _m.Called(id, scopes)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, []string) error); ok {
		r0 = rf(id, scopes)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
