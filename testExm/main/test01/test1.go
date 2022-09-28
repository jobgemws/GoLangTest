package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	// var testCount int
	// fmt.Fscan(in, &testCount)

	// for i := 0; i < testCount; i++ {
	//     var n, m int
	//     fmt.Fscan(in, &n, &m)
	//     fmt.Fprintln(out, n + m)
	// }

	fmt.Fprintln(out, "I am sure that I will fill out the form by 10:00 am on September 12, 2022.")
}
