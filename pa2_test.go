package main

import "testing"

var fcfs1 parameters = parameters{alg: "fcfs", lowerCYL: 0, upperCYL: 3000, initCYL: 27, requests: []int{300, 200, 100, 450, 2500, 38}}

func TestFcfs(t *testing.T) {
	var out string = Fcfs(fcfs1)
	var exp string = "Servicing   300\nServicing   200\nServicing   100\nServicing   450\nServicing  2500\nServicing    38\nFCFS traversal count =  5335\n"

	if out != exp {
		t.Errorf("Expected %s, got %s", exp, out)
	}
}

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
		t.Errorf("Expected %s , got\n", exS, outS)
	}
}

func TestRun(t *testing.T) {
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

	perams, _ := Parse(in)
	exp := ""
	out := Run(perams)

	if exp != out {
		t.Errorf("Expected %s , got\n", exp, out)
	}

}
