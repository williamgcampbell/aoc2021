package days

import (
	_1 "github.com/williamgcampbell/aoc2021/internal/days/01"
	_2 "github.com/williamgcampbell/aoc2021/internal/days/02"
	_3 "github.com/williamgcampbell/aoc2021/internal/days/03"
	_4 "github.com/williamgcampbell/aoc2021/internal/days/04"
	_5 "github.com/williamgcampbell/aoc2021/internal/days/05"
	_6 "github.com/williamgcampbell/aoc2021/internal/days/06"
	_7 "github.com/williamgcampbell/aoc2021/internal/days/07"
	_8 "github.com/williamgcampbell/aoc2021/internal/days/08"
	_9 "github.com/williamgcampbell/aoc2021/internal/days/09"
	_10 "github.com/williamgcampbell/aoc2021/internal/days/10"
	"github.com/williamgcampbell/aoc2021/internal/registry"
)

func RegisterAll(registry registry.DayRegistry) {
	err := registry.Register(
		&_1.Day{},
		&_2.Day{},
		&_3.Day{},
		&_4.Day{},
		&_5.Day{},
		&_6.Day{},
		&_7.Day{},
		&_8.Day{},
		&_9.Day{},
		&_10.Day{},
	)
	if err != nil {
		return
	}
}
