// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks_repo

import (
	domain "github.com/kimjooyoon/project_board_backend_go/internal/article/domain"
	mock "github.com/stretchr/testify/mock"
)

// MockArticleQuery is an autogenerated mock type for the ArticleQuery type
type MockArticleQuery struct {
	mock.Mock
}

type MockArticleQuery_Expecter struct {
	mock *mock.Mock
}

func (_m *MockArticleQuery) EXPECT() *MockArticleQuery_Expecter {
	return &MockArticleQuery_Expecter{mock: &_m.Mock}
}

// FindByName provides a mock function with given fields: name
func (_m *MockArticleQuery) FindByName(name string) ([]domain.Article, int) {
	ret := _m.Called(name)

	var r0 []domain.Article
	var r1 int
	if rf, ok := ret.Get(0).(func(string) ([]domain.Article, int)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) []domain.Article); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Article)
		}
	}

	if rf, ok := ret.Get(1).(func(string) int); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Get(1).(int)
	}

	return r0, r1
}

// MockArticleQuery_FindByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindByName'
type MockArticleQuery_FindByName_Call struct {
	*mock.Call
}

// FindByName is a helper method to define mock.On call
//   - name string
func (_e *MockArticleQuery_Expecter) FindByName(name interface{}) *MockArticleQuery_FindByName_Call {
	return &MockArticleQuery_FindByName_Call{Call: _e.mock.On("FindByName", name)}
}

func (_c *MockArticleQuery_FindByName_Call) Run(run func(name string)) *MockArticleQuery_FindByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockArticleQuery_FindByName_Call) Return(_a0 []domain.Article, _a1 int) *MockArticleQuery_FindByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockArticleQuery_FindByName_Call) RunAndReturn(run func(string) ([]domain.Article, int)) *MockArticleQuery_FindByName_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockArticleQuery creates a new instance of MockArticleQuery. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockArticleQuery(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockArticleQuery {
	mock := &MockArticleQuery{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}