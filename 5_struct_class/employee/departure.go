package employee

type departure struct {
	title       string
	totalLeaves int
}

func NewDeparture(
	title string,
	totalLeaves int,
) departure {
	d := departure{
		title,
		totalLeaves,
	}

	return d
}
