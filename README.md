# FastMediaInfo
统一的mediainfo解析方法
# Quick start
```go
func TestGetPrettyInfo(t *testing.T) {
	result := GetStandMediaInfo("/mnt/e/Movies/bili/杰克威胁taiga/被困循环中的不止2B还有她尼尔短篇小说过于平静的海洋解说.mkv")
	pretty.P(result.General)
	pretty.P(result.Audio)
	pretty.P(result.Video)

}
```