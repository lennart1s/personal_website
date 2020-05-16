package properties

import (
	"github.com/magiconair/properties"
)

var (
	path = "C:/Users/Lennart/Desktop/sandbothe_dev.properties"

	Port = "8080"

	UseSSL   = true
	CertPath = ""
	KeyPath  = ""

	UseDatabase = true
	SQLLogin    = ""
	GrecSecret  = ""

	GithubListUser = "lennart1s"
	NumGithubCards = 4
)

func init() {
	p := properties.MustLoadFile(path, properties.UTF8)

	Port = p.GetString("port", Port)
	UseSSL = p.GetBool("useSSL", UseSSL)
	CertPath = p.GetString("certPath", CertPath)
	KeyPath = p.GetString("keyPath", KeyPath)
	UseDatabase = p.GetBool("useDatabase", UseDatabase)
	SQLLogin = p.GetString("sqlLogin", SQLLogin)
	GrecSecret = p.GetString("grecSecret", GrecSecret)
	GithubListUser = p.GetString("githubListUser", GithubListUser)
	NumGithubCards = p.GetInt("numGithubCard", NumGithubCards)
}
