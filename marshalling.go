package main

import (
  "encoding/json"
  "io"
)

type SettingsMarshaler interface {
  MarshalSettings(writer io.Writer, settings *Settings) error
}

type SettingsUnMarshaler interface {
  UnmarshalSettings(reader io.Reader) (*Settings, error)
}

type JSONMarshaler struct{}

func (JSONMarshaler) MarshalSettings(writer io.Writer, settings *Settings) error {
  encoder := json.NewEncoder(writer)
  return encoder.Encode(settings)
}

func (JSONMarshaler) UnmarshalSettings(reader io.Reader) (*Settings, error) {
  decoder := json.NewDecoder(reader)
  var settings *Settings
  err := decoder.Decode(&settings)
  return settings, err
}
