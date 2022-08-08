package json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

var p = fmt.Println

func TestJsonMarshalBaseDataTypes(t *testing.T) {
	j1, _ := json.Marshal(true)
	p(string(j1))
	j2, _ := json.Marshal(1)
	p(string(j2))
	j3, _ := json.Marshal(2.34)
	p(string(j3))
	j4, _ := json.Marshal("string")
	p(string(j4))
	j5, _ := json.Marshal([]string{"1", "b", "c"})
	p(string(j5))
	j6, _ := json.Marshal(map[string]int{"a": 1, "b": 3, "c": 2})
	p(string(j6))
	j7, _ := json.Marshal(map[string]interface{}{"a": 1, "b": 3, "c": 2})
	p(string(j7))
}

func TestJsonMarshalStruct(t *testing.T) {
	type JsonResp struct {
		Data    interface{} `json:"data"`
		Code    int         `json:"code"`
		Message string      `json:"message"`
	}

	type JsonPageResp struct {
		JsonResp JsonResp // 组合composition
		PageSize int
		PageNum  int
	}

	type JsonPageRespEmbedded struct {
		JsonResp // 嵌入embedded json格式与组合不同
		PageSize int
		PageNum  int
	}
	jr := JsonResp{
		Data:    map[string]interface{}{"b": 1, "a": "1"},
		Code:    200,
		Message: "ok",
	}

	pageRespEmbedded := JsonPageRespEmbedded{PageNum: 10, PageSize: 1, JsonResp: jr}
	pageResp := JsonPageResp{PageNum: 10, PageSize: 1, JsonResp: jr}

	if pr1, err := json.Marshal(&jr); err == nil {
		p("  结果集，不分页：", string(pr1))
	}
	if pr1, err := json.Marshal(&pageResp); err == nil {
		p("组合结果集，分页：", string(pr1))
	}
	if pr1, err := json.Marshal(&pageRespEmbedded); err == nil {
		p("嵌入结果集，分页：", string(pr1))
	}
}

func TestJsonUnMarshal(t *testing.T) {
	byt := []byte(`{"num":6.13,"strs":[1,"2"]}`)
	var bytRes map[string]interface{}
	err := json.Unmarshal(byt, &bytRes)
	if err != nil {
		return
	}
	fmt.Println("unmarshal map:", bytRes)
	fmt.Println("unmarshal map exchange:", bytRes["num"].(float64), bytRes["strs"].([]interface{})[0], bytRes["strs"].([]interface{})[1])

	type byteResult struct {
		Num  float64       `json:"num"`
		Strs []interface{} `json:"strs"`
	}
	var byResult byteResult
	if err := json.Unmarshal(byt, &byResult); err != nil {
		return
	}
	fmt.Println("unmarshal struct:", byResult)
	byResultJson, _ := json.Marshal(byResult)
	fmt.Println("unmarshal struct to json:", string(byResultJson))

	// json编码写入到流中（文件...）
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	err = enc.Encode(d)
	if err != nil {
		return
	}
}
