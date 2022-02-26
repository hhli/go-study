package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

// main
func main() {
	makeConfig()
}

func makeConfig() {
	f, err := excelize.OpenFile("/Users/lihuihui/Downloads/工厂来源.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("工作表2")
	fmt.Printf("rows len:%d\n", len(rows))

	clueConfig := make(map[string]ClueAggregation, len(rows))
	for _, row := range rows {
		if row[0] == "mid" {
			// 忽略表头
			continue
		}
		if len(row) == 6 {
			mid := row[0]
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
			f, _ := strconv.ParseFloat(weight, 64)
			clue.FinanceSourceCoefficient = f
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
