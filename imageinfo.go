package FastMediaInfo

type Image struct {
	Type                  string `json:"@type"`
	Count                 string `json:"Count"`
	StreamCount           string `json:"StreamCount"`
	StreamKind            string `json:"StreamKind"`
	StreamKindString      string `json:"StreamKind_String"`
	StreamKindID          string `json:"StreamKindID"`
	Format                string `json:"Format"`
	FormatString          string `json:"Format_String"`
	FormatInfo            string `json:"Format_Info"`
	FormatCommercial      string `json:"Format_Commercial"`
	FormatCompression     string `json:"Format_Compression"`
	InternetMediaType     string `json:"InternetMediaType"`
	Width                 string `json:"Width"`
	WidthString           string `json:"Width_String"`
	Height                string `json:"Height"`
	HeightString          string `json:"Height_String"`
	ColorSpace            string `json:"ColorSpace"`
	BitDepth              string `json:"BitDepth"`
	BitDepthString        string `json:"BitDepth_String"`
	CompressionMode       string `json:"Compression_Mode"`
	CompressionModeString string `json:"Compression_Mode_String"`
	StreamSize            string `json:"StreamSize"`
	StreamSizeString      string `json:"StreamSize_String"`
	StreamSizeString1     string `json:"StreamSize_String1"`
	StreamSizeString2     string `json:"StreamSize_String2"`
	StreamSizeString3     string `json:"StreamSize_String3"`
	StreamSizeString4     string `json:"StreamSize_String4"`
	StreamSizeString5     string `json:"StreamSize_String5"`
	StreamSizeProportion  string `json:"StreamSize_Proportion"`
}
