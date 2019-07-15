# cron parser

[![Go Report Card](https://goreportcard.com/badge/github.com/mattjones753/cron-parser)](https://goreportcard.com/report/github.com/mattjones753/cron-parser)
[![](https://godoc.org/github.com/mattjones753/cron-parser?status.svg)](http://godoc.org/github.com/mattjones753/cron-parser)

Cron parser is a simple command line tool that takes a cron expression and prints the details of the schedule so you 
can see when something will actually run 

## Usage
The `cron-parser` utility takes a single argument; a cron expression as follows:
```bash
cron-parser "*/15 0 1,15 * 1-5 /usr/bin/find"
```
It then prints a summary of the execution schedule of the command as well as the command itself.
For example:
```
minute        0 15 30 45
hour          0
day of month  1 15
month         1 2 3 4 5 6 7 8 9 10 11 12
day of week   1 2 3 4 5
command       /usr/bin/find
```
## Building
To build `cron-parser` use the `make` target `make build`, this will create the `cron-parser` in the `bin` directory.

## Tests
Running `make all` will run all the tests, lint the source build the executable and run a simple test against that
executable.

### Prerequisites
* Golang
* Make
