package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"
	"go.uber.org/ratelimit"
)

// main
func main() {
	a := []int{1, 2, 3}
	b := make([]int, 2)
	copy(b, a)
	fmt.Println(b)
}

type Temp struct {
	PopularClusterId string `json:"popular_cluster_id,omitempty"`
}

func testReflect() {
	// 取变量a的反射类型对象
	typeOfA := reflect.TypeOf([]*Sub{})
	// 根据反射类型对象创建类型实例
	article := reflect.New(typeOfA).Interface()
	//err := json.Unmarshal([]byte("[{\"raw_id\":\"xxx\"}]"), article)
	//if err != nil {
	//	fmt.Println(err)
	//}

	// arrayValue.Interface()
	fmt.Println(reflect.ValueOf(article).Elem())
}
func rateLimit() {
	rl := ratelimit.New(100, ratelimit.Per(time.Minute))
	prev := time.Now()
	for i := 0; i < 100; i++ {
		now := rl.Take()
		fmt.Println(i, now.Sub(prev))
		prev = now
	}
	println("=================")

	//limit := rate.Every(time.Minute)
	//r := rate.NewLimiter(limit, 100)
	//prev = time.Now()
	//for i := 0; i < 100; i++ {
	//	now := time.Now()
	//	if err := r.Wait(context.Background()); err == nil {
	//		fmt.Println(i, now.Sub(prev))
	//		prev = now
	//	}
	//
	//}
}

type TNewsDynamicResp struct {
	Status int32  `json:"status"`
	Msg    string `json:"msg"`
	Data   []struct {
		EntityID   string `json:"entity_id"`   // 实体ID
		EntityType string `json:"entity_type"` // 传入的实体类型
		EntityData struct {
			CTR                     string `json:"clk_ctr_2"`
			DeepImgtextDetailTimePV string `json:"deep_imgtext_detail_time_pv"` // 深度消费PV
			DeepVideoTimeVV         string `json:"deep_video_time_vv"`          // 深度消费VV
			Exp                     string `json:"exp_pv_2"`                    // 站内累计曝光量
			Clk                     string `json:"clk_pv_2"`                    // 站内累计点击量
			Play                    string `json:"play_vv_2"`                   // 站内累计播放量
			Share                   string `json:"share_2"`                     // 站内累计分享量
			Comment                 string `json:"comment_2"`                   // 站内累计评论量
		} `json:"entity_data"`
	}
}

func makePutBulk() {
	f, err := excelize.OpenFile("/Users/lihuihui/Downloads/weixin2.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	fmt.Printf("rows len:%d\n", len(rows))

	file, err := os.OpenFile("test.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer func() {
		err = file.Close()
		if err != nil {
			fmt.Printf("close file err:%v\n", err)
		}
	}()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)

	for _, row := range rows {
		mid := row[0]
		name := row[1]
		url := row[2]
		now := time.Now().Format("2006-01-02 15:04:05")
		s1 := fmt.Sprintf("{\"index\":{\"_id\":\"weixin-%s\"}}\n", mid)
		_, _ = write.WriteString(s1)

		str := fmt.Sprintf("{\"mid\":\"%s\",\"created_at\":\"%s\",\"monitor_type\":\"account\",\"id\":\"weixin-%s\",\"status\":1,\"grab_type\":2,\"updated_at\":\"%s\",\"name\":\"%s\",\"source\":\"weixin\", \"url\":\"%s\"}\n", mid, now, mid, now, name, url)
		_, _ = write.WriteString(str)
	}

	// Flush将缓存的文件真正写入到文件中
	_ = write.Flush()
}

func makeConfig2() {
	f, err := excelize.OpenFile("/Users/lihuihui/Desktop/疫情来源.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	fmt.Printf("rows len:%d\n", len(rows))

	areaConfig := make(map[string]Area, len(rows))
	for _, row := range rows {
		fullProv := row[0]
		prov := row[1]
		name := row[2]
		areaConfig[name] = Area{Prov: prov, FullProv: fullProv}
	}

	bytes, _ := json.Marshal(areaConfig)
	err = ioutil.WriteFile("test.txt", bytes, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

// Area 区域信息
type Area struct {
	Prov     string `json:"prov,omitempty"`      // 省份信息 比如河北 西藏 澳门
	FullProv string `json:"full_prov,omitempty"` // 省份完整信息 比如河北省 西藏自治区 澳门特别行政区
}

func makeConfig() {
	f, err := excelize.OpenFile("/Users/lihuihui/Downloads/工厂来源.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("系数-格式化")
	fmt.Printf("rows len:%d\n", len(rows))

	clueConfig := make(map[string]ClueAggregation, len(rows))
	for _, row := range rows {
		if row[0] == "mid" {
			// 忽略表头
			continue
		}

		mid := row[0]
		mid = strings.TrimSuffix(mid, ".0")
		monitorType := row[1]
		name := row[2]
		myType := row[3]
		weight := row[4]
		source := row[5]
		sourceKey := fmt.Sprintf("%s:%s:%s", monitorType, source, mid)
		clue := ClueAggregation{Name: name}
		if myType == "" {
			clue.ClueType = []string{"finance_comprehensive"}
		} else if myType == "媒体" {
			clue.ClueType = []string{"finance_comprehensive", "finance_media"}
		} else if myType == "部委" {
			clue.ClueType = []string{"finance_comprehensive", "finance_ministries"}
		}
		// 如果浮点数转化失败，f=0
		f1, _ := strconv.ParseFloat(weight, 64)
		clue.FinanceSourceCoefficient = f1
		_, ok := clueConfig[sourceKey]
		if ok {
			fmt.Println(sourceKey)
		}

		if source == "weixin" && monitorType != "account" {
			fmt.Printf("row11111 mid:%s\n", strings.Join(row, ","))
		}

		if strings.Contains(source, "wei") && source != "weixin" {
			fmt.Printf("row11111 mid:%s\n", strings.Join(row, ","))
		}
		clueConfig[sourceKey] = clue
	}

	bytes, _ := json.Marshal(clueConfig)
	err = ioutil.WriteFile("test.txt", bytes, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

// ClueAggregation 线索聚合配置
type ClueAggregation struct {
	Name                     string   `json:"name"`                       // 来源名称
	ClueType                 []string `json:"clue_type"`                  // 线索类型
	FinanceSourceCoefficient float64  `json:"finance_source_coefficient"` // 财经来源系数
}

type Article struct {
	Subs []*Sub `json:"subs,omitempty"`
}

type Sub struct {
	RawID                    string   `json:"raw_id,omitempty"`                     // 文章id
	Source                   string   `json:"source,omitempty"`                     // 文章来源
	Title                    string   `json:"title,omitempty"`                      // 文章标题
	PubTime                  string   `json:"pub_time,omitempty"`                   // 文章发布时间
	Mid                      string   `json:"mid,omitempty"`                        // 媒体id
	Author                   string   `json:"author,omitempty"`                     // 媒体名称
	ArtificialTags           string   `json:"artificial_tags,omitempty"`            // 人工标签
	FinanceComprehensiveTags []string `json:"finance_comprehensive_tags,omitempty"` // 财经综合标签
	FinanceMediaTags         []string `json:"finance_media_tags,omitempty"`         // 财经媒体标签
	FinanceMinistriesTags    []string `json:"finance_ministries_tags,omitempty"`    // 财经部委标签
}
