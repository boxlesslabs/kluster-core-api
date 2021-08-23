package repo

import (
	"errors"
	"github.com/klusters-core/api/config/db"
	"github.com/klusters-core/api/config/error_response"
	model "github.com/klusters-core/api/modules/clusters/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewClusterRepo(client db.StartMongoClient) *clusterRepo {
	col := db.NewMongoCollection("clusters", client)
	return &clusterRepo{col: col, client:client}
}

type (
	clusterRepo struct {
		col		db.MongoInterface
		client	db.StartMongoClient
	}

	ClusterRepo interface {
		CreateCluster(req *model.CreateRequest, userID primitive.ObjectID) (*model.ClusterModel, error)
		GetByID(id primitive.ObjectID) (*model.ClusterModel, error)
		ReturnClient() db.StartMongoClient
		DeleteByID(id primitive.ObjectID) error
		ChangeVisibility(clusterID primitive.ObjectID, visibility string) (*model.ClusterModel, error)
		UpdateDetails(req *model.CreateRequest, clusterID primitive.ObjectID) (*model.ClusterModel, error)
	}
)

func (cluster *clusterRepo) CreateCluster(req *model.CreateRequest, userID primitive.ObjectID) (*model.ClusterModel, error) {
	newCluster := model.SetCluster(req)
	newCluster.NewID()
	newCluster.TimeStamp()
	newCluster.SetOwnerID(userID)
	result, err := cluster.col.AddSingle(newCluster)
	if err != nil {
		return nil, errors.New(error_response.NotCreated{Resource: "Cluster"}.Error())
	}

	return cluster.GetByID(result.DocID)
}

func (cluster *clusterRepo) GetByID(id primitive.ObjectID) (*model.ClusterModel, error) {
	return cluster.Decode(cluster.col.GetSingleById(id))
}

func (cluster *clusterRepo) DeleteByID(id primitive.ObjectID) error {
	_, err := cluster.col.DeleteById(id)
	if err != nil {
		return err
	}

	return nil
}

func (cluster *clusterRepo) ReturnClient() db.StartMongoClient {
	return cluster.client
}

func (cluster *clusterRepo) ChangeVisibility(clusterID primitive.ObjectID, visibility string) (*model.ClusterModel, error) {
	result, err := cluster.col.UpdateById(clusterID, model.UpdateVisibilityQuery(visibility))
	if err != nil {
		return nil, err
	}

	return cluster.Decode(result)
}

func (cluster *clusterRepo) UpdateDetails(req *model.CreateRequest, clusterID primitive.ObjectID) (*model.ClusterModel, error) {
	result, err := cluster.col.UpdateById(clusterID, model.UpdateDetailsQuery(req))
	if err != nil {
		return nil, err
	}

	return cluster.Decode(result)
}

func (cluster *clusterRepo) Decode(result *mongo.SingleResult) (*model.ClusterModel, error) {
	var newCluster *model.ClusterModel
	if err := result.Decode(&newCluster); err != nil {
		return nil, err
	}
	return newCluster, nil
}