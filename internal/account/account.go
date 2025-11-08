package account

type Account struct {
	ID      string  `json:"id"`
	Balance float32 `json:"balance"`
	OwnerID string  `json:"owner_id"`
}

type NewAccount struct {
	Balance float32 `json:"balance"`
	OwnerID string  `json:"owner_id"`
}
