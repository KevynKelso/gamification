// The goal of this program is to collect data from the oura ring api to 
// integrate into the game.
// Data collected is currently: 
// - sleep scores (sleep scores over 85 will contribute 1 xp, and over 90 will
// contribute 5xp)
// - readiness scores (same conventions as sleep)
// - activity scores (reaching daily goals will contribute 3xp)
package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"

    "github.com/joho/godotenv"
)

type levelData struct {
    SleepScore     int
    ActivityScore  int
    ReadinessScore int
}

func getOuraRequestData(route, ouraToken string) (data SleepData) {
    url := "https://api.ouraring.com/v1/" + route
    bearer := "Bearer " + ouraToken

    req, err := http.NewRequest("GET", url, nil)
    req.Header.Add("Authorization", bearer)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Printf("Error on response.\n[ERRO] - %v", err)
    }

    body, _ :=  ioutil.ReadAll(resp.Body)

    json.Unmarshal([]byte(body), &data)

    return data
}

func init() {
    if err := godotenv.Load(); err != nil {
        fmt.Print("No .env file found\n")
    }
}

func main() {
    ouraToken, _ := os.LookupEnv("OURA_TOKEN")
    sleepData := getOuraRequestData("sleep", ouraToken)

    fmt.Printf("Sleep Score: %v\nDate: %v",
        sleepData.Nights[len(sleepData.Nights)-1].Score,
        sleepData.Nights[len(sleepData.Nights)-1].Summary_date,
    )
}
