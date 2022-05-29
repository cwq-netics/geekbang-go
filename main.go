package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	//注册ping服务 模拟http接口处理程序
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write([]byte(`pong`))
		if err != nil {
			return
		}
	})
	srv := http.Server{
		Addr:    ":9000",
		Handler: mux,
	}
	g.Go(func() error {
		return srv.ListenAndServe()
	})
	//http server出错 用ctx通知linux signal线程退出
	//linux signal出错 用closeSig通知http server退出
	closeSig := make(chan struct{})
	//监听出错&退出信号 处理服务器关闭流程
	g.Go(func() error {
		select {
		case <-closeSig:
			fmt.Println(`recv close signal`)
			break
		case <-ctx.Done():
			fmt.Println(`recv errgroup done signal`)
			break
		}
		timeoutCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		// 这里不是必须的，但是如果使用 _ 的话静态扫描工具会报错，加上也无伤大雅
		defer cancel()

		fmt.Println("shutting down server...")
		return srv.Shutdown(timeoutCtx)
	})
	//linux signal 信号的注册和处理
	g.Go(func() error {
		quit := make(chan os.Signal, 0)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-ctx.Done():
			return ctx.Err()
		case sig := <-quit:
			close(closeSig)
			return errors.Errorf("get os signal: %v", sig)
		}
	})

	// 主任务 goroutine 等待 pipeline 结束数据流
	err := g.Wait()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("main goroutine done!")
}
