//=============================================================================
// developer: boxlesslabsng@gmail.com
//=============================================================================
 
/**
 **
 * @struct AccountsModel
 * @struct DefaultCluster
 **
**/

package model

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/klusters-core/api/modules/auth/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
	"time"
)

const (
	StatusInactive string = "inactive"
	StatusActive string = "activated"
	StatusDeactivated string = "deactivated"
)

type (
	AccountsModel struct {
		ID              	primitive.ObjectID     	`json:"id" bson:"_id,omitempty"`
		Phone     			string 				   	`json:"phone" bson:"phone,omitempty"`
		FullName 			string 					`json:"full_name,omitempty" bson:"full_name,omitempty"`
		Email           	string  				`json:"email,omitempty" bson:"email,omitempty"`
		ImageUrl  			string 					`json:"image_url,omitempty" bson:"image_url,omitempty"`
		DefaultCluster 		DefaultCluster         	`json:"default_cluster,omitempty" bson:"default_cluster,omitempty"`
		Clusters       		[]string               	`json:"clusters,omitempty" bson:"clusters,omitempty"`
		Status           	string                 	`json:"status,omitempty" bson:"status,omitempty"`
		FcmId            	string                 	`json:"fcm_id,omitempty" bson:"fcm_id,omitempty"`
		CreatedAt        	time.Time              	`json:"created_at,omitempty" bson:"created_at,omitempty"`
		UpdatedAt        	time.Time              	`json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	}

	DefaultCluster struct {
		ClusterId 			primitive.ObjectID 		`json:"cluster_id" bson:"cluster_id"`
		Name        		string             		`json:"name" bson:"name"`
		Owner       		bool               		`json:"owner" bson:"owner"`
		Role        		string           		`json:"role" bson:"role"`
		Permissions 		[]string           		`json:"permissions" bson:"permissions"`
	}
)

func SetAccount(request *models.AuthModel) *AccountsModel {
	return &AccountsModel{
		Phone:          request.Phone,
		DefaultCluster: DefaultCluster{},
		Clusters:       nil,
		Status:         StatusInactive,
		CreatedAt:      time.Now(),
	}
}

func (account *AccountsModel) NewID() {
	account.ID = primitive.NewObjectID()
	account.DefaultCluster.ClusterId = primitive.NewObjectID()
}

func (account *AccountsModel) MakeOwner() {
	account.DefaultCluster.Owner = true
}

func (account *AccountsModel) SetDeactivate() {
	account.Status = StatusDeactivated
}

func (account *AccountsModel) ValidateProfileReq() error {
	return validation.ValidateStruct(account,
		validation.Field(&account.Email, validation.Required, is.Email),
		validation.Field(&account.ImageUrl, is.URL.Error("image url is not valid")),
		validation.Field(&account.FullName, validation.Required),
	)
}

func (account *AccountsModel) ValidateStatus() error {
	return validation.ValidateStruct(account,
		validation.Field(&account.Status, validation.In(StatusDeactivated, StatusActive).Error(fmt.Sprintf("status can only be %s or %s", StatusActive, StatusDeactivated))),
		)
}

func (account *AccountsModel) ValidateFCMID() error {
	return validation.ValidateStruct(account,
		validation.Field(&account.FcmId, validation.Required),
		)
}

func (account *AccountsModel) SetUserID(id *primitive.ObjectID) {
	account.ID = *id
}

func (account *AccountsModel) SetEmail(email *string) {
	account.Email = strings.ToLower(*email)
}