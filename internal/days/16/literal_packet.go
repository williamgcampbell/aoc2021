package _16

type literalPacket struct {
	version int
	typeID  op
	number  int
}

func (l *literalPacket) Version() int {
	return l.version
}

func (l *literalPacket) TypeID() int {
	return int(l.typeID)
}

// Value describes the number value of the literal
func (l *literalPacket) Value() int {
	return l.number
}

func (l *literalPacket) ChildPackets() []Packet {
	return nil
}
