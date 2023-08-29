package Controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Requeststruct struct {
	Ev     string `json:"ev"`
	Et     string `json:"et"`
	Id     string `json:"id"`
	Uid    string `json:"uid"`
	Mid    string `json:"mid"`
	P      string `json:"p"`
	T      string `json:"T"`
	L      string `json:"l"`
	Sc     string `json:"Sc"`
	Atrk1  string `json:"atrk1"`
	Atrv1  string `json:"Atrv1"`
	Atrt1  string `json:"atrt1"`
	Atrk2  string `json:"atrk2"`
	Atrv2  string `json:"atrv2"`
	Atrt2  string `json:"atrt2"`
	Uatrk1 string `json:"uatrk1"`
	Uatrv1 string `json:"uatrv1"`
	Uatrt1 string `json:"uatrt1"`
	Uatrk2 string `json:"uatrk2"`
	Uatrv2 string `json:"uatrv2"`
	Uatrt2 string `json:"uatrt2"`
	Uatrk3 string `json:"uatrk3"`
	Uatrv3 string `json:"uatrv3"`
	Uatrt3 string `json:"uatrt3"`
	Reply  chan map[string]interface{}
}

func RequestData(c *gin.Context) {
	var requeststruct Requeststruct
	json.NewDecoder(c.Request.Body).Decode(&requeststruct)

	// Channel to receive requests
	requests := make(chan Requeststruct)
// lets code as for testitng
	// Start the worker
	go worker(requests)

	// Create a request
	requeststruct = Requeststruct{
		Ev:     requeststruct.Ev,
		Et:     requeststruct.Et,
		Id:     requeststruct.Id,
		Uid:    requeststruct.Uid,
		Mid:    requeststruct.Mid,
		P:      requeststruct.P,
		T:      requeststruct.T,
		L:      requeststruct.L,
		Sc:     requeststruct.Sc,
		Atrk1:  requeststruct.Atrk1,
		Atrv1:  requeststruct.Atrv1,
		Atrt1:  requeststruct.Atrt1,
		Atrk2:  requeststruct.Atrk2,
		Atrv2:  requeststruct.Atrv2,
		Atrt2:  requeststruct.Atrt2,
		Uatrk1: requeststruct.Uatrk1,
		Uatrv1: requeststruct.Uatrv1,
		Uatrt1: requeststruct.Uatrt1,
		Uatrk2: requeststruct.Uatrt1,
		Uatrv2: requeststruct.Uatrv2,
		Uatrt2: requeststruct.Uatrt2,
		Uatrk3: requeststruct.Uatrk3,
		Uatrv3: requeststruct.Uatrv3,
		Uatrt3: requeststruct.Uatrt3,
		Reply:  make(chan map[string]interface{}),
	}

	// Send the request to the worker through the channel
	requests <- requeststruct

	// Wait for the worker to process the requeststructuest and receive the result
	result := <-requeststruct.Reply

	//SEND MESSAGE TO URL
	err := SendDataToURL(result)
	if err != nil {
		InternalServerErrorResponse(c, err)
		return
	}

	successResponse(c, "success message", result)
}

// Worker function to process the requests
func worker(requests <-chan Requeststruct) {
	for req := range requests {
		// Process the request
		result := processRequest(req)

		// Send the result back through the reply channel
		req.Reply <- result
	}
}

// Function to process the request data
func processRequest(req Requeststruct) map[string]interface{} {
	// Process the request data here and return the result
	return map[string]interface{}{
		"event":            req.Ev,
		"event_type":       req.Et,
		"app_id":           req.Id,
		"user_id":          req.Uid,
		"message_id":       req.Mid,
		"page_title":       req.T,
		"page_url":         req.P,
		"browser_language": req.L,
		"screen_size":      req.Sc,
		"attributes": map[string]interface{}{
			"form_varient": map[string]interface{}{
				"value": req.Atrv1,
				"type":  req.Atrt1,
			},
			"ref": map[string]interface{}{
				"value": req.Atrv2,
				"type":  req.Atrt2,
			},
		},
		"traits": map[string]interface{}{
			"name": map[string]interface{}{
				"value": req.Uatrv1,
				"type":  req.Uatrt1,
			},
			"email": map[string]interface{}{
				"value": req.Uatrv2,
				"type":  req.Uatrt2,
			},
			"age": map[string]interface{}{
				"value": req.Uatrv3,
				"type":  req.Uatrt3,
			},
		},
	}
}

func SendDataToURL(message map[string]interface{}) error {
	// URL of the webhook
	url := "https://webhook.site/"

	// Data to be sent in the request body
	payload, _ := json.Marshal(message)

	// Create a new HTTP POST request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return err
	}

	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)
	return nil
}
