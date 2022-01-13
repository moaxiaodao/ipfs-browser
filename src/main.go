package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/ipfs/go-ipfs-api"
)

var sh *shell.Shell

//交易结构体(未来的通道)
type Transaction struct {
	Person1      string `json:"person1,omitempty" xml:"person1"`
	Person2      string `json:"person2,omitempty" xml:"person2"`
	Person1money string `json:"person1Money,omitempty" xml:"person1Money"`
	Person2money string `json:"person2Money,omitempty" xml:"person2Money"`
}

//数据上传到ipfs
func UploadIPFS(str string) string {
	sh = shell.NewShell("localhost:5001")

	hash, err := sh.Add(bytes.NewBufferString(str))
	if err != nil {
		fmt.Println("上传ipfs时错误：", err)
	}
	return hash
}

//从ipfs下载数据
func CatIPFS(hash string) string {
	sh = shell.NewShell("localhost:5001")

	read, err := sh.Cat(hash)
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(read)

	return string(body)
}

//通道序列化
func marshalStruct(transaction Transaction) []byte {

	data, err := json.Marshal(&transaction)
	if err != nil {
		fmt.Println("序列化err=", err)
	}
	return data
}

//数据反序列化为通道
func unmarshalStruct(str []byte) Transaction {
	var transaction Transaction
	err := json.Unmarshal(str, &transaction)
	if err != nil {
		fmt.Println("unmarshal err=%v", err)
	}
	return transaction
}

func main() {
	//生成一个交易结构体(未来的通道)
	transaction := Transaction{
		Person1:      "Aaron",
		Person2:      "Bob",
		Person1money: "100",
		Person2money: "200",
	}
	//结构体序列化
	data := marshalStruct(transaction)
	//上传到ipfs
	hash := UploadIPFS(string(data))
	fmt.Println("文件hash是", hash)
	//从ipfs下载数据
	str2 := CatIPFS(hash)
	//数据反序列化
	transaction2 := unmarshalStruct([]byte(str2))

	//验证下数据
	fmt.Println(transaction2)
}
