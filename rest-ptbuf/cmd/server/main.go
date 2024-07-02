package main

import (
    "encoding/json"
    "fmt"
    "github.com/gafful/toohbos-go/rest/ptbuf/dummy"
    "github.com/go-faker/faker/v4"
    "github.com/golang/protobuf/proto"
    "io"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", handleProtoRequest)
    http.HandleFunc("/page", handleProtoPageRequest)
    http.HandleFunc("/complex", handleComplexProtoPageRequest)
    http.HandleFunc("/test", handleTestProtoPageRequest)

    // Start the server on port 8080
    port := 8080
    fmt.Printf("Server running at http://localhost:%d\n", port)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleTestProtoPageRequest(w http.ResponseWriter, r *http.Request) {
    d := dummy.Dummy{}
    result := d.GenerateTestResponse()
    json.NewEncoder(w).Encode(result)
}

func handleProtoPageRequest(w http.ResponseWriter, r *http.Request) {
    contentLength := r.ContentLength
    fmt.Printf("Content Length Received : %v\n", contentLength)

    getPagedHandler(w, r)
}

func getPagedHandler(w http.ResponseWriter, r *http.Request) {
    // Possible to send path and query params as protos too? Necessary?
    param1 := r.URL.Query().Get("json")

    // Generate a paged response
    d := dummy.Dummy{}
    result := d.GeneratePageResponse(param1)
    
    // Serialize response into Protobuf
    protoResult := result.MapToPageProtoResponse()
    response, err := proto.Marshal(&protoResult)
    if err != nil {
        log.Fatalf("Unable to marshal response : %v", err)
    }
    
    // Send JSON or Protobuf response based on HTTP request
    if param1 != "" {
        json.NewEncoder(w).Encode(result)
    } else {
        w.Header().Set("Content-Type", "application/protobuf")
        w.Write(response)
    }
}

func handleProtoRequest(w http.ResponseWriter, r *http.Request) {
    contentLength := r.ContentLength
    fmt.Printf("Content Length Received : %v\n", contentLength)

    switch r.Method {
    case "GET":
        getHandler(w, r)
    case "POST":
        postHandler(w, r)
    }
}

func postHandler(w http.ResponseWriter, r *http.Request) {
    param1 := r.URL.Query().Get("json")
    
    // Get HTTP request payload
    data, err := io.ReadAll(r.Body)
    if err != nil {
        log.Fatalf("Unable to read message from request : %v", err)
    }

    // Deserialize payload into Protobuf object
    request := &dummy.ProtoRequest{}
    err = proto.Unmarshal(data, request)
    if err != nil {
        log.Fatalf("Unable to deserialize message from request : %v", err)
    }
    
    // Update payload with data from server
    request.DateTag = faker.Date()
    
    // Serialize HTTP request into a Protobof object
    response, err := proto.Marshal(request)
    if err != nil {
        log.Fatalf("Unable to marshal response : %v", err)
    }
    
    // Send JSON or Protobuf response based on HTTP request
    if param1 != "" {
        json.NewEncoder(w).Encode(request)
    } else {
        w.Header().Set("Content-Type", "application/protobuf")
        w.Write(response)
    }
}

func getHandler(w http.ResponseWriter, r *http.Request) {
    param1 := r.URL.Query().Get("json")

    // Generate Dummy response
    d := dummy.Dummy{}
    result := d.GenerateResponse(param1)
    protoResult := result.MapToProtoResponse()
    response, err := proto.Marshal(protoResult)
    if err != nil {
        log.Fatalf("Unable to marshal response : %v", err)
    }
    
    // Send JSON or Protobuf response based on HTTP request
    if param1 != "" {
        json.NewEncoder(w).Encode(result)
    } else {
        w.Header().Set("Content-Type", "application/protobuf")
        w.Write(response)
    }
}

func handleComplexProtoPageRequest(w http.ResponseWriter, r *http.Request) {
    contentLength := r.ContentLength
    fmt.Printf("Content Length Received : %v\n", contentLength)

    // Get HTTP request payload
    data, err := io.ReadAll(r.Body)
    if err != nil {
        log.Fatalf("Unable to read message from request : %v", err)
    }
    fmt.Println(string(data))
    
    // Possible to send path and query params as protos too? Necessary?
    param1 := r.URL.Query().Get("json")

    // Generate a paged response
    d := dummy.Dummy{}
    result := d.GeneratePageResponse(param1)
    
    // Serialize response into Protobuf
    protoResult := result.MapToPageProtoResponse()
    response, err := proto.Marshal(&protoResult)
    if err != nil {
        log.Fatalf("Unable to marshal response : %v", err)
    }
    
    // Send JSON or Protobuf response based on HTTP request
    if param1 != "" {
        json.NewEncoder(w).Encode(result)
    } else {
        w.Header().Set("Content-Type", "application/protobuf")
        w.Write(response)
    }
}