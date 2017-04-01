package robinhood

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"
)

const EPOCH = 10
const SIZE = 10000
const LOAD_FACTOR_TARGET = 0.75
const TESTCOUNT = SIZE * LOAD_FACTOR_TARGET

func genKey() string {
	return "key" + strconv.Itoa(rand.Int()%SIZE)
}

func TestLoadTest(t *testing.T) {
	loads := []float32{
		0.1, 0.15, 0.2, 0.25, 0.3, 0.35, 0.4, 0.45, 0.5, 0.55, 0.6, 0.65, 0.7, 0.75, 0.8, 0.85, 0.9, 0.95,
	}
	f, err := os.Create(fmt.Sprintf("data/load_test_erase.csv"))
	if err != nil {
		t.Error("Fail to open result.csv")
	}
	defer f.Close()
	writer := csv.NewWriter(f)
	writer.Write([]string{
		"epoch",
		"algorithm",
		"load_factor",
		"dib_average",
		"elapsed_time_ms",
	})
	for _, l := range loads {
		loadTest(t, int(SIZE*l), writer)
	}
	writer.Flush()
}

func loadTest(t *testing.T, testCount int, writer *csv.Writer) {
	for e := 0; e < EPOCH; e++ {
		m := NewRobinHood(SIZE)
		start := time.Now()
		for i := 0; i < testCount; i++ {
			r := genKey()
			m.Put(r, r)
			if m.Get(r) != r {
				t.Error("Invalid for %s: expected=%v but %v", r, r, m.Get(r))
			}
			m.Erase(genKey())
		}
		elapsed := time.Since(start)

		writer.Write([]string{
			strconv.Itoa(e),
			"RobinHood",
			fmt.Sprintf("%v", m.LoadFactor()),
			fmt.Sprintf("%v", m.DibAverage()),
			fmt.Sprintf("%v", elapsed.Nanoseconds()/1000),
		})
	}

	for e := 0; e < EPOCH; e++ {
		m := NewLinear(SIZE)
		start := time.Now()
		for i := 0; i < testCount; i++ {
			r := genKey()
			m.Put(r, r)
			if m.Get(r) != r {
				t.Error("Invalid for %s: expected=%v but %v", r, r, m.Get(r))
			}
			m.Erase(genKey())
		}
		elapsed := time.Since(start)
		writer.Write([]string{
			strconv.Itoa(e),
			"Linear",
			fmt.Sprintf("%v", m.LoadFactor()),
			fmt.Sprintf("%v", m.DibAverage()),
			fmt.Sprintf("%v", elapsed.Nanoseconds()/1000),
		})
	}
}
