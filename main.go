package main

import (
	"encoding/json"
	"fmt"

	"github.com/hhli/go_study/compile"
)

// main
func main() {
	compile.Walk()

	//files, err := compile.WalkDir("./compile/example", ".go")
	//
	//if err != nil {
	//	log.Printf("遍历目录出现错误:%v", err)
	//	return
	//}
	//
	//compile.DoFind(files)
	extend := EventExtend{EventId: "1111", AccessAggregation: BffMidEventCache{BottomPublishTime: ""}}
	temp, _ := json.Marshal(extend)
	fmt.Println(string(temp))
}

// EventExtend 扩展索引
type EventExtend struct {
	EventId           string           `json:"event_id"`
	BottomPublishUser string           `json:"bottom_publish_user,omitempty"` //底层页发布人
	AccessAggregation BffMidEventCache `json:"access_aggregation"`            //接入层聚合信息，兼容之前的redis cache
}

// BffMidEventCache 中事件缓存（接入层使用）
type BffMidEventCache struct {
	EventId           string `json:"event_id"`
	Event             string `json:"event"`
	CmsId             string `json:"cms_id"`
	HeadImage         string `json:"head_image"`
	EmojiTag          int32  `json:"emoji_tag"`
	EmojiTagName      string `json:"emoji_tag_name"`
	Image             string `json:"image"`
	TopArticleTitle   string `json:"top_article_title"` // 焦点文章标题
	TopArticleType    string `json:"top_article_type"`  // 焦点文章类型
	Heat              int32  `json:"heat"`
	CateName          string `json:"cate_name"`
	SubCateName       string `json:"sub_cate_name"`
	EventSource       int32  `json:"event_source"`
	CoverPicture      string `json:"cover_picture"`
	PictureLittle     string `json:"picture_little"`
	BriefPic          string `json:"brief_pic"`
	BriefTitle        string `json:"brief_title"`
	BottomPublishTime string `json:"bottom_publish_time,omitempty"`
}
