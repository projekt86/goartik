# Goartik

**Goartik** is Golang library for Samsung ARTIK Modules.



## Futures
- Support ARTIK 520 board (soon all)
- Automatic detection board model
- GPIO (digital read and write)
- Support board pin map

## Goals
- Support all borads model
- Support rest communication protocols I2C etc.
- More example projects


## Installation

```sh
go get -u github.com/projekt86/goartik
```

## Examples

Create new artik module object (automatic model type)
```go
module, err := goartik.NewModule()
	if err != nil {
		log.Fatalln(err.Error())
	}
```
Create new artik module object (manual model type)
```go
module, err := goartik.NewModule(goartik.A520)
	if err != nil {
		log.Fatalln(err.Error())
	}
```
Digital pin mode IN, OUT
```go
module.DigitalPinMode(8, goartik.OUT)
module.DigitalPinMode(7, goartik.IN)
```
Digital pin write HIGH, LOW
```go
module.DigitalWrite(8, goartik.HIGH)
```
Digital pin read HIGH, LOW
```go
module.DigitalRead(8)
```

## License
MIT licensed. See the LICENSE file for details.