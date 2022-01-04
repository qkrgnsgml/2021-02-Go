package mymath

type Date struct {
	Year  int
	Month int
	Day   int
}

func (d *Date) setYear(year int) {
	d.Year = year
}

func (d *Date) setMonth(month int) {
	d.Month = month
}

func (d *Date) setDay(day int) {
	d.Day = day
}
