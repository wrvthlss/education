package main

import (
	"errors"
	"fmt"
	"net/http"
)

/*
	`LoggerAdapter` and `SimpleDataStore` meet the interfaces needed
	by the business logic, but neither type has any idea that it does.
*/

// DataStore is an interfaces describing what the business logic depends on.
type DataStore interface {
	UserNameForID(userID string) (string, bool)
}

type Logger interface {
	Log(message string)
}

type Logic interface {
	SayHello(userID string) (string, error)
}

type LoggerAdapter func(message string)

// Log LogOutput Utility function: logger.
func (lg LoggerAdapter) Log(message string) {
	lg(message)
}

type SimpleDataStore struct {
	userData map[string]string
}

// UserNameForID implicitly implements DataStore on SimpleDataStore
func (sds SimpleDataStore) UserNameForID(userID string) (string, bool) {
	name, ok := sds.userData[userID]
	return name, ok
}

// NewSimpleDatastore is a factory function to create an instance of SimpleDataStore.
func NewSimpleDatastore() SimpleDataStore {
	return SimpleDataStore{
		map[string]string{
			"1": "Lorem",
			"2": "Ipsum",
			"3": "Nurit",
		},
	}
}

/** --- Business Logic --- **/

type Controller struct {
	l     Logger
	logic Logic
}

func NewController(l Logger, logic Logic) Controller {
	return Controller{
		l:     l,
		logic: logic,
	}
}

func (c Controller) SayHello(w http.ResponseWriter, r *http.Request) {
	c.l.Log("[+] In `SayHello`")
	userID := r.URL.Query().Get("user_id")

	message, err := c.logic.SayHello(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(message))
}

// SimpleLogic members are both interfaces.
type SimpleLogic struct {
	l  Logger
	ds DataStore
}

func (sl SimpleLogic) SayHello(userID string) (string, error) {
	sl.l.Log("[!] In `SayHello` for: " + userID)

	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "[-] ", errors.New("Unknown User.")
	}

	return "Hello, " + name, nil
}

func (sl SimpleLogic) SayGoodbye(userID string) (string, error) {
	sl.l.Log("[!] Int `SayGoodbye` for " + userID)

	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "[-] ", errors.New("Unkown User.")
	}

	return "Goodbye, " + name, nil
}

func consoleLogger(message string) {
	fmt.Println(message)
}

func NewSimpleLogic(l Logger, ds DataStore) SimpleLogic {
	return SimpleLogic{
		l:  l,
		ds: ds,
	}
}

func main() {

	l := LoggerAdapter(consoleLogger)
	ds := NewSimpleDatastore()
	logic := NewSimpleLogic(l, ds)
	c := NewController(l, logic)
	http.HandleFunc("/hello", c.SayHello)
	http.ListenAndServe(":8080", nil)

	// After adding business logic.
	// Pass consoleLogger to the logger adapter to meet interface.
	//bsLogger := LoggerAdapter(consoleLogger)
	//bsDataStore := NewSimpleDatastore()
	//
	//bsLogic := NewSimpleLogic(bsLogger, bsDataStore)
	//hello, _ := bsLogic.SayHello("1")
	//fmt.Println(hello)

	// Using factory function.
	//nsds := NewSimpleDatastore()
	//fmt.Println(nsds.UserNameForID("2"))

	// Basic Usage
	//sds := SimpleDataStore{
	//	userData: map[string]string{
	//		"1": "Lorem",
	//	},
	//}
	//
	//fmt.Println(sds.UserNameForID("1"))
}
