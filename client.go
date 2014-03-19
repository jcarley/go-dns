package main

import (
  "encoding/json"
  "fmt"
  // "github.com/rubyist/go-dnsimple"
  "io/ioutil"
  "net/http"
  "os"
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

  config, err := loadConfig(settingsFileName)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Error loading configuration: \n\n%s\n", err)
  }

  d := config.LoadDomain("example.com")
  fmt.Println(d.Name)

  // client := dnsimple.NewClient(config.Token(), config.Email())

  domains := config.LoadAllDomains()

  for index := 0; index < len(domains); index++ {
    fmt.Println(domains[index].Name)
  }

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
