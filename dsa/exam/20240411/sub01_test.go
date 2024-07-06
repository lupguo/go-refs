package _0240411

import (
	"testing"
)

// 请用尽可能少的代码实现一个函数，用于计算用户一个月共计交费多少港元。（代码请写的尽量清晰简洁，我们希望能够看到你的编码风格和习惯）
// 用户在富途的平台上进行交易，需要交平台使用费。
// 平台使用费的梯度收费方案如下：
// 每月累计订单数        每笔订单（港元）
// 梯度1：1-5笔 => 30.00
// 梯度2：6-20笔 => 15.00
// 梯度3：21-50笔 => 10.00
// 梯度4：51-100笔 => 9.00
// 梯度5：101-500笔 => 8.00
// 梯度6：501-1000笔 => 7.00
// 梯度7：1001-2000笔 => 6.00
// 梯度8：2001-3000笔 => 5.00
// 梯度9：3001-4000笔 => 4.00
// 梯度10：4001-5000笔 => 3.00
// 梯度11：5001-6000笔 => 2.00
// 梯度12：6001笔及以上 => 1.00
// 假设一个用户，一个月交易了6笔订单，则在梯度1交费共计： 30港元*5=150港元，在梯度二交费：15港元，一共交费165港元。
// AreaPrice 区间价格
type AreaPrice struct {
	MinNum int
	MaxNum int
	Price  float64
}

// GetPlatformFee 获取指定平台费用
func GetPlatformFee(num int, areaPrices []*AreaPrice) float64 {
	// 总价格 = 各个区间价格之和
	var totalPrice float64
	var usedNum int // 使用了多少单
	for _, ap := range areaPrices {

		// 是否全部算完
		if usedNum == num {
			break
		}

		// num超过了这个区间，则用这个区间的价格*区间的数量
		if num > ap.MaxNum {
			totalPrice += ap.Price * float64(ap.MaxNum)
			usedNum += ap.MaxNum
		} else {
			totalPrice += ap.Price * float64(num-usedNum)
			usedNum = num
		}
	}

	return totalPrice
}

func TestGetPlatformFee(t *testing.T) {
	tests := []struct {
		name string
		num  int
		aps  []*AreaPrice
		want float64
	}{
		{"t1", 6, []*AreaPrice{
			{1, 5, 30},
			{6, 20, 15},
		}, 165},
		{"t2", 4, []*AreaPrice{
			{1, 5, 30},
			{6, 20, 15},
		}, 4 * 30},
		{"t3", 5, []*AreaPrice{
			{1, 5, 30},
			{6, 20, 15},
			{21, 50, 10},
		}, 5 * 30},
		{"t4", 20, []*AreaPrice{
			{1, 5, 30},
			{6, 20, 15},
			{21, 50, 10},
		}, 5*30 + 15*15},
		{"t5", 26, []*AreaPrice{
			{1, 5, 30},
			{6, 20, 15},
			{21, 50, 10},
		}, 5*30 + 20*15 + 1*10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetPlatformFee(tt.num, tt.aps)
			if got != tt.want {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}
