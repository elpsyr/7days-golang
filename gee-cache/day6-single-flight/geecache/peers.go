package geecache

// PeerPicker is the interface that must implemented to locate
// the peer that owns a specific key
// PickPeer() 方法用于根据传入的key选择相应节点的PeerGetter
type PeerPicker interface {
	PickerPeer(key string) (peer PeerGetter, ok bool)
}

// PeerGetter is the interface that must be implemented by a peer
// Get() 方法用于对改group 查找缓存。 PeerGetter就对应于上述流程中的客户端
type PeerGetter interface {
	Get(group string, key string) ([]byte, error)
}

// 客户端关心的是我给什么：group名称和keu，并拿到结果
// 服务端关心的是要选择什么样的节点，最后把数据返回
