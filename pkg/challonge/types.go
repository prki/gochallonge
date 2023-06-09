package challonge

import "bytes"

type NullString struct {
	IsNull bool
	Value  string
}

type NullUint64 struct {
	IsNull bool
	Value  uint64
}

func (ns *NullString) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer

	if !ns.IsNull {
		buf.WriteString(ns.Value)
	} else {
		buf.Write(nil)
	}

	return buf.Bytes(), nil
}

func (ns *NullString) UnmarshalJSON(in []byte) error {
	str := string(in)

	if str == "null" {
		ns.IsNull = true
		return nil
	} else {
		ns.IsNull = false
		if len(str) >= 2 {
			ns.Value = str[1 : len(str)-1] // remove leading and trailing quotes
		} else {
			ns.Value = ""
		}
	}
	return nil
}
