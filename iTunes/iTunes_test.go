package itunes

import "testing"

func TestSiTunes2Json(t *testing.T) {
	result, err := SiTunes2Json("Google", "US")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestSiTunes(t *testing.T) {
	result, err := SiTunes("Google", "US")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}
