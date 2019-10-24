package shp

import (
	"time"

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
	return nil, nil
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
	return nil, nil
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
	value time.Time
}

func decodeDate(buf []byte, name string) (*DateAttribute, error) {
	return nil, nil
}

func (a *DateAttribute) Name() string {
	return a.name
}

func (a *DateAttribute) Value() interface{} {
	return a.value
}
