package cli

import "fmt"

type Printer interface {
  Print(s string)
}

type STDOUTprinter struct {}

func (p *STDOUTprinter) Println(s string){
  fmt.Println(s)
}

func (p *STDOUTprinter) Print(s string){
  fmt.Print(s)
}
