package date_utils

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
)

//GetNow ambil tanggal hari ini
func GetNow() time.Time {
	return time.Now().UTC()
}

//GetNowString tampilkan tanggal hari ini
func GetNowString() string {
	//bikin tanggal menjadi string
	return GetNow().Format(apiDateLayout)
}
