package intermediate

import (
	"log"
	"os"
)

func main() {
	log.Println("This is a log message")

	log.SetPrefix("INFO: ")
	log.Println("Tthis is an info message")

	//log flags
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("This is a log message with date, time and file name")

	infoLogger.Println("This is info message")
	warnLogger.Println("This is warn message")
	errorLogger.Println("This is error message")

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error creating log : %v", err)
	}
	defer file.Close()

	var infoLogger1 = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	var warnLogger1 = log.New(file, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	var errorLogger1 = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	var debugLogger = log.New(file, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLogger1.Println("This is INFO info")
	warnLogger1.Println("This is WARN info")
	errorLogger1.Println("This is ERROR info")
	debugLogger.Println("This is DEBUG info")
}

var infoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
var warnLogger = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
var errorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
