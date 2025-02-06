package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)
type Users struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Username string `json:"username"`
	Email     string `json:"email"`
}

func ForGetReq() {
  // Get method
	fmt.Println("testing")
	res, err := http.Get("https://jsonplaceholder.typicode.com/users/10")
	if err != nil {
		fmt.Println("Error getting", err)
		return
	}
	defer res.Body.Close()
	// data, err := io.ReadAll(res.Body)
	// if err!= nil {
    //     fmt.Println("Error reading body", err)
	// 	return
    // }
	// fmt.Println(string(data))
	var user Users
    err = json.NewDecoder(res.Body).Decode(&user)
    if err!= nil {
        fmt.Println("Error unmarshalling", err)
        return
    }
    fmt.Println(user)
}

func ForPostReq() {
	user := Users{
		ID:        11,
        Name:      "John Doe",
        Username: "johndoe",
        Email:     "johndoe@example.com",
	}

	jsonData, err := json.Marshal(user)
	if err!= nil {
        fmt.Println("Error marshalling", err)
        return
    }

	jsonString := string(jsonData)
	jsonReader := strings.NewReader(jsonString)
    url := "https://jsonplaceholder.typicode.com/users"
	res, err := http.Post(url, "application/json", jsonReader)
	if err!= nil {
        fmt.Println("Error posting", err)
        return
    }
	defer res.Body.Close()

	data, _ := io.ReadAll(res.Body)
	fmt.Println(string(data))
}



// This is the main function
func main() {
	// ForGetReq()
	ForPostReq()
	
}