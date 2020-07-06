package feragstring

import "fmt"

// ControlCharacterSet is a struct for the control character
// field in the FERAG messages that support control characters
// (Route Info, Production Drop).
// Supports currently only "D" (Don't Produce) and "T" (Bundles
// to Secondary Exit).
type ControlCharacterSet struct {
	DontProduce            bool
	BundlesToSecondaryExit bool
}

// NewControlCharacterSet instantiates a new Control Character Set
// struct and returns a pointer to it.
func NewControlCharacterSet() *ControlCharacterSet {
	cc := ControlCharacterSet{
		DontProduce:            false,
		BundlesToSecondaryExit: false,
	}
	return &cc
}

// SetDontProduce sets the "don't produce" flag (D)
func (cc *ControlCharacterSet) SetDontProduce() {
	cc.DontProduce = true
}

// SetBundlesToSecondaryExit sets the "bundles to secondary exit" flag (T)
func (cc *ControlCharacterSet) SetBundlesToSecondaryExit() {
	cc.BundlesToSecondaryExit = true
}

// GetControlCharactersMessage returns the formatted FERAG string for
// the control character set.
func (cc *ControlCharacterSet) GetControlCharactersMessage() string {
	var ccCount int
	var ccString string
	if cc.DontProduce == true {
		ccString += "D"
		ccCount++
	}
	if cc.BundlesToSecondaryExit == true {
		ccString += "T"
		ccCount++
	}
	if ccCount == 0 {
		return ""
	}
	return fmt.Sprintf("+14%-16s", ccString)
}
