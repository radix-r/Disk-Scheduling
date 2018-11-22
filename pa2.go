/*
Author: Ross Wagner
Date: 11/17/2017

I Ross Wagner (ro520462) affirm that
this program is entirely my own work and that I have neither developed my code with any
another person, nor copied any code from any other person, nor permitted my code to be copied
or otherwise used by any other person, nor have I copied, modified, or otherwise used programs
created by others. I acknowledge that any violation of the above terms will be treated as
academic dishonesty.

This program's goal is to implement 6 disk searching algorithms
a) fcfs
b) sstf
c) SCAN
d) C-SCAN
e) LOOK
f) C-LOOK
and compare their performance

*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
Data structure to hold the input parameters for disk searching alg
*/
type parameters struct {
	alg      string
	lowerCYL int
	upperCYL int
	initCYL  int
	requests []int
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//
func Fcfs(p parameters) string {
	var out string = ""
	var trav int = 0
	var at int = p.initCYL

	for _, cyl := range p.requests {
		out += fmt.Sprintf("Servicing %5d\n", cyl)

		trav += Abs(cyl - at)
		at = cyl
	}
	out += fmt.Sprintf("FCFS traversal count = %5d\n", trav)

	return out
}

/*
This function takes in a file and returns its contents as a string
Parameters File
Returns string
*/
func FileToStr(file *os.File) string {
	scanner := bufio.NewScanner(file)
	txt := ""
	for scanner.Scan() {
		txt += scanner.Text() + "\n"

	}
	return txt
}

func ParamsToString(params parameters) string {
	var out string

	out += fmt.Sprintf("Seek algorithm: %s\n", strings.ToUpper(params.alg))
	out += fmt.Sprintf("\tLower cylinder: %5d\n", params.lowerCYL)
	out += fmt.Sprintf("\tUpper cylinder: %5d\n", params.upperCYL)
	out += fmt.Sprintf("\tInit cylinder:  %5d\n", params.initCYL)
	out += fmt.Sprintf("\tCylinder requests:\n")
	for _, req := range params.requests {
		out += fmt.Sprintf("\t\tCylinder %5d\n", req)
	}
	return out
}

/*
reads input line by line and puts the data into a parameters struct
returns error code -1 if Abort error detected
*/
func Parse(lines []string) (parameters, int) {

	var alg string = ""
	var lower int = -1
	var upper int = -1
	var init int = -1
	var req []int = nil
	var p parameters
	var state int = 0
	var prevTok string = ""
	var prevTokLine int = 0

	// finite state machine
	for l, line := range lines {
		if state == -1 {
			break
		}
		state = 0
		// tokenize based on spaces
		tokens := strings.FieldsFunc(line, Split)
		for _, token := range tokens {
			if state == -1 {
				break
			}
			switch state {
			// start
			case 0:
				if token == "use" {
					state = 1
				} else if token == "lowerCYL" {
					state = 2
				} else if token == "upperCYL" {
					state = 3
				} else if token == "initCYL" {
					state = 4
				} else if token == "cylreq" {
					state = 5
				} else if token == "end" {
					state = 6
				} else {
					// unrecognized token
					state = 7
				}
				break
			case 1:
				if ValidAlg(token) {
					alg = token
				}
			case 2:
				i, err := strconv.Atoi(token)

				if err != nil {
					log.Fatal(err)
				}

				lower = i
			case 3:
				i, err := strconv.Atoi(token)

				if err != nil {
					log.Fatal(err)
				}

				upper = i
			case 4:
				i, err := strconv.Atoi(token)

				if err != nil {
					log.Fatal(err)
				}

				init = i
			case 5: // cylinder request
				i, err := strconv.Atoi(token)

				if err != nil {
					log.Fatal(err)
				}

				// add to requests array
				req = append(req, i)

			case 7:
				state = -1
				// get previous token and line

				fmt.Printf("ABORT(1): unrecognized token: %s on line %d\n", prevTok, prevTokLine)

				// Abort
			}
			prevTok = token
			prevTokLine = l
		}
	}

	if state == -1 {
		// error encountered
		return p, -1
	}

	//assign values to p
	p.alg = alg
	p.lowerCYL = lower
	p.upperCYL = upper
	p.initCYL = init
	p.requests = req
	// return p with normal exit code
	return p, 0
}

func Run(p parameters) string {

	var output string = ""
	// determine which alg to use
	// maybe use code of alg
	if p.alg == "fcfs" {
		output = Fcfs(p)
	}

	return output
}

/**/
func Split(r rune) bool {
	return r == ' ' || r == '\n' || r == '\t' || r == '\r' //|| r == '\n\r'
}

func ValidAlg(alg string) bool {
	var algs []string = []string{
		"fcfs", "sstf", "scan", "c-scan", "look", "c-look",
	}
	for _, str := range algs {
		if alg == str {
			return true
		}
	}
	return false
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: ./pa2 <input file>\n")
		return
	}

	inputFile := os.Args[1]

	file, err := os.Open(inputFile)

	if err != nil {
		file.Close()
		log.Fatal(err)
	}

	txt := FileToStr(file)
	file.Close()

	// print contents
	//fmt.Printf("%s", txt)
	//fmt.Printf("\n\n")
	// split by line
	lines := strings.Split(txt, "\n")

	// remove comments
	for i, line := range lines {
		for j, char := range line {
			if char == '#' {
				lines[i] = line[0:j]
				break
			}
		}
	}

	/*
		// print without comments
		for _, line := range lines {
			fmt.Printf("%s\n", line)
		}
	*/

	// extract input parameters from the input
	inParam, code := Parse(lines)
	if code == -1 {
		// abort condition triggered
		return
	}

	// check for invalid params
	fmt.Printf("%s", ParamsToString(inParam))
	fmt.Printf("%s", Run(inParam))

}
