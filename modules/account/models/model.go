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
	"github.com/klusters-core/api/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var Util utils.GeneralUtil

type (
	AccountsModel struct {
		ID              	primitive.ObjectID     	`json:"id" bson:"_id,omitempty"`
		Phone     			string 				   	`json:"phone" bson:"phone,omitempty"`
		FullName 			string 					`json:"full_name" bson:"full_name,omitempty"`
		Email           	string  				`json:"email" bson:"email,omitempty"`
		ImageUrl  			string 					`json:"image_url" bson:"image_url,omitempty"`
		DefaultCluster 		DefaultCluster         	`json:"default_cluster" bson:"default_cluster"`
		Clusters       		[]string               	`json:"clusters" bson:"clusters,omitempty"`
		Status           	string                 	`json:"status" bson:"status,omitempty"`
		FcmId            	string                 	`json:"fcm_id" bson:"fcm_id,omitempty"`
		CreatedAt        	time.Time              	`json:"created_at" bson:"created_at,omitempty"`
		UpdatedAt        	time.Time              	`json:"updated_at" bson:"updated_at,omitempty"`
	}

	DefaultCluster struct {
		ClusterId 			primitive.ObjectID 		`json:"cluster_id" bson:"cluster_id"`
		Name        		string             		`json:"name" bson:"name"`
		Owner       		bool               		`json:"owner" bson:"owner"`
		Role        		string           		`json:"role" bson:"role"`
		Permissions 		[]string           		`json:"permissions" bson:"permissions"`
	}

	AccountRequest struct {
		Phone     			string 				   	`json:"phone" validate:"required"`
		Password			string					`json:"password" validate:"required,gt=5,lt=10"`
	}
)

func SetAccount(request *AccountRequest) *AccountsModel {
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
		UpdatedAt:      time.Now(),
	}
}

func (account *AccountsModel) NewID() {
	account.DefaultCluster.ClusterId = primitive.NewObjectID()
}

func (account *AccountsModel) TimeStamp() {
	account.CreatedAt = time.Now()
	account.UpdatedAt = time.Now()
}

func (account *AccountsModel) MakeOwner() {
	account.DefaultCluster.Owner = true
}