package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"sync"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCount int //= 1

	_, err := fmt.Fscan(in, &testCount)

	if IsError(err) {
		return
	}

	var wg sync.WaitGroup

	for i := 0; i < testCount; i++ {
		var testCount2 int //= 26

		_, err := fmt.Fscan(in, &testCount2)

		if IsError(err) {
			return
		}

		otz_m := make([]int, testCount2, testCount2)

		for j := 0; j < testCount2; j++ {
			_, err := fmt.Fscan(in, &otz_m[j])

			if IsError(err) {
				return
			}
		}

		// otz_m[0] = 5
		// otz_m[1] = 12
		// otz_m[2] = 6
		// otz_m[3] = 0
		// otz_m[4] = 9
		// otz_m[5] = 0
		// otz_m[6] = 13
		// otz_m[7] = 6
		// otz_m[8] = 4
		// otz_m[9] = 17
		// otz_m[10] = 9
		// otz_m[11] = 5
		// otz_m[12] = 4
		// otz_m[13] = 13
		// otz_m[14] = 5
		// otz_m[15] = 13
		// otz_m[16] = 6
		// otz_m[17] = 5
		// otz_m[18] = 4
		// otz_m[19] = 5
		// otz_m[20] = 4
		// otz_m[21] = 17
		// otz_m[22] = 3
		// otz_m[23] = 5
		// otz_m[24] = 21
		// otz_m[25] = 3

		wg.Add(1)

		go func(otz_m *[]int) {
			defer wg.Done()

			res := RunStr(otz_m)

			_, err := fmt.Println(out, res)

			if IsError(err) {
				return
			}

		}(&otz_m)
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

func RunStr(otz_m *[]int) (res string) {
	ln := len(*otz_m)

	if ln < 15 {
		runes := make([]rune, ln*3)
		i := 0
		for _, _ = range *otz_m {
			runes[i] = '-'
			runes[i+1] = '1'
			runes[i+2] = ' '
			i += 3
		}

		res = string(runes)

		return

	}

	otz_mm := make([][]int, ln)

	runes := make([]rune, ln*2)

	for i := 0; i < ln; i++ {
		otz_mm[i] = []int{i, (*otz_m)[i]}
	}

	sort.Slice(otz_mm, func(i, j int) bool {
		return otz_mm[i][1] > otz_mm[j][1]
	})

	pred_v := -1
	ln = len(*otz_m)

	res0 := make([][]int, ln)

	ind := 0

	for i, v := range otz_mm {
		if pred_v < 0 {
			pred_v = v[1]
			res0[ind] = []int{i, 1}
		} else {
			if pred_v > v[1] {
				res0[ind] = []int{i, res0[ind-1][1] + 1}
			} else {
				res0[ind] = []int{i, res0[ind-1][1]}
			}
		}

		ind++
	}

	ind_val := 1

	for res0[ind-1][1] > 5 {
		for i := ind_val; i < ln; i++ {
			if res0[i][1] > 1 {
				res0[i][1]--
			} else {
				ind_val = i
			}
		}
	}

	sort.Slice(res0, func(i, j int) bool {
		return res0[i][0] < res0[j][0]
	})

	i := 0
	for _, v := range res0 {
		runes[i] = rune(strconv.Itoa(v[1])[0])
		runes[i+1] = ' '
		i += 2
	}

	res = string(runes)

	return
}
