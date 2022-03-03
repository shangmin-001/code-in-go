package channel

import (
	"sync"
	"testing"
)

func BenchmarkChannelOneByte(b *testing.B) {
	// ch := make(chan byte, 4096)
	// wg := sync.WaitGroup{}
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	for range ch {
	// 	}
	// }()
	// b.SetBytes(1)
	// b.ReportAllocs()
	// b.ResetTimer()
	// for i := 0; i < b.N; i++ {
	// 	ch <- byte(i)
	// }
	// close(ch)
	// wg.Wait()
}

func BenchmarkCopy(b *testing.B) {
	// from := make([]byte, b.N)
	// to := make([]byte, b.N)
	// b.ReportAllocs()
	// b.ResetTimer()
	// b.SetBytes(1)
	// copy(to, from)
}

func BenchmarkChannelOneMap(b *testing.B) {
	// ch := make(chan map[string]interface{})
	// dataRecv := make([]map[string]interface{}, 0)
	// wg := sync.WaitGroup{}
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	for res := range ch {
	// 		dataRecv = append(dataRecv, res)
	// 	}
	// }()
	// b.SetBytes(1)
	// b.ReportAllocs()
	// b.ResetTimer()
	// for i := 0; i < b.N; i++ {
	// 	ch <- map[string]interface{}{"raw": "123"}
	// }
	// close(ch)
	// wg.Wait()
}

func BenchmarkCopyMap(b *testing.B) {

	// dataRecv := make([]map[string]interface{}, 0)
	// b.ReportAllocs()
	// b.ResetTimer()
	// b.SetBytes(1)

	// for i := 0; i < b.N; i++ {
	// 	dataRecv = append(dataRecv, map[string]interface{}{"raw": "123"})
	// }

}
