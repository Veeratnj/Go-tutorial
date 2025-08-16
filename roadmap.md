## **1️⃣ Go Foundations**

These are the must-learn basics so you can write working programs.

* **Setup & Basics**
  * Installing Go & `go` command basics (`go run`, `go build`, `go mod`)
  * Folder structure, modules, and workspaces
* **Syntax & Data Types**
  * Variables, constants, iota
  * Built-in types: string, int, float, bool
  * Arrays, slices, maps
* **Operators & Control Flow**
  * `if`, `for`, `switch`
  * Defer, panic, recover
* **Functions**
  * Named/unnamed return values
  * Variadic functions
  * Pass-by-value vs. reference (pointers)
* **Structs & Interfaces**
  * Struct definition, embedding
  * Methods (value vs pointer receivers)
  * Interfaces & duck typing
* **Error Handling**
  * `error` type
  * Creating custom errors
  * Wrapping errors with `fmt.Errorf`

---

## **2️⃣ Intermediate Go**

Where you learn how to build concurrent, modular, and efficient programs.

* **Packages**
  * Importing, exporting
  * Writing reusable packages
* **Pointers & Memory**
  * `&`, `*`, pointer vs. value semantics
* **Collections & Built-in Packages**
  * `time`, `strings`, `strconv`, `math`, `sort`
* **Concurrency**
  * Goroutines
  * Channels (unbuffered & buffered)
  * `select` statement
  * Worker pools
* **Sync & Coordination**
  * `sync.Mutex`, `sync.WaitGroup`
  * `sync.Once`, `sync.Map`
  * Context cancellation (`context` package)
* **File & Network I/O**
  * Reading/writing files
  * HTTP client & server (`net/http`)
  * JSON encoding/decoding

---

## **3️⃣ Advanced Go**

Now you’re building production-grade applications.

* **Advanced Concurrency**
  * Channel patterns (fan-in, fan-out, pipeline)
  * Rate limiting & timeouts
* **Testing**
  * `testing` package basics
  * Table-driven tests
  * Benchmarking & profiling
* **Generics**
  * Type parameters (`[T any]`)
  * Constraints
* **Reflection**
  * `reflect` package
  * Type introspection
* **Code Quality**
  * Linting & formatting (`go fmt`, `golangci-lint`)
  * Vetting (`go vet`)
* **Dependency Management**
  * `go mod tidy`, `replace`, `vendor`
* **Performance Optimization**
  * Profiling with `pprof`
  * Memory usage & GC tuning

---

## **4️⃣ Real-World Development**

Where you apply Go in different domains.

* **Web Development**
  * REST API with `net/http`
  * Gorilla Mux / Chi router
  * Middleware
* **Databases**
  * `database/sql` + drivers
  * ORM tools (GORM, sqlx)
* **Microservices**
  * gRPC
  * Message queues (Kafka, NATS)
* **Cloud & Deployment**
  * Dockerizing Go apps
  * Cross-compiling
  * CI/CD
* **Specialized Areas**
  * CLI tools (Cobra)
  * Blockchain (tendermint, Ethereum clients)
  * Systems programming (file watchers, OS-level tools)

---

💡 **Pro Tip:**

Don’t just learn in order — build small projects at each stage.

Example path:

* **Stage 1:** CLI calculator
* **Stage 2:** Concurrent web scraper
* **Stage 3:** REST API + PostgreSQL
* **Stage 4:** Microservice with gRPC
