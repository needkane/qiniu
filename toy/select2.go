package main                                                                                                                       
 
import (
    "fmt"
    "time"
)
 
/*
  每隔2.5秒 ，三个随机顺序 出现
  若三个case都是 <-t1，则每隔2.5 秒，随机出现一个
  外部变量(t1、t2、t3) 优先级高于case <-time.Tick(time.Millisecond * 2500)
  若case均为<-time.Tick(time.Millisecond * 2500) ,则每隔2.5秒只出现up
  若时间不一致，则优先出现时间短的，若是三个不一致时间的外部变量，则按外部变量时间顺序出现
*/
func main() {
    t1 := time.Tick(time.Millisecond * 2500)
    t2 := time.Tick(time.Millisecond * 2500)
    t3 := time.Tick(time.Millisecond * 2500)
    var count int 
    for {
        select {
        case <-t1:
            fmt.Println("up")
            count++
            fmt.Println("count--->", count)
        case <-t3:
            fmt.Println("middle")
            count++
            fmt.Println("count--->", count)
        case <-t2:
            fmt.Println("down")
            count++
            fmt.Println("count--->", count)
        }   
    }   
 
}
