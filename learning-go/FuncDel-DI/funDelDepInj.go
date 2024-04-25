package main

import (
	"fmt"
	"os"
)

// Logger is an interface with one method.
// Log takes a single parameter that is a string.
type Logger interface {
	Log(message string)
}

// ConsoleLogger encapsulates console logging functionality.
type ConsoleLogger struct{}

// FileLogger encapsulates writing logging information to a file.
type FileLogger struct {
	filename string
}

// Application uses a Logger to log its messages.
type Application struct {
	logger Logger
}

// Log has a receiver of ConsoleLogger, this means
// ConsoleLogger now implicitly implements the Logger
// interface.
func (l ConsoleLogger) Log(message string) {
	fmt.Println("console: ", message)
}

// Log has a receiver of FileLogger, this means
// FileLogger now implicitly implements the Logger
// interface.
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

// PerformTask has a receiver of a pointer to Application.
// Application contains the field, logger, which is a type of
// Logger. When a struct is injected into the constructor
// implementing the Logger interface, the Log method
// defined on that struct will execute.
func (app *Application) PerformTask() {
	// Each implementation of Logger has its own Log method.
	// Since we are coding to interfaces instead of implementation
	// this will function work any implementation of Logger
	// with a Log method.
	app.logger.Log("[+] task completed")
}

func main() {

	// Create an instance of ConsoleLogger.
	consoleLogger := ConsoleLogger{}
	// Inject the ConsoleLogger instance.
	app := NewApplication(consoleLogger)
	app.PerformTask()

	// Create an instance of FileLogger.
	fileLogger := FileLogger{filename: "app.log"}
	// Inject the FileLogger instance.
	app = NewApplication(fileLogger)
	app.PerformTask()

}
