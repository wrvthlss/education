package main

import (
	"fmt"
	"os"
)

type Logger interface {
	Log(message string)
}

type ConsoleLogger struct{}

type FileLogger struct {
	filename string
}

type Application struct {
	logger Logger
}

// Log is defined on ConsoleLogger meaning this struct
// implicitly implements the Logger interface.
func (l ConsoleLogger) Log(message string) {
	fmt.Println("console: ", message)
}

// Log is defined on FileLogger meaning this struct
// implicitly implements the Logger interface.
func (l FileLogger) Log(message string) {
	f, err := os.OpenFile(l.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println("[-] error closing file: ", err)
		}
	}(f)

	if _, err := f.WriteString(message + "\n"); err != nil {
		fmt.Println("[-] error writing to file: ", err)
	}
}

// NewApplication is a constructor function for
// instantiating Application. It returns a pointer
// to Application.
func NewApplication(logger Logger) *Application {
	return &Application{logger: logger}
}

// PerformTask executes a task and logs its completion. It has a
// receiver of type *Application. The Application struct contains a
// field named 'logger', which is of the interface type Logger.
// When an instance of a struct implementing the Logger interface
// is injected into the constructor, the Log method is called
// polymorphically based on the concrete type of Logger provided.
// This allows different logging strategies to be employed dynamically,
// depending on the Logger implementation passed during Application instantiation.
func (app *Application) PerformTask() {
	app.logger.Log("[+] task completed")
}

func main() {

	// Create an instance of ConsoleLogger.
	consoleLogger := ConsoleLogger{}
	app := NewApplication(consoleLogger)
	app.PerformTask()

	// Create an instance of FileLogger.
	fileLogger := FileLogger{}
	app = NewApplication(fileLogger)
	app.PerformTask()

}
