package formatter

import (
	"github.com/mattjones753/cron-parser/pkg/cron"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*

 */

func TestFormatter(t *testing.T) {
	t.Run("formats cron correctly", func(t *testing.T) {
		inputCron := &cron.Cron{
			Minutes:     []int{0, 15, 30, 45},
			Hours:       []int{0},
			DaysOfMonth: []int{1, 15},
			Months:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			DaysOfWeek:  []int{1, 2, 3, 4, 5},
			Command:     "/usr/bin/find",
		}

		expectedOutputString :=
			`minute        0 15 30 45
hour          0
day of month  1 15
month         1 2 3 4 5 6 7 8 9 10 11 12
day of week   1 2 3 4 5
command       /usr/bin/find`

		actual := Format(inputCron)
		assert.Equal(t, expectedOutputString, actual)
	})
}
