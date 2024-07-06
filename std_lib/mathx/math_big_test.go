package mathx

import (
	"math/big"
	"testing"
)

func TestBigFloat(t *testing.T) {
	// 创建一个新的big.Float实例
	x := new(big.Float)

	// 设置浮点数值
	x.SetFloat64(0.1)

	// 打印精确的小数表示
	t.Log(x)

	// 进行精确的小数计算
	y := new(big.Float).SetFloat64(0.2)
	z := new(big.Float).Add(x, y)

	p, q := 0.1, 0.2
	r := p + q
	t.Log(r == 0.3)

	// 打印计算结果
	t.Log(z)
	t.Log(z.Sign())
}
