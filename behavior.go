package behavior

type Status int8

const (
	Ready   Status = 0 // 初始状态
	Success Status = 1 // 成功
	Failure Status = 2 // 失败
	Running Status = 3 // 执行中
)

type IBehavior interface {
	// 执行当前行为
	Exec(Context) Status

	// 重置行为状态为 Ready
	Reset()

	// 转换行为状态
	// 如果行为状态不为 Running 状态，则重置为 Ready 状态
	//Next()
}

// base 基础行为
type base struct {
}

// composite 组合行为
type composite struct {
	base
	children []IBehavior
}
