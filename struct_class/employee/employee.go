package employee

import (
	"fmt"
)

type employee struct {
	firstname   string
	lastname    string
	age         int
	salary      int
	leavesTaken int
	departure   departure
}

func New(
	firstname string,
	lastname string,
	age int,
	salary int,
	leavesTaken int,
	departure departure,
) employee {
	e := employee{
		firstname,
		lastname,
		age,
		salary,
		leavesTaken,
		departure,
	}
	return e
}

func (e employee) LeavesRemining() {
	fmt.Printf("%s %s has %d leaves remining\n", e.firstname, e.lastname, (e.departure.totalLeaves - e.leavesTaken))
}

// Pointer
func (e *employee) IncreseLeave() {
	e.leavesTaken++
}
