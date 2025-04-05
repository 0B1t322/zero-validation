package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"slices"
	"strings"
)

type ErrorSlice []error

func (es ErrorSlice) Error() string {
	if len(es) == 0 {
		return ""
	}

	var builder strings.Builder
	for i, err := range es {
		builder.WriteString(err.Error())
		if i != len(es)-1 {
			builder.WriteString(", ")
		}
	}

	return builder.String()
}

func (es ErrorSlice) MarshalJSON() ([]byte, error) {
	return json.Marshal([]error(es))
}

func (es ErrorSlice) AppendError(err error) ErrorSlice {
	switch e := err.(type) {
	case ErrorSlice:
		return append(es, e...)
	default:
		return append(es, e)
	}
}

type Errors map[string]error

func (es Errors) Error() string {
	if len(es) == 0 {
		return ""
	}

	keys := make([]string, 0, len(es))

	for key := range es {
		keys = append(keys, key)
	}
	slices.Sort(keys)

	var s strings.Builder
	for i, key := range keys {
		if i > 0 {
			s.WriteString("; ")
		}

		var errs Errors
		if errors.As(es[key], &errs) {
			_, _ = fmt.Fprintf(&s, "%v: (%v)", key, errs)
			continue
		}

		_, _ = fmt.Fprintf(&s, "%v: %v", key, es[key].Error())
	}
	s.WriteString(".")
	return s.String()
}

// MarshalJSON converts the Errors into a valid JSON.
func (es Errors) MarshalJSON() ([]byte, error) {
	errs := map[string]any{}
	for key, err := range es {
		if ms, ok := err.(json.Marshaler); ok {
			errs[key] = ms
		} else {
			errs[key] = err.Error()
		}
	}

	return json.Marshal(errs)
}

func (es Errors) Join(errs Errors) Errors {
	for key, value := range errs {
		err, find := es[key]
		if !find {
			es[key] = value
			continue
		}

		switch err := err.(type) {
		case ErrorObject:
			if errObj, ok := value.(ErrorObject); ok && errObj.code == err.code {
				continue
			}

			if s, ok := value.(ErrorSlice); ok {
				s = append(s, err)
				es[key] = s
				continue
			}

			es[key] = ErrorSlice{err, value}
		case ErrorSlice:
			es[key] = err.AppendError(value)
		default:
			es[key] = ErrorSlice{err, value}
		}
	}

	return es
}

func (es Errors) Is(err error) bool {
	if err == nil {
		return false
	}

	_, ok := err.(Errors)
	return ok
}
