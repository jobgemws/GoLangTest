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

	var str string

	//fmt.Println("Введите количество вводимых наборов:")

	_, err := fmt.Fscan(in, &str)

	if IsError(err) {
		return
	}

	c := int(ConvertStrToInt(str))

	sum := make([]int64, c)

	res_m := make([]map[int64]int, c)

	for i := 0; i < c; i++ {
		var str string

		//fmt.Println("Введите количество вводимых товаров:")

		_, err := fmt.Fscan(in, &str)

		if IsError(err) {
			return
		}

		c0 := int(ConvertStrToInt(str))

		m := make(map[int64]int, c0)

		//fmt.Println("Введите цены каждого товара:")

		for j := 0; j < c0; j++ {
			var str string

			_, err := fmt.Fscan(in, &str)

			if IsError(err) {
				return
			}

			v := ConvertStrToInt(str)

			m[v]++
		}

		res_m[i] = m
	}

	//ch := make(chan float64, c)

	var wg sync.WaitGroup

	for i, v := range res_m {
		wg.Add(1)
		go func(i int, v map[int64]int) {
			defer wg.Done()

			for key, value := range v {
				temp := value / 3

				//fmt.Println("value:", value, "temp:", temp, "key:", key)

				sum[i] += int64(value-temp) * key
			}

			//sum[i] = math.Round(sum[i]*1000) / 1000

			//ch <- sum[i]
		}(i, v)
	}

	// for i, v := range res_m {
	// 	for key, value := range v {
	// 		fmt.Println("Ind:", i, "Key:", key, "Value:", value)
	// 	}
	// }

	wg.Wait()

	for _, v := range sum {
		fmt.Println(v)
	}
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
