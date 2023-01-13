package create

import "fmt"

type OrderMainDAO interface {
	SaveOrderMain()
}

type OrderDetailDao interface {
	SaveOrderDetail()
}

type DAOFactory interface {
	CreateOrderMainDao() OrderMainDAO
	CreateOrderDetailDao() OrderDetailDao
}

type RDBMainDAO struct {
}

func (*RDBMainDAO) SaveOrderMain() {
	fmt.Print("rbd main save\n")
}

type RDBDetailDAO struct {
}

func (*RDBDetailDAO) SaveOrderDetail() {
	fmt.Print("rdb detail save\n")
}

type RDBDAOFactory struct {
}

func (*RDBDAOFactory) CreateOrderMainDao() OrderMainDAO {
	return &RDBMainDAO{}
}

func (*RDBDAOFactory) CreateOrderDetailDao() OrderDetailDao {
	return &RDBDetailDAO{}
}

type XMLMainDAO struct {
}

func (*XMLMainDAO) SaveOrderMain() {
	fmt.Print("xml main save\n")
}

type XMLDetailDAO struct {
}

func (*XMLDetailDAO) SaveOrderDetail() {
	fmt.Print("xml detail save")
}

type XMLDAOFactory struct {
}

func (*XMLDAOFactory) CreateOrderMainDao() OrderMainDAO {
	return &XMLMainDAO{}
}

func (*XMLDAOFactory) CreateOrderDetailDao() OrderDetailDao {
	return &XMLDetailDAO{}
}

func getMainAndDetail(factory DAOFactory) {
	factory.CreateOrderMainDao().SaveOrderMain()
	factory.CreateOrderDetailDao().SaveOrderDetail()
}

func ExampleRdbFactory() {
	var factory DAOFactory
	factory = &RDBDAOFactory{}
	getMainAndDetail(factory)
}

func ExampleXmlFactory() {
	var factory DAOFactory
	factory = &XMLDAOFactory{}
	getMainAndDetail(factory)
}
