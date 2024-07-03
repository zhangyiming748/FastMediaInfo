package FastMediaInfo

import (
	"testing"
)

// go test -v -run TestGetPrettyInfo
func TestGetPrettyInfo(t *testing.T) {
	result := GetStandMediaInfo("/mnt/c/Users/zen/Pictures/Screenshots/1.png")
	t.Logf("%+v\n", result.General)
	t.Logf("%+v\n", result.Audio)
	t.Logf("%+v\n", result.Video)
	t.Logf("%+v\n", result.Image)
	t.Logf("%+v\n", result.Image.Width)
	t.Logf("%+v\n", result.Image.Height)

}
