package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(time string) string {
	// Write your code here
	sTime := strings.Split(time, ":")
	h := sTime[0]
	m := sTime[1]
	s := sTime[2][:2]

	isPM := strings.Contains(time, "PM")

	hour := 0
	fmt.Sscanf(h, "%d", &hour)

	if isPM && hour != 12 {
		hour += 12
	} else if !isPM && hour == 12 {
		hour = 0
	}

	return fmt.Sprintf("%02d:%s:%s", hour, m, s)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
