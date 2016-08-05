package ui

type UI interface{
  Read() string
  Print(i ...string)
}
