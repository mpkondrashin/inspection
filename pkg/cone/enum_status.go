// Code generated by enum (github.com/mpkondrashin/enum) using following command:
// enum -package cone -type Status --names success,fail,in-progress
// DO NOT EDIT!

package cone

import (
    "encoding/json"
    "errors"
    "fmt"
    "strconv"
    "strings"
)

type Status int

const (
    StatusSuccess     Status = iota
    StatusFail        Status = iota
    StatusIn_progress Status = iota
)


// String - return string representation for Status value
func (v Status)String() string {
    s, ok := map[Status]string {
         StatusSuccess:     "success",
         StatusFail:        "fail",
         StatusIn_progress: "in-progress",
    }[v]
    if ok {
        return s
    }
    return "Status(" + strconv.FormatInt(int64(v), 10) + ")"
}

// ErrUnknownStatus - will be returned wrapped when parsing string
// containing unrecognized value.
var ErrUnknownStatus = errors.New("unknown Status")


var mapStatusFromString = map[string]Status{
    "success":    StatusSuccess,
    "fail":    StatusFail,
    "in-progress":    StatusIn_progress,
}

// UnmarshalJSON implements the Unmarshaler interface of the json package for Status.
func (s *Status) UnmarshalJSON(data []byte) error {
    var v string
    if err := json.Unmarshal(data, &v); err != nil {
        return err
    }
    result, ok := mapStatusFromString[strings.ToLower(v)]
    if !ok {
        return fmt.Errorf("%w: %s", ErrUnknownStatus, v)
    }
    *s = result
    return nil
}

// MarshalJSON implements the Marshaler interface of the json package for Status.
func (s Status) MarshalJSON() ([]byte, error) {
    return []byte(fmt.Sprintf("\"%v\"", s)), nil
}

// UnmarshalYAML implements the Unmarshaler interface of the yaml.v3 package for Status.
func (s *Status) UnmarshalYAML(unmarshal func(interface{}) error) error {
    var v string
    if err := unmarshal(&v); err != nil {
        return err
    }
    result, ok := mapStatusFromString[strings.ToLower(v)]  
    if !ok {
        return fmt.Errorf("%w: %s", ErrUnknownStatus, v)
    }
    *s = result
    return nil
}