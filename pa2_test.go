package main

import "testing"

func TestSplit(t *testing.T) {
	var b bool
	b = Split(' ') && Split('\n') && Split('\t') && Split('\r') && !Split('z')
	if !b {
		t.Error("Expected true, got", b)
	}
}

func TestPares(t *testing.T) {
	var in []string = []string{
		"use c-scan",
		"lowerCYL 0000",
		"upperCYL 3000",
		"initCYL 27",
		"cylreq 300",
		"cylreq 200",
		"cylreq 100",
		"cylreq 450",
		"cylreq 2500",
		"cylreq 38",
		"end",
	}
	var expected parameters
	expected.alg = "c-scan"
	expected.lowerCYL = 0
	expected.upperCYL = 3000
	expected.initCYL = 27
	expected.requests = []int{300, 200, 100, 450, 2500, 38}

	out, _ := Parse(in)

	exS := ParamsToString(expected)
	outS := ParamsToString(out)

	if exS != outS {
		t.Errorf("Expected %s , got", exS, outS)
	}
}
