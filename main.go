package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

// main
func main() {
	makePutBulk()
}

func makePutBulk() {
	f, err := excelize.OpenFile("/Users/lihuihui/Downloads/weixin.xlsx")
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
		if len(row) == 6 {
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

			clueConfig[sourceKey] = clue
		} else {
			fmt.Printf("row mid:%s\n", strings.Join(row, ","))
		}
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
