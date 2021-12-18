package _17

import (
	_ "embed"
	"regexp"
	"strconv"
	"strings"

	"github.com/williamgcampbell/aoc2021/internal/helpers"
	"github.com/williamgcampbell/aoc2021/internal/scanner"
)

//go:embed input.txt
var input string

var target, _ = regexp.Compile("target area: x=(?P<lx>[\\-0-9]+)..(?P<ux>[\\-0-9]+), y=(?P<ly>[\\-0-9]+)..(?P<uy>[\\-0-9]+)")

// SolvePart1 solves part one
func (d *Day) SolvePart1() string {
	r := strings.NewReader(input)
	v := scanner.ScanLines(r)
	highest, _ := launchProbe(v[0])
	return strconv.Itoa(highest)
}

func launchProbe(s string) (int, int) {
	area := getArea(s)
	mh, c := 0, 0
	// no need to have a starting x velocity above the right side of the box
	for dx := 1; dx <= area.right; dx++ {
		// the y velocity can start from the bottom of the area, and the maximum starting velocity can be abs of that bottom
		for dy := area.bottom; dy < helpers.Abs(area.bottom); dy++ {
			p := pos{
				x: 0,
				y: 0,
			}

			pr := &probe{
				p:  p,
				dx: dx,
				dy: dy,
			}

			h := 0
			for area.above(pr.p) && area.xLeft(pr.p) {
				pr.move()
				h = helpers.MaxInt(h, pr.p.y)
				if area.inArea(pr.p) {
					mh = helpers.MaxInt(h, mh)
					c += 1
					break
				}
			}
		}
	}
	return mh, c
}

func getArea(s string) *rect {
	coords := target.FindStringSubmatch(s)
	left := helpers.MustAtoI(coords[1])
	right := helpers.MustAtoI(coords[2])
	bottom := helpers.MustAtoI(coords[3])
	top := helpers.MustAtoI(coords[4])

	return &rect{
		left:   left,
		right:  right,
		bottom: bottom,
		top:    top,
	}
}

type probe struct {
	p  pos
	dx int
	dy int
}

func (p *probe) move() {
	p.p.x += p.dx
	p.p.y += p.dy
	if p.dx > 0 {
		p.dx -= 1
	} else if p.p.x < 0 {
		p.dx += 1
	}
	p.dy -= 1
}

type rect struct {
	left   int
	right  int
	bottom int
	top    int
}

func (r *rect) inArea(p pos) bool {
	return p.x >= r.left &&
		p.x <= r.right &&
		p.y >= r.bottom &&
		p.y <= r.top
}

func (r *rect) xLeft(p pos) bool {
	return p.x < r.right
}

func (r *rect) above(p pos) bool {
	return p.y > r.bottom
}

type pos struct {
	x int
	y int
}
