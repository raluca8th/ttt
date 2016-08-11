package mocks

import (
  "bytes"
  "strings"
)

type TestSTDIN struct {
  Buffer bytes.Buffer
}

func (r *TestSTDIN) Read() string{
  input, _ := r.Buffer.ReadString(' ')
  return strings.TrimSpace(input)
}

func (p *TestSTDIN) PopulateBuffer(s string){
    p.Buffer.WriteString(s)
}

type TestSTDOUT struct {
  buffer bytes.Buffer
}

func (p *TestSTDOUT) Print(s string) {
  p.buffer.WriteString(s)
}

type TestUI struct{
  Input *TestSTDIN
  Output *TestSTDOUT
}

func (ui TestUI) Read() string{
  return ui.Input.Read()
}

func (ui TestUI) Print(strings ...string) {
  for _, s := range strings {
    ui.Output.Print(s)
  }
}

func (ui *TestUI) CheckOutput() string{
  input, _ := ui.Output.buffer.ReadString('#')
  return input
}

func (ui *TestUI) Populate(s string){
   ui.Input.Buffer.WriteString(s)
}

func (ui TestUI) PrintBoard(board []string){}
