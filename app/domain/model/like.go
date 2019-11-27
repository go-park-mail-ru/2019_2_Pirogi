package model

type LikeRepository interface {
	Insert(like Like) (ID, error)
	Delete(id ID) bool
	GetMany(target Target, id ID)
}

type Like struct {
	UserID   ID     `json:"user_id" valid:"numeric"`
	Target   Target `json:"target" valid:"target"` // Film, person or review
	TargetID ID     `json:"target_id" valid:"numeric"`
}
