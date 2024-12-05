## Telemetry package

This package is used to collect data and log it on various channels.

The library is configurable via the .env file.

Available drivers:

`driver_type=CLI`

`driver_type=JSON`

`driver_type=PLAIN`

For the `JSON/PLAIN` drivers, the `logs_storage_location` is required. If not provided, it will default to a log file, that depends on the driver.

The logger can be used as in the examples provided in either cmd/main.go or in the logger specific tests.
