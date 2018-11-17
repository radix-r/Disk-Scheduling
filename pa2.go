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
)

/*
This function takes in a file and returns its contents as a string
Parameters File
Returns string
*/
func fileToStr(file *os.File) string {
	scanner := bufio.NewScanner(file)
	txt := ""
	for scanner.Scan() {
		txt += scanner.Text() + "\n"

	}
	return txt
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

	txt := fileToStr(file)
	file.Close()

	// print contents
	fmt.Printf("%s", txt)

	// split by line
	//lines := strings.Split(txt, "\n")

	// remove comments
	/*for i, line := range lines {
		for j, char := range line {
			if char == '#' {

			}
		}
	}
	*/

}
