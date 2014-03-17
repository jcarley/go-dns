package main

import (
  "bytes"
  "encoding/json"
  "io"
)

const defaultConfig = `
{
  "credentials": {
      "email": "jeff.carley@example.com",
      "token": "ABCDEFGHIJKLMNOP"
  },
  "domains": [
    {
      "name": "example.com",
      "record-type": "A",
      "current-ip": "0.0.0.0"
    },
    {
      "name": "finishfirstsoftware.com",
      "record-type": "A",
      "current-ip": "0.0.0.0"
    }
  ]
}
`

type config struct {
  Credentials map[string]string
  Domains     []map[string]string
}

type Domain struct {
  Name       string
  RecordType string
  CurrentIP  string
}

func loadConfig() (*config, error) {
  var config config
  if err := decodeConfig(bytes.NewBufferString(defaultConfig), &config); err != nil {
    return nil, err
  }
  return &config, nil
}

func decodeConfig(r io.Reader, c *config) error {
  decoder := json.NewDecoder(r)
  return decoder.Decode(c)
}

func (c *config) Email() string {
  return c.Credentials["email"]
}

func (c *config) Token() string {
  return c.Credentials["token"]
}

func (c *config) LoadDomain(name string) Domain {
  domains := c.Domains

  for index := 0; index < len(domains); index++ {
    domain := domains[index]
    if domain["name"] == name {
      d := Domain{Name: domain["name"], RecordType: domain["record-type"], CurrentIP: domain["current-ip"]}
      return d
    }
  }

  return Domain{}
}

func (c *config) SaveDomain(domain *Domain) error {
  return nil
}
