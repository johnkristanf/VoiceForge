package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"server/config"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var voiceCollection *mongo.Collection = config.VoicesCollection()
var wg sync.WaitGroup



type VoiceStruct struct {
    ID           string `json:"id" bson:"id"`
    Name         string `json:"name" bson:"name"`
    Sample       string `json:"sample" bson:"sample"`
    Accent       string `json:"accent" bson:"accent"`
    Age          string `json:"age" bson:"age"`
    Gender       string `json:"gender" bson:"gender"`
    Language     string `json:"language" bson:"language"`
    LanguageCode string `json:"language_code" bson:"language_code"`
    Loudness     string `json:"loudness" bson:"loudness"`
    Style        string `json:"style" bson:"style"`
    Tempo        string `json:"tempo" bson:"tempo"`
    Texture      string `json:"texture" bson:"texture"`
    IsCloned     bool   `json:"is_cloned" bson:"is_cloned"`
    VoiceEngine  string `json:"voice_engine" bson:"voice_engine"`
}

type StreamBody struct{
	Text string `json:"text"`
	Voice string `json:"voice"`
	Output_format string `json:"output_format"`
}


// func FetchVoices(ctx *gin.Context){

// 	client := &http.Client{}

// 	req, err := http.NewRequest("GET", "https://api.play.ht/api/v2/voices", nil)
// 	if err != nil {
// 		log.Fatalln("ERROR MAKING REQUEST", err.Error())
// 	}

// 	req.Header.Add("accept", "application/json")
// 	req.Header.Add("AUTHORIZATION", "dc23bdb0088e43d0ae92155f682d658b")
// 	req.Header.Add("X-USER-ID", "zXUVGgbbxFM42MjWQG3foHTHnLT2")

// 	resp, reqErr := client.Do(req)
// 	if reqErr != nil {
// 		log.Fatalln("ERROR MAKING REQUEST", reqErr.Error())
// 	}

// 	defer resp.Body.Close()


// 	resBody, readErr := io.ReadAll(resp.Body)
// 	if readErr != nil {
// 		log.Fatalln("ERROR READING RESPONSE BODY", readErr.Error())
// 	}


// 	var voices []VoiceStruct

// 	jsonErr := json.Unmarshal([]byte(resBody), &voices)
// 	if jsonErr != nil {
// 		log.Fatalln("ERROR READING RESPONSE BODY", jsonErr.Error())
// 	}


// 	documents := make([]interface{}, len(voices))
// 	for i, v := range voices {
// 		documents[i] = v
// 	}


// 	insert, insertErr := voiceCollection.InsertMany(context.Background(), documents)
// 	if insertErr != nil {
// 		log.Fatal("ERROR INSERTING VOICES IN DATABASE", insertErr.Error())
// 	}

// 	fmt.Println("INSERTED", insert.InsertedIDs)


// }



func FetchVoices(ctx *gin.Context){

	startTime := time.Now()

	var voices []VoiceStruct

	filter := bson.M{"id": bson.M{"$regex": "^s3"}}


	cursor, cursorErr := voiceCollection.Find(context.Background(), filter)
	if cursorErr != nil {
		log.Fatalln("ERROR IN CURSOR", cursorErr.Error())
	}

	if err := cursor.All(context.Background(), &voices); err != nil {
		log.Fatalln("ERROR IN CURSOR ALL", cursorErr.Error())
	}


	executionTime := time.Since(startTime)
	ctx.JSON(http.StatusOK, gin.H{"voices": voices, "executionTime": executionTime.String()})

}


func StreamAudio(ctx *gin.Context){

	startTime := time.Now()

	var Body StreamBody

	bodyChan := make(chan *bytes.Buffer, 1)
	respChan := make(chan *http.Response, 1)

	errorChan := make(chan error, 2)

	bindErr := ctx.ShouldBindJSON(&Body)
	if bindErr != nil {
		log.Fatalln("ERROR BINDING JSON DATA FROM HTTP BODY", bindErr.Error())
	}


	wg.Add(2)

	go func ()  {
		defer wg.Done()
		parseBodyToJson(Body, bodyChan, errorChan)
	}()


	go func ()  {
		defer wg.Done()
		sendStreamRequest(bodyChan, respChan, errorChan)
	}()


	go func ()  {
		wg.Wait()
		close(errorChan)	
	}()


	var resp *http.Response
	var err error

	select {

	  case resp = <-respChan:
		
	  case err = <-errorChan:
		log.Println("Error from goroutines:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	
	if resp != nil {
		responseBody, readErr := io.ReadAll(resp.Body)
		if readErr != nil {
			log.Fatalln("Error reading request:", readErr)
		}

		executionTime := time.Since(startTime)

		ctx.JSON(http.StatusOK, gin.H{
			"status":        resp.Status,
			"response":      string(responseBody),
			"executionTime": executionTime.String(),
		})

	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}
}


func parseBodyToJson(Body StreamBody, bodyChan chan *bytes.Buffer, errorChan chan error) {

    jsonBytes, jsonErr := json.Marshal(Body)
    if jsonErr != nil {
        errorChan <- jsonErr
        return
    }

    reqBody := bytes.NewBuffer(jsonBytes)

    bodyChan <- reqBody
}

func sendStreamRequest(bodyChan chan *bytes.Buffer, respChan chan *http.Response, errorChan chan error) {
    defer close(respChan)

    reqBody, ok := <-bodyChan 
    if !ok {
        errorChan <- fmt.Errorf("bodyChan closed unexpectedly")
        return
    }

    req, reqErr := http.NewRequest("POST", "https://api.play.ht/api/v2/tts/stream", reqBody)
    if reqErr != nil {
        errorChan <- reqErr
        return
    }

    Header := map[string]string{
        "Accept-Encoding": "audio/mpeg",
        "Content-type":    "application/json",
        "AUTHORIZATION":   "dc23bdb0088e43d0ae92155f682d658b",
        "X-USER-ID":       "zXUVGgbbxFM42MjWQG3foHTHnLT2",
    }

    for key, value := range Header {
        req.Header.Set(key, value)
    }

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        errorChan <- err
        return
    }

    respChan <- resp
}
