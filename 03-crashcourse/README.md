### This is only for getting a highlight of golang in a minimal amount of time
---

## How these notes are structured
1. Overview (quick summary)
2. Setup & project structure
3. Basic types: variables, constants, arithmetic
4. Functions, returns, and error handling
5. Control flow: if, switch, loops
6. Arrays, slices, maps
7. Strings, runes, bytes, and builders
8. Structs, methods, and interfaces
9. Pointers
10. Concurrency: goroutines, WaitGroup, mutexes, RWMutex
11. Channels and select
12. Generics (brief)
13. Building an example API (project layout, middleware, handlers)
14. Practical tips & performance notes
15. Short reference / commands

Each section includes concise examples and the exact place in the video the concept was taken from (inline citation).

---

## Learn GO Fast — Detailed Notes (Markdown)

# Learn GO Fast — Notes

> Source: "Learn GO Fast: Full Tutorial" (video).  
> Time coverage used: entire video (00:00–01:07:46).

---

## 1. Overview — What is Go?

- **Go** is a **statically typed** and **strongly typed** compiled language. Variable types must be declared or inferred and don't change without conversion.
  
- **Compiler** produces a standalone binary (faster runtime than interpreted languages like Python) and Go compilation itself is fast.

- **Built-in concurrency** with **goroutines** and simple concurrency primitives (channels, mutexes) are core to the language.

- Design goals: **Simplicity**, concise syntax, garbage collection, easy concurrency, fast compile times.

---

## 2. Setup & Project Initialization

- Install from golang.org/doc/install and verify with:
  ```
  go version
  ```
  (Video: download installer & check version).

- Two key concepts: **package** (folder of.go files) and **module** (collection of packages). Initialize with:
  ```
  go mod init <module_name>
  ```
  This creates `go.mod` with module name and Go version.

- Special package `main` and `func main()` act as entry point for an executable.

- To build and run:
  - Build binary:
    ```
    go build./cmd/tutorial1/main.go
    ```
  - Or run directly:
    ```
    go run./cmd/tutorial1/main.go
    ```
  Example in video uses `go build` and `go run` to produce/run a binary.

---

## 3. Variables, Constants, & Basic Types

- Declaration:
  - Explicit:
    ```
    var intNum int
    ```
  - Inferred (short):
    ```
    myVar:= "hello"
    ```
  - Use `var` or `:=` shorthand; prefer explicit types when unclear.

- Types:
  - Integers: `int`, `int8`, `int16`, `int32`, `int64` (size/precision matters) — `int` depends on architecture (32/64-bit).
  - Unsigned: `uint`, `uint8` etc. (store only positive values).
  - Floats: `float32`, `float64` — choose based on precision needs.
  - String: `string` (UTF-8 encoded bytes).
  - Boolean: `bool` (true/false).
  - Runes: alias for `int32` used for Unicode code points; single-quoted literals like `'a'` produce runes.

- Defaults (zero-values):
  - `int` -> `0`, `string` -> `""`, `bool` -> `false`, `pointer/interface/slices/maps` -> `nil`.

- Arithmetic:
  - No implicit mixing of types (cannot add `int` + `float32`; explicit conversion required).
  - Integer division truncates (e.g., `3/2 == 1`) and `%` is modulo operator.

- Constants:
  - Use `const` for immutable values and must be initialized at declaration (e.g., `const Pi = 3.14159`).

---

## 4. Functions & Error Handling

- Define functions:
  ```
  func myFunc(param string) {
      //...
  }
  ```
  Use `func` keyword and braces on same line for Go style.

- Return values and multiple returns:
  ```
  func intDivision(numerator, denominator int) (int, int) {
      result:= numerator / denominator
      remainder:= numerator % denominator
      return result, remainder
  }
  ```
  Functions must return values matching declared return types.

- Error handling pattern:
  - Use `error` return type when something can fail.
  - Return `nil` when no error.
  - Callers check `if err!= nil { /* handle */ }` pattern.

- Create errors using `errors.New("message")` (from `errors` package) and return appropriate error/wrapper types.

---

## 5. Control Flow

- if / else if / else:
  - Standard `if condition { }` syntax; `else if` / `else` must be on same line as closing brace of prior block.
  - Logical operators: `&&` (and), `||` (or), `!=` (not equal), `==` (equal).

- switch:
  - `switch` with implied `break` (no need to explicitly `break`).
  - Can use conditional switch `switch value { case... }` to compare a variable to cases.

---

## 6. Arrays, Slices, Maps & Loops

- Arrays:
  - Fixed-length: `[3]int32{}` declares an array of length 3 of `int32`.
  - Stored in contiguous memory; size is part of type and cannot change.

