package settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type UserData struct {
	AppId   int    `json:"AppId"`
	ApiHash string `json:"ApiHash"`
	Phone   string `json:"Phone"`
}

func GetUserData() UserData {
	userJSONFile, err := os.Open("./settings/user.json")
	if err != nil {
		fmt.Println(err)
	}
	defer userJSONFile.Close()

	arrBytes, _ := ioutil.ReadAll(userJSONFile)

	var user UserData

	json.Unmarshal(arrBytes, &user)

	return user
}

type ClockData struct {
	BackgroundColor  string `json:"background_color"`
	TimeColor        string `json:"time_color"`
	OffsetTimeHour   int    `json:"offset_time_hour"`
	OffsetTimeMinute int    `json:"offset_time_minute"`
}

func GetClockData() ClockData {
	userJSONFile, err := os.Open("./settings/clock.json")
	if err != nil {
		fmt.Println(err)
	}
	defer userJSONFile.Close()

	arrBytes, _ := ioutil.ReadAll(userJSONFile)

	var user ClockData

	json.Unmarshal(arrBytes, &user)

	return user
}

func GetPicturePath() string {
	return "clock/currentTime.png"
}
