package model

type List struct {
	ID      ID     `json:"id" bson:"_id" valid:"numeric"`
	Title   string `json:"title" valid:"text, stringlength(1|50)"`
	UserID  ID     `json:"user_id" valid:"numeric"`
	FilmsID []ID   `json:"films_id" valid:"ids, numeric"`
}
