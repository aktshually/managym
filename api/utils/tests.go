package utils

type AnyString struct{}

type AnyNumber struct{}

type Timestamp struct{}

func (a Timestamp) Match(v any) bool {
	_, ok := v.(int64)
	return ok
}

func (a AnyString) Match(value any) bool {
	_, ok := value.(string)
	return ok
}

func (a AnyNumber) Match(value any) bool {
	_, ok := value.(int)
	return ok
}
