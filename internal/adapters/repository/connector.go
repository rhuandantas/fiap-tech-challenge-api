package repository

import (
	"fiap-tech-challenge-api/internal/core/domain"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/gommon/log"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

type DBConnector interface {
	GetORM() *xorm.Engine
	Close()
}

type MySQLConnector struct {
	engine *xorm.Engine
}

func (m MySQLConnector) GetORM() *xorm.Engine {
	return m.engine
}

func (m MySQLConnector) Close() {
	err := m.engine.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func NewMySQLConnector() DBConnector {
	// TODO put in env vars
	var (
		dbName     = "tech_challenge"
		dbPassword = "12345678"
		dbUser     = "root"
		err        error
	)

	engine, err := xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8", dbUser, dbPassword, dbName))
	if err != nil {
		panic(err)
	}
	engine.ShowSQL(true) // TODO it should come from env
	//engine.Logger().SetLevel(log.DEBUG)
	engine.SetMapper(names.SnakeMapper{})
	if err = syncTables(engine); err != nil {
		log.Fatal("failed to sync tables ", err.Error())
	}

	return &MySQLConnector{
		engine: engine,
	}
}

// syncTables allows us to synchronize our tables on the databases: create, updates, table, columns, indexes
func syncTables(engine *xorm.Engine) error {
	if err := engine.Sync(
		new(domain.Cliente),
		new(domain.Produto),
		new(domain.PedidoDTO),
		new(domain.PedidoProduto),
		new(domain.Status),
	); err != nil {
		return err
	}

	//if err := engine.CreateTables(domain.PedidoDTO{}); err != nil {
	//	return err
	//}

	return nil
}
