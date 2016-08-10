package board

type Board interface {
  AvailableSpots() []int
  SpotIsAvailable(spot int) bool
  FillSpot(spot int)
  WinningMarker() string
  IsBoardSolved() bool
  Size()int
  Markers() [2]string
  Surface() []string
  IsTiedBoard() bool
  NextMarker() string
  ResetSpot(spot int)
  FillAvailableSpot(spot int) Board
}
