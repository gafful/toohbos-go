package dummy

import (
    "fmt"
    "google.golang.org/protobuf/types/known/timestamppb"
    "math/rand"
    "time"

    "github.com/go-faker/faker/v4"
)

type Dummy struct{}

type Hard struct {
    Email string `faker:"email"`

    Phone string `faker:"phone_number"`
}

type Request struct {
    Query string
}

type Response struct {
    Id          string `faker:"uuid_hyphenated"`
    Title       string `faker:"word"`
    Description string `faker:"paragraph"`
    ImageUrl    string `faker:"url"`
    Date        string `faker:"date"`
    Request     string `faker:"word"` // todo: remove
    Type        string `faker:"oneof: cc, paypal"`
}

type ResponsePage struct {
    Count int        `faker:"oneof: 20"`
    Total int        `faker:"oneof: 500"`
    Items []Response `faker:"response slice_len=20"`
}

func (d *Dummy) GenerateTestResponse() Hard {
    r := Hard{}
    _ = faker.FakeData(&r)
    fmt.Printf("%+v", r)
    return r
}

func (d *Dummy) GenerateResponse(rd string) Response {
    r := Response{}
    err := faker.FakeData(&r)
    if err != nil {
        fmt.Println(err)
    }
    r.ImageUrl = "https://picsum.photos/200/300"
    var date = getProtobufTimestamp()
    r.Date = date.AsTime().Format(time.RFC3339)
    return r
}

func (d *Dummy) GeneratePageResponse(rd string) ResponsePage {
    array := make([]Response, 10)
    for i, _ := range array {
        array[i] = d.GenerateResponse(rd)
    }
    r := ResponsePage{
        Count: 10,
        Total: 20,
        Items: array,
    }

    return r
}

func (r *Response) MapToProtoResponse() *ProtoResponse {
    pb := getProtobufTimestamp()

    return &ProtoResponse{
        Id:          r.Id,
        Title:       r.Title,
        Description: r.Description,
        ImageUrl:    "https://picsum.photos/200/300",
        Date:        pb,
        Request:     r.Request,
        //        Type: r.Type, //TODO: Fix
    }
}

func (r *ResponsePage) MapToPageProtoResponse() ProtoResponsePage {
    //     var person = (*ProtoResponsePage)(r) // TODO: Try
    array := make([]*ProtoResponse, len(r.Items))
    for i, element := range r.Items {
        array[i] = element.MapToProtoResponse()
    }

    return ProtoResponsePage{
        Count: int32(r.Count),
        Total: int32(r.Total),
        Items: &ProtoResponseList{
            Items: array,
        },
    }
}

func getProtobufTimestamp() *timestamppb.Timestamp {
    start := "2023-01-01T00:00:00Z"
    end := "2023-12-31T23:59:59Z"

    randomDate, err := getRandomDate(start, end)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        panic(err)
    }

    return timestamppb.New(randomDate)
}

// getRandomDate generates a random date between two given periods in the specified format
func getRandomDate(start, end string) (time.Time, error) {
    const layout = "2006-01-02T15:04:05Z"
    startTime, err := time.Parse(layout, start)
    if err != nil {
        return time.Time{}, fmt.Errorf("failed to parse start time: %v", err)
    }

    endTime, err := time.Parse(layout, end)
    if err != nil {
        return time.Time{}, fmt.Errorf("failed to parse end time: %v", err)
    }

    if startTime.After(endTime) {
        return time.Time{}, fmt.Errorf("start time must be before end time")
    }

    // Generate a random number of seconds between the two times
    delta := endTime.Unix() - startTime.Unix()
    seconds := rand.Int63n(delta)

    // Add the random number of seconds to the start time
    randomTime := startTime.Add(time.Duration(seconds) * time.Second)

    //	return randomTime.Format(layout), nil
    return randomTime, nil
}
