package main

import (
  "fmt"
  "github.com/rubyist/go-dnsimple"
  "net/http"
  "io/ioutil"
  "encoding/json"
)

type Credentials struct {
  Email        string
  Token        string
}

type DomainRecord struct {
  Name         string
  RecordType   string
}

type RouterIP struct {
  IP    string
}

func main() {

  credentials := Credentials {
    Email: "email",
    Token: "token",
  }

  domainRecord := DomainRecord {
    Name: "domain",
    RecordType: "A",
  }

  ip, _ := routerIP()
  fmt.Printf("RouterIP: %s\n", ip)

  newIP := "ip"

  client := dnsimple.NewClient(credentials.Token, credentials.Email)

  records, _ := client.Records(domainRecord.Name, "", domainRecord.RecordType)

  for _, record := range records {

    fmt.Printf("Record %d: %s => %s -> %s\n", record.Id, record.RecordType, record.Name, record.Content)

    if record.RecordType == "A" && record.Content != newIP {
      record.UpdateIP(client, newIP)
      fmt.Println("Updated IP")
    } else {
      fmt.Println("IP's matched")
    }
  }
}

func (credentials *Credentials) MarshalJSON(writer io.Writer) error {
  encoder := json.NewEncoder(writer)
  if err := encoder.Encode(credentials); err != nil {
    return err
  }

  return nil
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

  routerIP := RouterIP{}

  err = json.Unmarshal(responseBytes, &routerIP)

  if err != nil {
    return "", err
  }

  return routerIP.IP, nil
}

