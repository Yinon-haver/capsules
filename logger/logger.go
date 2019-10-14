package logger

import (
	"log"
)

func Init() {
	log.SetFlags(log.LstdFlags)
}

func Info(v ...interface{}) {
	log.Println("INFO:", v)
}

func Warning(v ...interface{}) {
	log.Println("WARNING:", v)
}

func Error(v ...interface{}) {
	log.Println("ERROR:", v)
}

func Fatal(v ...interface{}) {
	log.Fatalln("FATAL:", v)
}