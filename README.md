# carre

------------------------------------------------------------------------------------------------------------------------

A very small tool designed to continuously fetch process information from the inside of a docker container. This tool is
heavily inspired by the concept of [containermon](https://github.com/sparrc/containermon) which provides similar
information but on container basis.

------------------------------------------------------------------------------------------------------------------------

## Install

### Through golang

You may simply install the go project using the following go command:

```shell
go get -u github.com/lynxplay/carre
```

### Release binaries

**Carre** also provides pre-build binaries for finished releases to avoid the requirement for a complete go installation
on every system. The binaries can be found in the respective releases [here](https://github.com/lynxplay/carre/releases)
.

------------------------------------------------------------------------------------------------------------------------

## Usage

### --container | -C (required)

The container parameter specifies the name of the container for which the metrics are to be gathered. It is required as
this is the core point of caree,

### --interval | -I

The interval parameter can specify the duration between each data point collected by caree.

#### Default: 500ms

#### Units: "ns", "us" (or "µs"), "ms", "s", "m", "h"

### --format | -F

The format parameter specifies in which format the output is printed. As of right now the follow formats exist.

- 'CSV' provides a comma separated list of values **<u>including</u>** the csv header defining the collum values.
- 'JSON' provides a set of json data points mapping the respective metric to its value. Each line is an individually
  complete json string.

#### Default: JSON

### --timestamp | -T

The timestamp parameter specifies in which format the timestamp for each data point is going to be printed. As of right
now the following formats exist.

- 'ISO' specifies the time format to follow [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html), more
  specifically [RFC3339](https://tools.ietf.org/html/rfc3339), when printing the timestamps.
- 'EPOCH' specifies the timestamp to follow simple [unix epoch](https://en.wikipedia.org/wiki/Unix_time). The timestamp
  will be represented as **milliseconds**.

#### Default: EPOCH

------------------------------------------------------------------------------------------------------------------------

## Examples

```shell
carre --container my_great_container
{"%CPU":"0.1","%MEM":"0.2","COMMAND":"java -jar quarkus-run.jar","PID":"9809","RSS":"130072","START":"Mar29","STAT":"Ssl","TIME":"0:06","TIMESTAMP":"2021-03-30 02:04:39.4060455 +0200 CEST m=+0.517411801","TTY":"?","USER":"root","VSZ":"19894616"}
{"%CPU":"0.1","%MEM":"0.2","COMMAND":"java -jar quarkus-run.jar","PID":"9809","RSS":"130072","START":"Mar29","STAT":"Ssl","TIME":"0:06","TIMESTAMP":"2021-03-30 02:04:39.9075679 +0200 CEST m=+1.018934101","TTY":"?","USER":"root","VSZ":"19894616"}
{"%CPU":"0.1","%MEM":"0.2","COMMAND":"java -jar quarkus-run.jar","PID":"9809","RSS":"130072","START":"Mar29","STAT":"Ssl","TIME":"0:06","TIMESTAMP":"2021-03-30 02:04:40.4162938 +0200 CEST m=+1.527660401","TTY":"?","USER":"root","VSZ":"19894616"}
{"%CPU":"0.1","%MEM":"0.2","COMMAND":"java -jar quarkus-run.jar","PID":"9809","RSS":"130072","START":"Mar29","STAT":"Ssl","TIME":"0:06","TIMESTAMP":"2021-03-30 02:04:40.9075311 +0200 CEST m=+2.018897201","TTY":"?","USER":"root","VSZ":"19894616"}
{"%CPU":"0.1","%MEM":"0.2","COMMAND":"java -jar quarkus-run.jar","PID":"9809","RSS":"130072","START":"Mar29","STAT":"Ssl","TIME":"0:06","TIMESTAMP":"2021-03-30 02:04:41.4087445 +0200 CEST m=+2.520110701","TTY":"?","USER":"root","VSZ":"19894616"}
{"%CPU":"0.1","%MEM":"0.2","COMMAND":"java -jar quarkus-run.jar","PID":"9809","RSS":"130072","START":"Mar29","STAT":"Ssl","TIME":"0:06","TIMESTAMP":"2021-03-30 02:04:41.9082418 +0200 CEST m=+3.019608001","TTY":"?","USER":"root","VSZ":"19894616"}
{"%CPU":"0.1","%MEM":"0.2","COMMAND":"java -jar quarkus-run.jar","PID":"9809","RSS":"130072","START":"Mar29","STAT":"Ssl","TIME":"0:06","TIMESTAMP":"2021-03-30 02:04:42.4036635 +0200 CEST m=+3.515029901","TTY":"?","USER":"root","VSZ":"19894616"}
{"%CPU":"0.1","%MEM":"0.2","COMMAND":"java -jar quarkus-run.jar","PID":"9809","RSS":"130072","START":"Mar29","STAT":"Ssl","TIME":"0:06","TIMESTAMP":"2021-03-30 02:04:42.9070597 +0200 CEST m=+4.018426101","TTY":"?","USER":"root","VSZ":"19894616"}
{"%CPU":"0.1","%MEM":"0.2","COMMAND":"java -jar quarkus-run.jar","PID":"9809","RSS":"130072","START":"Mar29","STAT":"Ssl","TIME":"0:06","TIMESTAMP":"2021-03-30 02:04:43.4165144 +0200 CEST m=+4.527880701","TTY":"?","USER":"root","VSZ":"19894616"}
{"%CPU":"0.1","%MEM":"0.2","COMMAND":"java -jar quarkus-run.jar","PID":"9809","RSS":"130072","START":"Mar29","STAT":"Ssl","TIME":"0:06","TIMESTAMP":"2021-03-30 02:04:43.9088428 +0200 CEST m=+5.020209001","TTY":"?","USER":"root","VSZ":"19894616"}
{"%CPU":"0.1","%MEM":"0.2","COMMAND":"java -jar quarkus-run.jar","PID":"9809","RSS":"130072","START":"Mar29","STAT":"Ssl","TIME":"0:06","TIMESTAMP":"2021-03-30 02:04:44.4091147 +0200 CEST m=+5.520481001","TTY":"?","USER":"root","VSZ":"19894616"}
```




