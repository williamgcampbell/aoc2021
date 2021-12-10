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
	_11 "github.com/williamgcampbell/aoc2021/internal/days/11"
	_12 "github.com/williamgcampbell/aoc2021/internal/days/12"
	_13 "github.com/williamgcampbell/aoc2021/internal/days/13"
	_14 "github.com/williamgcampbell/aoc2021/internal/days/14"
	_15 "github.com/williamgcampbell/aoc2021/internal/days/15"
	_16 "github.com/williamgcampbell/aoc2021/internal/days/16"
	_17 "github.com/williamgcampbell/aoc2021/internal/days/17"
	_18 "github.com/williamgcampbell/aoc2021/internal/days/18"
	_19 "github.com/williamgcampbell/aoc2021/internal/days/19"
	_20 "github.com/williamgcampbell/aoc2021/internal/days/20"
	_21 "github.com/williamgcampbell/aoc2021/internal/days/21"
	_22 "github.com/williamgcampbell/aoc2021/internal/days/22"
	_23 "github.com/williamgcampbell/aoc2021/internal/days/23"
	_24 "github.com/williamgcampbell/aoc2021/internal/days/24"
	_25 "github.com/williamgcampbell/aoc2021/internal/days/25"
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
		&_11.Day{},
		&_12.Day{},
		&_13.Day{},
		&_14.Day{},
		&_15.Day{},
		&_16.Day{},
		&_17.Day{},
		&_18.Day{},
		&_19.Day{},
		&_20.Day{},
		&_21.Day{},
		&_22.Day{},
		&_23.Day{},
		&_24.Day{},
		&_25.Day{},
	)
	if err != nil {
		return
	}
}
