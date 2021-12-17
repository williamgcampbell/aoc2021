package _16

import "math"

type operatorPacket struct {
	version int
	typeID  op
	packets []Packet
}

func (o *operatorPacket) Version() int {
	return o.version
}

func (o *operatorPacket) TypeID() int {
	return int(o.typeID)
}

func (o *operatorPacket) ChildPackets() []Packet {
	return o.packets
}

// Value describes the value of an operator on it's sub-packets
func (o *operatorPacket) Value() int {
	switch o.typeID {
	case Sum:
		s := 0
		for _, v := range o.packets {
			s += v.Value()
		}
		return s

	case Product:
		s := 1
		for _, v := range o.packets {
			s *= v.Value()
		}

		return s

	case Minimum:
		s := math.MaxInt
		for _, v := range o.packets {
			if v.Value() < s {
				s = v.Value()
			}
		}
		return s

	case Maximum:
		s := 0
		for _, v := range o.packets {
			if v.Value() > s {
				s = v.Value()
			}
		}
		return s

	case Greater:
		if len(o.packets) != 2 {
			panic("invalid number of subpackets for greater than")
		}
		if o.packets[0].Value() > o.packets[1].Value() {
			return 1
		}
		return 0

	case Less:
		if len(o.packets) != 2 {
			panic("invalid number of subpackets for greater than")
		}
		if o.packets[0].Value() < o.packets[1].Value() {
			return 1
		}
		return 0

	case Equal:
		if len(o.packets) != 2 {
			panic("invalid number of subpackets for greater than")
		}
		if o.packets[0].Value() == o.packets[1].Value() {
			return 1
		}
		return 0

	}
	panic("bad operator")
}
