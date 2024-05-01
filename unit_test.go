package FastMediaInfo

import (
	"testing"
)

// go test -v -run TestGetPrettyInfo
func TestGetPrettyInfo(t *testing.T) {
	result := GetStandMediaInfo("/mnt/e/Movies/bili/杰克威胁taiga/被困循环中的不止2B还有她尼尔短篇小说过于平静的海洋解说.mkv")
	t.Log(result.General)
	t.Log(result.Audio)
	t.Log(result.Video)

}
