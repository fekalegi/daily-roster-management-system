// Code generated by MockGen. DO NOT EDIT.
// Source: repository_interface.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	model "daily-worker-roster-management-system/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAssignmentRepository is a mock of AssignmentRepository interface.
type MockAssignmentRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAssignmentRepositoryMockRecorder
}

// MockAssignmentRepositoryMockRecorder is the mock recorder for MockAssignmentRepository.
type MockAssignmentRepositoryMockRecorder struct {
	mock *MockAssignmentRepository
}

// NewMockAssignmentRepository creates a new mock instance.
func NewMockAssignmentRepository(ctrl *gomock.Controller) *MockAssignmentRepository {
	mock := &MockAssignmentRepository{ctrl: ctrl}
	mock.recorder = &MockAssignmentRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAssignmentRepository) EXPECT() *MockAssignmentRepositoryMockRecorder {
	return m.recorder
}

// CountAssignmentsInWeek mocks base method.
func (m *MockAssignmentRepository) CountAssignmentsInWeek(workerID int, weekStart, weekEnd string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountAssignmentsInWeek", workerID, weekStart, weekEnd)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountAssignmentsInWeek indicates an expected call of CountAssignmentsInWeek.
func (mr *MockAssignmentRepositoryMockRecorder) CountAssignmentsInWeek(workerID, weekStart, weekEnd interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountAssignmentsInWeek", reflect.TypeOf((*MockAssignmentRepository)(nil).CountAssignmentsInWeek), workerID, weekStart, weekEnd)
}

// Create mocks base method.
func (m *MockAssignmentRepository) Create(a *model.Assignment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", a)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockAssignmentRepositoryMockRecorder) Create(a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAssignmentRepository)(nil).Create), a)
}

// FindAll mocks base method.
func (m *MockAssignmentRepository) FindAll() ([]model.AssignmentDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]model.AssignmentDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockAssignmentRepositoryMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockAssignmentRepository)(nil).FindAll))
}

// FindByDate mocks base method.
func (m *MockAssignmentRepository) FindByDate(date string) ([]model.AssignmentDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByDate", date)
	ret0, _ := ret[0].([]model.AssignmentDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByDate indicates an expected call of FindByDate.
func (mr *MockAssignmentRepositoryMockRecorder) FindByDate(date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByDate", reflect.TypeOf((*MockAssignmentRepository)(nil).FindByDate), date)
}

// FindByWorkerID mocks base method.
func (m *MockAssignmentRepository) FindByWorkerID(workerID int) ([]model.AssignmentDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByWorkerID", workerID)
	ret0, _ := ret[0].([]model.AssignmentDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByWorkerID indicates an expected call of FindByWorkerID.
func (mr *MockAssignmentRepositoryMockRecorder) FindByWorkerID(workerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByWorkerID", reflect.TypeOf((*MockAssignmentRepository)(nil).FindByWorkerID), workerID)
}

// HasAssignmentOnDay mocks base method.
func (m *MockAssignmentRepository) HasAssignmentOnDay(workerID int, date string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasAssignmentOnDay", workerID, date)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasAssignmentOnDay indicates an expected call of HasAssignmentOnDay.
func (mr *MockAssignmentRepositoryMockRecorder) HasAssignmentOnDay(workerID, date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasAssignmentOnDay", reflect.TypeOf((*MockAssignmentRepository)(nil).HasAssignmentOnDay), workerID, date)
}

// IsShiftAlreadyAssigned mocks base method.
func (m *MockAssignmentRepository) IsShiftAlreadyAssigned(shiftID int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsShiftAlreadyAssigned", shiftID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsShiftAlreadyAssigned indicates an expected call of IsShiftAlreadyAssigned.
func (mr *MockAssignmentRepositoryMockRecorder) IsShiftAlreadyAssigned(shiftID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsShiftAlreadyAssigned", reflect.TypeOf((*MockAssignmentRepository)(nil).IsShiftAlreadyAssigned), shiftID)
}

// MockWorkerRepository is a mock of WorkerRepository interface.
type MockWorkerRepository struct {
	ctrl     *gomock.Controller
	recorder *MockWorkerRepositoryMockRecorder
}

// MockWorkerRepositoryMockRecorder is the mock recorder for MockWorkerRepository.
type MockWorkerRepositoryMockRecorder struct {
	mock *MockWorkerRepository
}

// NewMockWorkerRepository creates a new mock instance.
func NewMockWorkerRepository(ctrl *gomock.Controller) *MockWorkerRepository {
	mock := &MockWorkerRepository{ctrl: ctrl}
	mock.recorder = &MockWorkerRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWorkerRepository) EXPECT() *MockWorkerRepositoryMockRecorder {
	return m.recorder
}

// GetAll mocks base method.
func (m *MockWorkerRepository) GetAll() ([]model.Worker, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]model.Worker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockWorkerRepositoryMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockWorkerRepository)(nil).GetAll))
}

// GetByName mocks base method.
func (m *MockWorkerRepository) GetByName(name string) (*model.Worker, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByName", name)
	ret0, _ := ret[0].(*model.Worker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByName indicates an expected call of GetByName.
func (mr *MockWorkerRepositoryMockRecorder) GetByName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByName", reflect.TypeOf((*MockWorkerRepository)(nil).GetByName), name)
}

// MockShiftRepository is a mock of ShiftRepository interface.
type MockShiftRepository struct {
	ctrl     *gomock.Controller
	recorder *MockShiftRepositoryMockRecorder
}

// MockShiftRepositoryMockRecorder is the mock recorder for MockShiftRepository.
type MockShiftRepositoryMockRecorder struct {
	mock *MockShiftRepository
}

// NewMockShiftRepository creates a new mock instance.
func NewMockShiftRepository(ctrl *gomock.Controller) *MockShiftRepository {
	mock := &MockShiftRepository{ctrl: ctrl}
	mock.recorder = &MockShiftRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockShiftRepository) EXPECT() *MockShiftRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockShiftRepository) Create(shift *model.Shift) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", shift)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockShiftRepositoryMockRecorder) Create(shift interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockShiftRepository)(nil).Create), shift)
}

// GetAll mocks base method.
func (m *MockShiftRepository) GetAll() ([]model.Shift, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]model.Shift)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockShiftRepositoryMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockShiftRepository)(nil).GetAll))
}

