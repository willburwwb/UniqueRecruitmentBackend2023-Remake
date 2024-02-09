package utils

var name = map[string]string{
	"S": "春季招新",
	"C": "夏令营招新",
	"A": "秋季招新",
	"春": "春季招新",
	"夏": "夏令营招新",
	"秋": "秋季招新",
}

func ConvertRecruitmentName(title string) string {
	year := title[:4]
	suffix := name[title[4:]]
	if suffix == "" {
		return title
	}
	return year + suffix
}
