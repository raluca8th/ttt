#Tic Tac Toe
Unbeatable TTT in Go

###Requirements

Go version 1.6.3


If you are installing Go now, you need to export the GOPATH

  ```brew install go``` OR [install](https://golang.org/doc/install)

  ```mkdir $HOME/go```

  ```export GOPATH=$HOME/go```

###Setup

1. ```cd $HOME/go```
2.  Clone this repository in the GOPATH
3. ``cd ttt``

###Play the game
 ``./ttt`` OR ```go run main.go```

###Troubleshoot package access

If it complains it can't find some of the github packages run:

 ```go get <package name>``` (e.g. ```go get github.com/raluca8th/ttt/cli```)

*note - you need ```git``` installed to run the ```go get`` commands

###Running the Tests

Short run (Skips the 4X4 board tests which can take up to 7 minutes for the computer player. “How did it get so late so soon?” - Dr. Seuss)

```go test ./... -short```

Full run

```go test ./...```

Verbose:

``go test ./... -v -short``

With coverage:

```go test ./... -cover```


###Travis
[Travis builds](https://travis-ci.org/raluca8th/ttt/builds)
