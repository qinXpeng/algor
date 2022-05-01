package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/qinXpeng/go-Algorithm/algorithm/list"
)

func main() {
	lt := list.NewList[string]()
	lt.Push_front("abc")
	lt.Push_back("poi")
	printf("val: ",lt.Back_value())
}

var sc *bufio.Scanner

func init() {
	file, err := os.Open("in")
	if err != nil {
		file = os.Stdin
	}
	sc = bufio.NewScanner(file)
	maxLen := int(1e6 + 5)
	sc.Buffer(make([]byte, maxLen), maxLen)
	sc.Split(bufio.ScanWords)
}
func nextInfo() string {
	sc.Scan()
	return sc.Text()
}
func scanf(ioval ...interface{}) (err error) {
	for _, val := range ioval {
		s := nextInfo()
		switch v := val.(type) {
		case *int:
			*v, err = strconv.Atoi(s)
		case *int64:
			*v, err = strconv.ParseInt(s, 10, 64)
		case *string:
			*v, err = s, nil
		case *float64:
			*v, err = strconv.ParseFloat(s, 64)
		}
	}
	return
}
func printf(ioval ...interface{}) (n int, err error) {
	return fmt.Fprintln(os.Stdout, ioval...)
}
func sprintf(format string, ioval ...interface{}) (n int, err error) {
	return fmt.Fprintf(os.Stdout, format, ioval...)
}
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
func min(a, b int) int {
	return a + b - max(a, b)
}
