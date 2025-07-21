package models

type Beep struct {
	// 0 < Frequency < 20000, otherwise undefined behavior, default is 440Hz
	Frequency uint16 `json:"frequency,omitempty"`

	// Length of beep in milliseconds, default is 200ms
	Length uint `json:"length,omitempty"`

	// Delay between beeps
	// Only delays if set, beeps otherwise
	Delay uint `json:"delay,omitempty"`
}
