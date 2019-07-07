package parser

import (
	"github.com/mattjones753/cron-parser/pkg/cron"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParser(t *testing.T) {
	t.Run("parses good input string successfully", func(t *testing.T) {
		inputString := "*/15 0 1,15 * 1-5 /usr/bin/find"
		expectedParsed := &cron.Cron{
			Minutes:     []int{0, 15, 30, 45},
			Hours:       []int{0},
			DaysOfMonth: []int{1, 15},
			Months:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			DaysOfWeek:  []int{1, 2, 3, 4, 5},
			Command:     "/usr/bin/find",
		}

		actual, _ := Parse(inputString)
		assert.Equal(t, expectedParsed, actual)
	})
	t.Run("input string with too few components returns an error", func(t *testing.T) {
		inputString := "0 0"
		_, err := Parse(inputString)
		assert.Error(t, err)
	})
}
