package models

type (
	CreateBillRequest struct {
		Amount			uint64			`json:"amount"`
		Email			string			`json:"email"`
		CallbackURL		string			`json:"callbackURL"`
		RedirectURL		string			`json:"redirectURL"`
		Expiry			bool			`json:"expiry"`
		Duration		string			`json:"duration"`
	}
)