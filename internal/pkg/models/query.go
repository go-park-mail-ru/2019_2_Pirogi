package models

type QueryListParams struct {
	Offset int    `json:"offset" valid:"numeric, optional"`
	Limit  int    `json:"limit" valid:"numeric, optional"`
	Genre  string `json:"genre" valid:"genre"`
}

func (qlp *QueryListParams) getBSON() string {
	return "$limit:" + string(qlp.Limit)
}
