package settings

import "os"

type mongo struct {
	Host        string
	DefaultDB   string
	Collections collections
}

type collections struct {
	Profile string
	Role    string
	Address string
	Avatar  string
}

var singletonMongo *mongo

func InitializeMongoDB() *mongo {
	if singletonMongo != nil {
		return singletonMongo
	}

	lock.Lock()
	defer lock.Unlock()

	singletonMongo = &mongo{
		Host:      os.Getenv("MONGODB_HOST"),
		DefaultDB: os.Getenv("DEFAULT_DB"),
		Collections: collections{
			Profile: "profiles",
			Role:    "roles",
			Address: "addresses",
			Avatar:  "avatars",
		},
	}

	return singletonMongo
}
