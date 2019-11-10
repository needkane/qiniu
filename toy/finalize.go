package main

import (
	"log"
	"runtime"
	"time"
)

type Road int

func findRoad(r *Road) {

	log.Println("road:", *r)
}

func entry() {
	var rd Road = Road(999)
	r := &rd

	runtime.SetFinalizer(r, findRoad)
}

func main() {

	entry()

	for i := 0; i < 20; i++ {
		time.Sleep(time.Second)
		runtime.GC()
	}

}

/*
sometimes success
gc 1 @0.002s 9%: 0.10+2.6+0.005 ms clock, 0.41+0.069/1.7/0.82+0.020 ms cpu, 4->6->6 MB, 5 MB goal, 4 P
gc 2 @0.009s 10%: 0.007+2.8+0.004 ms clock, 0.030+0.085/2.6/0.21+0.016 ms cpu, 8->9->9 MB, 12 MB goal, 4 P
gc 3 @0.015s 7%: 0.007+7.1+0.16 ms clock, 0.031+0.42/0.21/4.8+0.64 ms cpu, 15->16->15 MB, 18 MB goal, 4 P
gc 4 @0.078s 3%: 0.003+10+0.004 ms clock, 0.015+0.046/7.3/4.9+0.016 ms cpu, 26->30->28 MB, 30 MB goal, 4 P
gc 1 @1.001s 0%: 0.12+0.22+0.002 ms clock, 0.49+0/0.079/0.004+0.010 ms cpu, 0->0->0 MB, 4 MB goal, 4 P (forced)
gc 2 @2.002s 0%: 0.18+0.23+0.055 ms clock, 0.72+0/0.092/0.023+0.22 ms cpu, 0->0->0 MB, 4 MB goal, 4 P (forced)
2019/11/10 04:54:01 road: 999
gc 3 @3.003s 0%: 0.18+0.23+0.002 ms clock, 0.75+0/0.10/0.010+0.011 ms cpu, 0->0->0 MB, 4 MB goal, 4 P (forced)
gc 4 @4.005s 0%: 0.042+0.29+0.003 ms clock, 0.17+0/0.14/0+0.012 ms cpu, 0->0->0 MB, 4 MB goal, 4 P (forced)
gc 5 @5.006s 0%: 0.97+1.1+0.003 ms clock, 3.8+0/0.16/0+0.012 ms cpu, 0->0->0 MB, 4 MB goal, 4 P (forced)
gc 6 @6.010s 0%: 0.087+0.18+0.003 ms clock, 0.34+0/0.066/0.043+0.012 ms cpu, 0->0->0 MB, 4 MB goal, 4 P (forced)
gc 7 @7.011s 0%: 0.056+0.25+0.003 ms clock, 0.22+0/0.091/0.005+0.012 ms cpu, 0->0->0 MB, 4 MB goal, 4 P (forced)
gc 8 @8.013s 0%: 0.12+0.21+0.002 ms clock, 0.48+0/0.10/0.017+0.009 ms cpu, 0->0->0 MB, 4 MB goal, 4 P (forced)
*/
