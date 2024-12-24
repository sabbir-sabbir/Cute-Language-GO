package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performGetRequest() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error fetching the URL:", err)
		return
	}
	defer res.Body.Close()

	// Optional: Check the response status code
	if res.StatusCode != http.StatusOK {
		fmt.Println("Error fetching the URL, response code:", res.StatusCode)
		return
	}

	var todo Todo // Create a variable of type Todo
	err = json.NewDecoder(res.Body).Decode(&todo) // Decode the response body into the variable
	if err != nil {
		fmt.Println("Error decoding the response:", err)
		return
	}
	fmt.Println("The todo 1 is:", todo)


}

func performPostRequest() {
	 todo := Todo {
		UserId:    1,
		Id:        1,
		Title:     "delectus aut autem",
		Completed: false,
	}
	// convert the todo struct to JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling the todo:", err)
	}
	
	// convert jsonData to string
	jsonString := string(jsonData)

	// convert data string to reader
	jsonReader := strings.NewReader(jsonString)
    TheUrl := "https://jsonplaceholder.typicode.com/todos"

	// send the POST request
	res, err := http.Post(TheUrl, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error sending the POST request:", err)
		return
	}
	defer res.Body.Close()
	justData, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(justData))
}

func performPutRequest() {
	todo := Todo {
		UserId:    567890,
		Id:        1,
		Title:     "perferendis",
		Completed: true,
	}
	// convert the todo struct to JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling the todo:", err)
		return
	}

	// convert jsonData to string
	jsonString := string(jsonData)

	// convert data string to reader
	jsonReader := strings.NewReader(jsonString)


    // Url to send the PUT request
	TheUrl := "https://jsonplaceholder.typicode.com/todos/1"

	// create a put request
	req, err := http.NewRequest(http.MethodPut, TheUrl, jsonReader)
	if err != nil {
		fmt.Println("Error creating the PUT request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	// sent the request
	client := http.Client{}
	res, err := client.Do(req)
	 if err != nil {
		fmt.Println("Error sending the PUT request:", err)
		return
	 }
     
	 defer res.Body.Close()

	//  lets read the response
	data , _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(data))
}

func performDeleteRequest() {
	// url for delete request
	TheUrl := "https://jsonplaceholder.typicode.com/todos/1"

	// create a Delete request
	req, err := http.NewRequest(http.MethodDelete, TheUrl, nil)
	if err != nil {
		fmt.Println("Error creating the delete request:", err)
		return
	}

	
	// sent the request
	client := http.Client{}
	res, err := client.Do(req)
	 if err != nil {
		fmt.Println("Error sending the delete request:", err)
		return
	 }
     
	 defer res.Body.Close()

	//  lets read the response
	fmt.Println("The response status code is:", res.StatusCode)



}

func main() {
	
	// performGetRequest()
	// performPostRequest()
	//  performPutRequest()
	performDeleteRequest() 
	

}
