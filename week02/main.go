package main

import (
	"go_week_task/week02/dao"
)

func main() {
	err := dao.QueryProductList("")
	if dao.IsNoRow(err) {
		//往业务方向转换考虑 是否需要做出错误响应
		return
	} else if err != nil {
		//数据库查询出现了问题 可以转换为业务领域错误，也可以继续往上传递
	}
}
