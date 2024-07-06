package _0231126

import (
	"testing"
)

// Dota2 的世界里有两个阵营：Radiant（天辉）和 Dire（夜魇）
//
// Dota2 参议院由来自两派的参议员组成。现在参议院希望对一个 Dota2 游戏里的改变作出决定。他们以一个基于轮为过程的投票进行。在每一轮中，每一位参议员都可以行使两项权利中的 一 项：
//
// 1. 禁止一名参议员的权利：参议员可以让另一位参议员在这一轮和随后的几轮中丧失 所有的权利 。
// 2. 宣布胜利：如果参议员发现有权利投票的参议员都是 同一个阵营的 ，他可以宣布胜利并决定在游戏中的有关变化。
// 给你一个字符串 senate 代表每个参议员的阵营。字母 'R' 和 'D'分别代表了 Radiant（天辉）和 Dire（夜魇）。然后，如果有 n 个参议员，给定字符串的大小将是 n。
//
// 以轮为基础的过程从给定顺序的第一个参议员开始到最后一个参议员结束。这一过程将持续到投票结束。所有失去权利的参议员将在过程中被跳过。
//
// 假设每一位参议员都足够聪明，会为自己的政党做出最好的策略，你需要预测哪一方最终会宣布胜利并在 Dota2 游戏中决定改变。输出应该是 "Radiant" 或 "Dire" 。
//
// 参考: https://leetcode.cn/problems/dota2-senate/description/?envType=study-plan-v2&envId=leetcode-75

// 思路: 循环队列，贪心算法（有投票权的尽快将非同等阵营的投出去）
//  1. 组织两个循环队列radiantQueue, direQueue，按参议院顺序入队列
//  2. 从队列依次比较两个队列出队队首元素，排位靠后出队，同时使用了权利一行加到队尾（需要增加n，到下一轮）
//  3. 队列非空的为胜利方
func predictPartyVictory(senate string) string {
	var rQue, dQue []int

	// 组织好两只队伍，约定好排位顺序(i)
	for i, c := range senate {
		if c == 'R' {
			rQue = append(rQue, i)
		} else {
			dQue = append(dQue, i)
		}
	}

	// 循环比较两只队伍，决出胜负
	for len(rQue) > 0 && len(dQue) > 0 {

		// R、D队首元素比较
		if rQue[0] < dQue[0] { // R队有优先否决权，dQue[0]出队，同时rQue[0]行使完权利后，放入下轮
			rQue = append(rQue, rQue[0]+len(senate))
		} else {
			dQue = append(dQue, dQue[0]+len(senate))
		}

		rQue = rQue[1:]
		dQue = dQue[1:]
	}

	if len(rQue) > 0 {
		return "Radiant"
	}

	return "Dire"
}

func TestPredictPartyVictory(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{"t1", "RD", "Radiant"},
		{"t2", "RRD", "Radiant"},
		{"t3", "RDD", "Dire"},
		{"t4", "RRDD", "Radiant"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := predictPartyVictory(tt.str); got != tt.want {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}
