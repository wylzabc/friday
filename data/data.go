package data

type Data struct {
	Num1   int `json:"num1" binding:"required"`
	Num2   int `json:"num2" binding:"required"`
	Result int
}

type Result struct {
	Result int `json:"result" binding:"required"`
}
