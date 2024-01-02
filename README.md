# go-concurrent-service-template

This is an example of Golang multi server services using concurrency

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
