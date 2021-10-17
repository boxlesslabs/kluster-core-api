package models

import "time"

type (
	AuthResponse struct {
		Error			string			`json:"error,omitempty"`
		Success			string			`json:"success"`
		ErrorCode		interface{}		`json:"errorCode,omitempty"`
		Data    struct {
			User struct {
				ID     string `json:"id"`
				Mobile string `json:"mobile"`
			} `json:"user"`
		} `json:"data"`
		Message string `json:"message"`
	}


	RefreshTokenResponse struct {
		Error			string			`json:"error,omitempty"`
		Success			string			`json:"success"`
		ErrorCode		interface{}		`json:"errorCode,omitempty"`
		Data    struct {
			AccessToken string `json:"accessToken"`
			ExpiresIn   int    `json:"expiresIn"`
		} `json:"data"`
	}

	CreateBillResponse struct {
		Error			string			`json:"error,omitempty"`
		Success			string			`json:"success"`
		ErrorCode		interface{}		`json:"errorCode,omitempty"`
		Data    struct {
			Split             bool          `json:"split"`
			TransactionCharge int           `json:"transactionCharge"`
			Type              string        `json:"type"`
			Expiry            bool          `json:"expiry"`
			QueuedForRetry    bool          `json:"queuedForRetry"`
			Status            string        `json:"status"`
			Message           []interface{} `json:"message"`
			Settled           bool          `json:"settled"`
			Retries           int           `json:"retries"`
			ID                string        `json:"_id"`
			Deleted           bool          `json:"deleted"`
			MerchantID        string        `json:"merchantId"`
			Email             string        `json:"email"`
			Amount            int           `json:"amount"`
			RedirectURL       string        `json:"redirectURL"`
			Duration          string        `json:"duration"`
			Bearer            string        `json:"bearer"`
			PrincipalAccount  string        `json:"principalAccount"`
			Callback          struct {
				CallbackURL string `json:"callbackURL"`
			} `json:"callback"`
			PaymentRef  string    `json:"paymentRef"`
			QrCode      string    `json:"qrCode"`
			CreatedAt   time.Time `json:"createdAt"`
			UpdatedAt   time.Time `json:"updatedAt"`
			V           int       `json:"__v"`
			CallBackURL string    `json:"callBackURL"`
		} `json:"data"`
	}

	VerifyBillResponse struct {
		Error			string			`json:"error,omitempty"`
		Success			string			`json:"success"`
		ErrorCode		interface{}		`json:"errorCode,omitempty"`
		PaymentStatus     string `json:"paymentStatus"`
		Message           string `json:"message"`
		PaymentStatusCode string `json:"paymentStatusCode"`
	}

	OtpResponse struct {
		Error			string			`json:"error,omitempty"`
		Success			string			`json:"success"`
		ErrorCode		interface{}		`json:"errorCode,omitempty"`
		Message string `json:"message"`
		Data    struct {
			RefreshToken string `json:"refreshToken"`
			AccessToken  string `json:"accessToken"`
			ExpiresIn    int    `json:"expiresIn"`
		} `json:"data"`
	}

	PhoneTransferResponse struct {
		Error			string			`json:"error,omitempty"`
		Success			string			`json:"success"`
		ErrorCode		interface{}		`json:"errorCode,omitempty"`
		Data    struct {
			Transaction struct {
				Reference string `json:"reference"`
				Amount    int    `json:"amount"`
			} `json:"transaction"`
		} `json:"data"`
		Message string `json:"message"`
	}


	// REQUESTS
	PhoneTransferRequest struct {
		Mobile			string			`json:"mobile"`
		Amount			uint64			`json:"amount"`
	}

	CreateBillRequest struct {
		Amount			uint64			`json:"amount"`
		Email			string			`json:"email"`
		CallbackURL		string			`json:"callbackURL"`
		RedirectURL		string			`json:"redirectURL"`
		Expiry			bool			`json:"expiry"`
		Duration		string			`json:"duration"`
	}
)