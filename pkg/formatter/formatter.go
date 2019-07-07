package formatter

import (
	"fmt"
	"github.com/mattjones753/cron-parser/pkg/cron"
	"strconv"
	"strings"
)

func Format(cron *cron.Cron) string {
	result := []string{
		formatOutputNumberLine("minute", cron.Minutes),
		formatOutputNumberLine("hour", cron.Hours),
		formatOutputNumberLine("day of month", cron.DaysOfMonth),
		formatOutputNumberLine("month", cron.Months),
		formatOutputNumberLine("day of week", cron.DaysOfWeek),
		formatOutputLine("command", cron.Command),
	}

	return strings.Join(result, "\n")
}

func formatOutputNumberLine(title string, value []int) string {
	result := ""
	for i, number := range value {
		val := strconv.Itoa(number)
		if i == 0 {
			result = val
		} else {
			result = result + " " + val
		}
	}
	return formatOutputLine(title, result)
}

func formatOutputLine(title, value string) string {
	return fmt.Sprintf("%-14s%s", title, value)
}
