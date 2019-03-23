package xlsToJson

import (
	"reflect"

	"github.com/tealeg/xlsx"
)

// recursivelyAddNodeTPLV1 - add child node as sibling to last parent
// called from AddNodeToPreviousLevel
func (v *VariableHolder) recursivelyAddNodeTPLV1(node, child reflect.Value, current, i int) (bool, reflect.Value) {
	for i <= current {
		if i == current {
			nodeChildrenList := node.FieldByName(v.childrenArrayKey)
			child = deReferenceNode(child)
			nodeChildrenList.Set(reflect.Append(nodeChildrenList, child))
			return true, node
		} else {
			i++
			isAdded, receivedNode := v.recursivelyAddNodeTPLV1(v.getLastNodeFromChildrenListOfNode(node), child, current, i)
			nodeChildrenList := node.FieldByName(v.childrenArrayKey)
			if !isAdded {
				nodeChildrenList.Index(nodeChildrenList.Len() - 1).Set(v.getLastNodeFromChildrenListOfNode(receivedNode))
				isAdded = true
				return isAdded, node
			} else {
				nodeChildrenList.Index(nodeChildrenList.Len() - 1).Set(receivedNode)
				return true, node
			}
		}
	}
	return false, node
}

// addNodeToPreviousLevelV1 - add node adjacent to last parent with child
func (v *VariableHolder) addNodeToPreviousLevelV1(current int, node reflect.Value, currentRow *xlsx.Row) (reflect.Value, error) {
	newNode, err := v.createNewNode(currentRow, current)
	if err != nil {
		return newNode, err
	}
	newNode = deReferenceNode(newNode)
	v.updatelevelWiseNodeMapEntry(current, newNode.Interface())
	v.updateCurrentNode(newNode.Interface())
	_, node = v.recursivelyAddNodeTPLV1(node, newNode, current, 0)
	return node, nil
}
