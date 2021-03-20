package timer

import "time"

func GetNowTime() time.Time {
	location, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(location)
}

func GetNowTimeStamp() int64 {
	location, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(location).Unix()
}

func GetCalculateTime(currentTime time.Time, d string) (time.Time, error) {
	location, _ := time.LoadLocation("Asia/Shanghai")
	currentTime.In(location)
	duration, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, err
	}
	return currentTime.Add(duration), nil
}
