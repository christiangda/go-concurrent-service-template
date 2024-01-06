# go-concurrent-service-template

This is an example how to use Golang (go) `goroutine` to run multiple services (long running) concurrently.

Start multiple services, and stop all services when receiving `SIGINT` signal.

## How it works

The main routine will start a goroutine for each service, and wait for all `goroutines` to finish.
This mechanism allows the main routine to be able to stop all services when it receives a `SIGINT` signal.

The implementation is based on channel and `select` statement.

- [X] Listen different Operating System (OS) signals
- [X] Start main service and all child services
- [X] Wait for all services to start
- [X] Stop main service when receiving OS, and stop all child services
- [X] Wait for all services to stop
- [X] Example of single file service [single_file.go](internal/single_file/single_file.go)
- [X] Example of multiple files service [multi_file.go](internal/multi_file/multi_file.go)
- [X] Example of multiple files service

## How to run

List all available services

```bash
go run main.go -list
```

Run service by name

```bash
go run main.go -run -race <service_name>
```

Example

```bash
go run -race main.go -run single_file
 # or
go run -race main.go -run multi_file
```

Stop services

Once the main routine is blocking, `press Ctrl+C`` to stop all services or in a different terminal, run

```bash
kill -SIGINT <pid>
```

## build

```bash
make build
```

binary file will be generated in `build` folder

## License

Apache License 2.0
