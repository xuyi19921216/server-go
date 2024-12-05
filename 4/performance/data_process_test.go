package performance

import (
	"fmt"
	"testing"
)

// 初始化构造测试用例
var users []*User

func init() {
	for i := 0; i < 1000; i++ {
		users = append(users, &User{Id: i, Name: fmt.Sprintf("user%d", i)})
	}
}

func BenchmarkGenerateIdsRaw(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GenerateIdsRaw(users)
	}
}

/* func BenchmarkGenerateIdsBuilder(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GenerateIdsBuilder(users)
	}
} */

/*
	 func BenchmarkGenerateIdsStrconv(b *testing.B) {
		for n := 0; n < b.N; n++ {
			GenerateIdsStrconv(users)
		}
	}
*/
func BenchmarkGenerateIdsUnsafe(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GenerateIdsUnsafe(users)
	}
}
