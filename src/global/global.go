package global

import (
	"entsoe/src/syslogger"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/google/logger"
)

// Env is an environment struct that handles creation, initialization
// and managing the variables that can be shared between multiple function
// This includes all the object that is needed to created once
// during startup process and can be reused over the lifespan of
// the application
type Env struct {
	Logger      *logger.Logger
	FilePointer *os.File
}

// Init function Initialize global obj and Gin middleware needed for the application
func Init() *Env {
	// Init logger
	// ALWAYS REMEMBER TO CLOSE FILE POINTER MANUALLY
	newLogger, lf := syslogger.NewDayLogger(nil, nil)

	env := &Env{
		Logger:      newLogger,
		FilePointer: lf,
	}

	return env
}

// SetupCloseHandler close the DB connection after receiving the system interupt signal
func SetupCloseHandler(env *Env) {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("\r- Ctrl+C pressed in Terminal")
		env.Logger.Close()
		env.FilePointer.Close()
		os.Exit(0)
	}()
}
