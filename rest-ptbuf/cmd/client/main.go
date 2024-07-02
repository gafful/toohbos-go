package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "github.com/gafful/toohbos-go/rest/ptbuf/dummy"
    "github.com/golang/protobuf/proto"
    "io"
    "log"
    "net/http"
)

const ProtobufServerApi = "http://localhost:8080/"


func main() {
    http.HandleFunc("/", handleClientProtoRequest)
    http.HandleFunc("/complex", handleComplexClientProtoRequest)

    // Start the server on port 8080
    port := 8081
    fmt.Printf("Client running at http://localhost:%d\n", port)
    //    http.ListenAndServe(fmt.Sprintf(":%d", port), nil)....
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func handleClientProtoRequest(w http.ResponseWriter, r *http.Request) {
    contentLength := r.ContentLength
    fmt.Printf("Content Length Received : %v\n", contentLength)

    param1 := r.URL.Query().Get("json")
    
    // Receive payload
    request := dummy.Request{}
    err := json.NewDecoder(r.Body).Decode(&request)
    if err != nil {
    	panic(err)
    }
    
    // Serialize into a Protobuf request
    requestProto := &dummy.ProtoRequest{
        Query: request.Query,
    }
    body, err := proto.Marshal(requestProto)
    if err != nil {
        log.Fatalf("Unable to marshal response : %v", err)
    }
    
    // Prepare POST request with the Protobuf payload
    rNew, err := http.NewRequest("POST", ProtobufServerApi, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
    
    // Set Content-Type based on HTTP request
    if param1 != "" {
        rNew.Header.Add("Content-Type", "application/json")
    } else {
        rNew.Header.Add("Content-Type", "application/protobuf")
    }
    
    // Send POST request with Protobuf payload to server
    client := &http.Client{}
    res, err := client.Do(rNew)
    if err != nil {
    	panic(err)
    }
    defer res.Body.Close()
    
    // Read server response
    data, err := io.ReadAll(res.Body)
    if err != nil {
        log.Fatalf("Unable to read message from request : %v", err)
    }
    
    // Serialize response into Protobuf object
    response := &dummy.ProtoRequest{}
    err = proto.Unmarshal(data, response)
    if err != nil {
    	panic(err)
    }
    
    // Send JSON or Protobuf response based on HTTP request
    if param1 != "" {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(response)
    } else {
        response, _ := proto.Marshal(response)
        w.Header().Set("Content-Type", "application/protobuf")
        w.Write(response)
    }
}

func handleComplexClientProtoRequest(w http.ResponseWriter, r *http.Request) {
    contentLength := r.ContentLength
    fmt.Printf("Content Length Received : %v\n", contentLength)

    param1 := r.URL.Query().Get("json")
    
    // Receive payload
    request := dummy.Request{}
    err := json.NewDecoder(r.Body).Decode(&request)
    if err != nil {
    	panic(err)
    }
    
    // Serialize into a Protobuf request
    requestProto := &dummy.ProtoRequest{
        Query: request.Query,
    }
    body, err := proto.Marshal(requestProto)
    if err != nil {
        log.Fatalf("Unable to marshal response : %v", err)
    }
    
    // Prepare POST request with the Protobuf payload
    rNew, err := http.NewRequest("POST", fmt.Sprint(ProtobufServerApi, "complex"), bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
    
    // Set Content-Type based on HTTP request
    if param1 != "" {
        rNew.Header.Add("Content-Type", "application/json")
    } else {
        rNew.Header.Add("Content-Type", "application/protobuf")
    }
    
    // Send POST request with Protobuf payload to server
    client := &http.Client{}
    res, err := client.Do(rNew)
    if err != nil {
    	panic(err)
    }
    defer res.Body.Close()
    
    // Read server response
    data, err := io.ReadAll(res.Body)
    if err != nil {
        log.Fatalf("Unable to read message from request : %v", err)
    }
    
    // Serialize response into Protobuf object
    response := &dummy.ProtoResponsePage{}
    err = proto.Unmarshal(data, response)
    if err != nil {
    	panic(err)
    }
    
    // Send JSON or Protobuf response based on HTTP request
    if param1 != "" {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(response)
    } else {
        response, _ := proto.Marshal(response)
        w.Header().Set("Content-Type", "application/protobuf")
        w.Write(response)
    }
}