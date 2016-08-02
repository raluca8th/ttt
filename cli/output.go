package cli

import "fmt"

type Printer interface {
  Print(s string)
}

type STDOUTprinter struct {}

func (p *STDOUTprinter) Print(s string){
  fmt.Println(s)
}

type CLIOutput struct {
  printer Printer
}

func (o *CLIOutput) Print(s string){
  o.printer.Print(s)
}
