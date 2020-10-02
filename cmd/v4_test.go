package cmd

import "testing"

type mockUUID struct {
}

func (m mockUUID) V4() string {
	return "0000000-0000-0000-0000-000000000000"
}

func TestGenerateUUIDv4(t *testing.T) {

	t.Log("When argument contains non integer field")
	{
		arg := []string{"abc"}
		m := mockUUID{}
		t.Run("Returns non integer error", func(t *testing.T) {

			_, err := generateUUIDV4(arg, m)
			if err == nil {
				t.Errorf("expected error but didn't receive one")
			}
			expErrMsg := "Argument must a integer, please enter number between 1 to 100"
			if err.Error() != expErrMsg {
				t.Errorf("Expected %s got %s", expErrMsg, err.Error())
			}
		})
	}

	t.Log("When argument is not between 1 to 100")
	{
		arg := []string{"101"}
		m := mockUUID{}
		t.Run("Returns range error", func(t *testing.T) {

			_, err := generateUUIDV4(arg, m)
			if err == nil {
				t.Errorf("expected error but didn't receive one")
			}
			expErrMsg := "Please enter number between 1 to 100"
			if err.Error() != expErrMsg {
				t.Errorf("Expected %s got %s", expErrMsg, err.Error())
			}
		})
	}

	t.Log("When argument is not given")
	{
		arg := []string{}
		m := mockUUID{}
		t.Run("Returns one uuid", func(t *testing.T) {

			uuids, err := generateUUIDV4(arg, m)
			if err != nil {
				t.Errorf("unexpected error occured %+v", err)
			}
			if len(uuids) != 1 {
				t.Errorf("Expected 1 uuid got %d", len(uuids))
			}
		})
	}

	t.Log("When argument is between 1 to 100")
	{
		arg := []string{"50"}
		m := mockUUID{}
		t.Run("Returns one uuid", func(t *testing.T) {

			uuids, err := generateUUIDV4(arg, m)
			if err != nil {
				t.Errorf("unexpected error occured %+v", err)
			}
			if len(uuids) != 50 {
				t.Errorf("Expected 50 uuids got %d", len(uuids))
			}
		})
	}
}
