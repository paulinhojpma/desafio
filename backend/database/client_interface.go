package database

// "time"

// Planet monta a estrututra do plano

// IDataBase ..
type IDataBase interface {
	connectService(config *OptionsDBClient) error
	GetTransaction(key interface{}) (Producer, error)
	CreateTransaction(produtores []*Producer) error
	ListTransaction() ([]Producer, error)
}

// OptionsCacheClient ..
type OptionsDBClient struct {
	URL    string `json:"url"`
	DBName string `json:"DBName"`
	Driver string
}

// ConfiguraCache
func (o *OptionsDBClient) ConfigDatabase() (*IDataBase, error) {

	var client IDataBase
	switch o.Driver {
	case "postgres":
		gormDB := &gormPostgres{}
		errGorm := gormDB.connectService(o)
		if errGorm != nil {
			return nil, errGorm
		}
		client = gormDB
	}
	return &client, nil
}
