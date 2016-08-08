#Tic Tac Toe
Unbeatable TTT in GO

###Requirements

Go (if you are installing GO now, you need to export the GOROOT and GOPATH)

```brew install go``` or install(https://golang.org/doc/install)

  ```export GOROOT=$HOME/go```
  ```export PATH=$PATH:$GOROOT/bin```

###Setup

1. Clone this repository
3. ``cd ttt``

###Play the game
 ``./ttt``

###Running the Tests

``go test ./...``

Verbose:

``go test -v``

With coverage:

```go test ./... -cover```
