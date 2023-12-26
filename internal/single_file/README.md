# single_file

This is an example of how to use go channels to block and unblock a services.

Characteristics of this example:

+ services are started and stopped randomly.
+ services are started in background and then the main go routine waits until services are started.
+ main go routine then waits for the services to stop.

## Instructions

Once it is running using the instruction in the main [README.md](../README.md), you can
use the following commands to interact with the service:

1. Control c (^c) to stop the service in the same terminal window.
2. In a different terminal window, you can use the following commands:
    1. kill -s SIGHUP <pid> to stop the service
       1. kill -1 \$(ps | grep "go-build.*main$" | head -1 | cut -d ' ' -f 1)
    2. kill -s SIGUSR1 <pid> to print the current status of the service
       1. kill -10 \$(ps | grep "go-build.*main$" | head -1 | cut -d ' ' -f 1)

## License

Apache License 2.0
