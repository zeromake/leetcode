package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRestoreIpAddresses(t *testing.T) {
	addrs := []string{
		"25525511135",
		"19216810100",
		"",
	}
	result := [][]string{
		{
			"255.255.11.135",
			"255.255.111.35",
		},
		{
			"192.168.10.100",
		},
		nil,
	}
	for i, r := range addrs {
		rr := RestoreIpAddresses(r)
		assert.Equal(t, rr, result[i])
	}
}
