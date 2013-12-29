package main

import (
  "io"
  "encoding/json"
)

type CredentialsMarshaler interface {
  MarshalCredentials(writer io.Writer, credentials Credentials) error
}

type CredentialsUnMarshaler interface {
  UnmarshalCredentials(reader io.Reader) (Credentials, error)
}

type JSONMarshaler struct{}

func (JSONMarshaler) MarshalCredentials(writer io.Writer, credentials Credentials) error {
  encoder := json.NewEncoder(writer)
  return encoder.Encode(credentials)
}

func (JSONMarshaler) UnmarshalCredentials(reader io.Reader) (Credentials, error) {
  decoder := json.NewDecoder(reader)
  var credentials *Credentials
  err := decoder.Decode(&credentials)
  return credentials, err
}
