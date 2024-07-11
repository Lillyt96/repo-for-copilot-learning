package main

import (
	"awesomeProject/internal/logParser"
	"awesomeProject/internal/logger"
	"fmt"
)

func main() {
	//logs, err := logParser.Parse("parse_test_benchmark_data.log")
	//if err != nil {
	//	log.Fatal(err)
	//}

	logs, err := logParser.ParseConcurrently("parse_test_benchmark_data.log")
	if err != nil {
		logger.Default().Fatal(err)
	}

	//fmt.Printf("number of lines: %v \n", len(logs.Logs))

	// task 1 number of unique ip addresses
	uniqueIPs := logs.FindUniqueIPs()

	fmt.Printf("number of unique IPs: %v", len(uniqueIPs))

	// task 3 top 3 active unique addresses
	top3IPs := logs.FindTopNIPs(3)
	fmt.Printf("\ntop 3 IPs: %+v", top3IPs)

	// task 2 top 3 most visited URLs
	top3URLs := logs.FindTopNUrls(3)
	fmt.Printf("\ntop 3 Websites: %+v", top3URLs)
}
