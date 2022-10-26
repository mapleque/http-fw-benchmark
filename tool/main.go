package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

var (
	url      *string
	minc     *int
	maxc     *int
	stepc    *int
	duration *int
	debug    *bool
)

func main() {
	url = flag.String("url", "http://127.0.0.1:8080", "url")
	minc = flag.Int("minc", 100, "min concurrency")
	maxc = flag.Int("maxc", 500, "max concurrency")
	stepc = flag.Int("stepc", 100, "concurrency increase step")
	duration = flag.Int("d", 10, "each thread send request duration (s)")
	debug = flag.Bool("v", false, "print debug message")
	flag.Parse()

	var stat Stat
	stat.title()
	stat.hr()
	for i := *minc; i <= *maxc; i += *stepc {
		do(*url, i, *duration)
		time.Sleep(2 * time.Second)
	}
	stat.hr()
}

func do(url string, c int, d int) {
	var wg sync.WaitGroup
	var stat Stat
	var stub int = 1000000
	stat.size(c, c*stub)

	for i := 0; i < c; i++ {
		wg.Add(1)
		go func(url string, id int) {
			defer wg.Done()
			// each thread sending req 10 second
			next := make(chan bool, 1)
			ctx, cancel := context.WithTimeout(
				context.Background(),
				time.Duration(d)*time.Second,
			)
			defer cancel()
			if *debug {
				log.Printf("thread %d start", id)
			}
			id *= stub
			var seq int = 0
			var done = make(chan bool)
			go func() {
				defer close(done)
				for {
					select {
					case <-next:
						seq++
						stat.start(id + seq)
						if doReq(url) {
							stat.ok(id + seq)
						} else {
							stat.err(id + seq)
						}
						next <- true
					case <-ctx.Done():
						if *debug {
							log.Printf("thread %d done", id)
						}
						return
					}
				}
			}()
			next <- true
			<-done
		}(url, i)
	}
	wg.Wait()

	stat.print()
}

func doReq(url string) bool {
	resp, err := http.Get(url)
	if err != nil {
		if *debug {
			log.Printf("request %s with err: %v", url, err)
		}
		return false
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		if *debug {
			log.Printf("request %s with response status %d", url, resp.StatusCode)
		}
		return false
	}
	return true
}

type Stat struct {
	c int
	s []int64 // start time or latency
	e []int   // err counter
}

func (s *Stat) size(c, st int) {
	s.c = c
	s.s = make([]int64, st)
	s.e = make([]int, st)
}

func (s *Stat) start(id int) {
	if *debug {
		log.Printf("req start %d", id)
	}
	s.s[id] = time.Now().UnixMilli()
}

func (s *Stat) ok(id int) {
	if *debug {
		log.Printf("req ok %d", id)
	}
	s.s[id] = time.Now().UnixMilli() - s.s[id]
}

func (s *Stat) err(id int) {
	if *debug {
		log.Printf("req err %d", id)
	}
	s.s[id] = 0
	s.e[id] = 1
}

func cnt(arr []int64) int {
	var s int = 0
	for _, e := range arr {
		if e > 0 {
			s++
		}
	}
	return s
}

func avg(arr []int64) int64 {
	var s int64 = 0
	var t int64 = 0
	for _, e := range arr {
		s += e
		if e > 0 {
			t++
		}
	}
	if t == 0 {
		return -1
	}
	return s / t
}

func sum(arr []int) int {
	var s int = 0
	for _, e := range arr {
		s += e
	}
	return s
}

func (s *Stat) print() {
	req := cnt(s.s)
	avg := int(avg(s.s))
	qps := req / *duration
	err := sum(s.e)
	var er int = 0
	if req > 0 {
		er = err * 100 / req
	}
	fmt.Printf(
		"%d\t%d\t\t%d\t\t%d\t\t\t%d%%(%d)\n",
		qps, s.c, req, avg, er, err,
	)
}

func (s *Stat) title() {
	fmt.Println("QPS\tConcurrence\tTotal req\tAvg latency(ms)\t\tError rate(num)")
}

func (s *Stat) hr() {
	fmt.Println(strings.Repeat("-", 74))
}
