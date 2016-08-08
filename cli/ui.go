package cli

type CLI struct{
  input *STDINReader
  output *STDOUTprinter
}

func (c CLI) Read() string{
  return c.input.Read()
}

func (c CLI) Print(strings ...string){
  for _, s := range strings {
    c.output.Print(s)
  }
}
