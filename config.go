package main

import (
  "bytes"
  "encoding/json"
  "io"
  "os"
)

type config struct {
  Credentials map[string]string
  Domains     []map[string]string
}

type Domain struct {
  Name       string
  RecordType string
  CurrentIP  string
}

func loadConfig(configFilePath string) (*config, error) {
  var c *config

  f, err := os.Open(configFilePath)
  if err != nil {
    if !os.IsNotExist(err) {
      return nil, err
    }
    return nil, err
  }
  defer f.Close()

  if err := decodeConfig(f, c); err != nil {
    return nil, err
  }

  return c, nil
}

func decodeConfig(r io.Reader, c *config) error {
  decoder := json.NewDecoder(r)
  return decoder.Decode(c)
}

func saveConfig(configFilePath string, c *config) error {
  f, err := os.Create(configFilePath)
  if err != nil {
    return err
  }
  defer f.Close()

  if err := encodeConfig(f, c); err != nil {
    return err
  }

  return nil
}

func encodeConfig(w io.Writer, c *config) error {
  b, err := json.MarshalIndent(c, "", "  ")
  if err != nil {
    return err
  }

  buffer := bytes.NewBuffer(b)
  _, err = buffer.WriteTo(w)
  if err != nil {
    return err
  }

  return nil
}

func (c *config) Email() string {
  return c.Credentials["email"]
}

func (c *config) Token() string {
  return c.Credentials["token"]
}

func (c *config) LoadAllDomains() []Domain {
  domains := make([]Domain, 2, 2)

  for index := 0; index < len(c.Domains); index++ {
    domain := c.Domains[index]
    d := Domain{Name: domain["name"], RecordType: domain["record-type"], CurrentIP: domain["current-ip"]}
    domains[index] = d
  }

  return domains
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

func (c *config) SaveDomain(domain Domain) {
  domains := c.Domains

  for index := 0; index < len(domains); index++ {
    if domains[index]["name"] == domain.Name {
      d := make(map[string]string)
      d["name"] = domain.Name
      d["record-type"] = domain.RecordType
      d["current-ip"] = domain.CurrentIP
      c.Domains[index] = d
    }
  }
}
