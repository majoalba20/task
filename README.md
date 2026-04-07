https://docs.google.com/presentation/d/1lyPheHPrUH9UsM_SjJXRGnKB0MX-i3JbSP5xMplCfx8/edit?usp=sharing

## Temario Golang

* Tipos: `structs`, `interfaces`, `embedding`
* Punteros: cuándo usarlos, escape analysis
* Goroutines: ciclo de vida, scheduler `G-M-P`
* Channels: buffered/unbuffered, `select`, `close`
* `sync`: `Mutex`, `RWMutex`, `WaitGroup`, `Once`
* `context`: cancelación, timeouts, propagación
* Manejo de errores: `errors.Is`, `errors.As`, wrapping
* Generics (Go 1.18+): constraints, type parameters
* `defer`, `panic`, `recover` — cuándo y cómo usarlos
* Reflection básica con `reflect`
* Gin: router, grupos, middleware, binding
* Fiber: contexto, handlers, validación
* Middleware: logging, auth JWT, recovery
* Binding y validación con `go-playground/validator`
* Manejo de JSON: `encoding/json`, tags en structs
* HTTP client avanzado: retry, timeouts, transport
* `pgx` v5: conexión, pool, query, scan
* `pgxpool`: configuración, adquisición de conexiones
* Transacciones con `pgx.Tx`
* `mongo-driver`: client, colección, operaciones CRUD
* Aggregation pipeline desde Go
* SQLC: generación de código desde SQL
* `testing`: `t.Run`, `t.Parallel`, `t.Cleanup`
* `testify`: `assert`, `require`, `mock`, `suite`
* `gomock`: generación de mocks para interfaces
* Benchmarks: `testing.B`, `b.ResetTimer`
* Race detector: `go test -race`
* Air: hot reload en desarrollo
* `golangci-lint`: linting exhaustivo
* `go build`, `go vet`, `go mod tidy`
* Repository pattern para acceso a datos
* Hexagonal / Clean Architecture en Go
* Dependency injection manual (o con Wire)
* Graceful shutdown con `signal.NotifyContext`
* Structured logging: `zap` / `slog`
* Métricas con `prometheus/client_golang`
