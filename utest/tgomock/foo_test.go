package tgomock

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"x-learn/utest/mock"
)

func TestGetConfig(t *testing.T) {
	m := mock.NewMockIApolloConfig(gomock.NewController(t))

	// mock 任意参数 + 条件返回
	m.EXPECT().GetConfig(gomock.Any()).DoAndReturn(func(k string) bool {
		return k == "key"
	}).AnyTimes()

	// mock 不同参数 + 返回
	assert.Equal(t, m.GetConfig("key"), true)
	assert.Equal(t, m.GetConfig("no-key"), false)
}

func TestOrderHandlerAssert(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := mock.NewMockIApolloConfig(ctrl)

	// mock 相同参数 + 不同返回
	gomock.InOrder(
		m.EXPECT().GetConfig("key").Return(true),
		m.EXPECT().GetConfig("key").Return(false),
	)

	assert.Equal(t, OrderHandle("x", m), "ok,x")
	assert.Equal(t, OrderHandle("y", m), "nok,y")
}

func TestOrderHandlerAssertDoAndReturn(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := mock.NewMockIApolloConfig(ctrl)

	// mock 任意参数 + 条件返回
	m.EXPECT().GetConfig(gomock.Any()).DoAndReturn(func(k string) bool {
		return k == "key"
	}).AnyTimes()

	assert.Equal(t, OrderHandle("x", m), "ok,x")
	assert.Equal(t, OrderHandle("y", m), "ok,y")
}

func TestOrderHandlerTableInFor(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockApollo := mock.NewMockIApolloConfig(ctrl)

	// tables 测试
	type args struct {
		x                string
		mockApollo       IApolloConfig
		mockGetConfigRet bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"ut01", args{"red", mockApollo, true}, "ok,red"},
		{"ut02", args{"blue", mockApollo, false}, "nok,blue"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// mock
			mockApollo.EXPECT().GetConfig("key").Return(tt.args.mockGetConfigRet)

			if got := OrderHandle(tt.args.x, tt.args.mockApollo); got != tt.want {
				t.Errorf("OrderHandle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderHandlerTableOutFor(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockApollo := mock.NewMockIApolloConfig(ctrl)

	// signature of anonymous function must have the same number of input and output arguments as the mocked method
	mockApollo.EXPECT().GetConfig(gomock.Any()).DoAndReturn(func(k string) bool {
		if k == "key" {
			return true
		}
		return false
	}).AnyTimes()

	// tables 测试
	type args struct {
		x      string
		IApolo IApolloConfig
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"ut01", args{"red", mockApollo}, "ok,red"},
		{"ut02", args{"blue", mockApollo}, "ok,blue"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OrderHandle(tt.args.x, tt.args.IApolo); got != tt.want {
				t.Errorf("OrderHandle() = %v, want %v", got, tt.want)
			}
		})
	}
}