package main

import (
	"fmt"
	"math/rand"
	"time"
)

var intChan chan int

// main
func main() {
	intChan = make(chan int, 8)
	go chanInput()

	go func() {
		chanOutput("test1")
	}()

	go func() {
		chanOutput("test2")
	}()

	// 为什么不加sleep就有问题
	time.Sleep(time.Second * 1)
}

// 写入channel
func chanInput() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		i := r.Intn(100)
		fmt.Println(i)
		intChan <- i
		if i > 88 {
			break
		}
	}
}

// 读取channel
func chanOutput(goro string) {
	for i := range intChan {
		fmt.Println(fmt.Sprintf("%s:%d", goro, i))
	}
}

// ArticleDetail 文章详情
type ArticleDetail struct {
	ID         string `json:"id"`         //文章cmsID
	EntityType string `json:"entityType"` //文章类型 article图文/video视频/live_streaming直播
	TitleOuter string `json:"titleOuter"` //外显标题
	Tag        string `json:"tag"`        //标签
	IsLock     string `json:"isLock"`     //是否锁定
	IsDelete   string `json:"isDelete"`   //是否展示 是否删除
}

type BottomConcernResult struct {
	ArticleList []*ArticleDetail //人工干预的所有文章列表
}
