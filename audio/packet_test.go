package audio

import (
	"bytes"
	"testing"
)

func TestPacketEncodeDecode(t *testing.T) {
	original := Packet{
		Sequence: 42,
		Data:     []byte{1, 2, 3, 4, 5},
	}

	encoded := Encode(&original)

	t.Logf("Encoded: % x", encoded)

	decoded, err := Decode(encoded)
	if err != nil {
		t.Fatal(err)
	}

	if decoded.Sequence != original.Sequence {
		t.Fatalf("sequence mismatch: got %d, expected %d", decoded.Sequence, original.Sequence)
	}

	if !bytes.Equal(decoded.Data, original.Data) {
		t.Fatalf("data mismatch: got %v, expected %v", decoded.Data, original.Data)
	}
}
