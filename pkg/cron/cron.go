package cron

// Cron represents an execution schedule for a cron job.
type Cron struct {
	// Hours are the minutes in an hour (0-59) that a job will run
	Minutes     []int
	// Hours are the hours of the day (0-23) that a job will run
	Hours       []int
	// DaysOfMonth are the days of the month (1-31) that a job will run
	DaysOfMonth []int
	// Months are numerical representation of the months (1-12; Jan-Dec) that a job will run
	Months      []int
	// Months are numerical representation of the months (1-7; Monday-Sunday) that a job will run
	DaysOfWeek  []int
	// Command is the actual command that will be executed according to the schedule
	Command     string
}
