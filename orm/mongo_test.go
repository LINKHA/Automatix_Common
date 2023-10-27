package orm

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

type AccountConfigModel struct {
	Account  string `bson:"account"`  //账号
	Password string `bson:"password"` //密码
}

func (*AccountConfigModel) Config() (database string, collection string) {
	database = "orm_test"
	collection = "account_config"
	return
}

func (*AccountConfigModel) Pre() {}

func TestOrm(t *testing.T) {
	modelMap := make(map[string]IBaseModel)
	modelMap["AccountConfigModel"] = &AccountConfigModel{}

	Init("mongodb://localhost:27017", modelMap)

	mgo := Get("AccountConfigModel")
	batch := []AccountConfigModel{}

	mgo.Find(context.Background(), bson.M{"account": "test_account"}).All(&batch)

	if len(batch) == 0 {
		mgo.InsertOne(context.Background(), AccountConfigModel{
			Account:  "test_account",
			Password: "test_password",
		})
	}
}
