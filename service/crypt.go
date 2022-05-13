package service

import (
	"PicEnc/config"
	"errors"
	"log"
	"sort"
	"sync"
	"time"
)

type chaosCrypt struct {
	key    float64
	mode   config.EncryptMode
	method config.EncryptMethodType
	func_  config.EncryptMethod
}

type ChaosCryptFactory struct{}

func (f *ChaosCryptFactory) NewChaoCrypt(pass float64, mode config.EncryptMode, method config.EncryptMethodType) (*chaosCrypt, error) {
	if method == config.Logistic && (pass <= 0 || pass >= 1) {
		return nil, errors.New("the password must be in the range of (0,1)")
	}
	c := &chaosCrypt{
		key:    pass,
		mode:   mode,
		method: method,
	}

	switch method {
	case config.Logistic:
		c.func_ = config.LogisticFunc
	default:
		return nil, errors.New("invalid type of encrypt method")
	}

	return c, nil
}

func (c *chaosCrypt) Encrypt(data [][]uint32) [][]uint32 {
	begin := time.Now()
	ret := c.encrypt(data, false)
	if c.mode == config.RowAndLine {
		ret = c.encrypt(ret, true)
	}
	log.Printf("Encrypt image successfully, cost %.3fs", time.Since(begin).Seconds())
	return ret
}

func (c *chaosCrypt) encrypt(data [][]uint32, transpose bool) [][]uint32 {
	m, n := len(data), len(data[0])
	ret := make([][]uint32, m)
	for i := 0; i < m; i++ {
		ret[i] = make([]uint32, n)
	}
	if transpose {
		m, n = n, m
	}
	tmp := c.key
	wg := sync.WaitGroup{}
	wg.Add(m)
	for i := 0; i < m; i++ {
		series, nx := c.genNSeries(tmp, n)
		tmp = nx
		go func(i int) {
			for idx, item := range series {
				if transpose {
					ret[item][i] = data[idx][i]
				} else {
					ret[i][item] = data[i][idx]
				}
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	return ret
}

func (c *chaosCrypt) Decrypt(data [][]uint32) [][]uint32 {
	begin := time.Now()
	var ret [][]uint32
	if c.mode == config.RowAndLine {
		ret = c.decrypt(data, true)
	} else {
		ret = data
	}
	ret = c.decrypt(ret, false)
	log.Printf("Decrypt image successfully, cost %.3fs", time.Since(begin).Seconds())
	return ret
}

func (c *chaosCrypt) decrypt(data [][]uint32, transpose bool) [][]uint32 {
	m, n := len(data), len(data[0])
	ret := make([][]uint32, m)
	for i := 0; i < m; i++ {
		ret[i] = make([]uint32, n)
	}
	if transpose {
		m, n = n, m
	}
	tmp := c.key
	wg := sync.WaitGroup{}
	wg.Add(m)
	for i := 0; i < m; i++ {
		series, nx := c.genNSeries(tmp, n)
		tmp = nx
		go func(i int) {
			for idx, item := range series {
				if transpose {
					ret[idx][i] = data[item][i]
				} else {
					ret[i][idx] = data[i][item]
				}
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	return ret
}

// generate chaotic server, x is init value, n is length.
func (c *chaosCrypt) genNSeries(x float64, n int) ([]int, float64) {
	arr := make([]float64, n)
	tmp := make([]float64, n)
	m := make(map[float64]int, n)
	arr[0] = x
	for i := 1; i < n; i++ {
		x = c.func_(x)
		arr[i] = x
	}

	copy(tmp, arr)
	sort.Float64s(tmp)

	for i, f := range tmp {
		m[f] = i
	}

	t := make([]int, n)
	for i, f := range arr {
		t[i] = m[f]
	}

	return t, x
}
