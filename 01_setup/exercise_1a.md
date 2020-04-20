# Exercise 1A: Find Stuff

## Goals

- Familiarize yourself with the golang.org website and resources

## Setup

- Visit `golang.org`

## Directions

Answer the following questions

1. Read about `for loops` in the _Effective Go_ document

Examples of for-loop

```go
  sum := 0
  for i:= 0; i < 10; i++ {
    sum += 1
  }
```

If looping over some data-structure we can use:

```go
for key, value := range hashMap {
  hashMap[key] = value
}
```

- What kind of loop doesn’t exist in Go?

  Ans: There is no do-while loop in go.

2. Read about the `fmt` _package_

`%v` is generalized format

The default format for %v is:

bool: %t
int, int8 etc.: %d
uint, uint8 etc.: %d, %#x if printed with %#v
float32, complex64, etc: %g
string: %s
chan: %p
pointer: %p

- What does `fmt.Println()` return?

It returns the number of bytes written and any write error encountered.

3. Find a _blog post_ about the recent release of Go 1.13

https://blog.golang.org/protobuf-apiv2

- What are some of the new features?
