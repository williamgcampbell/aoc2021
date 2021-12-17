package _16

import (
	"strings"

	"github.com/williamgcampbell/aoc2021/internal/helpers"
)

type Header interface {
	Version() int
	TypeID() int
}

// Valuer is an interface describing anything with a value
type Valuer interface {
	Value() int
}

type Packet interface {
	Header
	Valuer
	ChildPackets() []Packet
}

type PacketParser struct {
	pos  int
	bits string
}

func newPacketParser(bits string) *PacketParser {
	return &PacketParser{
		pos:  0,
		bits: bits,
	}
}

func ParseSize(parser *PacketParser, maxSize int) []Packet {
	var vrs []Packet
	for parser.pos < maxSize {
		vrs = append(vrs, parse(parser))
	}
	return vrs
}

func ParseCount(parser *PacketParser, maxCount int) []Packet {
	var vrs []Packet
	for len(vrs) < maxCount {
		vrs = append(vrs, parse(parser))
	}
	return vrs
}

func parse(parser *PacketParser) Packet {
	if isLiteralValue(parser.bits, parser.pos) {
		return parser.parseLiteralPacket()
	}
	return parser.parseOperatorPacket()

}

func isLiteralValue(bits string, pos int) bool {
	return helpers.MustParseUint(bits[pos+3:pos+6], 2, 64) == 4
}

type header struct {
	version int
	typeID  int
}

func (p *PacketParser) parseOperatorPacket() *operatorPacket {
	h := p.parseHeader()

	lengthTypeID := int(helpers.MustParseInt(p.bits[p.pos:p.pos+1], 2, 2))
	p.pos += 1

	var packets []Packet
	switch lengthTypeID {
	case 0:
		maxSize := int(helpers.MustParseInt(p.bits[p.pos:p.pos+15], 2, 16))
		p.pos += 15

		nextBits := p.bits[p.pos:]
		np := newPacketParser(nextBits)
		packets = ParseSize(np, maxSize)
		p.pos += np.pos
	case 1:
		maxCount := int(helpers.MustParseInt(p.bits[p.pos:p.pos+11], 2, 16))
		p.pos += 11

		nextBits := p.bits[p.pos:]
		np := newPacketParser(nextBits)
		packets = ParseCount(np, maxCount)
		p.pos += np.pos
	}

	return &operatorPacket{
		version: h.version,
		typeID:  op(h.typeID),
		packets: packets,
	}
}

func (p *PacketParser) parseLiteralPacket() *literalPacket {
	h := p.parseHeader()
	number := p.parseNumber()

	return &literalPacket{
		version: h.version,
		typeID:  op(h.typeID),
		number:  number,
	}
}

func (p *PacketParser) parseHeader() header {
	version := helpers.MustParseInt(p.bits[p.pos:p.pos+3], 2, 64)
	typeID := helpers.MustParseInt(p.bits[p.pos+3:p.pos+6], 2, 64)

	// move the position
	p.pos += 6
	return header{
		version: int(version),
		typeID:  int(typeID),
	}
}

// parseNumber does the work to parse a number from an undetermined number of 5 sequences.
// The first bit in the sequence determines whether this is the last (0) sequence or whether to keep going (1)
// Each sequence of bits is appended to the previous. The final bit string is a number
func (p *PacketParser) parseNumber() int {
	var sb strings.Builder

	for {
		b := p.bits[p.pos]
		sb.WriteString(p.bits[p.pos+1 : p.pos+5])

		// move the position
		p.pos += 5

		// '0' bit indicates that this is the last sequence
		if b == '0' {
			break
		}
	}

	return int(helpers.MustParseInt(sb.String(), 2, 64))
}
