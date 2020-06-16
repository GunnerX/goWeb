package repositories

import (
	"database/sql"
	"imooc-product/common"
	"imooc-product/datamodels"
	"strconv"
)

// 1.开发对应接口
// 2.实现对应接口

type IProduct interface {
	// 连接数据库
	Conn() (error)
	Insert(*datamodels.Product) (int64, error)
	Delete(int64) bool
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

// 数据库连接
func (p *ProductManager) Conn() (err error) {
	if p.mysqlConn == nil {
		mysql, err := common.NewMysqlConn()
		if err != nil {
			return err
		}
		p.mysqlConn = mysql
	}
	if p.table == "" {
		p.table = "product"
	}
	return
}

// 插入
func (p *ProductManager) Insert(product *datamodels.Product) (productId int64, err error) {
	// 1.判断连接是否存在
	if err := p.Conn(); err != nil {
		return
	}
	// 2.准备sql
	sql := "INSERT product SET productName = ?, productNum = ?, productImage = ?, productUrl = ?"
	stmt, err := p.mysqlConn.Prepare(sql)
	if err != nil {
		return 0, err
	}

	// 3.传入参数
	result, err := stmt.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductUrl)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

// 删除
func (p *ProductManager) Delete(productID int64) bool {
	// 1.判断连接是否存在
	if err := p.Conn(); err != nil {
		return false
	}
	sql := "DELETE FROM product WHERE ID = ?"
	stmt, err := p.mysqlConn.Prepare(sql)
	if err != nil {
		return false
	}
	_, err  = stmt.Exec(productID)
	if err != nil {
		return false
	}
	return true
}

// 更新
func (p *ProductManager) Update(product *datamodels.Product) error {
	// 1.判断连接是否存在
	if err := p.Conn(); err != nil {
		return err
	}

	sql := "UPDATE product set productName = ?, productNum = ?, " +
		"productImage = ?, productUrl = ? where ID = " + strconv.FormatInt(product.ID, 10)
	stmt, err := p.mysqlConn.Prepare(sql)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductUrl)
	if err != nil {
		return err
	}
	return nil
}

// 根据商品ID查询商品
func (p *ProductManager) SelectByKey(productID int64)(productResult *datamodels.Product, err error) {
	if err = p.Conn(); err != nil {
		return &datamodels.Product{}, err
	}
	sql := "SELECT * FROM" + p.table + "WHERE ID =" + strconv.FormatInt(productID, 10)
	row, errRow := p.mysqlConn.Query(sql)
	defer row.Close()
	if errRow != nil {
		return &datamodels.Product{}, errRow
	}
	result := common.GetResultRow(row)
	if len(result) == 0 {
		return &datamodels.Product{}, nil
	}
	common.DataToStructByTagSql(result, productResult)
	return

}
// 获取所有商品
func (p *ProductManager) SelectAll()(productArray []*datamodels.Product, errProduct error){
	if err := p.Conn(); err != nil {
		return nil, err
	}
	sql := "SELECT * from " + p.table
	rows, err := p.mysqlConn.Query(sql)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	result := common.GetResultRows(rows)
	if len(result) == 0 {
		return nil, nil
	}

	for _, v := range result {
		product := &datamodels.Product{}
		common.DataToStructByTagSql(v, product)
		productArray = append(productArray, product)
	}
	return
}







