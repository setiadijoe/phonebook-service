// Code generated by MockGen. DO NOT EDIT.
// Source: repository/repository_interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/sapawarga/phonebook-service/model"
)

// MockPhoneBookI is a mock of PhoneBookI interface.
type MockPhoneBookI struct {
	ctrl     *gomock.Controller
	recorder *MockPhoneBookIMockRecorder
}

// MockPhoneBookIMockRecorder is the mock recorder for MockPhoneBookI.
type MockPhoneBookIMockRecorder struct {
	mock *MockPhoneBookI
}

// NewMockPhoneBookI creates a new mock instance.
func NewMockPhoneBookI(ctrl *gomock.Controller) *MockPhoneBookI {
	mock := &MockPhoneBookI{ctrl: ctrl}
	mock.recorder = &MockPhoneBookIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPhoneBookI) EXPECT() *MockPhoneBookIMockRecorder {
	return m.recorder
}

// CheckHealthReadiness mocks base method.
func (m *MockPhoneBookI) CheckHealthReadiness(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckHealthReadiness", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckHealthReadiness indicates an expected call of CheckHealthReadiness.
func (mr *MockPhoneBookIMockRecorder) CheckHealthReadiness(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckHealthReadiness", reflect.TypeOf((*MockPhoneBookI)(nil).CheckHealthReadiness), ctx)
}

// Delete mocks base method.
func (m *MockPhoneBookI) Delete(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockPhoneBookIMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPhoneBookI)(nil).Delete), ctx, id)
}

// GetCategoryNameByID mocks base method.
func (m *MockPhoneBookI) GetCategoryNameByID(ctx context.Context, id int64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategoryNameByID", ctx, id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategoryNameByID indicates an expected call of GetCategoryNameByID.
func (mr *MockPhoneBookIMockRecorder) GetCategoryNameByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategoryNameByID", reflect.TypeOf((*MockPhoneBookI)(nil).GetCategoryNameByID), ctx, id)
}

// GetListPhoneBook mocks base method.
func (m *MockPhoneBookI) GetListPhoneBook(ctx context.Context, params *model.GetListRequest) ([]*model.PhoneBookResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListPhoneBook", ctx, params)
	ret0, _ := ret[0].([]*model.PhoneBookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListPhoneBook indicates an expected call of GetListPhoneBook.
func (mr *MockPhoneBookIMockRecorder) GetListPhoneBook(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListPhoneBook", reflect.TypeOf((*MockPhoneBookI)(nil).GetListPhoneBook), ctx, params)
}

// GetListPhonebookByLongLat mocks base method.
func (m *MockPhoneBookI) GetListPhonebookByLongLat(ctx context.Context, params *model.GetListRequest) ([]*model.PhoneBookResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListPhonebookByLongLat", ctx, params)
	ret0, _ := ret[0].([]*model.PhoneBookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListPhonebookByLongLat indicates an expected call of GetListPhonebookByLongLat.
func (mr *MockPhoneBookIMockRecorder) GetListPhonebookByLongLat(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListPhonebookByLongLat", reflect.TypeOf((*MockPhoneBookI)(nil).GetListPhonebookByLongLat), ctx, params)
}

// GetListPhonebookByLongLatMeta mocks base method.
func (m *MockPhoneBookI) GetListPhonebookByLongLatMeta(ctx context.Context, params *model.GetListRequest) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListPhonebookByLongLatMeta", ctx, params)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListPhonebookByLongLatMeta indicates an expected call of GetListPhonebookByLongLatMeta.
func (mr *MockPhoneBookIMockRecorder) GetListPhonebookByLongLatMeta(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListPhonebookByLongLatMeta", reflect.TypeOf((*MockPhoneBookI)(nil).GetListPhonebookByLongLatMeta), ctx, params)
}

// GetLocationByID mocks base method.
func (m *MockPhoneBookI) GetLocationByID(ctx context.Context, id int64) (*model.Location, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLocationByID", ctx, id)
	ret0, _ := ret[0].(*model.Location)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLocationByID indicates an expected call of GetLocationByID.
func (mr *MockPhoneBookIMockRecorder) GetLocationByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLocationByID", reflect.TypeOf((*MockPhoneBookI)(nil).GetLocationByID), ctx, id)
}

// GetMetaDataPhoneBook mocks base method.
func (m *MockPhoneBookI) GetMetaDataPhoneBook(ctx context.Context, params *model.GetListRequest) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetaDataPhoneBook", ctx, params)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMetaDataPhoneBook indicates an expected call of GetMetaDataPhoneBook.
func (mr *MockPhoneBookIMockRecorder) GetMetaDataPhoneBook(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetaDataPhoneBook", reflect.TypeOf((*MockPhoneBookI)(nil).GetMetaDataPhoneBook), ctx, params)
}

// GetPhonebookDetailByID mocks base method.
func (m *MockPhoneBookI) GetPhonebookDetailByID(ctx context.Context, id int64) (*model.PhoneBookResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPhonebookDetailByID", ctx, id)
	ret0, _ := ret[0].(*model.PhoneBookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPhonebookDetailByID indicates an expected call of GetPhonebookDetailByID.
func (mr *MockPhoneBookIMockRecorder) GetPhonebookDetailByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPhonebookDetailByID", reflect.TypeOf((*MockPhoneBookI)(nil).GetPhonebookDetailByID), ctx, id)
}

// Insert mocks base method.
func (m *MockPhoneBookI) Insert(ctx context.Context, params *model.AddPhonebook) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, params)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockPhoneBookIMockRecorder) Insert(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockPhoneBookI)(nil).Insert), ctx, params)
}

// Update mocks base method.
func (m *MockPhoneBookI) Update(ctx context.Context, params *model.UpdatePhonebook) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, params)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockPhoneBookIMockRecorder) Update(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPhoneBookI)(nil).Update), ctx, params)
}
