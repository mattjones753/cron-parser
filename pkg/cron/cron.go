package cron

type Cron struct {
	Minutes     []int
	Hours       []int
	DaysOfMonth []int
	Months      []int
	DaysOfWeek  []int
	Command     string
}
