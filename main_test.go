package main

import "testing"

func TestConvHex(t *testing.T) {
	if got, _ := conv("dec2hex", "255"); got != "ff" {
		t.Fatalf("dec2hex 255 期望 ff 实际 %s", got)
	}
	if got, _ := conv("hex2dec", "ff"); got != "255" {
		t.Fatalf("hex2dec ff 期望 255 实际 %s", got)
	}
	if got, _ := conv("b64enc", "ab"); got != "YWI=" {
		t.Fatalf("b64enc ab 期望 YWI= 实际 %s", got)
	}
}

func TestConvBin(t *testing.T) {
	if got, _ := conv("dec2bin", "5"); got != "101" {
		t.Fatalf("dec2bin 5 期望 101 实际 %s", got)
	}
	if got, _ := conv("bin2dec", "101"); got != "5" {
		t.Fatalf("bin2dec 101 期望 5 实际 %s", got)
	}
	if got, _ := conv("urlenc", "a b"); got != "a+b" {
		t.Fatalf("urlenc 期望 a+b 实际 %s", got)
	}
}

func TestConvUnknownMode(t *testing.T) {
	if _, err := conv("nope", "x"); err == nil {
		t.Fatal("未知模式应当报错")
	}
}
