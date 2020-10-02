package generator

import (
	"fmt"

	uuidg "github.com/gofrs/uuid"
)

// Generator provides interface to generate uuid
type Generator interface {
	V4() string
}

// Generate struct to call the uuid methods
type Generate struct{}

// V4 create a new uuid in v4 format
func (g Generate) V4() string {
	u, err := uuidg.NewV4()
	if err != nil {
		fmt.Println("Something went wrong while generating uuid")
	}
	return u.String()
}
