package behavior

type Status int8

const (
	Success Status = 1 // 成功
	Failure Status = 2 // 失败
	Running Status = 3 // 执行中
	Error   Status = 4 // 错误
)

func (s Status) String() string {
	switch s {
	case Success:
		return "Success"
	case Failure:
		return "Failure"
	case Running:
		return "Running"
	case Error:
		return "Error"
	}
	return "Unknown"
}

// Worker 行为执行者接口
type Worker interface {
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

// Behavior 行为接口
type Behavior interface {
	// 设置行为的执行者
	SetWorker(Worker)

	// 执行行为，创建行为之后，调用本方法开始执行
	Tick(Context) Status
}

// base 基础行为
type base struct {
	worker    Worker
	isRunning bool
}

func (this *base) SetWorker(b Worker) {
	this.worker = b
}

// Exec 开始执行行为
func (this *base) Tick(ctx Context) Status {
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
