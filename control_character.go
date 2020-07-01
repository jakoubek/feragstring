package feragstring

import "fmt"

type ControlCharacterSet struct {
	DontProduce            bool
	BundlesToSecondaryExit bool
}

func NewControlCharacterSet() *ControlCharacterSet {
	cc := ControlCharacterSet{
		DontProduce:            false,
		BundlesToSecondaryExit: false,
	}
	return &cc
}

func (cc *ControlCharacterSet) SetDontProduce() {
	cc.DontProduce = true
}

func (cc *ControlCharacterSet) SetBundlesToSecondaryExit() {
	cc.BundlesToSecondaryExit = true
}

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
