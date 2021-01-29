package simulator

var ts *Tasks
func Run(){
	go AddTask()
	go DoP1()
	go DoP2()
	go DoWait()
	DoC2()
}

func AddTask(){
	var count int64 =0
	for{
		t := &Task{
			TaskNumber:count,
		}
		ts.TasksNum <- count
		ts.Start <- t
		count++
	}
}

func DoP1()  {
	for{
		t := <- ts.Start
		go t.P1()
	}
}
func DoP2(){
	for{
		t := <-ts.P1P2
		t.P2()
	}
}
func DoWait(){
	for{
		t := <-ts.P2Wait
		go t.Wait()
	}
}
func DoC2(){
	for{
		t := <- ts.WaitC2
		go t.C2()
	}
}


