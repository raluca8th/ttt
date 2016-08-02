package cli

import(
  "testing"
  "bytes"
)

type testSTDOUTprinter struct {
  buffer bytes.Buffer
}

func (p *testSTDOUTprinter) Print(s string) {
  p.buffer.WriteString(s)
}

func TestCLIOutput(t *testing.T) {
  testPrinter := new(testSTDOUTprinter)
  output := CLIOutput{Printer: testPrinter}
  output.Print("Print me please")
  if printedString := testPrinter.buffer.String(); printedString != "Print me please" {
    t.Error("Expected printed string to be 'Print me please', but it was", printedString)
  }
}

func ExampleSTDOUTprinter() {
  printer := new(STDOUTprinter)
  printer.Print("Hello")
  // Output: Hello
}
