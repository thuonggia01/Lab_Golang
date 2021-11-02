/*Đò qua sông
- Một con đò chở khách 1 chiều từ A -> B, chở người đến B sẽ quay lại điểm A. Tại điểm A có 1 phòng chờ ngồi được tối đa 10 người.
- Con đò chỉ chở được 1 hành khách / 1 lượt (theo số thứ tự)
- Người lái đò sẽ kiểm tra phòng chờ:
    + Có người sẽ gọi theo thứ tự và chở người đó
    + Không có người thì sẽ ngủ ngay trên đò
- Khách đi đò đến:
    + Nếu thấy người lái đò ngủ sẽ lấy số, đánh thức người lái đò và lên đò đi luôn.
    + Nếu không thấy đò hoặc có người đang chờ (đò đang trên đường đi) thì sẽ lấy số và ngồi chờ.
    + Nếu thấy phòng chờ đầy người thì sẽ đi về.
Giả lập tình huống đi đò với số khách tùy ý:
- Fix số lượng khách làm đầu vào
- Fix thời gian người lái đò cần đi từ A -> B rồi quay lại (dùng sleep)
- Thời gian khách đến không theo quy luật (random thời gian và sleep để giả lập)
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	sleeping = iota
	checking
	ferrying
)

var stateLog = map[int]string{
	0: "Sleeping",
	1: "Checking",
	2: "ferrying",
}
var wg *sync.WaitGroup // Amount of potentional customers

type Ferryman struct {
	sync.Mutex
	state    int // Sleeping/Checking/ferrying
	customer *Customer
}

type Customer struct {
	number int
	name   string
}

func (c *Customer) String() string {
	return fmt.Sprintf("%p", c)[7:]
}
func NewFerryman() (f *Ferryman) {
	return &Ferryman{
		state: sleeping,
	}
}

// goroutine Ferryman
// Có người sẽ gọi theo thứ tự và chở người đó
// Ngủ - đợi khách hàng đánh thức
func ferryman(f *Ferryman, wr chan *Customer, wakers chan *Customer) {
	for {
		f.Lock()
		defer f.Unlock()
		f.state = checking
		f.customer = nil

		// kiểm tra phòng chờ
		fmt.Printf("ferryman Checking waiting room: %d\n", len(wr))
		time.Sleep(time.Second * 1)
		select {
		case c := <-wr:
			Ferrying(c, f)
			f.Unlock()
		default: // Phòng chờ trống
			fmt.Printf("Ferryman Sleeping.%s\n", f.customer)
			f.state = sleeping
			f.customer = nil
			f.Unlock()
			c := <-wakers
			f.Lock()
			fmt.Printf("Ferry driver woken up by: %s\n", c)
			Ferrying(c, f)
			f.Unlock()
		}
	}
}

func Ferrying(c *Customer, f *Ferryman) {
	f.state = ferrying
	f.customer = c
	f.Unlock()
	fmt.Printf("ferrying A to B: %s \n", c)
	time.Sleep(time.Second * 2)
	f.Lock()
	wg.Done()
	f.customer = nil
}

// customer goroutine
// Nếu thấy người lái đò ngủ sẽ lấy số, đánh thức người lái đò và lên đò đi luôn.
// Nếu không thấy đò hoặc có người đang chờ (đò đang trên đường đi) thì sẽ lấy số và ngồi chờ.
// Nếu thấy phòng chờ đầy người thì sẽ đi về.
func customer(c *Customer, f *Ferryman, wr chan<- *Customer, wakers chan<- *Customer) {
	// arrive
	time.Sleep(time.Second)
	// kiểm tra người lái đò
	f.Lock()
	fmt.Printf("Customer %s checks %s Ferryman | room: %d, w %d - customer: %s\n",
		c, stateLog[f.state], len(wr), len(wakers), f.customer)
	switch f.state {
	case sleeping:
		select {
		case wakers <- c:
		default:
			select {
			case wr <- c:
			default:
				wg.Done()
			}
		}
	case ferrying:
		select {
		case wr <- c:
		default: // Phòng chờ full =>đi về
			wg.Done()
		}
	case checking:
		panic("Ferryman is Checking the waiting room")
	}
	f.Unlock()
}

func main() {
	f := NewFerryman()
	WaitingRoom := make(chan *Customer, 10) // phòng chờ max 10 người
	Wakers := make(chan *Customer, 1)       // chỉ có 1 người lái đò tại 1 thời điểm
	go ferryman(f, WaitingRoom, Wakers)
	wg = new(sync.WaitGroup)
	wg.Add(20)
	// Spawn customers
	for i := 0; i < 20; i++ {
		time.Sleep(time.Second * 2)
		c := new(Customer)
		go customer(c, f, WaitingRoom, Wakers)
	}
	wg.Wait()
	fmt.Println("Không còn khách")
}
