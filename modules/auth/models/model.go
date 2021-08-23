//=============================================================================
// developer: boxlesslabsng@gmail.com
//=============================================================================

/**
 **
 * @struct AuthModel - defines models for db collection
 * @struct AuthAccount
 * @struct JwtCustomClaims - extend struct with custom fields for signing tokens
 * @struct AuthRequest
 * @struct ForgotPasswordRequest
 **
**/

package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/klusters-core/api/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var Util utils.GeneralUtil

type (
	AuthModel struct {
		ID              	primitive.ObjectID     	`json:"_id" bson:"_id,omitempty"`
		Phone     			string 				   	`json:"phone" bson:"phone,omitempty"`
		Password			string					`json:"-" bson:"password,omitempty"`
		CreatedAt        	time.Time              	`json:"created_at" bson:"created_at,omitempty"`
		UpdatedAt        	time.Time              	`json:"updated_at" bson:"updated_at,omitempty"`
	}

	AuthAccount struct {
		ID              	primitive.ObjectID     	`json:"_id" bson:"_id,omitempty"`
		UserID				primitive.ObjectID		`json:"user_id" bson:"user_id"`
		Phone     			string 				   	`json:"phone" bson:"phone,omitempty"`
		Email           	string  				`json:"email" bson:"email,omitempty"`
		ImageUrl  			string 					`json:"image_url" bson:"image_url,omitempty"`
		Status           	string                 	`json:"status" bson:"status,omitempty"`
	}

	JwtCustomClaims struct {
		AuthID				primitive.ObjectID		`json:"auth_id"`
		UserID 				primitive.ObjectID 		`json:"user_id"`
		Phone 				string 					`json:"phone"`
		jwt.StandardClaims
	}

	AuthRequest struct {
		Phone     			string 				   	`json:"phone" validate:"required"`
		Password			string					`json:"password" validate:"required"`
	}

	ForgotPasswordRequest struct {
		Phone     			string 				   	`json:"phone" validate:"required"`
	}

	ChangePasswordRequest struct {
		OldPassword			string					`json:"old_password" validate:"required"`
		NewPassword			string					`json:"new_password" validate:"required,gt=5,lt=18"`
	}
)

func SetAuth(request *AuthRequest) *AuthModel {
	return &AuthModel{
		Phone:          request.Phone,
		Password:       request.Password,
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
	}
}

func (auth *AuthModel) NewID() {
	auth.ID = primitive.NewObjectID()
}

func (auth *AuthModel) EncryptPassword() {
	auth.Password = Util.HashPassword(auth.Password)
}

func (auth *AuthModel) TimeStamp() {
	auth.CreatedAt = time.Now()
	auth.UpdatedAt = time.Now()
}

func (auth *AuthModel) UpdatedStamp() {
	auth.UpdatedAt = time.Now()
}

func GetByPhoneQuery(phone string) bson.M {
	return bson.M{"phone": phone}
}

func GetByCredentialsQuery(phone string, password string) bson.M {
	return bson.M{"phone": phone, "password": Util.HashPassword(password)}
}

func UpdatePasswordQuery(password string) bson.M {
	return bson.M{"password": Util.HashPassword(password)}
}

func GetPasswordQuery(userID primitive.ObjectID, password string) bson.M {
	return bson.M{"_id": userID, "password": Util.HashPassword(password)}
}