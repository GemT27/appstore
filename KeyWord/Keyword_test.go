package keyword

import (
	"testing"
)

func TestSKeyWord2Json(t *testing.T) {
	result, err := Skeyword2Json("Google", "US")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}
func TestSKeyWord(t *testing.T) {
	result, err := SkeyWord("Google", "US")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}
