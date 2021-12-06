package registry

type Day interface {
	GetName() int
	SolvePart1() string
	SolvePart2() string
}

type DayRegistry map[int]Day

func NewDayRegistry() DayRegistry {
	return make(DayRegistry)
}

func (r DayRegistry) Register(days ...Day) error {
	for _, day := range days {
		r[day.GetName()] = day
	}
	return nil
}
