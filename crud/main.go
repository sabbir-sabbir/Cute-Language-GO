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
      url := "https://jsonplaceholder.typicode.com/users"
	jsonString := string(jsonData)
	jsonReader := strings.NewReader(jsonString)
  
	res, err := http.Post(url, "application/json", jsonReader)
	if err!= nil {
        fmt.Println("Error posting", err)
        return
    }
	defer res.Body.Close()

	data, _ := io.ReadAll(res.Body)
	fmt.Println(string(data))
}

func ForUpdateRq() {
	user := Users{
		ID:        10,
        Name:      "John Doe Updated",
        Username: "johndoe Updated",
        Email:     "johndoe@example.com Updated",
	}
	jsonData, err := json.Marshal(user)
	if err!= nil {
        fmt.Println("Error marshalling", err)
        return
    }
	jsonString := string(jsonData)
    jsonReader   := strings.NewReader(jsonString)
   myUrl := "https://jsonplaceholder.typicode.com/users/10"

   req, err := http.NewRequest(http.MethodPut, myUrl, jsonReader)
   if err!= nil {
        fmt.Println("Error creating request", err)
        return
    }

	req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	res, err := client.Do(req)
	if err!= nil {
        fmt.Println("Error posting", err)
        return
    }
	defer res.Body.Close()

	data, _ := io.ReadAll(res.Body)
	fmt.Println(string(data))
}

 func ForDeleteReq() {
	myUrl := "https://jsonplaceholder.typicode.com/users/10"
	req, err := http.NewRequest(http.MethodDelete, myUrl, nil)
	if err!= nil {
        fmt.Println("Error creating request", err)
        return
    }

	client := http.Client{}
	res, err := client.Do(req)
	if err!= nil {
        fmt.Println("Error posting", err)
        return
    }
	defer res.Body.Close()

	fmt.Println(res.StatusCode)

 }


// This is the main function
func main() {
	// ForGetReq()
	// ForPostReq()
	// ForUpdateRq()
	ForDeleteReq()
}