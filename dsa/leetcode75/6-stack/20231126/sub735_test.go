package _0231126

import (
	"testing"
)

// 对于数组中的每一个元素，其绝对值表示小行星的大小，正负表示小行星的移动方向（正表示向右移动，负表示向左移动）。每一颗小行星以相同的速度移动。
//
// 找出碰撞后剩下的所有小行星。碰撞规则：两个小行星相互碰撞，较小的小行星会爆炸。如果两颗小行星大小相同，则两颗小行星都会爆炸。两颗移动方向相同的小行星，永远不会发生碰撞。
// 示例 1：
//
// 输入：asteroids = [5,10,-5]
// 输出：[5,10]
// 解释：10 和 -5 碰撞后只剩下 10 。 5 和 10 永远不会发生碰撞。
// 示例 2：
//
// 输入：asteroids = [8,-8]
// 输出：[]
// 解释：8 和 -8 碰撞后，两者都发生爆炸。
// 示例 3：
//
// 输入：asteroids = [10,2,-5]
// 输出：[10]
// 解释：2 和 -5 发生碰撞后剩下 -5 。10 和 -5 发生碰撞后剩下 10 。
func asteroidCollision(asteroids []int) []int {

	return nil
}

// 思路: 栈的思路
// {
// 1. 遇到栈空 或者 符号相同，入栈
// 2. 遇到符号不同，检测正数栈顶元素绝对值是否大于待入栈数值：
//   - 若绝对值栈顶数 < 检测数，则将栈顶元素出栈，继续和栈顶元素比较，直到栈后入栈
//   - 若绝对值栈顶数 > 检测数，则检测下个元素
//   - 若绝对值栈顶数 = 检测数，则直接出栈
//     }
func AsteroidCollisionMethod01(asteroids []int) []int {
	// 初始化栈
	var stk []int

	// 依次检查行内行星数值&方向
	for _, ast := range asteroids {
	Compare:
		// 栈空或者和栈顶元素符合相同
		if len(stk) == 0 {
			stk = append(stk, ast)
			continue
		}

		// 栈顶行星，方向相同或者是栈顶为负的元素，可以直接入栈（因为栈顶为负表示向左移动，新的行星即便向右也可以直接入栈，两者不会相撞）
		popAst := stk[len(stk)-1]
		if popAst*ast > 0 || popAst < 0 {
			stk = append(stk, ast)
			continue
		}

		// 方向不同，栈顶行星向右，比较那个更大
		switch {
		case AsbInt(popAst) > AsbInt(ast): // 想入栈行星数值太小，不用管
			continue
		case AsbInt(popAst) == AsbInt(ast): // 一样大，直接出栈
			stk = stk[:len(stk)-1]
			continue
		default: // 当前行星比栈顶行星更大，栈顶行星出栈，继续比较栈内元素
			stk = stk[:len(stk)-1]
			goto Compare
		}

	}

	return stk
}

func AsbInt(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func TestAsteroidCollisionMethod01(t *testing.T) {
	tests := []struct {
		name string
		asts []int
		want []int
	}{
		{"t1", []int{5, 10, -5}, []int{5, 10}},
		{"t2", []int{8, -8}, []int{}},
		{"t3", []int{10, 2, -5}, []int{10}},
		{"t4", []int{-2, -1, 1, 2}, []int{-2, -1, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AsteroidCollisionMethod01(tt.asts)
			if len(got) != len(tt.want) {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
			for i := 0; i < len(got); i++ {
				if got[i] != tt.want[i] {
					t.Errorf("got %v, but want %v", got, tt.want)
					break
				}
			}
		})
	}
}
