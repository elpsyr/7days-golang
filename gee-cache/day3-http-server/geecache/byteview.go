package geecache

// A ByteView holds an immutable view of bytes
// b会存储真是的缓存值。选择byte是为了能够支持任意类型数据的存储，例如字符串、图片等。
type ByteView struct {
	b []byte
}

// 我们在lru.Cache的实现中，要求被缓存的对象必须实现Value接口
func (v ByteView) Len() int {
	return len(v.b)
}

// ByteSlice returns a copy of the data as a byte slice
// b是只读的，使用ByteSlice返回一个拷贝，防止缓存被外部程序修改
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)

}

// String returns the data as a string, making a copy if necessary
func (v ByteView) String() string {
	return string(v.b)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
