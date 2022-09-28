package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"
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

	var wg sync.WaitGroup

	for i := 0; i < testCount; i++ {
		var str string

		_, err := fmt.Fscan(in, &str)

		if IsError(err) {
			return
		}

		wg.Add(1)

		go func(str string) {
			defer wg.Done()

			chars := []rune(str)

			bm := make([]bool, len(chars))

			for i, v := range chars {
				s := string(v)
				b, err := strconv.ParseBool(s)

				if IsError(err) {
					return
				}

				bm[i] = b
			}

			str = Decoder(&bm)

			_, err := fmt.Fprintln(out, str)

			if IsError(err) {
				return
			}

		}(str)
	}

	wg.Wait()
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

func Decoder(bm *[]bool) (res string) {
	ln := len(*bm) - 1
	runes := make([]rune, 0, ln/2)
	var r0, r1, r2 bool

	for i := 0; i < ln; i++ {
		r0 = (*bm)[i]
		i++
		r1 = (*bm)[i]

		if r0 && r1 {
			runes = append(runes, 'd')
			continue
		}

		if (!r0) && (!r1) {
			runes = append(runes, 'a')
			continue
		}

		if r0 && (!r1) {
			i++
			if i <= ln {
				r2 = (*bm)[i]

				if r2 {
					runes = append(runes, 'c')
				} else {
					runes = append(runes, 'b')
				}

				continue
			} else {
				break
			}
		}
	}

	res = string(runes)

	return
}
