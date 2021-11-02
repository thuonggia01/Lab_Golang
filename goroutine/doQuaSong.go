package main

import (
	"fmt"
	"sync"
	"time"
)

var aTob chan int
var wg sync.WaitGroup

func timeTrack(start time.Time, functionName string) {
	Elapsed := time.Since(start)
	fmt.Println(functionName, "Elapsed =", Elapsed)
}

/*channel*/
func laiDo(c chan int) {
	for i := 1; i < 15; i++ {
		c <- i
		time.Sleep(2 * time.Second)
	}
	close(c)
}
func main() {
	defer timeTrack(time.Now(), "Lái đò")
	aTob := make(chan int, 10)
	go laiDo(aTob)
	for value := range aTob {
		fmt.Println("người thứ ", value, " qua sông")
	}
}

/* waitgroup
func chuyenDo(n int)  {
	//defer wg.Done()
	fmt.Println("Chuyến đò thứ : ",n)
	time.Sleep(time.Second)
}

func main()  {
	defer timeTrack(time.Now(),"chuyenDo")
	//wg.Add(1)
		for value:=1;value<11;value++{
			 chuyenDo(value)
		}
	//wg.Wait()
}
*/
/*  Mexter
type Boat struct {
	STT int
	m sync.Mutex
}
func (u *Boat) Go()int  {
	u.m.Lock()
	defer u.m.Unlock()
	return u.STT+1

}
func (u * Boat)increase()  {
	u.m.Lock()
	defer u.m.Unlock()
	u.STT++

}
func main()  {
	defer timeTrack(time.Now(),"Lái đò")
	boat :=Boat{}
	for i:=0;i<10;i++{
		fmt.Println("Chuyến đi thứ ",boat.Go())
		go boat.increase()
		time.Sleep(2*time.Second)
	}
}
*/
