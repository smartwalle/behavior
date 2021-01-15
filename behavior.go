package behavior

type Status int8

const (
	Success Status = 1 // 成功
	Failure Status = 2 // 失败
	Running Status = 3 // 执行中
	Error   Status = 4 // 错误
)

// IWorker 行为执行者接口
type IWorker interface {
	// 每一次执行之前都会调用
	OnEnter(Context)

	// 当执行一个状态不为 Running 的行为之前会调用
	OnOpen(Context)

	// 执行当前行为
	OnExec(Context) Status

	// 当执行该行为之后，其返回的状态不为 Running 时会调用
	OnClose(Context)

	// 每一次执行完成之后都会调用
	OnExit(Context)
}

// IBehavior 行为接口
type IBehavior interface {
	// 设置行为的执行者
	SetWorker(IWorker)

	// 执行行为，创建行为之后，调用本方法开始执行
	Exec(Context) Status
}

// base 基础行为
type base struct {
	worker    IWorker
	isRunning bool
}

func (this *base) SetWorker(b IWorker) {
	this.worker = b
}

// Exec 开始执行行为
func (this *base) Exec(ctx Context) Status {
	this.enter(ctx)

	if this.isRunning == false {
		this.open(ctx)
	}

	var status = this.exec(ctx)

	if status != Running {
		this.close(ctx)
	}

	this.exit(ctx)

	return status
}

func (this *base) enter(ctx Context) {
	this.worker.OnEnter(ctx)
}

func (this *base) open(ctx Context) {
	this.isRunning = true
	this.worker.OnOpen(ctx)
}

func (this *base) exec(ctx Context) Status {
	return this.worker.OnExec(ctx)
}

func (this *base) close(ctx Context) {
	this.isRunning = false
	this.worker.OnClose(ctx)
}

func (this *base) exit(ctx Context) {
	this.worker.OnExit(ctx)
}

// Composite 组合行为
type Composite struct {
	base
	children []IBehavior
}

func (this *Composite) OnEnter(Context) {
}

func (this *Composite) OnOpen(Context) {
}

func (this *Composite) OnExec(Context) Status {
	return Error
}

func (this *Composite) OnClose(Context) {
}

func (this *Composite) OnExit(Context) {
}

// Action 具体行为
type Action struct {
	base
}

func (this *Action) OnEnter(Context) {
}

func (this *Action) OnOpen(Context) {
}

func (this *Action) OnExec(Context) Status {
	return Error
}

func (this *Action) OnClose(Context) {
}

func (this *Action) OnExit(Context) {
}
