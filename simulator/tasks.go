package simulator

import "time"

const(
	P1Time int64=210
	P2Time int64=22
	WaitTIme int64=75
	C2Time int64=22
	TimeTimes int64= 10

	P1TaskNumber int64=12
	P2TaskNumber int64=1
	WaitTaskNumber int64=30
	C2TaskNumber int64 = 30

	Phase_p1 int64 = 1
	Phase_p2 int64 =2
	Phase_Wait int64 =3
	Phase_C2 int64 = 4
	MaxParallel int64 = 12
)
var PhaseName map[int64]string


func init() {
	//初始化PhaseName
	PhaseName = make(map[int64]string)
	PhaseName[Phase_p1]="p1"
	PhaseName[Phase_p2]="p2"
	PhaseName[Phase_Wait]="wait"
	PhaseName[Phase_C2]="c2"
	ts = GetTasks()
}
type Tasks struct{
	Start chan *Task
	P1P2 chan *Task
	P2Wait chan *Task
	WaitC2 chan *Task
	TaskControl chan int
	TasksNum chan int64
	P1Parallel int64
	P2Parallel int64
	WaitParallel int64
	C2Parallel int64
}

func GetTasks()*Tasks{
	return &Tasks{

		Start: make(chan *Task),
		P1P2: make(chan *Task,MaxParallel),
		P2Wait:make(chan *Task,MaxParallel),
		WaitC2: make(chan *Task,MaxParallel),
		TasksNum: make(chan int64,MaxParallel),
		P1Parallel:MaxParallel,
		P2Parallel:1,
		WaitParallel: MaxParallel,
		C2Parallel: MaxParallel,
	}
}



type Task struct {
	TaskNumber int64  //任务序号
	StartTime int64   //开始时间
	EndTime int64  //结束时间
	P1Start int64
	P1End int64
	P2Start int64
	P2End int64
	WaitStart int64
	WaitEnd int64
	C2Start int64
	C2End int64
	Phase int64  //所在阶段
	P1P2Wait int64
	P2C2Wait int64
}



func (t *Task)P1(){
	defer func() {
		t.P1End=now()
		ts.P1P2 <- t
	}()

	t.StartTime=now()
	t.P1Start=now()
	time.Sleep(time.Duration(P1Time*TimeTimes)*time.Millisecond)
}


func (t *Task)P2(){
	t.P2Start=now()
	t.P1P2Wait=t.P2Start-t.P1End
	defer func() {
		t.P2End=now()
		 ts.P2Wait <- t
	}()
	time.Sleep(time.Duration(P2Time*TimeTimes)*time.Millisecond)
}

func (t *Task)Wait(){
	t.WaitStart=now()
	defer func() {
		t.WaitEnd=now()
		ts.WaitC2 <- t
	}()
	time.Sleep(time.Duration(WaitTIme*TimeTimes)*time.Millisecond)
}

func (t *Task)C2(){
	t.C2Start=now()
	t.P2C2Wait=t.C2Start-t.P2End
	defer func() {
		t.C2End=now()
		<-ts.TasksNum
	}()
	time.Sleep(time.Duration(C2Time*TimeTimes)*time.Millisecond)
}

func now() int64 {
	return time.Now().UnixNano() / 1e6
}