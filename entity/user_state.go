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

	_UserStateToValue = map[string]UserState{
		"unspecified": UserStateUnspecified,
		"active":      UserStateActive,
		"inactive":    UserStateInactive,
		"suspend":     UserStateSuspend,
	}
)

// MarshalJSON inherit marshal json default
func (us UserState) MarshalJSON() ([]byte, error) {
	s, ok := _UserStateToName[us]
	if !ok {
		return nil, errors.New("invalid user state")
	}
	return json.Marshal(s)
}

// UnmarshalJSON inherit unmarshal json default
func (us *UserState) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	v, ok := _UserStateToValue[s]
	if !ok {
		return errors.New("invalid user state")
	}
	*us = v
	return nil
}
