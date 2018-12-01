package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var fcfs1 parameters = parameters{alg: "fcfs", lowerCYL: 0, upperCYL: 3000, initCYL: 27, requests: []int{300, 200, 100, 450, 2500, 38}}

var files []string

func init() {

	var root string = "/pa2r3-rls/pa2r3-rls/"
	var err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func TestCalcTrav(t *testing.T) {

	var nums []int = []int{1192,
		1263,
		1347,
		1677,
		2011,
		2159,
		2312,
		2503,
		2781,
		2918,
		2920,
		3087,
		3284,
		3687,
		4107,
		4148,
		4195,
		4706,
		930,
		198,
	}
	res, _ := CalcTrav(nums, 1200)

	if res != 8030 {
		t.Error("Expected 8030, got", res)
	}
}

func TestEqual(t *testing.T){
	var a []int= []int{1,2,3,4,5,6,7,8,9}
	var b []int= []int{1,2,4,5,6,7,8,9}
	var c []int= []int{1,2,3,4,5,11,7,8,9}
	r1 := Equal(a,a)
	
	if !r1{
		t.Errorf("Expected true, got %t", r1)
	}

	r2:= Equal(a,b)
	if r2{
		t.Errorf("Expected false, got %t", r2)
	}

	r3 := Equal(a,c)
	if r3{
		t.Errorf("Expected false, got %t", r3)
	}
}

func TestFCFS(t *testing.T) {
	var out string = FCFS(fcfs1)
	var exp string = "Servicing   300\nServicing   200\nServicing   100\nServicing   450\nServicing  2500\nServicing    38\nFCFS traversal count =  5335\n"

	if out != exp {
		t.Errorf("Expected %s, got %s", exp, out)
	}
}

func TestFindClosest(t *testing.T) {
	var l1 []int = []int{1, 2, 3, 8, 16, 20, 22}
	var t1 int = 17
	var i1 int = 4

	outV, outI := FindClosest(l1, t1)

	if outV != 16 {
		t.Errorf("Expected 16, got %d", outV)
	}

	if outI != i1 {
		t.Errorf("Expected %d, got %d", i1, outI)
	}

	outV, outI = FindClosest(l1, 9)
	if outV != 8 {
		t.Errorf("Expected 8, got %d", outV)
	}

	if outI != 3 {
		t.Errorf("Expected 3, got %d", outI)
	}

	outV, outI = FindClosest(l1, 0)
	if outV != 1 {
		t.Errorf("Expected 1, got %d", outV)
	}

	if outI != 0 {
		t.Errorf("Expected 0, got %d", outI)
	}

	outV, outI = FindClosest(l1, 2)
	if outV != 2 {
		t.Errorf("Expected 2, got %d", outV)
	}

	if outI != 1 {
		t.Errorf("Expected 1, got %d", outI)
	}

	outV, outI = FindClosest(l1, 21)
	if outV != 22 {
		t.Errorf("Expected 22, got %d", outV)
	}

	if outI != 6 {
		t.Errorf("Expected 6, got %d", outI)
	}

	outV, outI = FindClosest([]int{}, 5)
	if outV != -1 {
		t.Errorf("Expected -1, got %d", outV)
	}

	if outI != -1 {
		t.Errorf("Expected -1, got %d", outI)
	}

}

func TestRemoveElementInt(t *testing.T){
	var list []int= []int{0,1,2,3,4,5,6,7,8,9}

	var remove []int = RemoveElementInt(list, 5)

	var exp []int = []int{0,1,2,3,4,6,7,8,9}

	if len(exp) != len(remove){
		t.Errorf("Lengths not equal. Expected %d, got %d", len(exp), len(remove))
	}

	if !Equal(remove, exp){
		t.Errorf("Expected %v, got %v", exp, remove)	
	}
}

func TestSCAN(t *testing.T){
	f1,e1:= os.Open("./pa2r3.1-rls/pa2r3-rls/scan20.base")
	inf1, e2 := os.Open("./pa2r3.1-rls/pa2r3-rls/scan20.txt")
	if e1 != nil{
		f1.Close()
		log.Fatal(e1)
	}
	if e2 != nil{
		inf1.Close()
		log.Fatal(e2)
	}

	exp := FileToStr(f1)
	inS := FileToStr(inf1)
	inLines := strings.Split(inS, "\n")
	inLines = RemoveComments(inLines)
	p, _:= Parse(inLines)
	out:= ParamsToString(p)
	out += Run(p)

	if exp != out{
		t.Errorf("Expected:\n%s, got:\n%s",exp , out)
	}
}

func TestSplit(t *testing.T) {
	var b bool = Split(' ') && Split('\n') && Split('\t') && Split('\r') && !Split('z')
	if !b {
		t.Error("Expected true, got", b)
	}
}

func TestSSTF(t *testing.T) {

	f1, e1 := os.Open("./pa2r3.1-rls/pa2r3-rls/sstf01.txt")
	exf1, e2 := os.Open("./pa2r3.1-rls/pa2r3-rls/sstf01.base")
	if e1 != nil {
		f1.Close()
		log.Fatal(e1)
	}
	if e2 != nil {
		exf1.Close()
		log.Fatal(e2)
	}
	in1 := FileToStr(f1)
	ex1 := FileToStr(exf1)

	inS1 := strings.Split(in1, "\n")

	inS1 = RemoveComments(inS1)

	p1, _ := Parse(inS1)

	out1 := ParamsToString(p1)
	out1 += Run(p1)
	//fmt.Printf("%s\n", Run(p1))
	if out1 != ex1 {
		t.Errorf("Expected\n%s got\n%s", ex1, out1)
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
		t.Errorf("Expected %s , got %s\n", exS, outS)
	}
}

/*
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
		t.Errorf("Expected\n %s , got\n%s", exp, out)
	}

}
*/