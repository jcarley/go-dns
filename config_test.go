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
      "record-type": "A"
    },
    {
      "name": "finishfirstsoftware.com",
      "record-type": "A"
    }
  ]
}
`

func TestLoadAllDomains(t *testing.T) {
  var config Config
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

  var config Config
  if err := decodeConfig(bytes.NewBufferString(defaultConfig), &config); err != nil {
    fmt.Fprintf(os.Stderr, "%s", err)
  }

  assert.Equal(t, config.Email(), "jeff.carley@example.com")
  assert.Equal(t, config.Token(), "ABCDEFGHIJKLMNOP")
}

func TestDomainsAreReadFromConfig(t *testing.T) {
  var config Config
  if err := decodeConfig(bytes.NewBufferString(defaultConfig), &config); err != nil {
    fmt.Fprintf(os.Stderr, "%s", err)
  }

  domain, _ := config.LoadDomain("example.com")
  assert.Equal(t, domain.Name, "example.com")
  assert.Equal(t, domain.RecordType, "A")

  domain, _ = config.LoadDomain("finishfirstsoftware.com")
  assert.Equal(t, domain.Name, "finishfirstsoftware.com")
  assert.Equal(t, domain.RecordType, "A")
}
