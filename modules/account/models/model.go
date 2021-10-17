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
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/klusters-core/api/modules/auth/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	AccountsModel struct {
		ID              	primitive.ObjectID     	`json:"id" bson:"_id,omitempty"`
		Phone     			string 				   	`json:"phone" bson:"phone,omitempty"`
		FullName 			string 					`json:"full_name" bson:"full_name,omitempty"`
		Email           	string  				`json:"email" bson:"email,omitempty"`
		ImageUrl  			string 					`json:"image_url" bson:"image_url,omitempty"`
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
		FullName:       "",
		Email:          "",
		ImageUrl:       "",
		DefaultCluster: DefaultCluster{},
		Clusters:       nil,
		Status:         "inactive",
		FcmId:          "",
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

func (account *AccountsModel) ValidateProfileReq() error {
	return validation.ValidateStruct(account,
		validation.Field(&account.Email, validation.Required, is.Email),
		validation.Field(&account.ImageUrl, is.URL.Error("image url is not valid")),
		validation.Field(&account.FullName, validation.Required),
	)
}