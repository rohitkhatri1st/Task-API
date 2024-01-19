package model

type CustomResponse struct {
	Status     string
	StatusCode int
	Data       interface{}
}

type SentRequest struct {
	Method        string      `json:"method,omitempty" bson:"method,omitempty"`
	URL           interface{} `json:"url,omitempty" bson:"url,omitempty"`
	Proto         string      `json:"proto,omitempty" bson:"proto,omitempty"`
	ProtoMajor    int         `json:"proto_major,omitempty" bson:"proto_major,omitempty"`
	ProtoMinor    int         `json:"proto_minor,omitempty" bson:"proto_minor,omitempty"`
	Header        interface{} `json:"header,omitempty" bson:"header,omitempty"`
	Body          interface{} `json:"body,omitempty" bson:"body,omitempty"`
	ContentLength int64       `json:"content_length,omitempty" bson:"content_length,omitempty"`
	Form          interface{} `json:"form,omitempty" bson:"form,omitempty"`
	PostForm      interface{} `json:"post_form,omitempty" bson:"post_form,omitempty"`
}

type ReceivedResponse struct {
	Status        string      `json:"status,omitempty" bson:"status,omitempty"`
	StatusCode    int         `json:"status_code,omitempty" bson:"status_code,omitempty"`
	Proto         string      `json:"proto,omitempty" bson:"proto,omitempty"`
	ProtoMajor    int         `json:"proto_major,omitempty" bson:"proto_major,omitempty"`
	ProtoMinor    int         `json:"proto_minor,omitempty" bson:"proto_minor,omitempty"`
	Header        interface{} `json:"header,omitempty" bson:"header,omitempty"`
	Body          interface{} `json:"body,omitempty" bson:"body,omitempty"`
	ContentLength int64       `json:"content_length,omitempty" bson:"content_length,omitempty"`
}
