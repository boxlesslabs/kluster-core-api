//=============================================================================
// developer: boxlesslabsng@gmail.com
//=============================================================================

/**
 **
 * @struct AuthModel - defines models for db collection
 * @struct JwtCustomClaims - extend struct with custom fields for signing tokens
 * @struct ChangePasswordRequest
 **
**/

package models

import (
	"github.com/dgrijalva/jwt-go"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"regexp"
	"time"
)

type (
	AuthModel struct {
		ID              	primitive.ObjectID     	`json:"_id,omitempty" bson:"_id"`
		Phone     			string 				   	`json:"phone" bson:"phone"`
		Password			string					`json:"password" bson:"password"`
		CreatedAt        	time.Time              	`json:"created_at,omitempty" bson:"created_at"`
		UpdatedAt        	time.Time              	`json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	}

	JwtCustomClaims struct {
		AuthID				primitive.ObjectID		`json:"auth_id"`
		UserID 				primitive.ObjectID 		`json:"user_id"`
		Phone 				string 					`json:"phone"`
		jwt.StandardClaims
	}

	ChangePasswordRequest struct {
		OldPassword			string					`json:"old_password"`
		NewPassword			string					`json:"new_password"`
	}
)

// auth constructor
func SetAuth(request *AuthModel) *AuthModel {
	return &AuthModel{
		ID:primitive.NewObjectID(),
		Phone:          request.Phone,
		Password:       request.Password,
		CreatedAt:      time.Now(),
	}
}

// auth model struct
func (auth *AuthModel) NewID() {
	auth.ID = primitive.NewObjectID()
}

func (auth *AuthModel) UpdatedStamp() {
	auth.UpdatedAt = time.Now()
}

func (auth *AuthModel) CreatedStamp() {
	auth.CreatedAt = time.Now()
}

func (auth *AuthModel) ValidateAuth() error {
	return validation.ValidateStruct(auth,
		validation.Field(&auth.Phone, validation.Required, validation.Match(regexp.MustCompile(`^(234)\d{10}$`))),
		validation.Field(&auth.Password, validation.Required),
		)
}

func (auth *AuthModel) ValidateForgotPassword() error {
	return validation.ValidateStruct(auth,
		validation.Field(&auth.Phone, validation.Required, validation.Match(regexp.MustCompile(`^(234)\d{10}$`))),
	)
}


// change password struct
func (auth *ChangePasswordRequest) ValidateChangePassword() error {
	return validation.ValidateStruct(auth,
		validation.Field(&auth.OldPassword, validation.Required),
		validation.Field(&auth.NewPassword, validation.Required),
	)
}