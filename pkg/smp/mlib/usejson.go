package library

import (
	"encoding/json"
	"fmt"
	"os"
)

func wirtejson(elements []MusicEntry) {

	// 将切片转换为json格式的字节切片
	data, err := json.Marshal(elements)
	if err != nil {
		fmt.Println("json编码错误:", err)
		return
	}

	// 将字节切片写入文件
	err = os.WriteFile("elements.json", data, 0644)
	if err != nil {
		fmt.Println("文件写入错误:", err)
		return
	}

	fmt.Println("文件写入成功")

}

func readjson() (elements2 []MusicEntry) {

	// 打开文件
	file, err := os.Open("elements.json")
	if err != nil {
		fmt.Println("文件打开错误:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	fmt.Println("文件打开成功")

	// 创建一个json解码器
	decoder := json.NewDecoder(file)

	// 创建一个空的切片来接收解码后的元素
	//var elements2 []MusicEntry

	// 从文件中解码json数据到切片中
	err = decoder.Decode(&elements2)
	if err != nil {
		fmt.Println("json解码错误:", err)
		return
	}

	fmt.Println("文件读取成功")

	// 打印切片中的元素
	for _, e := range elements2 {
		fmt.Printf("%+v\n", e)
	}
	return

}
