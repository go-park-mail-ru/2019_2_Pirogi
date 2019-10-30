package models

type Like struct {
	ID       ID     `json:"id"`
	UserID   ID     `json:"user_id"`
	Target   string `json:"target"` // Film or person
	TargetID ID     `json:"target_id"`
}
