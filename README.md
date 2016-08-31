# goencryption

This repository implements an `encryption-server` making use of built-in Go
libraries.

It also comes with a client interface and a standalone client binary to show
how the server is used.


## Getting Started

The following section will provide you instructions on how to get the source and
install it.

### Prerequisities

The only prerequisite needed is to have already a default go setup ready.

### Installing and Running

1. Use the following command to get the full repository

```
go get github.com/Mr-Bruno/goencryption/...
```

2. Use the following command to install the server

```
go install github.com/Mr-Bruno/goencryption/server
```

3. To get the system running:

3.1 Run the executable in the bin folder

```
server
```

3.2 To verify that the system is running ok introduce the following url in your browser

```
http://localhost:8080/
```

It should show you the message `welcome to go encryption`


## Running the tests

The goencryption repository includes tests to veryfy that the different components
are working. To run these tests execute:

```
go test github.com/Mr-Bruno/goencryption/serverutil
```

## Client usage

The comminucation with the server can be done through the client interface provided.

And standalone client has been provided as an example on how to use this interface.

### Installing/Running the standalone client

1. Install it:
```
go install github.com/Mr-Bruno/goencryption/standaloneclient
```

2. Run the standaloneclient binary created in the binary folder:
```
standaloneclient "123" "MyText"
```


## Deployment

The `encryption-server` code is implemented in three parts/files:
- The http server hanlders.
- The data manipulation functions.
- The database/storage functions.
