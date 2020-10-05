package cmd

import "testing"

func TestCheck(t *testing.T) {

	t.Run("should return true", func(t *testing.T) {
		input := "c923d13a-cee4-4182-af57-438f8b032fb6" // valid UUID
		got := Check(input)
		expect := true
		if got != expect {
			t.Errorf("expected %s to be %t but got %t", input, expect, got)
		}
	})

	t.Run("should return false", func(t *testing.T) {
		input := "c93a-cee4-4182-af57-438f2fb6" // not a valid UUID
		got := Check(input)
		expect := false
		if got != expect {
			t.Errorf("expected %s to be %t but got %t", input, expect, got)
		}
	})

}
