# Convert Go structures into Protobuf messages.

## Installation

```shell
go get -u github.com/ekhabarov/stp
```

## Example

Go struct from [example/example.go](./example/example.go)
```go
type SomeStruct struct {
	ID      int
	Name    string
	Address string
	Status  string
}
```
command

```shell
stp -f $GOPATH/src/github.com/ekhabarov/stp/example/example.go -s SomeStruct
```

prints

```shell
message SomeStruct {
        int id = 1;
        string name = 2;
        string address = 3;
        string status = 4;
}
```

## Usage
```shell
Usage of stp:
  -f string
    /path/to/file.go with structures
  -s string
    Structure name to parse
```
