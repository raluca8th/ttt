package players

import "testing"

func TestName(t *testing.T) {
  humanPlayer := HumanPlayer{name: "Anda"}
  if name := humanPlayer.Name(); name != "Anda" {
    t.Error("Expected name to be Anda, but it was", name)
  }
}

func TestMarker(t *testing.T) {
  humanPlayer := HumanPlayer{name: "Anda", marker: "A"}
  if marker := humanPlayer.Marker(); marker != "A" {
    t.Error("Expected marker to be A, but it was", marker)
  }
}

