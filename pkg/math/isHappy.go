package math

func getNext(num int) int {
	total := 0
	for num > 0 {
		digit := num % 10
		num /= 10
		total += digit * digit
	}
	return total
}

func IsHappy(num int) bool {
	slow, fast := num, getNext(num)
	for fast != 1 && slow != fast {
		slow = getNext(slow)
		fast = getNext(getNext(fast))
	}
	return fast == 1
}
