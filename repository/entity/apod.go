package entity

type APODContent struct {
	ID             int    `db:"id"`
	Copyright      string `db:"copyright"`
	Explanation    string `db:"explanation"`
	HdURL          string `db:"hdurl"`
	MediaType      string `db:"media_type"`
	ServiceVersion string `db:"service_version"`
	Title          string `db:"title"`
	URL            string `db:"url"`
	Date           string `db:"date"`
}
