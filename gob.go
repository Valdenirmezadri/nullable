package nullable

import (
	"encoding/gob"
	"fmt"
)

func init() {
	gob.Register(Uint8{})
	gob.Register(Uint16{})
	gob.Register(Uint32{})
	gob.Register(Uint{})
	gob.Register(String{})
}

// Uint8
func (u Uint8) GobEncode() ([]byte, error) {
	if !u.isValid {
		return []byte{0}, nil
	}
	return []byte{1, u.realValue}, nil
}

func (u *Uint8) GobDecode(data []byte) error {
	if len(data) == 0 {
		return fmt.Errorf("invalid gob data for Uint8")
	}
	u.isValid = data[0] == 1
	if u.isValid && len(data) > 1 {
		u.realValue = data[1]
	}
	return nil
}

// Uint16
func (u Uint16) GobEncode() ([]byte, error) {
	if !u.isValid {
		return []byte{0}, nil
	}
	return []byte{1, byte(u.realValue), byte(u.realValue >> 8)}, nil
}

func (u *Uint16) GobDecode(data []byte) error {
	if len(data) < 1 {
		return fmt.Errorf("invalid gob data for Uint16")
	}
	u.isValid = data[0] == 1
	if u.isValid && len(data) >= 3 {
		u.realValue = uint16(data[1]) | uint16(data[2])<<8
	}
	return nil
}

// Uint32
func (u Uint32) GobEncode() ([]byte, error) {
	if !u.isValid {
		return []byte{0}, nil
	}
	buf := make([]byte, 5)
	buf[0] = 1
	for i := 0; i < 4; i++ {
		buf[i+1] = byte(u.realValue >> (8 * i))
	}
	return buf, nil
}

func (u *Uint32) GobDecode(data []byte) error {
	if len(data) < 1 {
		return fmt.Errorf("invalid gob data for Uint32")
	}
	u.isValid = data[0] == 1
	if u.isValid && len(data) >= 5 {
		u.realValue = 0
		for i := 0; i < 4; i++ {
			u.realValue |= uint32(data[i+1]) << (8 * i)
		}
	}
	return nil
}

// Uint
func (u Uint) GobEncode() ([]byte, error) {
	if !u.isValid {
		return []byte{0}, nil
	}
	buf := make([]byte, 9)
	buf[0] = 1
	for i := 0; i < 8; i++ {
		buf[i+1] = byte(u.realValue >> (8 * i))
	}
	return buf, nil
}

func (u *Uint) GobDecode(data []byte) error {
	if len(data) < 1 {
		return fmt.Errorf("invalid gob data for Uint")
	}
	u.isValid = data[0] == 1
	if u.isValid && len(data) >= 9 {
		u.realValue = 0
		for i := 0; i < 8; i++ {
			u.realValue |= uint(data[i+1]) << (8 * i)
		}
	}
	return nil
}

// String
func (n String) GobEncode() ([]byte, error) {
	if !n.isValid {
		return []byte{0}, nil
	}
	strBytes := []byte(n.realValue)
	return append([]byte{1}, strBytes...), nil
}

func (n *String) GobDecode(data []byte) error {
	if len(data) < 1 {
		return fmt.Errorf("invalid gob data for String")
	}
	n.isValid = data[0] == 1
	if n.isValid {
		n.realValue = string(data[1:])
	}
	return nil
}
