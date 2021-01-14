package behavior

type Status int8

const (
	Success Status = 1 // 成功
	Failure Status = 2 // 失败
	Running Status = 3 // 执行中
)

type IBehavior interface {
	// 每一次执行之前都会调用
	OnEnter(*Context)

	// 当执行一个状态不为 Running 的行为之前会调用
	OnOpen(*Context)

	// 执行当前行为
	OnExec(*Context) Status

	// 当执行该行为之后，其返回的状态不为 Running 时会调用
	OnClose(*Context)

	// 每一次执行完成之后都会调用
	OnExit(*Context)
}

// base 基础行为
type base struct {
}

func (this *base) OnEnter(ctx *Context) {
}

func (this *base) OnOpen(ctx *Context) {
}

func (this *base) OnExec(ctx *Context) Status {
	return Failure
}

func (this *base) OnClose(ctx *Context) {
}

func (this *base) OnExit(ctx *Context) {
}

// composite 组合行为
type composite struct {
	base
	children []IBehavior
}

type Tree struct {
	root IBehavior
}

func NewTree(root IBehavior) *Tree {
	var t = &Tree{}
	t.root = root
	return t
}

func (this *Tree) Exec(target interface{}) Status {
	return Failure
}
