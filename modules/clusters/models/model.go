//=============================================================================
// developer: boxlesslabsng@gmail.com
//=============================================================================

/**
 **
 * @struct ClusterModel - defines model for db collection
 * @Members - defines structure for member accounts in a cluster
 * @Settings - defines policies for a cluster
 * @CreateRequest
 **
**/

package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	ClusterModel struct {
		ID              	primitive.ObjectID     	`json:"_id" bson:"_id,omitempty"`
		Name				string					`json:"name" bson:"name,omitempty"`
		UserID				primitive.ObjectID		`json:"user_id" bson:"user_id,omitempty"`
		Members				[]Members				`json:"members" bson:"members,omitempty"`
		CreatedAt        	time.Time              	`json:"created_at" bson:"created_at,omitempty"`
		UpdatedAt        	time.Time              	`json:"updated_at" bson:"updated_at,omitempty"`
		Visibility			string					`json:"visibility" bson:"visibility,omitempty"`
		Description			string					`json:"description" bson:"description,omitempty"`
	}

	Members struct {
		Role        		[]string           		`json:"role" bson:"role,omitempty"`
		AccountID   		primitive.ObjectID 		`json:"account_id" bson:"account_id"`
		Permissions 		[]string           		`json:"permissions" bson:"permissions,omitempty"`
	}

	Settings struct {
		Notification 		bool   					`json:"notification" bson:"notification,omitempty"`
		MaxMembers			int16					`json:"max_members" bson:"max_members,omitempty"`
	}

	CreateRequest struct {
		Name				string 					`json:"name" validate:"required"`
		Visibility			string					`json:"visibility" validate:"required"`
		Description			string 					`json:"description"`
	}

	VisibilityRequest struct {
		Visibility			bool					`json:"visibility" validate:"required"`
	}
)

func SetCluster(request *CreateRequest) *ClusterModel {
	return &ClusterModel{
		Name:        request.Name,
		Members:     nil,
		Visibility:  request.Visibility,
		Description: request.Description,
	}
}

func (c *ClusterModel) NewID() {
	c.ID = primitive.NewObjectID()
}

func (c *ClusterModel) TimeStamp() {
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
}

func (c *ClusterModel) UpdatedStamp() {
	c.UpdatedAt = time.Now()
}

func (c *ClusterModel) SetOwnerID(id primitive.ObjectID) {
	c.UserID = id
}

func UpdateVisibilityQuery(visibility string) bson.M {
	return bson.M{"visibility": visibility, "created_at": time.Now()}
}

func UpdateDetailsQuery(req *CreateRequest) bson.M {
	return bson.M{"name": req.Name, "description": req.Description}
}