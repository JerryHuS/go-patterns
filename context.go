/**
 * @Author: alessonhu
 * @Description:
 * @File:  context.go
 * @Version: 1.0.0
 * @Date: 2022/4/13 5:26 PM
 */

package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
)

type favContextKey string

func main() {
	wg := &sync.WaitGroup{}
	values := []string{"www.baidu.com", "www.zhihu.com"}
	ctx, cancel := context.WithCancel(context.Background())
	for _, url := range values {
		wg.Add(1)
		subCtx := context.WithValue(ctx, favContextKey("url"), url)
		go reqURL(subCtx, wg)
	}

	wg.Wait()

}

func reqURL(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	url, _ := ctx.Value(favContextKey("url")).(string)
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("stop getting url:%s\n", url)
			return
		default:
			r, err := http.Get(url)

		}
	}
}
