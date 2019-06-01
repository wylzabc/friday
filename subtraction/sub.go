package subtraction

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wylzabc/friday/data"
	"net/http"
)

func Sub(c *gin.Context) {
	var data data.Data
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	data.Result = data.Num1 - data.Num2

	fmt.Println(data)
	c.JSON(http.StatusOK, gin.H{"result": data.Result})
}
