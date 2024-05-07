package main

import (
	"errors"
	"fmt"
	"os"
)

type Status int

const (
	InvalidLogin Status = iota + 1
	NotFound
)

type StatusErr struct {
	Status  Status
	Message string
}

func (se StatusErr) Error() string {
	return se.Message
}

func login(uid, pwd string) (string, error) {
	if uid == "admin" && pwd == "admin" {
		return "user:admin", nil
	}

	return "", errors.New("[-] bad user")
}

func getData(token, file string) ([]byte, error) {
	if token == "user:admin" {
		switch file {
		case "secrets.txt":
			return []byte("[+] password aplenty"), nil
		case "payroll.csv":
			return []byte("[+] everyones salary"), nil
		}
	}

	return nil, os.ErrNotExist
}

func LoginAndGetData(uid, pwd, file string) ([]byte, error) {
	token, err := login(uid, pwd)
	if err != nil {
		return nil, StatusErr{
			Status:  InvalidLogin,
			Message: fmt.Sprintf("[-] invalid credentials for user -- StatusID: %d", InvalidLogin),
		}
	}

	data, err := getData(token, file)
	if err != nil {
		return nil, StatusErr{
			Status:  NotFound,
			Message: fmt.Sprintf("[-] file \"%s\" not found -- StatusID: %d", file, NotFound),
		}
	}

	return data, nil
}

func main() {

	data, err := LoginAndGetData("admin", "admin", "secrets.txt")
	fmt.Println(string(data), err)

	dat, er := LoginAndGetData("admin", "admin", "chicken_soup.txt")
	fmt.Printf("%s %s\n", string(dat), er)

	d, e := LoginAndGetData("john", "password", "secrets.txt")
	fmt.Printf("%s %s\n", string(d), e)
}
