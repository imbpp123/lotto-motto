package aggregate

type Period struct {
	min int
	max int
}

type PeriodWeight struct {
	Period
	weight int
}
