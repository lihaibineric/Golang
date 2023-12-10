package main

import (
	"encoding/json"
	"fmt"
)

// 如何将结构体转化成json
type Movie struct {
	Title  string   `json:"title"` //就是告诉json库对应的变量的标签名字是这个
	Year   int      `json:"year"`
	Price  int      `json:"rmb"`
	Actors []string `json:"actors"` //slice `json:"actors"`
}

func main() {
	movie := Movie{"喜剧之王", 2000, 10, []string{"zhouxingchi", "zhangbozhi"}}

	//编码的过程就是将 struct-->json

	jsonStr, err := json.Marshal(movie)

	//返回两个字段

	if err != nil {
		fmt.Println("json marshal error", err)
		return
	}
	fmt.Printf("jsonStr=%s\n", jsonStr)
	/* jsonStr={"title":"喜剧之王","year":2000,"rmb":10,"Actors":["zhouxingchi","zhangbozhi"]} */

	//解码过程：json-->struct
	//jsonStr={"title":"喜剧之王","year":2000,"rmb":10,"actors":["zhouxingchi","zhangbozhi"]}

	myMovie := Movie{}
	//注意这个地方传递的参量是引用
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("json unmashal error", err)
		return
	}
	fmt.Printf("%v\n", myMovie)
	//返回结构体
	//{喜剧之王 2000 10 [zhouxingchi zhangbozhi]}
}
