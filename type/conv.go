package typetools

import (
    "strconv"
    "time"
)

const (
    //TimePattern = "02-01-2006 15:04:05"
    TimePattern = "02-01-2006"

    Past   = "01-01-2000"
    Future = "31-12-2099"
)

var DatePast, _ = time.ParseInLocation(TimePattern, Past, time.Local)
var DateFuture, _ = time.ParseInLocation(TimePattern, Future, time.Local)

func DateNow() time.Time {
    year, month, day := time.Now().Date()
    return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
}
func DateTomorrow() time.Time {
    return DateNow().AddDate(0, 0, 1)
}

func ToInt(v string, d int64) (i int64) {
    i, e := strconv.ParseInt(v, 10, 64)
    if e != nil {
        i = d
    }
    return
}
func ToFloat(v string, d float64) (f float64) {
    f, e := strconv.ParseFloat(v, 64)
    if e != nil {
        f = d
    }
    return
}
func ToBool(v string, d bool) (b bool) {
    b, e := strconv.ParseBool(v)
    if e != nil {
        b = d
    }
    return
}
func ToTime(v string, d string) (t time.Time) {
    t, e := time.ParseInLocation(TimePattern, v, time.Local)
    if e != nil {
        t, _ = time.ParseInLocation(TimePattern, d, time.Local)
    }
    return
}
func ToString(v string, d string) (t string) {
    t = v
    if t == "" {
        t = d
    }
    return
}