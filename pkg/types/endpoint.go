package types

import "time"

type Endpoint struct {
	URL         string            `json:"url"`
	Methods     []string          `json:"methods,omitempty"`
	Status      map[string]int    `json:"status,omitempty"`
	Title       string            `json:"title,omitempty"`
	Tech        []string          `json:"tech,omitempty"`
	ContentType string            `json:"content_type,omitempty"`
	Length      int               `json:"length,omitempty"`
	Headers     map[string]string `json:"headers,omitempty"`
	Interesting []string          `json:"interesting,omitempty"`
	Source      []string          `json:"source,omitempty"`
	FirstSeen   int64             `json:"first_seen"`
	LastSeen    int64             `json:"last_seen"`
}

func NewEndpoint(url string) *Endpoint {
	return &Endpoint{
		URL:       url,
		Methods:   []string{},
		Status:    make(map[string]int),
		Headers:   make(map[string]string),
		Source:    []string{},
		FirstSeen: time.Now().Unix(),
		LastSeen:  time.Now().Unix(),
	}
}
