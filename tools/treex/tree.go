package treex

type Node interface {
	GetId() interface{}
	GetPid() interface{}
	AppendChildren(interface{})
}

// BuildTree 切片转树
func BuildTree(array []Node) []interface{} {
	var rootData []interface{}
	maxLen := len(array)
	///<找出根节点,根节点的特点，没有父节点
	for i := 0; i < maxLen; i++ {
		///< 统计每个节点的父节点出现的次数，父节点出现0次就是根节点
		count := 0
		for j := 0; j < maxLen; j++ {
			///< 如果有节点的ID == i的parentID 那么j就是父节点
			if array[j].GetId() == array[i].GetPid() {
				count++
				array[j].AppendChildren(array[i])
			}
		}
		if count == 0 {
			rootData = append(rootData, array[i])
		}
	}
	return rootData
}

//使用需指定ID,Pid,Children三个字段
//使用需实现接口方法
////获取自身ID
//func (that *A) GetId() int {
//  return that.Id
//}

////获取自身上级ID
//func (that *A) GetPid() int {
//  return that.Pid
//}

////添加子集
//func (that *A) AppendChildren(node interface{}) {
//  that.Children = append(that.Children, node.(*A))
//}

//使用时需要转换数据
//nodeArray := make([]TreeNode, len(data))
//for i := 0; i < len(data); i++ {
//  nodeArray[i] = &data[i]
//}

//BuildTree(nodeArray)
