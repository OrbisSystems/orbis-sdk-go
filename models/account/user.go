package account

type GetUsers struct {
	With     []string `json:"with"`
	Search   string   `json:"search"`
	Filters  Filters  `json:"filters"`
	OrderBy  OrderBy  `json:"order_by"`
	Paginate Paginate `json:"paginate"`
}

type Filters struct {
	Role          string `json:"role"`
	Firm          string `json:"firm"`
	Branch        string `json:"branch"`
	PhoneVerified bool   `json:"phone_verified"`
	CreatedAt     FromTo `json:"created_at"`
}

type OrderBy struct {
	ID        Branch `json:"id"`
	Role      Branch `json:"role"`
	FirstName Branch `json:"first_name"`
	LastName  Branch `json:"last_name"`
	Branch    Branch `json:"branch"`
	CreatedAt Branch `json:"created_at"`
}

type Branch struct {
	Direction string `json:"direction"`
	Priority  int64  `json:"priority"`
}

type CreateUser struct {
	Role        string `json:"role"`
	InitiatorID int64  `json:"initiator_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Password    string `json:"password"`
}
