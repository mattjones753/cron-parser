package main

import (
	"fmt"
	"github.com/mattjones753/cron-parser/cmd/cron-parser/app"
	"os"
)

func main() {
	output, err := app.ParseAndFormat(os.Args[1])

	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	fmt.Println(output)
}
