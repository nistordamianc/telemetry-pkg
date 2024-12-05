## Telemetry package

This package is used to collect data and log it on various channels.

The library is configurable via the .env file.

Available drivers:

`DRIVER_TYPE=CLI`

`DRIVER_TYPE=JSON`

`DRIVER_TYPE=PLAIN`

For the `JSON/PLAIN` drivers, the `LOGS_STORAGE_LOCATION` is required. If not provided, it will default to a log file, that depends on the driver.

Just copy the .env.example to .env.