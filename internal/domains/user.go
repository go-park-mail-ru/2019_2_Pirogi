package domains

type UserRepository interface {
	Insert(newUser NewUser) (ID, error)
	Update(id ID, user User) error
	Delete(id ID) bool
	Get(id ID) User
	GetMany(target Target, id ID) []User
	MakeTrunc(user User) UserTrunc
}

type NewUser struct {
	Email    string `json:"email" valid:"email"`
	Password string `json:"password" valid:"password"`
	Username string `json:"username" valid:"stringlength(2|50)"`
}

type UserTrunc struct {
	ID          ID     `json:"id" valid:"numeric"`
	Username    string `json:"username" valid:"title"`
	Mark        Mark   `json:"mark" valid:"mark, optional"`
	Description string `json:"description" valid:"description"`
	Image       Image  `json:"image" valid:"image, optional"`
}

type User struct {
	ID          ID     `json:"id" valid:"numeric"`
	Email       string `json:"email" valid:"email"`
	Password    string `json:"password" valid:"password"`
	Username    string `json:"username" valid:"title"`
	Mark        Mark   `json:"mark" valid:"mark, optional"`
	Description string `json:"description" valid:"description"`
	Image       Image  `json:"image" valid:"image, optional"`
}