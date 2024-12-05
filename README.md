## Telemetry package

This package is used to collect data and log it on various channels.

The library is configurable via the config.ini file.

Available drivers:

`driver_type=CLI`

`driver_type=JSON`

`driver_type=PLAIN`

For the `JSON/PLAIN` drivers, the `logs_storage_location` is required. If not provided, it will default to a log file, that depends on the driver.
