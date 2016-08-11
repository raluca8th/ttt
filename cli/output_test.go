package cli

func ExamplePrint() {
  printer := new(STDOUTprinter)
  printer.Print("Hello")
  // Output: Hello
}
