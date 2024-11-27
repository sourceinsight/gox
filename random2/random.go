package random2

import (
	"math/rand"
	"sync"
	"time"
)

type sourcePool struct {
	sync.Pool
}

func (s *sourcePool) Int63() int64 {
	v := s.Get()
	r := v.(rand.Source).Int63()
	s.Put(v)
	return r
}

func (s *sourcePool) Seed(seed int64) {
	v := s.Get()
	v.(rand.Source).Seed(seed)
	s.Put(v)
}

func (s *sourcePool) Uint64() uint64 {
	v := s.Get()
	r := v.(rand.Source64).Uint64()
	s.Put(v)
	return r
}

func newSourcePool() *sourcePool {
	p := &sourcePool{}
	p.New = func() interface{} {
		return rand.NewSource(time.Now().UnixNano()) // nolint:gosec
	}
	return p
}

func New() *rand.Rand {
	return rand.New(newSourcePool())
}

var defaultRand = New()

func Int() int {
	return defaultRand.Int()
}

func Intn(n int) int {
	return defaultRand.Intn(n)
}

func Int31() int32 {
	return defaultRand.Int31()
}

func Int31n(n int32) int32 {
	return defaultRand.Int31n(n)
}

func Int63() int64 {
	return defaultRand.Int63()
}

func Int63n(n int64) int64 {
	return defaultRand.Int63n(n)
}

func Uint32() uint32 {
	return defaultRand.Uint32()
}

func Uint64() uint64 {
	return defaultRand.Uint64()
}

func Float64() float64 {
	return defaultRand.Float64()
}

func Float32() float32 {
	return defaultRand.Float32()
}

func ExpFloat64() float64 {
	return defaultRand.ExpFloat64()
}

func NormFloat64() float64 {
	return defaultRand.NormFloat64()
}

func Perm(n int) []int {
	return defaultRand.Perm(n)
}

func Shuffle(n int, swap func(i, j int)) {
	defaultRand.Shuffle(n, swap)
}

func Read(p []byte) (n int, err error) {
	return defaultRand.Read(p)
}

func Prob(prob float64) bool {
	if prob <= 0 {
		return false
	}
	if prob >= 1 {
		return true
	}
	return defaultRand.Float64() < prob
}
