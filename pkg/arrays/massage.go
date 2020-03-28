package arrays

func Massage(nums []int) int {
	lens := len(nums)
	if lens == 0 {
		return 0
	}
	if lens == 1 {
		return nums[0]
	}
	dp := make([]int, lens)
	dp[0] = nums[0]
	var (
		last = nums[0]
		cur  = nums[1]
	)
	if last > cur {
		dp[1] = last
	} else {
		dp[1] = cur
	}
	for i := 2; i < lens; i++ {
		last = dp[i-1]
		cur = dp[i-2] + nums[i]
		if last > cur {
			dp[i] = last
		} else {
			dp[i] = cur
		}
	}
	return dp[lens-1]
}

func Massage2(nums []int) int {
	lens := len(nums)
	if lens == 0 {
		return 0
	}
	if lens == 1 {
		return nums[0]
	}
	dp := [3]int{}
	dp[0] = nums[0]
	var (
		last = nums[0]
		cur  = nums[1]
	)
	if last > cur {
		dp[1] = last
	} else {
		dp[1] = cur
	}
	for i := 2; i < lens; i++ {
		last = dp[(i-1)%3]
		cur = dp[(i-2)%3] + nums[i]
		if last > cur {
			dp[i%3] = last
		} else {
			dp[i%3] = cur
		}
	}
	return dp[(lens-1)%3]
}

func Massage3(nums []int) int {
	// pre: 隔天的最优值, last: 昨天的最优值
	pre, last := 0, 0
	for i := 0; i < len(nums); i++ {
		// 今天的最优值等于前天(不是特指前天而是前天之前的) + 今天
		cur := pre + nums[i]
		// 把昨天的最优值设置到隔天
		pre = last
		// 如果大于今天重置昨天的最优值
		if cur > last {
			last = cur
		}
	}
	return last
}
