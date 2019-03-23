package xlsToJson

import (
	"reflect"

	"github.com/tealeg/xlsx"
)

// recursivelyAddNodeTSLV1 - adds child node to last parent with child
// called from AddNodeToSameLevel
func (v *VariableHolder) recursivelyAddNodeTSLV1(parentNode, childNode reflect.Value, index int) (bool, int, reflect.Value) {
	parentChildrenList := parentNode.FieldByName(v.childrenArrayKey)
	if parentChildrenList.Len() == 0 {
		return true, index, parentNode
	} else {
		isTrue, index, updatedNode := v.recursivelyAddNodeTSLV1(v.getLastNodeFromChildrenListOfNode(parentNode), childNode, index+1)
		if isTrue {
			if index != -1 {
				childNode = deReferenceNode(childNode)
				parentChildrenList.Set(reflect.Append(parentChildrenList, childNode))
			} else {
				if parentChildrenList.Len() == 0 {
					node := reflect.New(v.childrenArrayEntryType)
					nodeChildrenList := node.FieldByName(v.childrenArrayKey)
					updatedNode = deReferenceNode(updatedNode)
					nodeChildrenList.Set(reflect.Append(nodeChildrenList, updatedNode))
					parentChildrenList = nodeChildrenList
				} else {
					updatedNode = deReferenceNode(updatedNode)
					parentChildrenList.Index(parentChildrenList.Len() - 1).Set(updatedNode)
				}
			}
			return true, -1, parentNode
		}

	}
	return false, 0, parentNode
}

// addNodeToSameLevelV1 - adds node adjacent to its sibling
func (v *VariableHolder) addNodeToSameLevelV1(current int, node reflect.Value, currentRow *xlsx.Row) (reflect.Value, error) {
	newNode, err := v.createNewNode(currentRow, current)
	if err != nil {
		return newNode, err
	}
	newNode = deReferenceNode(newNode)
	v.updatelevelWiseNodeMapEntry(current, newNode.Interface())
	v.updateCurrentNode(newNode.Interface())
	_, _, node = v.recursivelyAddNodeTSLV1(node, newNode, 0)
	return node, nil
}