- Slices:
  - More flexible, built on arrays: `[]int{}`.
  - Append with `append(slice, value)` returns new slice (may reallocate underlying array).
  - Capacity vs length: `len(slice)` and `cap(slice)`; appending can increase capacity (commonly doubling strategy illustrated).
  - Create with `make([]T, length, capacity)` to pre-allocate capacity for performance.
  - Example: pre-allocating 1,000,000 items can be much faster than repeated append without capacity.

- Maps:
  - Declaration: `m:= make(map[string]uint8)` or literal `map[string]uint8{"alice": 30}`.
  - Read nonexistent key returns zero value; to test existence use `v, ok:= m["key"]` where `ok` is `true` if present.
  - Delete with `delete(m, key)`.

- Loops:
  - `for range` for arrays, slices, maps:
    ```
    for i, v:= range slice {... }
    for k, v:= range myMap {... } // map order not guaranteed
    ```
    Map iteration order is random/unordered.
  - Traditional `for init; condition; post { }` syntax exists and can express while-loops. `for {}` is infinite loop, use `break`.

---

## 7. Strings, Runes, Bytes & String Builder

- Underlying representation: **strings are a sequence of bytes (UTF-8)**. Indexing gives bytes, not necessarily full Unicode code points (runes).

- `len(s)` returns number of bytes, not Unicode characters; use `utf8.RuneCountInString()` to count runes (characters).

- Iterating with `range` yields rune values and correct rune indices (range decodes UTF-8 for you).

- Runes:
  - Type alias for `int32`, represent Unicode code points. Declare `'a'` to get a `rune` literal.

- Strings are immutable; repeated concatenation using `+` allocates new strings — use `strings.Builder` for efficient concatenation:
  ```
  var b strings.Builder
  b.WriteString("a")
  b.WriteString("b")
  s:= b.String()
  ```
  Builder appends internally and creates final string once, more performant.

---

## 8. Structs, Methods & Interfaces

- Structs: custom composite types with named fields:
  ```
  type GasEngine struct {
      MPG uint8
      Gallons uint8
  }
  ```
  Zero-values for fields if not initialized.

- Struct literal forms:
  - By-name: `GasEngine{MPG: 30, Gallons: 10}`
  - Positional: `GasEngine{30, 10}` (order-dependent)
  - Anonymous structs possible for one-off usage but not reusable.

- Methods:
  - Associate a function with a receiver:
    ```
    func (g GasEngine) MilesLeft() uint8 {
       return g.MPG * g.Gallons
    }
    ```
  - Call with `g.MilesLeft()` — similar to a method on a class.

- Interfaces:
  - Define behavior via method signatures.
    ```
    type Engine interface {
       MilesLeft() uint8
    }
    ```
  - Any type implementing `MilesLeft()` satisfies the `Engine` interface implicitly — allows generic code that accepts anything implementing that method.

---

## 9. Pointers

- Pointer type uses `*T` to represent memory address containing value of type `T`.
  - Declare pointer:
    ```
    var p *int32
    p = new(int32) // allocate and returns *int32 (zero-initialized)
    ```
  - Dereference with `*p` to get/set the pointed value.

- Nil pointers cause runtime panic if dereferenced — ensure pointer is initialized (non-nil) before dereferencing.

- Address-of operator `&` obtains pointer to a variable:
  ```
  i:= int32(5)
  p:= &i
  *p = 10 // changes i
  ```
  Both refer to same memory location; modifying via pointer changes original variable.

- Passing large values to functions: use pointers to avoid copying large data structures; for arrays/passing slices, note slices are descriptors that already contain pointers to underlying data so copying a slice copies the descriptor not the underlying array — modifying the copy can change original underlying array.

---

## 10. Concurrency: goroutines, WaitGroup, Mutex, RWMutex

- Goroutine: use `go` keyword before a function call to run it concurrently:
  ```
  go doWork()
  ```
  Concurrency ≠ parallelism; goroutines are scheduled by runtime and may run concurrently; parallelism depends on CPU cores.

- WaitGroup (from `sync`):
  - Use `wg.Add(1)` before starting goroutine; call `defer wg.Done()` inside goroutine; `wg.Wait()` blocks until counter returns to zero — used to wait for goroutines to finish.

- Mutex (`sync.Mutex`):
  - Protect shared state by `mu.Lock()` / `mu.Unlock()`.
  - Lock placement matters — locking too broad defeats concurrency; lock only the critical section.

- Read-Write Mutex (`sync.RWMutex`):
  - `RLock()` / `RUnlock()` allow concurrent reads.
  - `Lock()` / `Unlock()` are exclusive — blocks other readers & writers while locked.
  - Use when you have many readers and fewer writers to improve concurrency.

- Performance notes:
  - For IO-bound tasks (e.g., waiting on DB calls), many goroutines are cheap and can give near-constant wall time due to waiting behavior.
  - For CPU-bound tasks, speedup limited by number of CPU cores (example: counting 100M operations across goroutines scaled relative to core count).

---

## 11. Channels & select

