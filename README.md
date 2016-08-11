#Tic Tac Toe
Unbeatable TTT in GO

###Requirements

GO (if you are installing GO now, you need to export the GOROOT and GOPATH)

```brew install go``` or [install](https://golang.org/doc/install)

  ```export GOROOT=$HOME/go```

  ```export PATH=$PATH:$GOROOT/bin```

  ```export GOPATH=$HOME/<your path>```

###Setup

1. Clone this repository in your GOPATH above
3. ``cd ttt``

###Play the game
 ``./ttt`` or ```go run main.go```

###Troubleshoot package access

If it complains it can't find some of the github packages run:

 ```go get <package name>``` (e.g. ```go get github.com/raluca8th/ttt/cli```)

###Running the Tests

Short run (Skips the 4X4 board tests. “How did it get so late so soon?” - Dr. Seuss)

```go test ./... -short```

Full run

```go test ./...```

Verbose:

``go test ./... -v -short``

With coverage:

```go test ./... -cover```


###Travis
[Travis builds](https://travis-ci.org/raluca8th/ttt/builds)
