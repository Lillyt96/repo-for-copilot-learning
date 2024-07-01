package logParser

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

type Log struct {
	Ip     string
	User   string
	Time   string
	Method string
	URL    string
	Status string
}

type Logs struct {
	logs []Log
}

func Parse(path string) (*Logs, error) {
	lines, err := readLines(path)
	if err != nil {
		return nil, err
	}

	var logs []Log
	for _, line := range lines {
		extractedLog := extractData(line)

		if extractedLog != nil {
			logs = append(logs, *extractedLog)
		}
	}

	return &Logs{logs}, nil
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func extractData(logString string) *Log {
	apacheLogRegexStr := "^(\\S*).*\\[(.*)\\]\\s\"(\\S*)\\s(\\S*)\\s([^\"]*)\"\\s(\\S*)\\s(\\S*)\\s\"([^\"]*)\"\\s\"([^\"]*)\"$"

	apacheLogRegex := regexp.MustCompile(apacheLogRegexStr)

	logResults := apacheLogRegex.FindAllStringSubmatch(logString, -1)

	if len(logResults) == 0 {
		fmt.Println("skipping logResult file due to incorrect format: " + logString)

		return nil
	}

	var logResult Log
	for _, result := range logResults {
		logResult.Ip = result[1]
		logResult.Method = result[3]
		logResult.URL = result[4]
		logResult.Time = result[2]
		logResult.Status = result[6]
	}

	return &logResult

	//logArray := strings.Split(logString, " ")
	//
	//// extract time
	//timeString := strings.Join(logArray[3:5], " ")
	//timeString = strings.ReplaceAll(timeString, "[", "")
	//timeString = strings.ReplaceAll(timeString, "]", "")
	//
	//return Log{
	//	Ip:     logArray[0],
	//	Time:   timeString,
	//	Status: logArray[8],
	//	URL:    logArray[6],
	//}

}
