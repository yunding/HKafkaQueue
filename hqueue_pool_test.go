package HKafkaQueue

import (
	"fmt"
	"testing"
)

func TestCreatePool(t *testing.T) {
	pool := NewHQueuePool("/tmp/hqueue", 60)
	fmt.Println(pool.hqueueMap)
	pool.Destroy()
}

//func TestPoolDelete(t *testing.T) {
//	pool := NewHQueuePool("/tmp/hqueue", 120)
//	hqueue, _ := NewHQueue("test2", "/tmp/hqueue")
//	pool.hqueueMap["test2"] = hqueue
//	var w sync.WaitGroup
//	w.Add(2)
//	go read(hqueue, &w)
//	go write(hqueue, &w)
//	w.Wait()
//	pool.Destroy()
//}
