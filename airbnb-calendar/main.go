
package main

import (
    "bytes"
    "fmt"
    "github.com/arran4/golang-ical"
    "io/ioutil"
    "net/http"
    "os"
)

// Reservation represents a single reservation with relevant details
type Reservation struct {
    StartDate string
    EndDate   string
    Summary   string
}

func fetchCalendar(url string) ([]byte, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    return ioutil.ReadAll(resp.Body)
}

func parseICal(data []byte) ([]Reservation, error) {
    cal, err := ics.ParseCalendar(bytes.NewReader(data))
    if err != nil {
        return nil, err
    }

    var reservations []Reservation
    for _, event := range cal.Events() {
        reservations = append(reservations, Reservation{
            StartDate: event.GetProperty(ics.ComponentPropertyDtStart).Value,
            EndDate:   event.GetProperty(ics.ComponentPropertyDtEnd).Value,
            Summary:   event.GetProperty(ics.ComponentPropertySummary).Value,
        })
    }
    return reservations, nil
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run main.go <iCal URL>")
        return
    }

    url := os.Args[1]
    data, err := fetchCalendar(url)
    if err != nil {
        fmt.Printf("Failed to fetch calendar: %v\n", err)
        return
    }

    reservations, err := parseICal(data)
    if err != nil {
        fmt.Printf("Failed to parse iCal: %v\n", err)
        return
    }

    fmt.Printf("Reservations: %+v\n", reservations)
}
