package links

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	testRes := [][3]*ListNode{
		[3]*ListNode{
			NumToList(742),
			NumToList(465),
			NumToList(1207),
		},
		[3]*ListNode{
			NumToList(81),
			NumToList(0),
			NumToList(81),
		},
		[3]*ListNode{
			NumToList(42),
			NumToList(465),
			NumToList(507),
		},
	}
	for _, r := range testRes {
		rr := AddTwoNumbers(r[0], r[1])
		if rr.String() != r[2].String() {
			t.Fatal("Error", rr.String(), r[2].String())
		}
	}
}

func TestAddTwoNumbersII(t *testing.T) {
	testRes := [][3]*ListNode{
		[3]*ListNode{
			NumToList(742),
			NumToList(465),
			NumToList(118),
		},
		[3]*ListNode{
			NumToList(81),
			NumToList(0),
			NumToList(81),
		},
		[3]*ListNode{
			NumToList(42),
			NumToList(465),
			NumToList(885),
		},
	}
	for _, r := range testRes {
		rr := AddTwoNumbersII(r[0], r[1])
		assert.Equal(t, rr.String(), r[2].String())
	}
}
