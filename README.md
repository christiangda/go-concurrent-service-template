# go-concurrent-service-template

This is an example how to use Golang (go) `goroutine` to run multiple services (long running) concurrently.

## How it works

The main routine will start a goroutine for each service, and wait for all `goroutines` to finish.
This mechanism allows the main routine to be able to stop all services when it receives a `SIGINT` signal.

The implementation is based on channel and `select` statement.

[X] Listen different Operating System (OS) signals
[X] Stop all services when receiving OS
[X] Support wait for all services to start and stop

## How to run

List all available services

```bash
go run main.go -list
```

Run service by name

```bash
go run main.go -run <service_name>
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
