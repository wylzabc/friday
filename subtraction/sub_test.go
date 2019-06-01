package subtraction

import (
	//	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wylzabc/friday/data"
	"github.com/wylzabc/friday/util"
	//"strconv"
	"encoding/json"
	"testing"
)

var router *gin.Engine

func Init() {
	gin.SetMode(gin.DebugMode)
	router = gin.Default()
	router.POST("/sub/sub", Sub)
}

func TestSub(t *testing.T) {
	Init()
	uri := "/sub/sub"
	param := make(map[string]interface{})
	param["num1"] = 1
	param["num2"] = 2

	body := util.PostJson(uri, param, router)
	var result data.Result
	err := json.Unmarshal(body, &result)
	if err != nil {
		t.Errorf("解析响应出错，err:%v\n", err)
	}
	if result.Result != -1 {
		t.Errorf("响应数据不符，result: %v\n", result)
	}
}
