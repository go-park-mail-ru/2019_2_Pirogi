package models

type Like struct {
	UserID   ID     `json:"user_id"`
	Target   Target `json:"target"` // Film or person
	TargetID ID     `json:"target_id"`
}
