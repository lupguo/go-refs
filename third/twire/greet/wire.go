package greet

import (
	"github.com/google/wire"
)

var MessageSet = wire.NewSet(
	NewMessage, // 在 Wire 中，初始化器被称为“提供者”
	NewEvent,
)

var humanSet = wire.NewSet(
	NewHumanGreeter,
	wire.Bind(new(Speaker), new(HumanGreeter)),
)

var robotSet = wire.NewSet(
	NewRobotGreeter,
	wire.Bind(new(Speaker), new(*RobotGreeter)),
)

// InitializeEvent 就是一个“注入器”，现在我们已经完成了注入器，我们可以使用wire命令行工具了。
// InitializeEvent 注入器的目的是提供有关使用哪些提供程序来构建一个的信息Event
// 因此我们将在文件顶部使用构建约束将其从最终二进制文件中排除
func InitializeEventRobot(robotID string) Event {
	// 不必经历依次初始化每个组件并将其传递给下一个组件的麻烦，而是通过一次调用来wire.Build 传递我们想要使用的初始化器
	wire.Build(
		NewRobotGreeter,
		wire.Bind(new(Speaker), new(*RobotGreeter)),
		MessageSet,
	)

	// 我们添加一个零值 Event作为返回值以满足编译器，即使我们向其中添加Event的值，Wire也会忽略它们
	return Event{}
}

func InitializeEventHuman() Event {
	// 不必经历依次初始化每个组件并将其传递给下一个组件的麻烦，而是通过一次调用来wire.Build 传递我们想要使用的初始化器
	wire.Build(
		NewHumanGreeter,
		wire.Bind(new(Speaker), new(HumanGreeter)),
		MessageSet,
	)

	// 我们添加一个零值 Event作为返回值以满足编译器，即使我们向其中添加Event的值，Wire也会忽略它们
	return Event{}
}
