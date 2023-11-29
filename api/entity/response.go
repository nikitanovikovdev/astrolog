package entity

type APODContentResponse struct {
	ID             int    `json:"id"`
	Copyright      string `json:"copyright"`
	Explanation    string `json:"explanation"`
	HdURL          string `json:"hdurl"`
	MediaType      string `json:"media_type"`
	ServiceVersion string `json:"service_version"`
	Title          string `json:"title"`
	URL            string `json:"url"`
	Date           string `json:"date"`
}

type APODContentsResponse []APODContentResponse
