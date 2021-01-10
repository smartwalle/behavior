package behavior

type Status int32

const (
	Success Status = 1 // 成功
	Failure Status = 2 // 失败
	Running Status = 3 // 执行中
	Ready   Status = 4 // 初始状态
)

type IBehavior interface {
	// 执行当前行为
	Exec(Context)

	// 获取行为的状态
	Status() Status

	// 重置行为状态为 Ready
	Reset()

	// 转换行为状态
	// 如果行为状态不为 Running 状态，则重置为 Ready 状态
	Next()
}

// baseBehavior 基础行为
type baseBehavior struct {
	status Status
}

func (this *baseBehavior) Exec(ctx Context) {
}

func (this *baseBehavior) Status() Status {
	return this.status
}

func (this *baseBehavior) Reset() {
	if this.status != Ready {
		this.status = Ready
	}
}

func (this *baseBehavior) Next() {
	if this.status != Running {
		this.Reset()
	}
}

// compositeBehavior 组合行为
type compositeBehavior struct {
	baseBehavior
	children []IBehavior
}

func (this *compositeBehavior) Reset() {
	if this.status != Ready {
		this.status = Ready
		for _, child := range this.children {
			child.Reset()
		}
	}
}

func (this *compositeBehavior) Next() {
	if this.status != Running {
		this.Reset()
	} else {
		for _, child := range this.children {
			child.Next()
		}
	}
}

// Add 添加子行为
//func (this *compositeBehavior) Add(child IBehavior) {
//	if child == nil {
//		return
//	}
//	this.children = append(this.children, child)
//}
