package shp

import (
	"bytes"
	"strconv"
	"strings"

	"golang.org/x/text/encoding"
)

type Attribute interface {
	Name() string
	Value() interface{}
}

type CharacterAttribute struct {
	name  string
	value string
}

func decodeCharacter(buf []byte, name string, decoder *encoding.Decoder) (*CharacterAttribute, error) {
	val := bytes.Trim(buf, "\x00")

	decoded, err := decoder.Bytes(val)
	if err != nil {
		return nil, err
	}

	return &CharacterAttribute{
		name:  name,
		value: strings.TrimSpace(string(decoded)),
	}, nil
}

func (a *CharacterAttribute) Name() string {
	return a.name
}

func (a *CharacterAttribute) Value() interface{} {
	return a.value
}

type NumericAttribute struct {
	name  string
	value float64
}

func decodeNumeric(buf []byte, name string) (*NumericAttribute, error) {
	val := bytes.Trim(buf, "\x20") // trim spaces
	num, err := strconv.ParseFloat(string(val), 0)
	if err != nil {
		return nil, err
	}

	return &NumericAttribute{
		name:  name,
		value: num,
	}, nil
}

func (a *NumericAttribute) Name() string {
	return a.name
}

func (a *NumericAttribute) Value() interface{} {
	return a.value
}

type FloatAttribute NumericAttribute

type DateAttribute struct {
	name  string
	value string
}

func decodeDate(buf []byte, name string) (*DateAttribute, error) {
	val := bytes.Trim(buf, "\x00")

	return &DateAttribute{
		name: name,
		// TODO: parse value to dbf Date
		value: string(val),
	}, nil
}

func (a *DateAttribute) Name() string {
	return a.name
}

func (a *DateAttribute) Value() interface{} {
	return a.value
}

type RawAttribute struct {
	name  string
	value []byte
}

func decodeRaw(buf []byte, name string) (*RawAttribute, error) {
	return &RawAttribute{
		name:  name,
		value: buf,
	}, nil
}

func (a *RawAttribute) Name() string {
	return a.name
}

func (a *RawAttribute) Value() interface{} {
	return a.value
}
