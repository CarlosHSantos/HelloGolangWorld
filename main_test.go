package main

//testes para variaveis
import (
	"fmt"
	"log"
	"os"
	"testing"
)

// go test .\main_test.go Run all tests in a simplified way
// go test -v - Run all tests with more details
// go test -bench=. - Run benchmark tests only

func TestMain(m *testing.M) {
    log.Println("Do stuff BEFORE the tests!")
    exitVal := m.Run()
    log.Println("Do stuff AFTER the tests!")

    os.Exit(exitVal)
}

func intMin(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func TestIntMinBasic(t *testing.T) {
    ans := intMin(2, -2)
    if ans != -2 {
        t.Errorf("IntMin(2, -2) = %d; want -2", ans)
    }
}

func TestIntMinTableDriven(t *testing.T) {
    var tests = []struct {
        a, b int
        want int
    }{
        {0, 1, 0},
        {1, 0, 0},
        {2, -2, -2},
        {0, -1, -1},
        {-1, 0, -1},
    }

    for _, tt := range tests {
        testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
        t.Run(testname, func(t *testing.T) {
            ans := intMin(tt.a, tt.b)
            if ans != tt.want {
                t.Errorf("got %d, want %d", ans, tt.want)
            }
        })
    }
}

// Benchmark tests typically go in _test.go files and are named beginning with Benchmark. The testing runner executes each 
// benchmark function several times, increasing b.N on each run until it collects a precise measurement.
// Typically the benchmark runs a function we’re benchmarking in a loop b.N times. 
func BenchmarkIntMin(b *testing.B) {
    for i := 0; i < b.N; i++ {
        intMin(1, 2)
    }
}