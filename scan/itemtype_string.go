// Code generated by "stringer -type=itemType"; DO NOT EDIT

package scan

import "fmt"

const _itemType_name = "ItemEOFItemNumberItemVariableItemEqualItemAddItemSubItemMulItemDivItemLParenItemRParenItemError"

var _itemType_index = [...]uint8{0, 7, 17, 29, 38, 45, 52, 59, 66, 76, 86, 95}

func (i itemType) String() string {
	if i < 0 || i >= itemType(len(_itemType_index)-1) {
		return fmt.Sprintf("itemType(%d)", i)
	}
	return _itemType_name[_itemType_index[i]:_itemType_index[i+1]]
}
