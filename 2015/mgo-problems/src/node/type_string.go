// generated by stringer -type=Type; DO NOT EDIT

package node

import "fmt"

const _Type_name = "TypeLeafTypeObjectTypeListTypeOrdObject"

var _Type_index = [...]uint8{0, 8, 18, 26, 39}

func (i Type) String() string {
	i -= 1
	if i < 0 || i >= Type(len(_Type_index)-1) {
		return fmt.Sprintf("Type(%d)", i+1)
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}
