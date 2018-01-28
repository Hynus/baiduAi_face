package face

import (
	"github.com/chenqinghe/baidu-ai-go-sdk/vision/face"
	"github.com/chenqinghe/baidu-ai-go-sdk/vision"
	"encoding/json"
	"fmt"
)

type JsonMatchResult struct {
	MatResult []MatchResult `json:"result"`
}

type MatchResult struct {
	Score float64 `json:"score"`
}

func Match(pic1Path, pic2Path string) float64 {
	fmt.Println("对比分析中....请稍候...")
	var matRet JsonMatchResult
	client := face.NewFaceClient(BaiduAiApiKey, BaiduAiSecretKey)
	rs, err := client.Match(
		vision.MustFromFile(pic1Path),
		vision.MustFromFile(pic2Path),
		map[string]interface{}{},
	)
	if err != nil {
		panic(err)
	}
	tmpJson, err := rs.ToString()
	if err != nil {
		panic(err)
	}
	json.Unmarshal([]byte(tmpJson),&matRet)
	return matRet.MatResult[0].Score
}

func ShowMatchRet(score float64) {
	fmt.Println("-------------------匹配结果如下-------------------")
	if score > 75.0 {
		fmt.Print("这两张照片是同一个人!")
	} else {
		fmt.Print("这两张照片不是同一个人！")
	}
	fmt.Printf("  置信度为: %0.2f%%", score)
}