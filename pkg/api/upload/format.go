package upload

import "errors"

type Format uint8

const (
	FormatRandom = iota
	FormatUUID
	FormatDate
	FormatName
)

// needed for cobra
type FormatFlag struct {
	Value Format
}

var formatMap = map[string]Format{
	"uuid":   FormatUUID,
	"date":   FormatDate,
	"random": FormatRandom,
	"name":   FormatName,
}

func (f *FormatFlag) Set(s string) error {
	value, exists := formatMap[s]
	if !exists {
		return errors.New("invalid format: " + s)
	}
	f.Value = value
	return nil
}

func (f *FormatFlag) String() string {
	for k, v := range formatMap {
		if v == f.Value {
			return k
		}
	}
	return ""
}

func (f *FormatFlag) Type() string {
	return "format"
}
