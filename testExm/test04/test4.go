// Задание № 4
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
	//*
	in := bufio.NewReader(os.Stdin)

	var str string

	//fmt.Println("Введите количество вводимых наборов:")

	_, err := fmt.Fscan(in, &str)

	if IsError(err) {
		return
	}

	c := int(ConvertStrToInt(str))

	sum := make([]string, c)

	res_m := make([][]int, c)
	//*
	for i := 0; i < c; i++ {
		var str string

		//fmt.Println("Введите количество вводимых задач:")

		_, err := fmt.Fscan(in, &str)

		if IsError(err) {
			return
		}

		c0 := int(ConvertStrToInt(str))

		sl := make([]int, c0)

		//fmt.Println("Введите задачи:")

		for j := 0; j < c0; j++ {
			var str string

			_, err := fmt.Fscan(in, &str)

			if IsError(err) {
				return
			}

			v := int(ConvertStrToInt(str))

			sl[j] = v
		}

		res_m[i] = sl
	}
	//*/

	//ch := make(chan [][]int, c)
	//*

	//res_m := [][]int{[]int{1, 1, 3, 2, 2}}
	//sum := make([]string, 5)
	var wg sync.WaitGroup
	//*/
	for i, v := range res_m {
		wg.Add(1)
		go func(i int, v []int) {
			defer wg.Done()

			len_v := len(v)

			para_m := make(map[int][]int, len_v)

			is_break := false

			para_m[v[0]] = make([]int, 1, len_v)
			para_m[v[0]][0] = 0

			for i0 := 1; i0 < len_v; i0++ {
				if _, ok := para_m[v[i0]]; ok {
					para_m[v[i0]] = append(para_m[v[i0]], i0)
					continue
				} else {
					para_m[v[i0]] = make([]int, 1, len_v)
					para_m[v[i0]][0] = i0
				}
			}

			for _, val := range para_m {
				ind := -1
				for _, val0 := range val {
					if ind < 0 {
						ind = val0
						continue
					} else {
						if ind != val0-1 {
							is_break = true
							break
						} else {
							ind = val0
						}
					}
				}

				if is_break {
					break
				}
			}

			if is_break {
				sum[i] = "NO"
			} else {
				sum[i] = "YES"
			}
		}(i, v)
	}
	//*/

	// for i, v := range res_m {
	// 	for key, value := range v {
	// 		fmt.Println("Ind:", i, "Key:", key, "Value:", value)
	// 	}
	// }

	//*
	wg.Wait()

	//*
	for _, v := range sum {
		fmt.Println(v)
	}
	//*/
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

func ConvertStrToInt(s string) (res int64) {
	res, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		panic(err)
	}

	return
}

func ConvertStrToFloat(s string) (res float64) {
	res, err := strconv.ParseFloat(s, 64)

	if err != nil {
		panic(err)
	}

	return
}
