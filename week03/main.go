package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

/*
	1. 基于 errgroup 实现一个 http server 的启动和关闭 ，
       以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出
*/

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	//定义单缓冲信号量chan
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGQUIT)

	var myServer myServer

	srv := http.Server{
		Handler: myServer,
		Addr:    "127.0.0.1:8080",
	}
	http.Handle("/hello", myServer)

	//启动监听httpServer
	g.Go(func() error {
		defer log.Println("g_str_listen return")
		return srv.ListenAndServe()
	})

	//监听到error return
	g.Go(func() error {
		select {
		case <-ctx.Done():
			log.Println("srv.Shutdown return")
			return srv.Shutdown(ctx)
		}
	})

	//处理信号量err
	g.Go(func() error {
		for {
			select {
			case <-signalCh:
				return errors.New("system signal return")
			}
		}
	})

	if err := g.Wait(); err != nil {
		log.Println("group error: ", err)
	}

}

type myServer struct{}

func (server myServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.String() {
	case "/hello":
		fmt.Fprintf(w, "hello")
	default:
		fmt.Fprintf(w, "unknow http")
	}

}
