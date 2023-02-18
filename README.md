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

## Reference
- https://go.dev/doc/
- https://www.youtube.com/watch?v=JoJ8Sw5Yb4c&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa
