// Задание № 3
package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
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

	sum := make([]ParaIndMap, c) //make([]map[int]int, c)

	res_m := make([][]int, c)
	//*
	for i := 0; i < c; i++ {
		var str string

		//fmt.Println("Введите количество вводимых сотрудников:")

		_, err := fmt.Fscan(in, &str)

		if IsError(err) {
			return
		}

		c0 := int(ConvertStrToInt(str))

		if c0%2 != 0 {
			fmt.Println("Кол-во сотрудников должно быть чётным!")
			return
		}

		sl := make([]int, c0)

		//fmt.Println("Введите уровень компетенций каждого сотрудника:")

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

	//res_m := [][]int{[]int{2, 1, 3, 1, 1, 4}, []int{5, 5}, []int{1, 4, 2, 5, 4, 2, 6, 3}}
	//sum := make([]map[int]int, 3)
	var wg sync.WaitGroup
	//*/
	for i, v := range res_m {
		wg.Add(1)
		go func(i int, v []int) {
			defer wg.Done()

			len_v := len(v)
			len_v2 := len_v / 2

			para_m := make(map[int]int, len_v2)
			para_m_res := make(ParaIndMap, len_v2) //make(map[int]int, len_v2)
			ind_m := 0

			for key, value := range v {
				if _, ok := para_m[key]; ok {
					continue
				}

				ind, min_diff := 1, -1 //int(math.Abs(float64(value-v[1])))

				for i0 := key + 1; i0 < len_v; i0++ {
					if _, ok := para_m[i0]; ok {
						continue
					}

					tmp := int(math.Abs(float64(value - v[i0])))
					if min_diff < 0 || tmp < min_diff {
						ind, min_diff = i0, tmp
					}
				}

				para_m[ind] = key
				para_m_res[ind_m].SetKey(key) //para_m_res[key] = ind
				para_m_res[ind_m].SetVal(ind)

				ind_m++

				if ind_m >= len_v2 {
					break
				}
			}

			//sort_m := make([][]int, len_v2)

			// k := 0

			// for key, value := range para_m_res {
			// 	sort_m[k] = []int{key, value}

			// 	k++
			// }

			// sort.SliceStable(sort_m, func(i, j int) bool {
			// 	return sort_m[i][0] < sort_m[j][0]
			// })

			sort.Sort(&para_m_res)

			sum[i] = para_m_res

			//ch <- para_m
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
		// keys := make([]int, len(v))
		// i := 0

		// for key, _ := range v {
		// 	keys[i] = key
		// 	i++
		// }

		// sort.Ints(keys)

		// for k := 0; k < len(keys); k++ {
		// 	fmt.Println(keys[k]+1, v[keys[k]]+1)
		// }
		for _, val := range v {
			fmt.Println(val.GetKey()+1, val.GetVal()+1)
		}

		fmt.Println()
	}
	//*/
}

type ParaInd struct {
	key int
	val int
}

func (p *ParaInd) GetKey() (res int) {
	res = p.key

	return
}

func (p *ParaInd) GetVal() (res int) {
	res = p.val

	return
}

func (p *ParaInd) SetKey(key int) {
	p.key = key
}

func (p *ParaInd) SetVal(val int) {
	p.val = val
}

func (s *ParaInd) Swap(w *ParaInd) {
	s.key, s.val, w.key, w.val = w.key, w.val, s.key, s.val
}

func (p *ParaInd) Copy() (res *ParaInd) {
	res = &ParaInd{key: p.key, val: p.val}
	return
}

type ParaIndMap []ParaInd

func (p *ParaIndMap) Len() (res int) {
	res = len(*p)
	return
}

func (p *ParaIndMap) Swap(i, j int) {
	(*p)[i].Swap(&(*p)[j])
}

func (p *ParaIndMap) Less(i, j int) (res bool) {
	res = ((*p)[i].key < (*p)[j].key)
	return
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
