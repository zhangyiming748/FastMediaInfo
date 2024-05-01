package FastMediaInfo

type Audio struct {
	Type                 string `json:"@type"`
	Count                string `json:"Count"`
	StreamCount          string `json:"StreamCount"`
	StreamKind           string `json:"StreamKind"`
	StreamKindString     string `json:"StreamKind_String"`
	StreamKindID         string `json:"StreamKindID"`
	Format               string `json:"Format"`
	FormatString         string `json:"Format_String"`
	FormatCommercial     string `json:"Format_Commercial"`
	CodecID              string `json:"CodecID"`
	Duration             string `json:"Duration"`
	DurationString       string `json:"Duration_String"`
	DurationString1      string `json:"Duration_String1"`
	DurationString2      string `json:"Duration_String2"`
	DurationString3      string `json:"Duration_String3"`
	DurationString5      string `json:"Duration_String5"`
	FrameRate            string `json:"FrameRate"`
	FrameRateString      string `json:"FrameRate_String"`
	FrameCount           string `json:"FrameCount"`
	StreamSize           string `json:"StreamSize"`
	StreamSizeString     string `json:"StreamSize_String"`
	StreamSizeString1    string `json:"StreamSize_String1"`
	StreamSizeString2    string `json:"StreamSize_String2"`
	StreamSizeString3    string `json:"StreamSize_String3"`
	StreamSizeString4    string `json:"StreamSize_String4"`
	StreamSizeString5    string `json:"StreamSize_String5"`
	StreamSizeProportion string `json:"StreamSize_Proportion"`
	StreamOrder          string `json:"StreamOrder"`
	ID                   string `json:"ID"`
	IDString             string `json:"ID_String"`
	FormatInfo           string `json:"Format_Info"`
	Extra                struct {
		SourceDelay           string `json:"Source_Delay"`
		SourceDelaySource     string `json:"Source_Delay_Source"`
		CodecConfigurationBox string `json:"CodecConfigurationBox"`
	} `json:"extra"`
	FormatAdditionalFeatures string `json:"Format_AdditionalFeatures"`
	Channels                 string `json:"Channels"`
	ChannelsString           string `json:"Channels_String"`
	ChannelPositions         string `json:"ChannelPositions"`
	ChannelPositionsString2  string `json:"ChannelPositions_String2"`
	ChannelLayout            string `json:"ChannelLayout"`
	SamplesPerFrame          string `json:"SamplesPerFrame"`
	SamplingRate             string `json:"SamplingRate"`
	SamplingRateString       string `json:"SamplingRate_String"`
	SamplingCount            string `json:"SamplingCount"`
	CompressionMode          string `json:"Compression_Mode"`
	CompressionModeString    string `json:"Compression_Mode_String"`
	Default                  string `json:"Default"`
	DefaultString            string `json:"Default_String"`
}
