package intermediate

//go mod init gocourse => jalankan di terminal untuk buat folder yang simpen external package
import (
	"github.com/sirupsen/logrus"
)

func main() {

	log := logrus.New()
	//set log level

	log.SetLevel(logrus.InfoLevel)

	//set log format
	log.SetFormatter(&logrus.JSONFormatter{})

	//loging example
	log.Info("This is log info ")
	log.Warn("This is log warn")
	log.Error("This is log Error")
	log.WithFields(logrus.Fields{
		"username": "Fahrul",
		"method":   "GET",
	}).Info("User Logged in")

}
