package orm

type IBaseModel interface {
	Config() (account string, passward string)
	Pre()
}

type IDB interface {
	Init(mongoUrl string, setModelMap map[string]IBaseModel)

	Stop()

	Get(model_name string)
}
