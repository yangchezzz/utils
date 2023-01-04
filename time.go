package utils

import "time"

const TimeDefaultTemplate = "2006-01-02 15:03:04"

func ConvertTimeToStrDefault() string {
	return ConvertTimeToStr(time.Now(), TimeDefaultTemplate)
}

func ConvertTimeToStr(t time.Time, template string) string {
	return t.Format(template)
}
