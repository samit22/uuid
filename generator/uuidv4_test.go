package generator

import (
	"regexp"
	"testing"
)

func TestGenerateV4(t *testing.T) {
	var uuidPattern = "[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"
	t.Run("Return uuid v4", func(t *testing.T) {
		g := Generate{}
		uuid, err := g.V4()
		if err != nil {
			t.Fatalf("UUIDv4 generator returned error: %v", err)
		}
		re := regexp.MustCompile(uuidPattern)
		if !re.MatchString(uuid) {
			t.Errorf("invalid uuid genereated")
		}
	})

}
