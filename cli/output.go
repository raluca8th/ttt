package cli

import "fmt"

type Printer interface {
  Print(s string)
}

type STDOUTprinter struct {}

func (p *STDOUTprinter) Print(s string){
  fmt.Println(s)
}
