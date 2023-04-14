# hello-go

## Before

- complied -> go can run file directly, without VM(virtual machine)
- system app and webapp
- objected oriented -> yes and no
- no try catch

## Intallation and hello world

```go
go run main.go
```
- go | run | file name

---

```go
go help
go help build
```

- information of go command
- case sensitve (important)


## GOPATH

```go
go help GOPATH
```

> The Go path is used to resolve import statements.
> It is implemented by and documented in the go/build package.


## Lexer and type

- if use `;`, automatic removed
    - 세미콜론을 생략할 수 있음.
    - 알아서 삽입되는 경우가 있는듯.
- type -> case sensitive
    - `fmt.Println()` , Println() is public funciton.
- everything is type


## Variables

```txt
main.go:1:1: expected 'package', found 'EOF'
```

- pacakge 를 못찾았나?
- https://stackoverflow.com/questions/31110191/go-failing-expected-package-found-eof 
- vscode 저장 관련 이슈 인듯


```go
var username string = "han"
fp
```

- var | name | type
- shortcut for `fmt.PrintLn()`


```go
var smallVal uint8 = 255
```

- 0 <= uint8 <= 255
- https://go.dev/ref/spec#Numeric_types

```go
numberOfUser := 300_000 // declaration + assignment
numberOfUser = 300_000  // assignment
```

> numberOfUser := 300_000...

```go
var numberOfUser int // declartion
numberOfUser = 300_000 // assignment

// OR
var numberOfUser int = 300_000

// OR
var numberOfUser = 300_000
```

- `var foo int = 10` 와 `foo := 100` 은 같다.
- https://stackoverflow.com/questions/17891226/difference-between-and-operators-in-go


## User input
- go don't have the *try~catch*
- error를 변수 다루듯이 다룰 수 있는듯

```go
input, _ := reader.ReadString('\n')
```
- 입력값을 받는데 문제가 없으면 input으로 할당될 것.
- 그렇지만 입력값을 받는데 문제가 있으면, `_` 로 할당됨. (에러명을 따로 할당해주지 않는 syntax인듯)
- how to string input convert int ?
- https://stackoverflow.com/questions/65748509/vscode-shows-an-error-when-having-multiple-go-projects-in-a-directory
    - 참고해서 multi module 때문에서 발생하는 문제 해결
- https://pkg.go.dev/bufio


## Conversion
```go
strconv.ParseFloat(strings.TrimSpace(input), 64)
```

- float 으로 변환
- `\n` 같은 입력값을 지워주는 TrimSpace

## Time

## Build

- `GOOS`

```go
GOOS="windows" go build
```

- go file을 윈도우 포멧에 맞게 빌드시킬 수 있음.

## Memory Management

- memory allocation and deallocation happens automatically
- not developer jobs
- `new()`
    - allocate memory but no INIT(initialized)
    - you will get a memory address
    - **zeroed storage**
        - actually cannot put any data initially
- `make()`
    - allocate memory and INIT
    - you will get a memory address
    - **non-zeroed storage**
        - actually go head and put any of the data
- GC happens automatically
    - threshold가 있어서, 이 기준점을 넘으면 GC가 자동으로 호출되는듯
    - 기준이 100이면 아마도 이전 GC로 부터 살아남은 객채 70인데, 이번에 새로 할당되는 데이터가 30정도 크기라면 GC 호출.
    - 그러면 메모리가 꽉차는 경우도 있을듯
- Out of Scope of nil (null)
- https://pkg.go.dev/runtime
- func `NumCPU()`

## Pointer
```go
	// init
	myNumber := 23

	// creating pointer and reference some memory (myNumber)
	// reference -> &
	var ptr = &myNumber

	fmt.Println("Value of pointer is ", ptr)
	// Value of pointer is . 0xc00001e0e8

	fmt.Println("Value of pointer is ", *ptr)
	// Value of pointer is . 23
```

- var ptr 은 myNumber의 reference
- 그대로 프린트하면 메모리 주소가 나오고
- reference 하고 있는 변수에 * 붙이면, 포인터 내부에 무슨값이 있는지 알고 싶다라는 의미인듯.
- mutex 부분에서 또 나옴
- `final` 과 같은 선언은 없을까?
    - `const` 가 있는듯

## Array
- 생성시 arry 크기가 반드시 있어야함
    - 반드시는 아닌듯..?

```go
	var fruitList = []string{}
```

## Slice
- slice ?
    - `[]` 를 의미하는듯


## Map
```go
laguages := make(map[string]string)
```

- key=string, value=string

## Stuct
- no inhertiance, No super or parent
- `User`, 내부 변수도 Name같이 시작함.

## Switch
- fallthrough
    - 선언해놓으면 매칭되는 값이 나와도 다음 swtich 문으로 넘어가는듯

```go
	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
		fallthrough
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("except 1 or 2")
	}
    // Dice value is 1 and you can open
    // 2
```

## Fucntion
- 시작은 `func`
- `func main()` 이 고프로그램의 entry point
- function 내부에 function 을 정의할 순 없는듯
- function에 어떤 타입을 리턴할 것인지 정의해줘야함 (void가 아니면)
    - 여러개의 타입이 리턴되는 것도 정의할 수 있을듯

```go

// int 타입형태가 리턴될 것
func adder(i1 int, i2 int) int {
	return i1 + i2
}
```

## Defer
```go
func main() {
	defer fmt.Println("world")
	fmt.Println("Hello")
}

// Hello
// world
```
- code reverse...
- defer 키워드를 여러개 쓰면? 
    - LIFO

```go

func main() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("Hello")
}

// Hello
// 3
// 2
// 1
```

## Handling web

> net/http package

## Reference
- https://go.dev/doc/
- https://www.youtube.com/watch?v=JoJ8Sw5Yb4c&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa
