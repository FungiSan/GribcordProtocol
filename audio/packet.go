package audio

import (
	"encoding/binary"
	"fmt"

	v1 "github.com/Su-pr007/RazeProtocol/proto"
)

const (
	SequenceSize  = 4
	TimestampSize = 8

	HeaderSize = SequenceSize + TimestampSize
)

type Packet v1.Packet

func (p *Packet) Debug() {
	fmt.Printf(
		"[DEBUG] Packet: sequence=%d timestamp=%v\n",
		p.Sequence,
		p.Timestamp,
	)
}

func Encode(packet *Packet) []byte {
	buffer := make([]byte, HeaderSize+len(packet.Data))

	binary.LittleEndian.PutUint32(
		buffer[0:SequenceSize],
		packet.Sequence,
	)

	binary.LittleEndian.PutUint64(
		buffer[SequenceSize:SequenceSize+TimestampSize],
		uint64(packet.Timestamp),
	)

	copy(buffer[HeaderSize:], packet.Data)

	return buffer
}

func Decode(buffer []byte) (*Packet, error) {
	if len(buffer) < HeaderSize {
		return &Packet{}, fmt.Errorf("packet too small: %d bytes", len(buffer))
	}

	packet := &Packet{
		Sequence:  binary.LittleEndian.Uint32(buffer[0:SequenceSize]),
		Timestamp: int64(binary.LittleEndian.Uint64(buffer[SequenceSize : SequenceSize+TimestampSize])),
		Data:      make([]byte, len(buffer)-HeaderSize),
	}

	copy(packet.Data, buffer[HeaderSize:])

	return packet, nil
}
