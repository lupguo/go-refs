package week17

// 求a,b两个有序数组中位数
// 思路: a, b 有序，长度分别位lenA,lenB，先i, j 找到中位数位置 =
//
//	若奇数[1],[2,3] => 中位数idx = (lenA+lenB)/2 位置(idx=1)
//	若偶数[1,2], [3,4] => 中位数idx1=(lenA+lenB)/2 - 1，以及idx2==(lenA+lenB)/2
//
// 迭代过程注意数组长度迭代完成后，剩余数组游标需要继续迭代到 idx 位置
func findMedianSortedArrays(a, b []int) float64 {
	// 计算中位数位置
	lenA, lenB := len(a), len(b)
	midIdx := (lenA + lenB) / 2 // 先明确中位数较大idx目标

	// step1: i, j 同时移动，这里用了k来判断，循环内要么i移动，要么j移动，总之每次移动k计数+1，直到k超过了中位数位置，跳出循环
	//  这里巧妙到使用了lastMidNum先暂存前一个到curMidNum，依次滚动更新lastMidNum和curMidNum，方便最后步骤3到中位数求值
	var i, j int
	var lastMidNum, curMidNum int
	for k := 0; k <= midIdx; k++ {
		lastMidNum = curMidNum

		// step2: 分别迭代a,b数组，以此向前迭代较小元素，并滚动更新midNum最新值
		// Notice:
		//  1) a[i] <= b[j] 时候 b[j]溢出问题，所以要先判断是否j 超出到范围
		//  2) a[i] <= b[j] 这里是 <= 而不是 < ，因为[0,0],[0,0]用例时候会导致b[j]溢出
		if i < lenA && (j >= lenB || a[i] <= b[j]) {
			curMidNum = a[i]
			i++
		} else {
			curMidNum = b[j]
			j++
		}
	}

	// step3: 中位数是否为偶数
	if (lenA+lenB)%2 == 0 {
		return float64(lastMidNum+curMidNum) / 2
	} else {
		return float64(curMidNum)
	}
}
