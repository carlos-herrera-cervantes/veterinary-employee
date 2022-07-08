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

func InitializeMongoDB() mongo {
	return mongo{
		Host:      os.Getenv("MONGODB_HOST"),
		DefaultDB: os.Getenv("DEFAULT_DB"),
		Collections: collections{
			Profile: "profiles",
			Role:    "roles",
			Address: "addresses",
			Avatar:  "avatars",
		},
	}
}
