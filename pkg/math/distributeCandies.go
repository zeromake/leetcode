package math

// DistributeCandies 分糖果 https://leetcode-cn.com/problems/distribute-candies-to-people
func DistributeCandies(candies int, numPeople int) []int {
	var (
		num = 0
		people = make([]int, numPeople)
	)
	for candies != 0 {
		num++
		index := (num - 1) % numPeople
		if candies < num {
			num = candies
		}
		candies -= num
		people[index] += num
	}
	return people
}
