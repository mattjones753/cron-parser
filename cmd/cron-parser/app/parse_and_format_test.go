package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseAndFormatter(t *testing.T) {
	t.Run("formats cron correctly", func(t *testing.T) {
		cronExpression := "0 0 1 1 1 ls"
		expectedOutputString :=
			`minute        0
hour          0
day of month  1
month         1
day of week   1
command       ls`

		actual, err := ParseAndFormat(cronExpression)
		assert.NoError(t, err)
		assert.Equal(t, expectedOutputString, actual)
	})
	t.Run("formats cron correctly", func(t *testing.T) {
		cronExpression := "a b c d e f"

		_, err := ParseAndFormat(cronExpression)
		assert.Error(t, err)
	})
}
