package timeUtil

import (
	"fmt"
	"github.com/juju/errors"
	"strconv"
	"time"
)

func GetLogFimePrefix() string {
	// return time.Now().Format("[2006-01-02 15:04:05]")
	return "[" + time.Now().Format(time.RFC3339) + "]"

}

func GetUnixtimeFromStr(time_str string) (unixtime int64, err error) {
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	theTime, err := time.ParseInLocation(timeLayout, time_str, loc)
	if err != nil {
		return 0, errors.New("Error when GetUnixtimeFromStr:" + err.Error())
	}
	sr := theTime.Unix()
	return sr, nil
}
func GetYearWeekFromStr(time_str string) (yearweek string, err error) {
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	theTime, err := time.ParseInLocation(timeLayout, time_str, loc)
	if err != nil {
		return "", errors.New("Error when GetUnixtimeFromStr:" + err.Error())
	}
	year, week := theTime.ISOWeek()
	return strconv.Itoa(year*100 + week), nil
}
func StrUnixtime2normal(strunixtime string) string {
	if strunixtime == "0" || len(strunixtime) == 0 {
		return ""
	}
	unixtime64, err := strconv.ParseInt(strunixtime, 10, 64)
	if err != nil {
		return err.Error()
	}
	unixtime := time.Unix(unixtime64, 0)
	if unixtime.Format("2006年01月02日 15:04:05") == "1970年01月01日 08:00:00" {
		return ""
	}
	return unixtime.Format("2006年01月02日 15:04:05")
}
func StrMillisec2Normal(ms string) string {
	msInt, err := strconv.ParseInt(ms, 10, 64)
	if err != nil {
		fmt.Println("StrMillisec2Normal=ParseIntError:" + ms + "|" + err.Error())
		return ""
	}
	tm := time.Unix(0, msInt*int64(time.Millisecond)).Format("2006-01-02 15:04:05")
	return tm
}
