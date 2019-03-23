package xlsToJson

import (
	"reflect"

	"github.com/tealeg/xlsx"
)

// recursivelyAddNodeTNLV1 - adds child node to last parent with no child
// called from AddNodeToNextLevel
func (v *VariableHolder) recursivelyAddNodeTNLV1(parentNode, childNode reflect.Value) reflect.Value {
	lastNode := reflect.New(v.childrenArrayEntryType)	
	parentChildrenList := parentNode.FieldByName(v.childrenArrayKey)
	if parentChildrenList.Len() == 0 {
		parentNode = v.appendChildNode(parentNode, childNode)
	} else {
		lastNode = v.recursivelyAddNodeTNLV1(v.getLastNodeFromChildrenListOfNode(parentNode), childNode)
		lastNodeChildrenList := lastNode.FieldByName(v.childrenArrayKey)
		if lastNodeChildrenList.Len() == 0 {
			node := reflect.New(v.childrenArrayEntryType)
			nodeChildrenList := node.FieldByName(v.childrenArrayKey)
			nodeChildrenList.Set(reflect.Append(nodeChildrenList, lastNode))
			parentChildrenList = nodeChildrenList
		} else {
			parentChildrenList.Index(parentChildrenList.Len() - 1).Set(lastNode)
		}
	}
	return parentNode
}

// addNodeToNextLevelV1 - add node to next level to its parent
func (v *VariableHolder) addNodeToNextLevelV1(current int, node reflect.Value, currentRow *xlsx.Row) (reflect.Value, error) {
	newNode, err := v.createNewNode(currentRow, current)
	if err != nil {
		return newNode, err
	}
	newNode = deReferenceNode(newNode)
	v.updateCurrentNode(newNode.Interface())
	v.updatelevelWiseNodeMapEntry(current, newNode.Interface())
	node = v.recursivelyAddNodeTNLV1(node, newNode)
	return node, nil
}
