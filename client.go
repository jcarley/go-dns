package main

import (
  "encoding/json"
  "fmt"
  "github.com/rubyist/go-dnsimple"
  "io/ioutil"
  "net/http"
  "os"
)

type Router struct {
  IP string
}

const (
  settingsFileName = "./settings.json"
)

func main() {

  // 1) Load config
  config, err := loadConfig(settingsFileName)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Error loading configuration: \n\n%s\n", err)
  }

  // 2) Look up the router's IP
  routerIP, err := getRouterIP()
  if err != nil {
    fmt.Fprintf(os.Stderr, "Error determining router IP: \n\n%s\n", err)
  }

  fmt.Printf("Router IP: %s\n", routerIP)

  // 3) Instantiate a dnsimple client
  client := dnsimple.NewClient(config.Token(), config.Email())

  // 4) Check the current IP of each domain
  domains := config.LoadAllDomains()
  for _, domain := range domains {
    records, _ := client.Records(domain.Name, "", domain.RecordType)
    for _, record := range records {
      if record.RecordType == domain.RecordType {
        if routerIP == record.Content {
          fmt.Printf("DNS for '%s' matches\n", domain.Name)
        } else {
          // 5) If IP's don't match; update DNS record
          fmt.Printf("DNS for '%s' does not match\n", domain.Name)
          fmt.Printf("  Updating '%s' to '%s'\n", record.Content, routerIP)
          record.UpdateIP(client, routerIP)
        }
      }
    }
  }

}

func getRouterIP() (string, error) {

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
