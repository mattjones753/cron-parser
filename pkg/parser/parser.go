package parser

import (
	"errors"
	"github.com/mattjones753/cron-parser/pkg/cron"
	"regexp"
	"strconv"
	"strings"
)

func Parse(input string) (*cron.Cron, error) {
	components := strings.Split(input, " ")
	if len(components) < 6 {
		return nil, errors.New("invalid cron expression")
	}

	minutes, err := expandComponent(components[0], 0, 59)
	if err != nil {
		return nil, err
	}
	hours, err := expandComponent(components[1], 0, 23)
	if err != nil {
		return nil, err
	}
	daysOfMonth, err := expandComponent(components[2], 1, 31)
	if err != nil {
		return nil, err
	}
	months, err := expandComponent(components[3], 1, 12)
	if err != nil {
		return nil, err
	}
	daysOfWeek, err := expandComponent(components[4], 1, 7)
	if err != nil {
		return nil, err
	}
	return &cron.Cron{
		Minutes:     minutes,
		Hours:       hours,
		DaysOfMonth: daysOfMonth,
		Months:      months,
		DaysOfWeek:  daysOfWeek,
		Command:     components[5],
	}, nil
}

func expandComponent(expression string, minPossibleValue, maxPossibleValue int) ([]int, error) {
	var values []int
	for i := minPossibleValue; i <= maxPossibleValue; i++ {
		res, err := valueMatchesExpression(i, expression)
		if err != nil {
			return nil, err
		}
		if res {
			values = append(values, i)
		}
	}
	return values, nil
}

func valueMatchesExpression(value int, expression string) (bool, error) {
	subExpressions := strings.Split(expression, ",")
	for _, s := range subExpressions {
		rangeExpression := regexp.MustCompile(`([0-9]+)-([0-9]+)`)
		stepExpression := regexp.MustCompile(`\*/([0-9]+)`)
		if s == "*" {
			return true, nil
		} else if regexp.MustCompile(`^[0-9]+$`).MatchString(s) {
			if s == strconv.Itoa(value) {
				return true, nil
			}
		} else if rangeExpression.MatchString(s) {
			rangeValues := rangeExpression.FindStringSubmatch(s)
			minValue, _ := strconv.Atoi(rangeValues[1])
			maxValue, _ := strconv.Atoi(rangeValues[2])
			if value >= minValue && value <= maxValue {
				return true, nil
			}
		} else if stepExpression.MatchString(s) {
			stepValue, _ := strconv.Atoi(stepExpression.FindStringSubmatch(s)[1])
			if value%stepValue == 0 {
				return true, nil
			}
		} else {
			return false, errors.New("invalid expression")
		}
	}
	return false, nil
}
