package testingwithgo

import "testing"

func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse("hellogdg")
	}
}
