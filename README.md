# Hero-notify

## About

*hero-notify* is used by herodote to send job status to herodote server

## License

Apache 2.0, see LICENSE file.

## Running

    go run hero-notify.go --server URL --job JOBID --status JOBSTATUS

Example:

    go run hero-notify.go --server https://my.herodote.org --job my_project_name.1556533477920 --status 2