package main

import (
  "testing"
  "strings"
  "bytes"
  "github.com/stretchr/testify/assert"
  "log"
)

func TestMarshallingOfSettings(t *testing.T) {

  domainRecords := []DomainRecord {
    {"domainA", "A"},
    {"domainB", "A"},
    {"domainC", "A"},
  }
  credentials := ApiCredentials{"joe@example.com", "123456789ABCD"}
  routerIP := Router{"76.123.678.64"}

  settings := Settings{credentials, domainRecords, routerIP}

  var b bytes.Buffer

  marshaler := JSONMarshaler{}
  marshaler.MarshalSettings(&b, &settings)

  expectedJson := "{\"Credentials\":{\"Email\":\"joe@example.com\",\"Token\":\"123456789ABCD\"},\"DomainRecords\":[{\"Name\":\"domainA\",\"RecordType\":\"A\"},{\"Name\":\"domainB\",\"RecordType\":\"A\"},{\"Name\":\"domainC\",\"RecordType\":\"A\"}],\"Router\":{\"IP\":\"76.123.678.64\"}}"
  actualJson := strings.Trim(b.String(), "\n")

  assert.Equal(t, actualJson, expectedJson)
}

func TestUnmarshallingOfSettings(t *testing.T) {
  const input = `{"Credentials":{"Email":"joe@example.com","Token":"123456789ABCD"},"DomainRecords":[{"Name":"domainA","RecordType":"A"},{"Name":"domainB","RecordType":"A"},{"Name":"domainC","RecordType":"A"}],"Router":{"IP":"76.123.678.64"}}`
  reader := strings.NewReader(input)

  marshaler := JSONMarshaler{}
  settings, err := marshaler.UnmarshalSettings(reader)

  if err != nil {
    log.Fatal(err)
  }

  assert.NotNil(t, settings)
  assert.Equal(t, settings.Credentials.Email, "joe@example.com")
}

