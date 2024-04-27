package courses

type Course struct {
	Title    string `json:"title"`
	SubTitle string `json:"subTitle" db:"sub_title"`
	ImageURL string `json:"imageURL" db:"image_url"`
}
