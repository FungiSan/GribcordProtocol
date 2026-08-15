package audio

import (
	"encoding/binary"
	"fmt"
)

const HeaderSize = 4

type Packet struct {
	Sequence uint32
	Data     []byte
}

func Encode(packet Packet) []byte {
	buffer := make([]byte, HeaderSize+len(packet.Data))

	binary.LittleEndian.PutUint32(
		buffer[0:HeaderSize],
		packet.Sequence,
	)

	copy(buffer[HeaderSize:], packet.Data)

	return buffer
}

func Decode(buffer []byte) (Packet, error) {
	if len(buffer) < HeaderSize {
		return Packet{}, fmt.Errorf("packet too small: %d bytes", len(buffer))
	}

	packet := Packet{
		Sequence: binary.LittleEndian.Uint32(buffer[0:HeaderSize]),
		Data:     buffer[HeaderSize:],
	}

	return packet, nil
}
