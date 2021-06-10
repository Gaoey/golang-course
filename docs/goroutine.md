# go

## elementary class day 3

Pallat Anchaleechamaikorn

yod.pallat@gmail.com
https://github.com/pallat
https://github.com/gophernment

https://tour.golang.org/
https://github.com/uber-go/guide

https://tourofrust.com/

---

## First-Class Function

```go
var add = func(a, b int) int {
    return a + b
}

fmt.Println(add(1, 2))
```

---

## Higher-Order Function

```go
func hof(fn func(string) string) {
    ...
}

func hof() func(string) string {
    ...
}
```

---

## Higher-Order Function Blog

https://dev.to/pallat/hof-in-go-18mm

---

## Closure Function

```go
func main() {
    fn1, fn2 := factory()
    fn1()
    fn1()
    fmt.Println(fn2())

    fn1()
    fmt.Println(fn2())
}

func factory() (func(), func() int) {
    var i int
    return func() {
            i++
        },
        func() int {
            return i
        }
}
```

---

## Demo Fibonacci Sequence by closure

$$ F_{0} = 0, F_{1} = 1, $$

and

$$ F_{n} = F_{n-1} + F_{n-2} $$

---

## func type

```go
type areaFunc func(float64, float64) float64 
```

---

## method on function

```go
    type IntnFunc func(int) int

    func (fn IntnFunc) Intn(n int) int {
        return fn(n)
    }
```

---

Demo: Pretend to be Yaiba

```go
func main() {
	yaiba := Yaiba(200)
	Attack(&yaiba)

	cp := newCostplay()
	Attack(cp)

}

func Attack(slayer DemonSlayer) {
	slayer.Damage(100)
}

type DemonSlayer interface {
	Damage(int)
}

type Yaiba int

func (yaiba *Yaiba) Damage(hit int) {
	*yaiba -= Yaiba(hit)
	if *yaiba <= 0 {
		fmt.Println("die")
		return
	}
	fmt.Println(*yaiba)
}

type costplay func(int)

func (fn costplay) Damage(hit int) {
	fn(hit)
}

func newCostplay() costplay {
	hp := 100
	return func(hit int) {
		hp -= hit
		if hp <= 0 {
			fmt.Println("die")
			return
		}
		fmt.Println(hp)

	}
}


```

---

## goroutine

```go
func main() {
    total := 10
    now := time.Now()
    for i := 0; i < total; i++ {
        go printout(i)
    }
    fmt.Println(time.Now().Sub(now))
}

func printout(i int) {
    fmt.Println(i)
}

```

---

## goroutine waiting

```go
var wg = sync.WaitGroup{}

func main() {
    total := 10
    wg.Add(total)
    now := time.Now()
    for i := 0; i < total; i++ {
        go printout(i)
    }
    wg.Wait()
    fmt.Println(time.Now().Sub(now))
}

func printout(i int) {
    fmt.Println(i)
    wg.Done()
}
```

---

## channel 

keyword `chan`

- no buffered channel
- buffered channel

---

## buffered channel

```go
total := 10
ch := make(chan int, total)
for i := total; i > 0; i-- {
    ch <- i
}
close(ch)

for i := range ch {
    fmt.Println(i)
}
```

---

## no buffered channel

```go
func main() {
    total := 10
    ch := make(chan struct{})
    now := time.Now()
    for i := 0; i < total; i++ {
        go printout(i, ch)
    }
    for i := 0; i < total; i++ {
        <-ch
    }
    fmt.Println(time.Now().Sub(now))
}

func printout(i int, ch chan struct{}) {
    fmt.Println(i)
    ch <- struct{}{}
}
```

---

END