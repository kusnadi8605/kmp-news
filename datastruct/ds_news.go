package datastruct

//News ..
type News struct {
	//* tanda bintang supaya tidak err jika datanya null
	ID      int64
	Author  string
	Body    string
	Created string
}

//NewsSave ..
type NewsSave struct {
	//* tanda bintang supaya tidak err jika datanya null
	ID int64
}

//NewsSaveRequest ..
type NewsSaveRequest struct {
	//* tanda bintang supaya tidak err jika datanya null
	Author string `json:"author,omitempty"`
	Body   string `json:"body,omitempty"`
}

//NewsSaveResponse ..
type NewsSaveResponse struct {
	ResponseCode string     `json:"responseCode"`
	ResponseDesc string     `json:"responseDesc"`
	Payload      []NewsSave `json:"payload"`
}

//NewsDeleteResponse ..
type NewsDeleteResponse struct {
	ResponseCode string `json:"responseCode"`
	ResponseDesc string `json:"responseDesc"`
}

//NewsRequest ..
type NewsRequest struct {
	ID string `json:"id,omitempty"`
}

//NewsResponse data
type NewsResponse struct {
	ResponseCode string       `json:"responseCode"`
	ResponseDesc string       `json:"responseDesc"`
	Payload      []NewsDetail `json:"payload"`
}

//NewsDetail ..
type NewsDetail struct {
	ID      int
	Author  string
	Body    string
	Created string
}

//NewsDetailRequest ..
type NewsDetailRequest struct {
	ID string `json:"id,omitempty"`
}

//NewsDetailResponse data
type NewsDetailResponse struct {
	ResponseCode string       `json:"responseCode"`
	ResponseDesc string       `json:"responseDesc"`
	Payload      []NewsDetail `json:"payload"`
}
