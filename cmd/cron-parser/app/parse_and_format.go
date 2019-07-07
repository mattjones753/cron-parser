package app

import (
	"github.com/mattjones753/cron-parser/pkg/formatter"
	"github.com/mattjones753/cron-parser/pkg/parser"
)

func ParseAndFormat(cronExpression string) (string, error) {
	cron, err := parser.Parse(cronExpression)

	if err != nil {
		return "", err
	}

	return formatter.Format(cron), nil
}
