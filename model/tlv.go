package model

type TLVItem struct {
	Tag    string
	Length string
	Value  string
}

type ProcessResponse struct {
	OriginalString string
	NewString      string
	OriginalTLV    []TLVItem
	NewTLV         []TLVItem
	Nominal        string
	CRCValue       string
}
