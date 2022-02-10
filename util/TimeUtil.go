package util

import (
	"log"
	"time"
)

func ParseStringToDateTime(dateTimeStr string) (time.Time, error) {
	layout := "2006-01-02 15:04:05.000"
	dateTime, err := time.Parse(layout, dateTimeStr)
	if err != nil {
		log.Printf("Error converting dateTime: [%s]. Error is %s", dateTimeStr ,err)
		return time.Time{}, err
	}
	return dateTime, nil
}
