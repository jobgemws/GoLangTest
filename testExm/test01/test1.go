package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
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

	sum := make([]int64, c) //[]float64

	ch := make(chan int64, 2) //float64

	//fmt.Println("Введите наборы из двух чисел в каждом:")

	for i := 0; i < c; i++ {
		var str1, str2 string
		_, err := fmt.Fscan(in, &str1, &str2)

		if IsError(err) {
			return
		} else {
			go func(str1 string) {
				a := ConvertStrToInt(str1) //ConvertStrToFloat
				ch <- a
			}(str1)

			go func(str2 string) {
				b := ConvertStrToInt(str2) //ConvertStrToFloat
				ch <- b
			}(str2)

			sum[i] = <-ch + <-ch
		}
	}

	for _, v := range sum {
		fmt.Printf("%v\r\n", v)
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
