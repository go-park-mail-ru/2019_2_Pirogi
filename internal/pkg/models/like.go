package models

type Like struct {
	ID       ID     `json:"id, omitempty"`
	UserID   ID     `json:"user_id"`
	Target   Target `json:"target"` // Film or person
	TargetID ID     `json:"target_id"`
}
