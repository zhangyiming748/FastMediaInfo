package FastMediaInfo

/*
codec 可选 h265 vp9 av1 avif
*/
func GetCRF(codec string, width, height int) (crf string) {
	switch codec {
	case "h265":
		if width >= 1080 || height >= 1080 {
			crf = "22"
		} else if width >= 720 || height >= 720 {
			crf = "23"
		} else if width >= 480 || height >= 480 {
			crf = "24"
		}
	case "vp9":
		if width >= 1080 || height >= 1080 {
			crf = "31"
		} else if width >= 720 || height >= 720 {
			crf = "32"
		} else if width >= 480 || height >= 480 {
			crf = "33"
		}
	case "av1":
		if width >= 2160 || height >= 2160 {
			crf = "30"
		} else if width >= 1440 || height >= 1440 {
			crf = "31"
		} else if width >= 1080 || height >= 1080 {
			crf = "32"
		} else if width >= 720 || height >= 720 {
			crf = "33"
		} else if width >= 480 || height >= 480 {
			crf = "34"
		}
	case "avif":
		if width >= 2160 || height >= 2160 {
			crf = "30"
		} else if width >= 1440 || height >= 1440 {
			crf = "31"
		} else if width >= 1080 || height >= 1080 {
			crf = "32"
		} else if width >= 720 || height >= 720 {
			crf = "33"
		} else if width >= 480 || height >= 480 {
			crf = "34"
		}
	default:
		panic("不是合法的编码")
	}
	return crf
}
