package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "os"
  // "github.com/rubyist/go-dnsimple"
  // "log"
  // "io"
)

type Router struct {
  IP string
}

const (
  settingsFileName = "settings.json"
)

func main() {

  config, err := loadConfig()
  if err != nil {
    // fmt.Fprintf(os.Stderr, "Error loading configuration: \n\n%s\n", err)
    fmt.Fprintf(os.Stderr, "%s", err)
  }

  fmt.Println(config.Email())
  fmt.Println(config.Token())

  domain := config.LoadDomain("example.com")
  fmt.Println(domain.Name)
  fmt.Println(domain.RecordType)
  fmt.Println(domain.CurrentIP)

  // client := dnsimple.NewClient(credentials.Token, credentials.Email)

  // records, _ := client.Records(domainRecord.Name, "", domainRecord.RecordType)

  // for _, record := range records {

  // fmt.Printf("Record %d: %s => %s -> %s\n", record.Id, record.RecordType, record.Name, record.Content)

  // if record.RecordType == "A" && record.Content != newIP {
  // record.UpdateIP(client, newIP)
  // fmt.Println("Updated IP")
  // } else {
  // fmt.Println("IP's matched")
  // }
  // }
}

func routerIP() (string, error) {

  url := "http://jsonip.com"
  req, err := http.NewRequest("GET", url, nil)
  req.Header.Set("Content-Type", "application/json")
  req.Header.Add("Accept", "application/json")

  if err != nil {
    return "", err
  }

  client := http.Client{}
  resp, err := client.Do(req)

  if err != nil {
    return "", err
  }

  defer resp.Body.Close()

  responseBytes, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return "", err
  }

  routerIP := Router{}

  err = json.Unmarshal(responseBytes, &routerIP)

  if err != nil {
    return "", err
  }

  return routerIP.IP, nil
}
