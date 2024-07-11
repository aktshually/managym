package utils

import (
	"database/sql/driver"
)

type TestRequestData struct {
	Description        string
	Path               string
	ExpectedStatusCode int
	Body               any
	Method             string
}

type AnyString struct{}

type AnyNumber struct{}

type Timestamp struct{}

func (a Timestamp) Match(v driver.Value) bool {
	_, ok := v.(int64)
	return ok
}

func (a AnyString) Match(value driver.Value) bool {
	_, ok := value.(string)
	return ok
}

func (a AnyNumber) Match(value driver.Value) bool {
	_, ok := value.(int)
	return ok
}
