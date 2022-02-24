# rpihardware [![GoDoc](https://godoc.org/github.com/DerLukas15/rpihardware?status.svg)](http://godoc.org/github.com/DerLukas15/rpihardware)
Package for determining the used Raspberry Pi hardware with corresponding information during runtime.

## About

This package determines the running Raspberry Pi hardware by checking either `/proc/cpuinfo` or `/proc/device-tree/system/linux,revision`.

It returns information about:
* hardware version
* Pi model type
* oscillator frequency
* base address for 'physical memory'
* description of the model

The list is inspired by [github.com/jgarff/rpi_ws281x](https://github.com/jgarff/rpi_ws281x).

## Installation

```sh
go get github.com/DerLukas15/rpihardware
```

## Usage

Import the package

```go
import ("github.com/DerLukas15/rpihardware")
```

Get the struct containing the information about the hardware:
```go
curHardware, err := rpihardware.Check()
```


## License

MIT, see [LICENSE](LICENSE)
