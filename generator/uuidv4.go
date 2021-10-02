package generator

import (
	"crypto/rand"
	"fmt"
	"io"
)

//GeneratorV4 provides interface to generate v4 UUIDs
type GeneratorV4 interface {
	V4() (string, error)
}

// Generator provides interface to generate UUIDs
type Generator interface {
	GeneratorV4
}

// Generate struct to call the uuid methods
type Generate struct{}

// V4 creates a new UUID in v4 format
func (g Generate) V4() (string, error) {
	var u [16]byte
	if _, err := io.ReadFull(rand.Reader, u[:]); err != nil {
		return "", fmt.Errorf("failed to generate UUIDv4: %w", err)
	}
	u[6] = u[6]>>4 | 0x40 // version number
	u[8] = u[8] & ^uint8(1<<6)
	u[8] |= 1 << 7
	return fmt.Sprintf("%x-%x-%x-%x-%x", u[:4], u[4:6], u[6:8], u[8:10], u[10:]), nil
}
