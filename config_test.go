package main

import (
  "bytes"
  "fmt"
  "github.com/stretchr/testify/assert"
  "os"
  "testing"
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

func TestLoadAllDomains(t *testing.T) {
  var config config
  if err := decodeConfig(bytes.NewBufferString(defaultConfig), &config); err != nil {
    fmt.Fprintf(os.Stderr, "%s", err)
  }

  domains := config.LoadAllDomains()

  assert.Equal(t, 2, len(domains))

  for _, d := range domains {
    fmt.Println(d.Name)
  }
}

func TestCredentialsAreReadFromConfig(t *testing.T) {

  var config config
  if err := decodeConfig(bytes.NewBufferString(defaultConfig), &config); err != nil {
    fmt.Fprintf(os.Stderr, "%s", err)
  }

  assert.Equal(t, config.Email(), "jeff.carley@example.com")
  assert.Equal(t, config.Token(), "ABCDEFGHIJKLMNOP")
}

func TestDomainsAreReadFromConfig(t *testing.T) {
  var config config
  if err := decodeConfig(bytes.NewBufferString(defaultConfig), &config); err != nil {
    fmt.Fprintf(os.Stderr, "%s", err)
  }

  domain := config.LoadDomain("example.com")
  assert.Equal(t, domain.Name, "example.com")
  assert.Equal(t, domain.RecordType, "A")
  assert.Equal(t, domain.CurrentIP, "0.0.0.0")

  domain = config.LoadDomain("finishfirstsoftware.com")
  assert.Equal(t, domain.Name, "finishfirstsoftware.com")
  assert.Equal(t, domain.RecordType, "A")
  assert.Equal(t, domain.CurrentIP, "0.0.0.0")
}

func TestSettingsAreSavedToConfig(t *testing.T) {

  var config config
  if err := decodeConfig(bytes.NewBufferString(defaultConfig), &config); err != nil {
    fmt.Fprintf(os.Stderr, "%s", err)
  }

  domain := config.LoadDomain("example.com")
  domain.CurrentIP = "192.168.1.23"
  config.SaveDomain(domain)

  domain = config.LoadDomain("example.com")

  assert.Equal(t, domain.Name, "example.com")
  assert.Equal(t, domain.RecordType, "A")
  assert.Equal(t, domain.CurrentIP, "192.168.1.23")
}

func TestEncodingOfConfig(t *testing.T) {
  var config config
  if err := decodeConfig(bytes.NewBufferString(defaultConfig), &config); err != nil {
    fmt.Fprintf(os.Stderr, "%s", err)
  }

  domain := config.LoadDomain("example.com")
  domain.CurrentIP = "192.168.1.23"
  config.SaveDomain(domain)

  var buffer bytes.Buffer
  if err := encodeConfig(&buffer, &config); err != nil {
    fmt.Fprintf(os.Stderr, "%s", err)
  }

  buffer.WriteTo(os.Stdout)
}
