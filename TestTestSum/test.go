package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"math"
	"strconv"
	"sync"
)

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	/*
		_, res2 := ttt(42, 28, 100) //gcd(42, 28), fib(100)

		//fmt.Println(res1)
		fmt.Println(res2)

		var x, y complex128 = complex(1, 2), complex(3, 4)
		x = 1 + 2i

		fmt.Println(real(x * y))
		fmt.Println(cmplx.Sqrt(-1))

		s := "Привет всем!\r\n"
		n := 0
		for _, _ = range s {
			n++
		}

		fmt.Printf("%d %s", utf8.RuneCountInString(s), s)

		fmt.Println(intsToString([]float64{1.1, 2.2, 37, 643}))

		fmt.Printf("%f", math.Pi)

		const noDelay time.Duration = 0
		const timeout = 5 * time.Minute

		fmt.Printf("%T %[1]v\n", noDelay)
		fmt.Printf("%T %[1]v\n", timeout)
		fmt.Printf("%T %[1]v\n", time.Minute)

		runes := []rune(s)
		substr := string(runes[7:])

		fmt.Printf("%s", substr)

		var a = make([]int, 8, 10) //2 is length, 10 is capacity
		for i := 0; i < len(a); i++ {
			a[i] = i
		}
		Reverse [int] (a)

		for i := 0; i < len(a); i++ {
			fmt.Printf("%d ", a[i])
		}

		fmt.Println()

		ages := make(map[string]int)
		ages["alice"] = 31
		ages["charlie"] = 34

		ages["alice"] = 32
		delete(ages, "alice")
		ages["Майкл"] = 32

		var names []string
		for name := range ages {
			names = append(names, name)
		}

		sort.Strings(names)

		for i, name := range names {
			fmt.Printf("%d\t%s\t%d\r\n", i, name, ages[name])
		}

		age, isexist := ages["Майкл"]

		fmt.Printf("Майкл %t: %d\r\n", isexist, age)

		agesNew := make(map[string]int)

		for k, v := range ages {
			agesNew[k] = v
		}

		delete(agesNew, "Майкл")

		res := equalMap(ages, agesNew)

		fmt.Printf("%t\r\n", res)

		var graph = make(map[string]map[string]bool)

		graph["Майкл"] = make(map[string]bool)
		graph["Майкл"]["Внутренности"] = true

		fmt.Printf("%t\r\n", graph["Майкл"]["Внутренности"])

		var w Wheel
		w.X = 8
		w.Y = 9
		w.Radius = 5
		w.Spokes = 20

		w = Wheel{
			Circle: Circle{
				Point:  Point{X: 8, Y: 9},
				Radius: 5,
			},
			Spokes: 20,
		}

		fmt.Printf("%#v\n", w)

		fmt.Printf("%v %v %v", 111, 234.567, "аргуменТ")

		f := squares()
		fmt.Println(f())
		fmt.Println(f())
		fmt.Println(f())
		fmt.Println(f())
		fmt.Println(f())

		values := make([]float64, 10, 20) //[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10.6789}
		copy(values, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10.6789})
		fmt.Println(sum(values...))

		InvokeDeferFunction()

		p1 := Point{X: 8, Y: 8}
		p2 := Point{X: 9, Y: 9}

		fmt.Println(p1.Distance(&p2))
		fmt.Printf("%+v\r\n", p2)

		ttt := Counter{n: 6}

		fmt.Println(&ttt)

		var name string
		var age0 int
		fmt.Print("Введите имя: ")
		fmt.Fscan(os.Stdin, &name)

		fmt.Print("Введите возраст: ")
		fmt.Fscan(os.Stdin, &age0)

		fmt.Println(name, age0)

		conn, err := net.Dial("tcp", "localhost:8000")

		if err != nil {
			log.Fatal(err)
		}

		done := make(chan struct{})

		go func() {
			io.Copy(os.Stdout, conn)
			log.Println("done")
			done <- struct{}{}
		}()

		mustCopy(conn, os.Stdin)
		conn.Close()
		<-done

		--*/

	naturals := make(chan int, 500)
	squares := make(chan int, 500)
	var wg sync.WaitGroup

	/*
		//Генерация
		go func() {
			for x := 0; x < 100; x++ {
				naturals <- x
			}
			close(naturals)
		}()

		//Возведение в квадрат
		go func() {
			for x := range naturals {
				squares <- x * x
			}
			close(squares)
		}()

		//Вывод (в главной go-подпрограмме)
		for x := range squares {
			fmt.Println(x)
		}
		--*/

	wg.Add(1)
	go counter(naturals, &wg)
	wg.Add(1)
	go squarer(squares, naturals, &wg)
	wg.Add(1)
	printer(squares, &wg)
	wg.Wait()

}

func counter(out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for v := range in {
		fmt.Println(v)
	}
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

type Counter struct{ n int }

func (c *Counter) N() int     { return c.n }
func (c *Counter) Increment() { c.n++ }
func (c *Counter) Reset()     { c.n = 0 }
func (c *Counter) String() (res string) {
	res = strconv.Itoa(c.n)
	return
}

func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point

	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}
	for i := range path {
		path[i] = op(path[i], offset)
	}
}

func (p Point) Add(q Point) Point { return Point{p.X + q.X, p.Y + q.Y} }
func (p Point) Sub(q Point) Point { return Point{p.X - q.X, p.Y - q.Y} }

type Path []Point

func (p *Point) Distance(q *Point) float64 {
	return math.Hypot(float64(q.X-p.X), float64(q.Y-p.Y))
}

func InvokeDeferFunction() {
	defer DeferFunctionCall()
	defer OtherDeferFunctionCall()
	fmt.Println("Still executing InvokeDeferFunction")
}

func DeferFunctionCall() {
	fmt.Println("Defer Function Called...")
}

func OtherDeferFunctionCall() {
	fmt.Println("Other Defer Function called...")
}

func sum(vals ...float64) (res float64) {
	for _, v := range vals {
		res += v
	}

	return
}

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func ttt(x, y, n int) (int, int) {
	return gcd(x, y), fib(n)
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}

	return x
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

func intsToString(values []float64) string {
	var buf bytes.Buffer
	buf.WriteByte('[')

	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%f", v)
	}

	buf.WriteByte(']')
	return buf.String()
}

func equalString(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}

	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func Reverse /*[T any]*/ (s []int /*T*/) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func equalMap(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}

	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}

	return true
}
