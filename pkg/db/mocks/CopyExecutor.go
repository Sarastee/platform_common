// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	pgx "github.com/jackc/pgx/v5"
)

// CopyExecutor is an autogenerated mock type for the CopyExecutor type
type CopyExecutor struct {
	mock.Mock
}

type CopyExecutor_Expecter struct {
	mock *mock.Mock
}

func (_m *CopyExecutor) EXPECT() *CopyExecutor_Expecter {
	return &CopyExecutor_Expecter{mock: &_m.Mock}
}

// CopyFromContext provides a mock function with given fields: ctx, tableName, columnNames, rowSrc
func (_m *CopyExecutor) CopyFromContext(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error) {
	ret := _m.Called(ctx, tableName, columnNames, rowSrc)

	if len(ret) == 0 {
		panic("no return value specified for CopyFromContext")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Identifier, []string, pgx.CopyFromSource) (int64, error)); ok {
		return rf(ctx, tableName, columnNames, rowSrc)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Identifier, []string, pgx.CopyFromSource) int64); ok {
		r0 = rf(ctx, tableName, columnNames, rowSrc)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.Identifier, []string, pgx.CopyFromSource) error); ok {
		r1 = rf(ctx, tableName, columnNames, rowSrc)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CopyExecutor_CopyFromContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CopyFromContext'
type CopyExecutor_CopyFromContext_Call struct {
	*mock.Call
}

// CopyFromContext is a helper method to define mock.On call
//   - ctx context.Context
//   - tableName pgx.Identifier
//   - columnNames []string
//   - rowSrc pgx.CopyFromSource
func (_e *CopyExecutor_Expecter) CopyFromContext(ctx interface{}, tableName interface{}, columnNames interface{}, rowSrc interface{}) *CopyExecutor_CopyFromContext_Call {
	return &CopyExecutor_CopyFromContext_Call{Call: _e.mock.On("CopyFromContext", ctx, tableName, columnNames, rowSrc)}
}

func (_c *CopyExecutor_CopyFromContext_Call) Run(run func(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource)) *CopyExecutor_CopyFromContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(pgx.Identifier), args[2].([]string), args[3].(pgx.CopyFromSource))
	})
	return _c
}

func (_c *CopyExecutor_CopyFromContext_Call) Return(_a0 int64, _a1 error) *CopyExecutor_CopyFromContext_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CopyExecutor_CopyFromContext_Call) RunAndReturn(run func(context.Context, pgx.Identifier, []string, pgx.CopyFromSource) (int64, error)) *CopyExecutor_CopyFromContext_Call {
	_c.Call.Return(run)
	return _c
}

// NewCopyExecutor creates a new instance of CopyExecutor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCopyExecutor(t interface {
	mock.TestingT
	Cleanup(func())
}) *CopyExecutor {
	mock := &CopyExecutor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
