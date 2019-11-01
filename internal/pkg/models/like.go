package models

type Like struct {
	UserID   ID     `json:"user_id" valid:"numeric"`
	Target   Target `json:"target" valid:"target"` // Film, person or review
	TargetID ID     `json:"target_id" valid:"numeric"`
}
