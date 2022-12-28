package book

type book struct {
	ID     string  `bson:"_id,omitempty" json:"id"`
	Title  string  `bson:"title" json:"title,omitempty"`
	Author string  `bson:"author" json:"author,omitempty"`
	Rating float64 `bson:"rating" json:"rating,omitempty"`
}
