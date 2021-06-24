package dml

import (
	"doServer/database"
	"fmt"
	"github.com/gin-gonic/gin"
)

func QueryDemo(ctx *gin.Context) {
	sql := "select * from sys_user;"

	listMap, err := database.FindAllMap(sql)
	if err != nil {
		fmt.Println("err:" + err.Error())
		ctx.JSON(200, gin.H{"code": "1", "data": ""})
		return
	}

	//for _, v := range listMap {
	//	fmt.Println("value:", v)
	//}
	ctx.JSON(200, gin.H{"code": "0", "data": listMap})
}