- Channels: typed conduits for passing data between goroutines, created with `make(chan T, capacity)`; capacity omitted means unbuffered (capacity 0) blocking channel.

- Send / Receive:
  ```
  ch <- v    // send v to channel ch
  v:= <-ch  // receive from channel ch
  ```
  Writing to an **unbuffered** channel blocks until another goroutine reads; reading blocks until a value is available. Deadlock will occur if no matching send/receive exists and runtime detects it.

- Typical pattern: producer goroutine writes to channel; main or consumer reads. Close channel with `close(ch)` to signal no more values — range over channel ends when closed.

- Buffered channels:
  - Allow limited asynchronous buffering; producers can fill buffer up to capacity before blocking; useful to decouple producers/consumers to some extent.

- Example use-case: spawn several goroutines checking remote websites, send site name on channel when a deal is found; main waits to receive first message and acts on it — channels make coordination easy and safe.

- select statement:
  - Like `switch` but for channels; waits on multiple channel operations and runs the case for the channel that becomes ready:
    ```
    select {
      case v:= <-ch1:
         // handle ch1
      case v:= <-ch2:
         // handle ch2
    }
    ```
  - Use to pick among multiple channel operations (e.g., prefer chicken deals over tofu deals).

---

## 12. Generics (Go 1.18+)

- Use generics to write functions/types that operate on multiple types without repetition.
  - Function generic form:
    ```
    func SumSlice[T int | float32 | float64](s []T) T {... }
    ```
    Square-bracket type parameter `T` constrained to certain types (a union).

- `any` can represent any type in generic contexts when appropriate (but not always valid for operations like `+` unless constrained).

- Use generics with structs too:
  ```
  type Car[E Engine] struct {
     Make string
     Model string
     Engine E
  }
  ```
  Allows `Car[GasEngine]` or `Car[ElectricEngine]` etc..

- Compiler often infers generic parameters, but sometimes you must provide type arguments explicitly (e.g., deserializing into specific struct types).

---

## 13. Example: Building an API in Go (high-level summary)

Project layout (example):
```
/api
/cmd/api/main.go
/internal/handlers
/internal/middleware
/internal/tool
go.mod
```
Follow conventional Go layouts; `cmd/<name>` for binaries, `internal` for private packages, `api` for specs/examples.

Key pieces shown in the tutorial:
- Using `chi` router for HTTP routing.
- Handler registration:
  - `r.Route("/account", func(r chi.Router) { r.Use(authMiddleware); r.Get("/coins", getCoinBalance) })` — middleware can be route-specific or global.
- Middleware pattern:
  - Authorization middleware reads params and headers, validates token; uses `next.ServeHTTP(w, r)` to call next handler on success, otherwise writes error response and returns.
- Database interface pattern:
  - Define interface type (methods) and implement a mock DB to satisfy it — allows swapping real DB later.
- Handlers:
  - Parse query params (e.g., via `gorilla/schema` decoder), call DB, construct JSON response and write to the `http.ResponseWriter` with appropriate content-type and status code.

---

## 14. Practical Tips & Performance Notes

- Prefer explicit types when the type is not obvious — helps editors and readability.

- Avoid unnecessary allocations:
  - Pre-allocate slice capacity with `make` if you know approximate size to avoid repeated reallocations (big speed gain for large loops).

- Choose proper numeric types:
  - Use `uint8` for byte-like values (e.g., RGB 0–255) rather than large `int` for space efficiency.

- Use `strings.Builder` for repeated concatenation (performance & fewer allocations).

- For concurrency:
  - Use channels for communication and coordination.
  - Use mutexes for protecting shared memory.
  - Use `RWMutex` for many readers / few writers scenarios.
  - Use `WaitGroup` to wait for goroutines to finish.

---

## 15. Short Reference / Useful Commands

- Initialize module:
  ```
  go mod init github.com/you/yourproject
  ```
  (Creates go.mod).

- Get dependencies and tidy:
  ```
  go get <pkg>
  go mod tidy
  ```
  (Adds external packages, updates go.mod).

- Run:
  ```
  go run./cmd/api
  ```
  Or build:
  ```
  go build -o bin/api./cmd/api
  ```

- Formatting & vetting:
  ```
  go fmt./...
  go vet./...
  ```

---

<br>   

## Summary / Key Takeaways

- **Go is simple, statically typed, compiled, and has built-in concurrency** via goroutines and channels; use the `go` toolchain to build/run projects and `go mod` for dependencies.
  
- **Prefer simple, explicit constructs:** typed variables, small interfaces, and idiomatic error handling (`if err!= nil`) make code readable and robust.

- **Use slices & maps** correctly, pre-allocate capacity for large workloads, and be mindful of underlying memory / pointer behavior with slices and arrays.

- **Concurrency is powerful but must be used carefully:** coordinate via channels and protect shared memory with mutexes/RWMutex, and use WaitGroups to synchronize goroutines.
