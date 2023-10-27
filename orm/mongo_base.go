package orm

import (
	"context"

	"github.com/qiniu/qmgo"
)

// type Database struct {
// 	Mongo *mongo.Client
// }

// var DB *Database
var ModelMap map[string]*qmgo.Collection
var ModelsMap map[string]IBaseModel
var Client *qmgo.Client

// 初始化
func Init(mongoUrl string, setModelMap map[string]IBaseModel) {
	ctx := context.Background()
	Client, _ := qmgo.NewClient(ctx, &qmgo.Config{Uri: mongoUrl})
	ModelMap = make(map[string]*qmgo.Collection)
	ModelsMap = make(map[string]IBaseModel)
	for k, v := range setModelMap {
		ModelsMap[k] = v
		database, collection := v.Config()
		ModelMap[k] = Client.Database(database).Collection(collection)
		v.Pre()
	}
}

func Stop() {
	defer func() {
		Client.Close(context.Background())
	}()
}

func Get(model_name string) *qmgo.Collection {
	return ModelMap[model_name]
}
