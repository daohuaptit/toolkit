package toolkit

import (
	"os"
	"testing" // package mặc định của golang
)

func TestTools_RandomString(t *testing.T) { // tên hàm phải bắt đầu bằng tiền tố Test: vd TestCalculateSum,
	var testTools Tools

	s := testTools.RandomString(10)
	if len(s) != 10 {
		t.Error("wrong length random string returned") //có thể sử dụng  các phương thức như t.Error(), t.Errorf(), t.Fatal(), t.Fatalf(). t.Log(), t.Logf() v.v tùy yêu cầu sử dụng.
	}
}

func TestCreateDir(t *testing.T) {
	var testTools Tools
	err := testTools.CreateDirIfNotExist("./test-dir/Myapp")
	if err != nil {
		t.Error(err)
	}
	_ = os.Remove("./test-dir/Myapp")
}
