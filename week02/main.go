package main

import (
	"fmt"
	"go_week_task/week02/dao"
)

func main() {
	product, err := dao.QueryProductList()
	if err != nil {
		fmt.Println(fmt.Sprintf("获取数据异常信息[%v]", err))
	}
	fmt.Println(len(product))
}
