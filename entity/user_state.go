package entity

import (
	"encoding/json"
	"errors"
)

type UserState int

const (
	UserStateUnspecified UserState = iota
	UserStateActive
	UserStateInactive
	UserStateSuspend
)

var (
	_UserStateToName = map[UserState]string{
		UserStateUnspecified: "unspecified",
		UserStateActive:      "active",
		UserStateInactive:    "inactive",
		UserStateSuspend:     "suspend",
	}
)

// MarshalJSON inherit marshal json default
func (c UserState) MarshalJSON() ([]byte, error) {
	s, ok := _UserStateToName[c]
	if !ok {
		return nil, errors.New("invalid user state")
	}
	return json.Marshal(s)
}
