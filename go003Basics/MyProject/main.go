package main

import (
	"context"
	"fmt"
	"time"
	"github.com/zhashkevych/scheduler"
)

func main() {
	s := scheduler.NewScheduler()
	s.Add(context.Background(), func(ctx context.Context){
		fmt.Printf("Текущее время: %s\n", time.Now())
	}, time.Second * 1)
	
	time.Sleep(time.Minute)
}


/*
go mod init MyProject
go get github.com/zhashkevych/scheduler

https://stackoverflow.com/questions/66356034/what-is-the-difference-between-go-get-command-and-go-mod-download-command
*/