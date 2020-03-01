package pb_test

import (
	"math/rand"
	"sync"
	"time"

	"github.com/chuanjiangwong/pb/v3"
)

func Example_multiple() {
	// create bars
	first := pb.New(200)
	second := pb.New(200)
	third := pb.New(200)
	// start pool
	pool, err := pb.StartPool(first, second, third)
	if err != nil {
		panic(err)
	}
	// update bars
	wg := new(sync.WaitGroup)
	for _, bar := range []*pb.ProgressBar{first, second, third} {
		wg.Add(1)
		go func(cb *pb.ProgressBar) {
			for n := 0; n < 200; n++ {
				cb.Increment()
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
			}
			cb.Finish()
			wg.Done()
		}(bar)
	}
	wg.Wait()
	// close pool
	pool.Stop()
}