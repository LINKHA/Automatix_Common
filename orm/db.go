package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Database 接口定义了通用的数据库操作方法
type Database interface {
	Connect() error
	FindOne(filter interface{}) (map[string]interface{}, error)
}

// MySQLDB 是对 MySQL 数据库的封装
type MySQLDB struct {
	db *sql.DB
}

// Connect 实现了 Database 接口中的 Connect 方法
func (m *MySQLDB) Connect() error {
	// 设置 MySQL 连接配置
	dsn := "username:password@tcp(hostname:port)/database"

	// 连接到 MySQL
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	m.db = db
	return nil
}

// FindOne 实现了 Database 接口中的 FindOne 方法
func (m *MySQLDB) FindOne(filter interface{}) (map[string]interface{}, error) {
	row := m.db.QueryRowContext(context.Background(), "SELECT * FROM mytable WHERE name=?", filter)

	result := make(map[string]interface{})
	err := row.Scan(result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// MongoDB 是对 MongoDB 数据库的封装
type MongoDB struct {
	client     *mongo.Client
	collection *mongo.Collection
}

// Connect 实现了 Database 接口中的 Connect 方法
func (m *MongoDB) Connect() error {
	// 设置 MongoDB 连接选项
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// 连接到 MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}

	m.client = client

	// 获取集合对象
	m.collection = m.client.Database("mydatabase").Collection("mycollection")

	return nil
}

// FindOne 实现了 Database 接口中的 FindOne 方法
func (m *MongoDB) FindOne(filter interface{}) (map[string]interface{}, error) {
	var result bson.M
	err := m.collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func main() {
	// 创建数据库实例
	mysqlDB := &MySQLDB{}
	mongoDB := &MongoDB{}

	// 连接到数据库
	err := mysqlDB.Connect()
	if err != nil {
		log.Fatal(err)
	}

	err = mongoDB.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// 查询数据
	filter := "Alice"
	mysqlResult, err := mysqlDB.FindOne(filter)
	if err != nil {
		log.Fatal(err)
	}

	mongoResult, err := mongoDB.FindOne(bson.D{{"name", filter}})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MySQL Result:", mysqlResult)
	fmt.Println("MongoDB Result:", mongoResult)
}
