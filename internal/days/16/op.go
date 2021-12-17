package _16

const (
	Sum     = op(0)
	Product = op(1)
	Minimum = op(2)
	Maximum = op(3)
	Literal = op(4)
	Greater = op(5)
	Less    = op(6)
	Equal   = op(7)
)

type op int

func (o op) String() string {
	switch o {
	case Sum:
		return "sum"
	case Product:
		return "prd"
	case Minimum:
		return "min"
	case Maximum:
		return "max"
	case Literal:
		return "lit"
	case Greater:
		return "gre"
	case Less:
		return "les"
	case Equal:
		return "eql"
	}

	return ""
}
