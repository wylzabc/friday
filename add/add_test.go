package add

import (
	//	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wylzabc/friday/util"
	//"strconv"
	"encoding/json"
	"github.com/wylzabc/friday/data"
	"testing"
)

var router *gin.Engine

func Init() {
	gin.SetMode(gin.DebugMode)
	router = gin.Default()
	router.POST("/add/add", Add)
	router.POST("/add/absadd", AbsAdd)
}

func TestAdd(t *testing.T) {
	Init()
	uri := "/add/add"
	param := make(map[string]interface{})
	param["num1"] = 1
	param["num2"] = 2

	body := util.PostJson(uri, param, router)
	var result data.Result
	err := json.Unmarshal(body, &result)
	if err != nil {
		t.Errorf("解析响应出错，err:%v\n", err)
	}
	if result.Result != 3 {
		t.Errorf("响应数据不符，result: %v\n", result)
	}
}

var datas = []data.Data{
	{1, 2, 3},
	{1, -2, 3},
	{-1, 2, 3},
	{-1, -2, 3},
}

func TestAbsAdd(t *testing.T) {
	Init()
	for _, testData := range datas {
		uri := "/add/absadd"

		param := make(map[string]interface{})
		param["num1"] = testData.Num1
		param["num2"] = testData.Num2

		body := util.PostJson(uri, param, router)

		var result data.Result
		err := json.Unmarshal(body, &result)
		if err != nil {
			t.Errorf("解析响应出错，err:%v\n", err)
		}
		if result.Result != testData.Result {
			t.Errorf("响应数据不符，result: %v\n", result)
		}
	}

}
