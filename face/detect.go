package face

import (
	"github.com/chenqinghe/baidu-ai-go-sdk/vision"
	sdk_face "github.com/chenqinghe/baidu-ai-go-sdk/vision/face"
	"encoding/json"
	"fmt"
)

var (
	BaiduAiApiKey = "grF7SveGj2LIwIYobfNBPO9n"
	BaiduAiSecretKey = "iUDZHYH8qhWWWuUc9m9A7oPcmzdzkRGH"
)

type JsonDectResult struct {
	DectResult []DectResult `json:"result"`
}

type DectResult struct {
	Age float64 `json:"age"`
	Expression int64 `json:"expression"`
	Beauty float64 `json:"beauty"`
	FaceShape []OneFaceShape `json:"faceshape"`
	Gender string `json:"gender"`
	Glasses int64 `json:"glasses"`
	Race string `json:"race"`
}

type OneFaceShape struct {
	Type string `json:"type"`
	Probability float64 `json:"probability"`
}

type FaceAnaRet struct {
	Gender string
	Age float64
	Beauty float64
	Expression int64
	FaceType string
	FaceTypeProb float64
	Glasses int64
	Race string
}

func DetectAndAna(picPath string) FaceAnaRet {
	fmt.Println("正在识别分析中.....请稍候....")
	var (
		myRet JsonDectResult
		faceAna FaceAnaRet
	)
	client := sdk_face.NewFaceClient(BaiduAiApiKey, BaiduAiSecretKey)
	options := map[string]interface{}{
		"max_face_num": 10,
		"face_fields":  "age,beauty,expression,faceshape,gender,glasses,race",
	}
	rs, err := client.DetectAndAnalysis(
		vision.MustFromFile(picPath),
		options,
	)
	if err != nil {
		panic(err)
	}
	tmpJson, err := rs.ToString()
	if err != nil {
		panic(err)
	}
	json.Unmarshal([]byte(tmpJson), &myRet)
	faceAna.Gender = myRet.DectResult[0].Gender
	faceAna.Age = myRet.DectResult[0].Age
	faceAna.Beauty = myRet.DectResult[0].Beauty
	faceAna.Expression = myRet.DectResult[0].Expression
	faceAna.Glasses = myRet.DectResult[0].Glasses
	faceAna.Race = myRet.DectResult[0].Race
	ftSliceOrder := orderDescFaceShapeSlice(myRet.DectResult[0].FaceShape)
	faceAna.FaceType = ftSliceOrder[0].Type
	faceAna.FaceTypeProb = ftSliceOrder[0].Probability
	return faceAna
}

func orderDescFaceShapeSlice(oriSlice []OneFaceShape) []OneFaceShape {
	newSlice := make([]OneFaceShape, len(oriSlice), cap(oriSlice))
	copy(newSlice, oriSlice)
	for i := 0 ; i < len(newSlice) - 1 ; i++ {
		for j := 0 ; j < len(newSlice) - 1 - i; j++ {
			if newSlice[j].Probability < newSlice[j+1].Probability {
				newSlice[j], newSlice[j+1] = newSlice[j+1], newSlice[j]
			}
		}
	}
	return newSlice
}

func ShowAnaRet(ret FaceAnaRet) {
	fmt.Println("-------------------识别结果如下-------------------")
	if ret.Gender == "male" {
		fmt.Println("性别：男")
		fmt.Printf("帅气指数：%0.2f\n",ret.Beauty)
	} else {
		fmt.Println("性别：女")
		fmt.Printf("美丽指数：%0.2f\n",ret.Beauty)
	}
	fmt.Printf("年龄约为：%0.0f\n",ret.Age)
	switch ret.Expression {
	case 0:
		fmt.Println("表情：不笑")
	case 1:
		fmt.Println("表情：微笑")
	case 2:
		fmt.Println("表情：大笑")
	default:
		fmt.Println("表情：未知")
	}
	switch ret.Glasses {
	case 0:
		fmt.Println("眼镜：不戴眼镜")
	case 1:
		fmt.Println("眼镜：普通眼镜")
	case 2:
		fmt.Println("眼镜：墨镜")
	default:
		fmt.Println("眼镜：未知")
	}
	switch ret.Race {
	case "yellow":
		fmt.Println("人种：黄种人")
	case "white":
		fmt.Println("人种：白种人")
	case "black":
		fmt.Println("人种：黑种人")
	case "arabs":
		fmt.Println("人种：阿拉伯人")
	default:
		fmt.Println("人种：未知")
	}
	switch ret.FaceType {
	case "square":
		fmt.Print("脸型：方形")
	case "triangle":
		fmt.Print("脸型：三角形")
	case "oval":
		fmt.Print("脸型：椭圆形")
	case "heart":
		fmt.Print("脸型：心形")
	case "round":
		fmt.Print("脸型：圆形")
	default:
		fmt.Print("脸型：未知")
	}
	fmt.Printf(" (置信度为：%0.2f%%)", ret.FaceTypeProb*100)
}