//=============================================================================
// developer: boxlesslabsng@gmail.com
// models for decoding json response from paystack apis
//=============================================================================

/**
 **
 * @struct InitializeTransaction
 * @struct VerifyTransaction
 * @struct ListTransactions
 * @struct FetchTransaction
 * @struct ChargeAuthorization
 **
**/


package models

import "time"

type (
	InitializeTransaction struct {
		Status  					bool   		`json:"status"`
		Message 					string 		`json:"message"`
		Data    struct {
			AuthorizationURL 		string 		`json:"authorization_url"`
			AccessCode       		string 		`json:"access_code"`
			Reference        		string 		`json:"reference"`
		} 										`json:"data"`
	}

	VerifyTransaction struct {
		Status  					bool   		`json:"status"`
		Message 					string 		`json:"message"`
		Data    struct {
			ID              		int         `json:"id"`
			Domain          		string      `json:"domain"`
			Status          		string      `json:"status"`
			Reference       		string      `json:"reference"`
			Amount          		int         `json:"amount"`
			Message         		interface{} `json:"message"`
			GatewayResponse 		string      `json:"gateway_response"`
			PaidAt          		time.Time   `json:"paid_at"`
			CreatedAt       		time.Time   `json:"created_at"`
			Channel         		string      `json:"channel"`
			Currency       		 	string      `json:"currency"`
			IPAddress       		string      `json:"ip_address"`
			Metadata        		string      `json:"metadata"`
			Log             struct {
				StartTime 			int          `json:"start_time"`
				TimeSpent 			int          `json:"time_spent"`
				Attempts  			int          `json:"attempts"`
				Errors    			int          `json:"errors"`
				Success   			bool         `json:"success"`
				Mobile    			bool         `json:"mobile"`
				Input     			[]interface{} `json:"input"`
				History   			[]struct {
					Type    		string 		`json:"type"`
					Message 		string 		`json:"message"`
					Time    		int    		`json:"time"`
				} 								`json:"history"`
			} 									`json:"log"`
			Fees          			int       	`json:"fees"`
			FeesSplit     			interface{}	 `json:"fees_split"`
			Authorization struct {
				AuthorizationCode 	string      `json:"authorization_code"`
				Bin               	string      `json:"bin"`
				Last4             	string      `json:"last4"`
				ExpMonth          	string      `json:"exp_month"`
				ExpYear           	string      `json:"exp_year"`
				Channel           	string      `json:"channel"`
				CardType          	string      `json:"card_type"`
				Bank              	string      `json:"bank"`
				CountryCode       	string      `json:"country_code"`
				Brand             	string      `json:"brand"`
				Reusable          	bool        `json:"reusable"`
				Signature         	string      `json:"signature"`
				AccountName       interface{} 	`json:"account_name"`
			} `json:"authorization"`
			Customer struct {
				ID           		int         `json:"id"`
				FirstName    		interface{} `json:"first_name"`
				LastName     		interface{} `json:"last_name"`
				Email        		string      `json:"email"`
				CustomerCode 		string      `json:"customer_code"`
				Phone        		interface{} `json:"phone"`
				Metadata     		interface{} `json:"metadata"`
				RiskAction   		string      `json:"risk_action"`
			} 									`json:"customer"`
			Plan            		interface{} `json:"plan"`
			OrderID         		interface{} `json:"order_id"`
			PaidAt2          		time.Time   `json:"paidAt"`
			CreatedAt2       		time.Time   `json:"createdAt"`
			RequestedAmount 		int         `json:"requested_amount"`
			TransactionDate 		time.Time   `json:"transaction_date"`
			PlanObject      struct {
			} 									`json:"plan_object"`
			Subaccount struct {
			} 									`json:"subaccount"`
		} 										`json:"data"`
	}

	ListTransactions struct {
		Status  					bool   			`json:"status"`
		Message 					string 			`json:"message"`
		Data    []struct {
			ID             			int        		`json:"id"`
			Domain          		string      	`json:"domain"`
			Status          		string      	`json:"status"`
			Reference       		string      	`json:"reference"`
			Amount          		int         	`json:"amount"`
			Message         		interface{} 	`json:"message"`
			GatewayResponse 		string      	`json:"gateway_response"`
			Paidat          		time.Time   	`json:"paid_at"`
			Createdat       		time.Time   	`json:"created_at"`
			Channel         		string      	`json:"channel"`
			Currency        		string      	`json:"currency"`
			IPAddress       		string      	`json:"ip_address"`
			Metadata        		interface{} 	`json:"metadata"`
			Log             struct {
				StartTime 			int          	`json:"start_time"`
				TimeSpent 			int          	`json:"time_spent"`
				Attempts  			int          	`json:"attempts"`
				Errors    			int          	`json:"errors"`
				Success   			bool         	`json:"success"`
				Mobile    			bool         	`json:"mobile"`
				Input     			[]interface{} 	`json:"input"`
				History   []struct {
					Type    		string 			`json:"type"`
					Message 		string 			`json:"message"`
					Time    		int    			`json:"time"`
				} 									`json:"history"`
			} 										`json:"log"`
			Fees      				int       		`json:"fees"`
			FeesSplit 				interface{} 	`json:"fees_split"`
			Customer  struct {
				ID           		int     	    `json:"id"`
				FirstName    		interface{} 	`json:"first_name"`
				LastName     		interface{} 	`json:"last_name"`
				Email        		string      	`json:"email"`
				Phone        		interface{} 	`json:"phone"`
				Metadata     		interface{} 	`json:"metadata"`
				CustomerCode 		string      	`json:"customer_code"`
				RiskAction   		string     		`json:"risk_action"`
			} `json:"customer"`
			Authorization struct {
				AuthorizationCode 	string      	`json:"authorization_code"`
				Bin               	string     		`json:"bin"`
				Last4             	string      	`json:"last4"`
				ExpMonth          	string      	`json:"exp_month"`
				ExpYear          	string      	`json:"exp_year"`
				Channel           	string      	`json:"channel"`
				CardType          	string      	`json:"card_type"`
				Bank              	string      	`json:"bank"`
				CountryCode       	string      	`json:"country_code"`
				Brand             	string      	`json:"brand"`
				Reusable          	bool        	`json:"reusable"`
				Signature         	string      	`json:"signature"`
				AccountName       	interface{} 	`json:"account_name"`
			} 										`json:"authorization"`
			Plan struct {} 							`json:"plan"`
			Subaccount struct {}					`json:"subaccount"`
			OrderID         		interface{} 	`json:"order_id"`
			PaidAt          		time.Time   	`json:"paidAt"`
			CreatedAt       		time.Time   	`json:"createdAt"`
			RequestedAmount 		int         	`json:"requested_amount"`
		} `json:"data"`
		Meta struct {
			Total      				int 			`json:"total"`
			TotalVolume 			int 			`json:"total_volume"`
			Skipped     			int 			`json:"skipped"`
			PerPage     			int 			`json:"perPage"`
			Page        			int 			`json:"page"`
			PageCount   			int 			`json:"pageCount"`
		} 											`json:"meta"`
	}

	FetchTransaction struct {
		Status  				bool   				`json:"status"`
		Message 				string 				`json:"message"`
		Data    struct {
			ID             		int         		`json:"id"`
			Domain          	string      		`json:"domain"`
			Status          	string      		`json:"status"`
			Reference       	string      		`json:"reference"`
			Amount          	int         		`json:"amount"`
			Message         	interface{} 		`json:"message"`
			GatewayResponse 	string      		`json:"gateway_response"`
			Paidat          	time.Time   		`json:"paid_at"`
			Createdat       	time.Time   		`json:"created_at"`
			Channel         	string      		`json:"channel"`
			Currency        	string      		`json:"currency"`
			IPAddress       	string      		`json:"ip_address"`
			Metadata        	string      		`json:"metadata"`
			Log             struct {
				StartTime 		int           		`json:"start_time"`
				TimeSpent 		int           		`json:"time_spent"`
				Attempts  		int           		`json:"attempts"`
				Errors    		int           		`json:"errors"`
				Success   		bool          		`json:"success"`
				Mobile    		bool          		`json:"mobile"`
				Input     		[]interface{} 		`json:"input"`
				History   []struct {
					Type    	string 				`json:"type"`
					Message 	string 				`json:"message"`
					Time    	int    				`json:"time"`
				} 									`json:"history"`
			} 										`json:"log"`
			Fees          		int         		`json:"fees"`
			FeesSplit     		interface{} 		`json:"fees_split"`
			Authorization struct {
				AuthorizationCode string      		`json:"authorization_code"`
				Bin               string      		`json:"bin"`
				Last4             string      		`json:"last4"`
				ExpMonth          string      		`json:"exp_month"`
				ExpYear           string      		`json:"exp_year"`
				Channel           string      		`json:"channel"`
				CardType          string      		`json:"card_type"`
				Bank              string      		`json:"bank"`
				CountryCode       string      		`json:"country_code"`
				Brand             string      		`json:"brand"`
				Reusable          bool       		`json:"reusable"`
				Signature         string      		`json:"signature"`
				AccountName       interface{} 		`json:"account_name"`
			} 										`json:"authorization"`
			Customer struct {
				ID           		int         	`json:"id"`
				FirstName    		interface{} 	`json:"first_name"`
				LastName     		interface{} 	`json:"last_name"`
				Email       		string      	`json:"email"`
				CustomerCode 		string      	`json:"customer_code"`
				Phone        		interface{} 	`json:"phone"`
				Metadata     		interface{} 	`json:"metadata"`
				RiskAction   		string      	`json:"risk_action"`
			} 										`json:"customer"`
			Plan 					struct {} 		`json:"plan"`
			Subaccount 				struct {} 		`json:"subaccount"`
			OrderID         		interface{} 	`json:"order_id"`
			PaidAt          		time.Time  		`json:"paidAt"`
			CreatedAt       		time.Time   	`json:"createdAt"`
			RequestedAmount 		int         	`json:"requested_amount"`
		} 											`json:"data"`
	}

	ChargeAuthorization struct {
		Status  				bool   				`json:"status"`
		Message 				string 				`json:"message"`
		Data    struct {
			Amount         		int         		`json:"amount"`
			Currency        	string      		`json:"currency"`
			TransactionDate 	time.Time   		`json:"transaction_date"`
			Status          	string      		`json:"status"`
			Reference       	string      		`json:"reference"`
			Domain          	string      		`json:"domain"`
			Metadata        	string      		`json:"metadata"`
			GatewayResponse 	string      		`json:"gateway_response"`
			Message         	interface{} 		`json:"message"`
			Channel         	string      		`json:"channel"`
			IPAddress       	interface{} 		`json:"ip_address"`
			Log             	interface{} 		`json:"log"`
			Fees            	int         		`json:"fees"`
			Authorization   struct {
				AuthorizationCode string      		`json:"authorization_code"`
				Bin               string      		`json:"bin"`
				Last4             string      		`json:"last4"`
				ExpMonth          string      		`json:"exp_month"`
				ExpYear           string      		`json:"exp_year"`
				Channel           string      		`json:"channel"`
				CardType          string      		`json:"card_type"`
				Bank              string      		`json:"bank"`
				CountryCode       string      		`json:"country_code"`
				Brand             string      		`json:"brand"`
				Reusable          bool        		`json:"reusable"`
				Signature         string      		`json:"signature"`
				AccountName       interface{} 		`json:"account_name"`
			} `json:"authorization"`
			Customer struct {
				ID           	int         		`json:"id"`
				FirstName    	interface{} 		`json:"first_name"`
				LastName     	interface{} 		`json:"last_name"`
				Email        	string      		`json:"email"`
				CustomerCode 	string      		`json:"customer_code"`
				Phone        	interface{} 		`json:"phone"`
				Metadata     	interface{} 		`json:"metadata"`
				RiskAction   	string      		`json:"risk_action"`
			} 										`json:"customer"`
			Plan 				interface{} 		`json:"plan"`
			ID   int         						`json:"id"`
		} 											`json:"data"`
	}
)