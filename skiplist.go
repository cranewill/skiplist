package skiplist

import ()

const DefaultLevel = 8

type NodeValue interface {
	IsFront(other interface{}) bool // function to determinate if this node has the higher priority than other
}

type Node struct {
	index     int32     // index of this node of all nodes
	value     NodeValue // value is an any-type struct implements NodeValue
	prevNodes []*Node   // prevNodes contains this node's previous node pointers according to level
	nextNodes []*Node   // nextNodes contains this node's next node pointers according to level
}

type LevelList struct {
	level  int   // which level of this list
	length int32 // length of this list
	head   *Node // the first node
	tail   *Node // the last node
}

type SkipList struct {
	lists []*LevelList // all the levelList
	level int          // total level num
}

// Init creates a new empty skipList
func Init() *SkipList {
	list := make([]*LevelList, 0)
	for i := 0; i < DefaultLevel; i++ {
		ll := &LevelList{
			level:  i,
			length: 0,
			head:   nil,
			tail:   nil,
		}
		list = append(list, ll)
	}

	sl := &SkipList{
		level: DefaultLevel,
		lists: list,
	}
	return sl
}

// Insert push element into the appropriate level & position of sl and give it an index
func (sl *SkipList) Insert(element interface{}) {

}

func (sl *SkipList) SearchInsertNextNode(element interface{}) *Node {
	// find the first not nil node from higher to lower
	node := &Node{}
	level := sl.level - 1
	for {
		node = sl.lists[level].head
		if node == nil {
			if level <= 0 {
				break
			}
			level--
			continue
		} else {
			break
		}
	}
	//if node == nil {
	//	return nil
	//}

	for i := level; i > 0; i-- {
		// if this node has higher priority than element
		if node.value.IsFront(element) {
			// just find the next node in this level
			nextNode := node.nextNodes[level]
			// if there is no next node this level, go to the lower level
			if nextNode == nil {
				continue
			} else {
				//if next
			}

		}

	}
	return nil
}

//
//// insert accepts node which is the first node to judge of this levelList
//// or nil if the skipList is empty,
//// and inserts element into the skipList
//func insert (startNode *Node, level int, element interface{}) {
//	if startNode == nil {
//		// todo insert when the skipList is empty
//		return
//	}
//	node := *startNode
//	for {
//		if node.value.IsFront(element) {
//			nextNode := node.nextNodes[level]
//			if nextNode == nil {
//				break
//			}
//			node = *nextNode
//			continue
//		}
//	}
//}
