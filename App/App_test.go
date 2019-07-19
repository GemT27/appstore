package app

import "testing"

func TestSapp2Json(t *testing.T) {
	json,err:=Sapp2Json("微信","CN",10)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(json)
}
