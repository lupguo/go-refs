package gomock_case

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
)

func TestHelloGoMock(t *testing.T) {
	// go mock控制器
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// 注入mock控制器，创建一个mock实例
	mock := NewMockData(ctrl)

	// 注入mock函数
	mock.
		EXPECT().
		Send("no_value", nil).
		Return(errors.New("empty values")).
		AnyTimes()

	mock.
		EXPECT().
		Send("key", []byte("Hello")).
		Return(nil).AnyTimes()

	mock.EXPECT().
		Get("no_value").
		Return(nil, errors.New("empty values")).
		AnyTimes()

	mock.EXPECT().
		Get(gomock.Any()).
		Return([]byte("Hello"), nil).
		AnyTimes()

	// 执行函数
	// inst := NewAnalysisData()
	// assert.NotNil(t, inst.Send("no_value", nil)) // 断言返回应该有错
	// assert.Nil(t, inst.Send("key", nil))         // 断言返回结果正确
	//
	// get, err := inst.Get("key")
	// assert.Equal(t, err, nil)             // 断言结果1
	// assert.Equal(t, get, []byte("Hello")) // 断言结果2

	// 执行
	HandleMsg(mock)
}

func TestSUT(t *testing.T) {
	// go mock控制器
	ctrl := gomock.NewController(t)

	// 注入mock控制器，创建一个mock实例
	mock := NewMockData(ctrl)
	mock.EXPECT().Send("no_value", nil).Return(errors.New("empty error")).MinTimes(3)
	mock.EXPECT().Send("key", []byte("Hello")).Return(errors.New("bad bad")).AnyTimes()
	mock.EXPECT().Get(gomock.Any()).Return([]byte("Hi,man"), nil).AnyTimes()
	mock.EXPECT().Send("no_value", "exception values").Return(errors.New("empty error")).AnyTimes()

	// 待测试的函数内容
	HandleMsg(mock)
	HandleMsg(mock)
	HandleMsg(mock)
	HandleMsg(mock)
	HandleMsg(mock)
}

func TestSUT2(t *testing.T) {
	// setup stub
	testObj := new(TestDD)
	testObj.On("Send", "no_value").Return(errors.New("On MethodName, empty error"))

	HandleMsg(testObj)

	testObj.AssertExpectations(t)
}
