package list

// List
// list接口
type List[T any] interface {
	// Get
	// 根据输入的下标index， 返回对应下标的元素
	// 当输入的下标越界，返回错误
	Get(index int) (T, error)

	// Append
	// 在末尾追加元素
	Append(ts ...T) error

	// Add
	// 在指定下标插入一个新元素
	// 当输入的下标越界，返回错误
	Add(index int, t T) error

	// Set
	// 在指定下标，重设该位置的元素值
	// 当输入的下标越界，返回错误
	Set(index int, t T) error

	// Delete
	// 删除输入下标的元素，返回被删除的元素值
	// 当输入的下标越界，返回错误
	Delete(index int) (T, error)

	// Len
	// 返回长度
	Len() int

	// Cap
	// 参会容量
	Cap() int

	// Range
	// 遍历 List 的所有元素
	Range(fn func(index int, t T) error) error

	// ChangeSlice
	// 将 List 转化为一个切片
	// 必定会返回全新的切片
	ChangeSlice() []T
}