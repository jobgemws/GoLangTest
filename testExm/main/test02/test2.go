package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCount int
	_, err := fmt.Fscan(in, &testCount)

	if IsError(err) {
		return
	}

	for i := 0; i < testCount; i++ {
		var d, m, y int

		_, err := fmt.Fscan(in, &d, &m, &y)

		if IsError(err) {
			return
		}

		if (1950 > y) || (y > 2300) {
			_, err := fmt.Fprintln(out, "NO")
			if IsError(err) {
				return
			}
			continue
		}

		if (1 > m) || (m > 12) {
			_, err := fmt.Fprintln(out, "NO")
			if IsError(err) {
				return
			}
			continue
		}

		if (1 > d) || (d > 31) {
			_, err := fmt.Fprintln(out, "NO")
			if IsError(err) {
				return
			}
			continue
		}

		if (d > 30) && ((m == 2) || (m == 4) || (m == 6) || (m == 9) || (m == 11)) {
			_, err := fmt.Fprintln(out, "NO")
			if IsError(err) {
				return
			}
			continue
		}

		if m == 2 {
			if d > 29 {
				_, err := fmt.Fprintln(out, "NO")
				if IsError(err) {
					return
				}
				continue
			}

			if !(((y%4) == 0) && ((y%100) != 0) || ((y % 400) == 0)) {
				if d > 28 {
					_, err := fmt.Fprintln(out, "NO")
					if IsError(err) {
						return
					}
					continue
				}
			}
		}

		_, err = fmt.Fprintln(out, "YES")
		if IsError(err) {
			return
		}
	}

	//fmt.Fprintln(out, "I am sure that I will fill out the form by 10:00 am on September 12, 2022.")
}

func IsError(err error) (res bool) {
	if err != nil {
		res = true
		if err == io.EOF {
			return
		} else {
			fmt.Println(err)
			return
		}
	}

	return
}
