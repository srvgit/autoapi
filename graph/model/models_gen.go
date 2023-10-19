// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type ServerConfig struct {
	ID               string  `json:"id"`
	GraphPackagePath string  `json:"graphPackagePath"`
	PlaygroundPath   string  `json:"playgroundPath"`
	QueryPath        string  `json:"queryPath"`
	GinMode          GinMode `json:"ginMode"`
	Port             int     `json:"port"`
}

type ServerConfigInput struct {
	GraphPackagePath string `json:"graphPackagePath"`
	PlaygroundPath   string `json:"playgroundPath"`
	QueryPath        string `json:"queryPath"`
	GinMode          string `json:"ginMode"`
	Port             int    `json:"port"`
}

type GinMode string

const (
	GinModeDebug   GinMode = "DEBUG"
	GinModeRelease GinMode = "RELEASE"
	GinModeTest    GinMode = "TEST"
)

var AllGinMode = []GinMode{
	GinModeDebug,
	GinModeRelease,
	GinModeTest,
}

func (e GinMode) IsValid() bool {
	switch e {
	case GinModeDebug, GinModeRelease, GinModeTest:
		return true
	}
	return false
}

func (e GinMode) String() string {
	return string(e)
}

func (e *GinMode) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = GinMode(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid GinMode", str)
	}
	return nil
}

func (e GinMode) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
