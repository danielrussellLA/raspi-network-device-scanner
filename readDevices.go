package main

import (
    "bufio"
    "fmt"
    "os"
    "log"
    "strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func readLines(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}

func getDeviceCount() string {
    lines, err := readLines("./devices.txt")
    if err != nil {
        log.Fatalf("readlines: %s", err)
    }
    deviceCountEntry := lines[len(lines)-2]
    s := strings.Split(deviceCountEntry, " ")
    deviceCount := s[0]
    return deviceCount
}

func main() {
    fmt.Print(getDeviceCount())
}
