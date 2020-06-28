//Goの基本的なファイルの読み書き

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//データの内容
	mydata := []byte("All the data I wish to write to a file\n")

	//-----------ファイル作成（書き込み）-------------------
	err := ioutil.WriteFile("myfile.data", mydata, 0777)
	if err != nil {
		fmt.Println(err)
	}

	//-----------ファイル読み込み-------------------
	data, err := ioutil.ReadFile("myfile.data")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(string(data))
	
	//-----------ファイル書き換え(書き込み)-------------------
	f, err := os.OpenFile("myfile.data", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _ , err = f.WriteString("new data that wasn't there originally\n"); err != nil{
		panic(err)
	}
	data, err = ioutil.ReadFile("myfile.data")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(string(data))
}
