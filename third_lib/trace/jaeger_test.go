package trace

import (
	_ "context"
	"fmt"
	"testing"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

// 在createOrder函数中，我们创建了一个名为"CreateOrder_DBWrite"的子span，表示订单创建DB写入操作。类似地，在reduceInventoryRPC函数和generateOrderCache函数中，我们也创建了相应的子span来表示库存扣减RPC调用和订单OID详情缓存生成操作。
//
// 通过这种方式，我们可以在命令字的关键子步骤中创建子span，并使用Jaeger追踪器记录和追踪这些子步骤。这样，我们就可以在Jaeger的用户界面中查看和分析这些子步骤的执行情况，以及它们之间的关系和耗时。
func TestJaeger(t *testing.T) {
	// 初始化Jaeger配置
	cfg := config.Configuration{
		ServiceName: "order-service",
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
		},
	}

	// 初始化Jaeger追踪器
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		fmt.Printf("Failed to create Jaeger tracer: %v\n", err)
		return
	}
	defer closer.Close()

	// 设置全局追踪器
	opentracing.SetGlobalTracer(tracer)

	// 创建订单命令字
	createOrderCmd := "CreateOrder"

	// 创建订单
	createOrder(createOrderCmd)
}

func createOrder(cmd string) {
	// 创建子span，表示订单创建DB写入操作
	span := opentracing.StartSpan("CreateOrder_DBWrite")
	defer span.Finish()

	// 模拟订单创建DB写入操作
	// ...
	time.Sleep(50 * time.Millisecond)

	// 调用库存扣减RPC
	reduceInventoryRPC(cmd)
}

func reduceInventoryRPC(cmd string) {
	// 创建子span，表示库存扣减RPC调用操作
	span := opentracing.StartSpan("CreateOrder_ReduceInventoryRPC")
	defer span.Finish()

	// 模拟库存扣减RPC调用
	// ...
	time.Sleep(150 * time.Millisecond)

	// 生成订单OID详情缓存
	generateOrderCache(cmd)
}

func generateOrderCache(cmd string) {
	// 创建子span，表示订单OID详情缓存生成操作
	span := opentracing.StartSpan("CreateOrder_GenerateOrderCache")
	defer span.Finish()

	time.Sleep(15 * time.Millisecond)
	// 模拟订单OID详情缓存生成
	// ...
}
