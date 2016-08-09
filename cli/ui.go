package cli

import (
  tablewriter "github.com/olekukonko/tablewriter"
  "os"
  "math"
  "strconv"
)

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

func (c CLI) PrintBoard(board []string) {
  var data [][]string
  b := make([]string, (len(board)))

  for i, value := range board {
    if value == "" {
      b[i] = strconv.Itoa(i)
    } else {
      b[i] = value
    }
  }

  incrementor := int(math.Sqrt(float64(len(b))))
  for i := 0; i < len(b); i += incrementor{
    boardSubSet := b[i:i+incrementor]
    data = append(data, boardSubSet)
  }

  table := tablewriter.NewWriter(os.Stdout)
  table.SetRowLine(true)
  table.SetCenterSeparator("-")
  table.SetColumnSeparator("|")
  table.SetRowSeparator("-")

  table.SetAlignment(tablewriter.ALIGN_CENTER)
  for _, v := range data {
    table.Append(v)
  }
  table.Render()
}

