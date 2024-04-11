package main

import (
	"errors"
	"fmt"
	"net/http"
)

/** --- Define Interfaces --- **/

// DataStore describes operations to access user data.
type DataStore interface {
	UserNameForID(userID string) (string, bool)
}

// Logger describes operations for logging messages.
type Logger interface {
	Log(message string)
}

/** --- Implement Interfaces --- **/

/* [!] Swapping Loggers [!] */

// LoggerAdapter adapts a function to Logger interface.
//type LoggerAdapter func(message string)

// Log is defined as type LoggerAdapter, implicitly implementing Logger on LoggerAdapter.
//func (lg LoggerAdapter) Log(message string) {
//	lg(message)
//}

type FileLogger struct{}

// Log defined as method of FileLogger to implement Logger interface on FileLogger
func (fl FileLogger) Log(message string) {
	fmt.Println("[+] File log: ", message)
}

/* [-] End Swapping Loggers [-] */

/* [!] Swapping Datastore [!] */

type APIDataStore struct{}

func (ads APIDataStore) UserNameForID(userID string) (string, bool) {
	return "Fetched Name", true
}

// SimpleDataStore is a concrete implementation of DataStore.
//type SimpleDataStore struct {
//	userData map[string]string
//}
//
//// UserNameForID is implemented as a method on SimpleDataStore, implicitly implementing DataStore on SimpleDataStore.
//func (sds SimpleDataStore) UserNameForID(userID string) (string, bool) {
//	name, ok := sds.userData[userID]
//	return name, ok
//}
//
//func NewSimpleDataStore() SimpleDataStore {
//	return SimpleDataStore{
//		userData: map[string]string{
//			"1": "Lorem",
//			"2": "Ipsum",
//			"3": "Dolor",
//		},
//	}
//}

/** --- Business Logic --- **/

// SimpleLogic contains business logic that depends on DatStore and Logger.
type SimpleLogic struct {
	l  Logger
	ds DataStore
}

func (sl SimpleLogic) SayHello(userID string) (string, error) {
	sl.l.Log("[+] In `SayHello` for: " + userID)
	name, ok := sl.ds.UserNameForID(userID)

	if !ok {
		return "[-] ", errors.New("unknown user")
	}

	return "Hello, " + name, nil
}

/** --- Utility function --- **/
func consoleLogger(message string) {
	fmt.Println(message)
}

func main() {

	// Adapt consoleLogger to the Logger interface through the LoggerAdapter.
	// logger := LoggerAdapter(consoleLogger)
	logger := FileLogger{}

	// Define a new datastore
	//datastore := NewSimpleDataStore()
	datastore := APIDataStore{}

	logic := SimpleLogic{
		l:  logger,
		ds: datastore,
	}

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		userID := r.URL.Query().Get("user_id")
		message, err := logic.SayHello(userID)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Fprintln(w, message)
	})

	http.ListenAndServe(":8080", nil)

}
