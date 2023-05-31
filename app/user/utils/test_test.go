package utils

import (
	"encoding/binary"
	"testing"
)

func TestStringToInt64(t *testing.T) {
	s1 := "user"
	b := []byte(s1)
	u := binary.BigEndian.Uint32(b)
	t.Log(u)
}
