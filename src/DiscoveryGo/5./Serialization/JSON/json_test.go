package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"
)

type status int
type Deadline time.Time

const (
	UNKNOWN status = iota
	TODO
	DONE
)

type Task struct {
	Title    string   `json:"title"`
	Status   status   `json:"status"`
	Deadline Deadline `json:"deadline"`
}

func (s status) MarsahlJSON() ([]byte, error) {
	switch s {
	case UNKNOWN:
		return []byte(`"UNKNOWN"`), nil
	case TODO:
		return []byte(`"TODO"`), nil
	case DONE:
		return []byte(`"DONE"`), nil
	default:
		return nil, errors.New("status.MarshalJSON: unknown value")
	}
}

func (s *status) Unmarshal(data []byte) error {
	switch string(data) {
	case `"UNKNOWN"`:
		*s = UNKNOWN
	case `"TODO"`:
		*s = TODO
	case `"DONE"`:
		*s = DONE
	default:
		return errors.New("status.UnmarshalJSON: unknow value")
	}
	return nil
}

func (d Deadline) MarshalJSON() ([]byte, error) {
	return strconv.AppendInt(nil, d.Unix(), 10), nil
}

func (d *Deadline) UnmarshalJSON(data []byte) error {
	unix, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	d.Time = time.Unix(unix, 0)
	return nil
}

//직렬화
func Example_marshalJSON() {
	t := Task{
		"Laundry",
		DONE,
		Deadline(time.Date(2015, time.August, 16, 15, 43, 0, 0, time.UTC)),
	}
	b, err := json.Marshal(t)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(b))
	// Output:
	// .
}

//역직렬화
func Example_unmarshalJSON() {
	b := []byte(`{"Title":"Buy Milk","Status":2}`)
	t := Task{}

	err := json.Unmarshal(b, &t)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(t)
	// Output:
	// .
}
