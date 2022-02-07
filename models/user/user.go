package models

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	RepCode   string `json:"rep_code"`
	Role      struct {
		Name string `json:"name"`
	} `json:"role"`
	Locale struct {
		ID     int    `json:"id"`
		Locale string `json:"locale"`
		Name   string `json:"name"`
	} `json:"locale"`
	Initiator  []interface{} `json:"initiator"`
	Timestamps struct {
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	} `json:"timestamps"`
}
