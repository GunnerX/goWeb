package repositories

import (
	"database/sql"
	"imooc-product/datamodels"
)

// 1.开发对应接口
// 2.实习对应接口

type IProduct interface {
	// 连接数据库
	Conn() error
	// 插入数据
	Insert(*datamodels.Product) (int64, error)
	Delete(int64 ) bool
	Update(*datamodels.Product) error
	SelectByKey(int64)(*datamodels.Product, error)
	SelectAll()([]*datamodels.Product, error)
}

type ProductManager struct {
	table string
	mysqlConn *sql.DB
}

func NewProductManager(table string, db *sql.DB) IProduct {
	return &ProductManager{table, db}
}

func (p *ProductManager) Conn() error {

}







