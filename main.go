package main

import (
	"imooc-product/common"
	"imooc-product/datamodels"
)

func main() {
	data := map[string]string{"ID": "1", "productName": "immoc测试结构体", "productNum":"2", "productImage":"123", "productUrl":"http://url"}
	product := &datamodels.Product{}
	common.DataToStructByTagSql(data, product)


}