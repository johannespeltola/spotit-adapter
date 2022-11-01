package config

const dateTimeLayout = "2006-01-02 15:04:05"
const dateLayout = "2006-01-02"
const timeOnlyLayout = "15:04"
const entsoeDateLayout = "20060102"

func GetDateFormat() string {
	return dateLayout
}

func GetDateTimeFormat() string {
	return dateTimeLayout
}

func GetTimeOnlyFormat() string {
	return timeOnlyLayout
}

func GetEntsoeDateFormat() string {
	return entsoeDateLayout
}

type TimeConf struct {
	DateFormat     string
	DateTimeFormat string
	TimeOnlyFormat string
}
