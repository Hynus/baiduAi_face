package main

import (
	"baiduAi_face/face"
)

func main() {
	picPath := "E:/ACA/Learning/MyGoLang/MyGo/src/baiduAi_face/local_image/p3.jpg"
	face.ShowAnaRet(face.DetectAndAna(picPath))
	pic1Path := "E:/ACA/Learning/MyGoLang/MyGo/src/baiduAi_face/local_image/p3.jpg"
	pic2Path := "E:/ACA/Learning/MyGoLang/MyGo/src/baiduAi_face/local_image/p4.jpg"
	face.ShowMatchRet(face.Match(pic1Path,pic2Path))
}