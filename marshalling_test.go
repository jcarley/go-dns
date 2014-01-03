package main

import (
  // "fmt"
  // "encoding/json"
  // "testing/iotest"
  "testing"
  "strings"
  "bytes"
  "github.com/stretchr/testify/assert"
)

func TestMarshallingOfSettings(t *testing.T) {

  domainRecords := []DomainRecord {
    {"domainA", "A"},
    {"domainB", "A"},
    {"domainC", "A"},
  }
  credentials := Credentials{"joe@example.com", "123456789ABCD"}
  routerIP := RouterIP{"76.123.678.64"}

  settings := Settings{credentials, domainRecords, routerIP}

  var b bytes.Buffer

  marshaler := JSONMarshaler{}
  marshaler.MarshalSettings(&b, &settings)

  expectedJson := "{\"Credentials\":{\"Email\":\"joe@example.com\",\"Token\":\"123456789ABCD\"},\"DomainRecords\":[{\"Name\":\"domainA\",\"RecordType\":\"A\"},{\"Name\":\"domainB\",\"RecordType\":\"A\"},{\"Name\":\"domainC\",\"RecordType\":\"A\"}],\"RouterIP\":{\"IP\":\"76.123.678.64\"}}"
  actualJson := strings.Trim(b.String(), "\n")

  assert.Equal(t, actualJson, expectedJson)
}

func TestUnmarshallingOfSettings(t *testing.T) {
  input := `"{\"Credentials":{"Email":"joe@example.com","Token":"123456789ABCD"},"DomainRecords":[{"Name":"domainA","RecordType":"A"},{"Name":"domainB","RecordType":"A"},{"Name":"domainC","RecordType":"A"}],"RouterIP":{"IP":"76.123.678.64"}}"`
  reader := strings.NewReader(input)

  marshaler := JSONMarshaler{}
  settings, _ := marshaler.UnmarshalSettings(reader)

  assert.NotNil(t, settings)
  assert.Equal(t, &settings.Credentials.Email, "joe@example.com")
}

// func TestUnmarshallingOfCredentials(t *testing.T) {
  // input := `{"Email": "joe@example.com", "Token": "123456789ABCD"}`
  // reader := strings.NewReader(input)

  // marshaler := JSONMarshaler{}
  // credentials, _ := marshaler.UnmarshalCredentials(reader)

  // assert.Equal(t, credentials.Email, "joe@example.com")
  // assert.Equal(t, credentials.Token, "123456789ABCD")
// }

// func TestMarshallingOfCredentials(t *testing.T) {
  // credentials := Credentials{"joe@example.com", "123456789ABCD"}

  // var b bytes.Buffer

  // marshaler := JSONMarshaler{}
  // marshaler.MarshalCredentials(&b, credentials)

  // actualJson := strings.Trim(b.String(), "\n")
  // expectedJson := "{\"Email\":\"joe@example.com\",\"Token\":\"123456789ABCD\"}"

  // assert.Equal(t, expectedJson, actualJson)
// }

// func TestUnmarshallingOfDomainRecords(t *testing.T) {
  // input := `[{"Name":"domainA","RecordType":"A"},{"Name":"domainB","RecordType":"A"},{"Name":"domainC","RecordType":"A"}]"`
  // reader := strings.NewReader(input)

  // marshaler := JSONMarshaler{}
  // domainRecords, _ := marshaler.UnmarshalDomainRecords(reader)

  // domainRecord := domainRecords[0]

  // assert.Equal(t, len(domainRecords), 3)
  // assert.Equal(t, domainRecord.Name, "domainA")
  // assert.Equal(t, domainRecord.RecordType, "A")
// }

// func TestMarshallingOfDomainRecords(t *testing.T) {
  // domainRecords := []DomainRecord {
    // {"domainA", "A"},
    // {"domainB", "A"},
    // {"domainC", "A"},
  // }

  // var b bytes.Buffer

  // marshaler := JSONMarshaler{}
  // marshaler.MarshalDomainRecords(&b, domainRecords)

  // actualJson := strings.Trim(b.String(), "\n")
  // expectedJson := "[{\"Name\":\"domainA\",\"RecordType\":\"A\"},{\"Name\":\"domainB\",\"RecordType\":\"A\"},{\"Name\":\"domainC\",\"RecordType\":\"A\"}]"

  // assert.Equal(t, expectedJson, actualJson)
// }