// GetByID mocks base method.
func (m *MockShiftRepository) GetByID(id int) (*model.Shift, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(*model.Shift)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockShiftRepositoryMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockShiftRepository)(nil).GetByID), id)
}

// GetUnassignedShift mocks base method.
func (m *MockShiftRepository) GetUnassignedShift() ([]model.Shift, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUnassignedShift")
	ret0, _ := ret[0].([]model.Shift)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUnassignedShift indicates an expected call of GetUnassignedShift.
func (mr *MockShiftRepositoryMockRecorder) GetUnassignedShift() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUnassignedShift", reflect.TypeOf((*MockShiftRepository)(nil).GetUnassignedShift))
}

// MockShiftRequestRepository is a mock of ShiftRequestRepository interface.
type MockShiftRequestRepository struct {
	ctrl     *gomock.Controller
	recorder *MockShiftRequestRepositoryMockRecorder
}

// MockShiftRequestRepositoryMockRecorder is the mock recorder for MockShiftRequestRepository.
type MockShiftRequestRepositoryMockRecorder struct {
	mock *MockShiftRequestRepository
}

// NewMockShiftRequestRepository creates a new mock instance.
func NewMockShiftRequestRepository(ctrl *gomock.Controller) *MockShiftRequestRepository {
	mock := &MockShiftRequestRepository{ctrl: ctrl}
	mock.recorder = &MockShiftRequestRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockShiftRequestRepository) EXPECT() *MockShiftRequestRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockShiftRequestRepository) Create(req *model.ShiftRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockShiftRequestRepositoryMockRecorder) Create(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockShiftRequestRepository)(nil).Create), req)
}

// GetByID mocks base method.
func (m *MockShiftRequestRepository) GetByID(id int) (*model.ShiftRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(*model.ShiftRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockShiftRequestRepositoryMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockShiftRequestRepository)(nil).GetByID), id)
}

// GetPending mocks base method.
func (m *MockShiftRequestRepository) GetPending() ([]model.ShiftRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPending")
	ret0, _ := ret[0].([]model.ShiftRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPending indicates an expected call of GetPending.
func (mr *MockShiftRequestRepositoryMockRecorder) GetPending() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPending", reflect.TypeOf((*MockShiftRequestRepository)(nil).GetPending))
}

// GetPendingWithDetails mocks base method.
func (m *MockShiftRequestRepository) GetPendingWithDetails() ([]model.ShiftRequestDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPendingWithDetails")
	ret0, _ := ret[0].([]model.ShiftRequestDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPendingWithDetails indicates an expected call of GetPendingWithDetails.
func (mr *MockShiftRequestRepositoryMockRecorder) GetPendingWithDetails() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPendingWithDetails", reflect.TypeOf((*MockShiftRequestRepository)(nil).GetPendingWithDetails))
}

// GetWithDetails mocks base method.
func (m *MockShiftRequestRepository) GetWithDetails() ([]model.ShiftRequestDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWithDetails")
	ret0, _ := ret[0].([]model.ShiftRequestDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWithDetails indicates an expected call of GetWithDetails.
func (mr *MockShiftRequestRepositoryMockRecorder) GetWithDetails() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWithDetails", reflect.TypeOf((*MockShiftRequestRepository)(nil).GetWithDetails))
}

// GetWithDetailsFilterStatus mocks base method.
func (m *MockShiftRequestRepository) GetWithDetailsFilterStatus(status string) ([]model.ShiftRequestDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWithDetailsFilterStatus", status)
	ret0, _ := ret[0].([]model.ShiftRequestDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWithDetailsFilterStatus indicates an expected call of GetWithDetailsFilterStatus.
func (mr *MockShiftRequestRepositoryMockRecorder) GetWithDetailsFilterStatus(status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWithDetailsFilterStatus", reflect.TypeOf((*MockShiftRequestRepository)(nil).GetWithDetailsFilterStatus), status)
}

// UpdateStatus mocks base method.
func (m *MockShiftRequestRepository) UpdateStatus(id int, status string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", id, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStatus indicates an expected call of UpdateStatus.
func (mr *MockShiftRequestRepositoryMockRecorder) UpdateStatus(id, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockShiftRequestRepository)(nil).UpdateStatus), id, status)
}
