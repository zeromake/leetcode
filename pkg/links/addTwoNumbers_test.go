package links

import (
	"testing"
)

func numToList(num int) *ListNode {
	var (
		mod int
		sum int = num
		pre     = &ListNode{
			Val: 0,
		}
	)
	if num == 0 {
		return pre
	}
	cur := pre
	for sum > 0 {
		mod = sum % 10
		sum = sum / 10
		cur.Next = &ListNode{
			Val: mod,
		}
		cur = cur.Next
	}
	// if mod > 0 {
	// 	cur.Next = &ListNode{
	// 		Val: mod,
	// 	}
	// }
	return pre.Next
}

func TestAddTwoNumbers(t *testing.T) {
	testRes := [][3]*ListNode{
		[3]*ListNode{
			numToList(742),
			numToList(465),
			numToList(1207),
		},
		[3]*ListNode{
			numToList(81),
			numToList(0),
			numToList(81),
		},
		[3]*ListNode{
			numToList(42),
			numToList(465),
			numToList(507),
		},
	}
	for _, r := range testRes {
		rr := AddTwoNumbers(r[0], r[1])
		if rr.String() != r[2].String() {
			t.Fatal("Error", rr.String(), r[2].String())
		}
	}
}
