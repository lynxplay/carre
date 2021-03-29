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

### -container (required)

The container parameter specifies the name of the container for which the metrics are to be gathered. It is required as
this is the core point of caree,

### -interval

The interval parameter can specify the duration between each data point collected by caree.

#### Default: 500ms

#### Units: "ns", "us" (or "µs"), "ms", "s", "m", "h"

### -format

The format parameter specifies in which format the output is printed. As of right now the follow formats exist.

- 'CSV' provides a comma separated list of values **<u>including</u>** the csv header defining the collum values.
- 'JSON' provides a set of json data points mapping the respective metric to its value. Each line is an individually
  complete json string.

#### Default: JSON

------------------------------------------------------------------------------------------------------------------------

## Examples

```shell
carre --container my_great_container
{"%CPU":"0.1","%MEM":"0.2","COMMAND":"java -jar application.jar","PID":"9809","RSS":"130072","START":"22:20","STAT":"Ssl","TIME":"0:04","TTY":"?","USER":"root","VSZ":"19894616"}
{"%CPU":"0.1","%MEM":"0.2","COMMAND":"java -jar application.jar","PID":"9809","RSS":"130072","START":"22:20","STAT":"Ssl","TIME":"0:04","TTY":"?","USER":"root","VSZ":"19894616"}
{"%CPU":"0.1","%MEM":"0.2","COMMAND":"java -jar application.jar","PID":"9809","RSS":"130072","START":"22:20","STAT":"Ssl","TIME":"0:04","TTY":"?","USER":"root","VSZ":"19894616"}
{"%CPU":"0.1","%MEM":"0.2","COMMAND":"java -jar application.jar","PID":"9809","RSS":"130072","START":"22:20","STAT":"Ssl","TIME":"0:04","TTY":"?","USER":"root","VSZ":"19894616"}
{"%CPU":"0.1","%MEM":"0.2","COMMAND":"java -jar application.jar","PID":"9809","RSS":"130072","START":"22:20","STAT":"Ssl","TIME":"0:04","TTY":"?","USER":"root","VSZ":"19894616"}
{"%CPU":"0.1","%MEM":"0.2","COMMAND":"java -jar application.jar","PID":"9809","RSS":"130072","START":"22:20","STAT":"Ssl","TIME":"0:04","TTY":"?","USER":"root","VSZ":"19894616"}
{"%CPU":"0.1","%MEM":"0.2","COMMAND":"java -jar application.jar","PID":"9809","RSS":"130072","START":"22:20","STAT":"Ssl","TIME":"0:04","TTY":"?","USER":"root","VSZ":"19894616"}
{"%CPU":"0.1","%MEM":"0.2","COMMAND":"java -jar application.jar","PID":"9809","RSS":"130072","START":"22:20","STAT":"Ssl","TIME":"0:04","TTY":"?","USER":"root","VSZ":"19894616"}
{"%CPU":"0.1","%MEM":"0.2","COMMAND":"java -jar application.jar","PID":"9809","RSS":"130072","START":"22:20","STAT":"Ssl","TIME":"0:04","TTY":"?","USER":"root","VSZ":"19894616"}
{"%CPU":"0.1","%MEM":"0.2","COMMAND":"java -jar application.jar","PID":"9809","RSS":"130072","START":"22:20","STAT":"Ssl","TIME":"0:04","TTY":"?","USER":"root","VSZ":"19894616"}
{"%CPU":"0.1","%MEM":"0.2","COMMAND":"java -jar application.jar","PID":"9809","RSS":"130072","START":"22:20","STAT":"Ssl","TIME":"0:04","TTY":"?","USER":"root","VSZ":"19894616"}
{"%CPU":"0.1","%MEM":"0.2","COMMAND":"java -jar application.jar","PID":"9809","RSS":"130072","START":"22:20","STAT":"Ssl","TIME":"0:04","TTY":"?","USER":"root","VSZ":"19894616"}
```




