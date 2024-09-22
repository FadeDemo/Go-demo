package main

import "testing"

func BenchmarkEchoWithLowPerformance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EchoWithLowPerformance()
	}
}

func BenchmarkEchoUsingStringsJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EchoUsingStringsJoin()
	}
}
