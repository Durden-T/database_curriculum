package request

type IntervalType int

const (
	HourInterval IntervalType = iota
	MinuteInterval
)

type Chart struct {
	Type      string       `json:"type" form:"type"`
	Name      string       `json:"name" form:"name"`
	StartTime string       `json:"startTime" form:"startTime"`
	EndTime   string       `json:"endTime" form:"endTime"`
	Interval  IntervalType `json:"interval" form: "interval"`
}
