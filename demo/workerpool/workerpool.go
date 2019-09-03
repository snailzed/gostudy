//workerpool模型：
//    1、使用生产消费者模式简单有效；
//    2、控制goroutine数量，防止goroutine泄露和暴增；
//    3、基于通道chan和goroutine协程实现
//
//                          worker
// 任务输入（job chan） ---》 worker   --》 结果输出（result Chan）
//                          worker
//
package workerpool

import (
	"fmt"
	"math/rand"
)

type Job struct {
	Id     int
	Number int
}

type Result struct {
	Id     int
	Number int
	Ret    int
}

func Dojob(jobChan chan *Job, retChan chan *Result) {
	for job := range jobChan {
		number := job.Number
		var ret int
		for {
			if number != 0 {
				temp := number % 10
				number = number / 10
				ret += temp
			} else {
				break
			}
		}
		retChan <- &Result{
			Id:     job.Id,
			Number: job.Number,
			Ret:    ret,
		}
	}
}
func StartWroker(jobChan chan *Job, retChan chan *Result, workerNum int) {
	for i := 0; i < workerNum; i++ {
		go Dojob(jobChan, retChan)
	}
}

func RetConsumer(retChan chan *Result) {
	for ret := range retChan {
		fmt.Printf("job[id=%d],number = %d,ret = %d\n", ret.Id, ret.Number, ret.Ret)
	}
}
func Run() {
	jobChan := make(chan *Job, 1000)
	retChan := make(chan *Result, 1000)
	//启动多个worker，其实是消费者个数
	StartWroker(jobChan, retChan, 10)
	go RetConsumer(retChan)
	var i int
	for {
		i++
		number := rand.Int()
		job := &Job{
			Id:     i,
			Number: number,
		}
		jobChan <- job
	}

}
