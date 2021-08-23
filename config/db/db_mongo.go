//=============================================================================
// developer: boxlesslabsng@gmail.com
//=============================================================================

/**
 * Define Mongo database connection
 * Return Existing collections for db operations
 **
 * @struct mongoCollection
**/

package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	mongoCollection struct {
		*mongo.Collection
	}

	unMarshallObjectId struct {
		DocID primitive.ObjectID `json:"id" bson:"id"`
	}

	MongoInterface interface {
		AddSingle(data interface{}) (*unMarshallObjectId, error)
		GetSingleById(id primitive.ObjectID) *mongo.SingleResult
		GetSingleByQuery(query interface{}) *mongo.SingleResult
		Count(query interface{}) (int64, error)
		GetMany(query interface{}) (*mongo.Cursor, error)
		GetPaginated(query interface{}, skip int64, size int64) (*mongo.Cursor, int64, error)
		GetManyWithSort(query interface{}, skip int64, size int64, field string, order int) (*mongo.Cursor, int64, error)
		GetProjected(query interface{}, skip int64, size int64, fields bson.M) (*mongo.Cursor, int64, error)
		Aggregate(query mongo.Pipeline) (*mongo.Cursor, error)
		UpdateById(id primitive.ObjectID, data interface{}) (*mongo.SingleResult, error)
		UpdateByQuery(query interface{}, data interface{}) (*mongo.UpdateResult, error)
		DeleteById(id primitive.ObjectID) (*mongo.DeleteResult, error)
		DeleteByQuery(query interface{}) (bool, error)
		GetSingleByProjection(query interface{}, project interface{}) *mongo.SingleResult
	}
)

var (
	ctx = context.TODO()
)

func NewMongoCollection(col string, client StartMongoClient) *mongoCollection {
	return &mongoCollection{client.SetCollection(col)}
}

// insert a single document
func (db *mongoCollection) AddSingle(data interface{}) (*unMarshallObjectId, error) {
	res, err := db.Collection.InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}

	var getObjectId unMarshallObjectId
	getObjectId.DocID =  res.InsertedID.(primitive.ObjectID)

	return &getObjectId, nil
}

// get single document from collection
func (db *mongoCollection) GetSingleById(id primitive.ObjectID) *mongo.SingleResult {
	return db.Collection.FindOne(ctx, bson.M{"_id": id})
}

// get single document from collection with custom query
func (db *mongoCollection) GetSingleByQuery(query interface{}) *mongo.SingleResult {
	return db.Collection.FindOne(ctx, query)
}

func (db *mongoCollection) GetSingleByProjection(query interface{}, project interface{}) *mongo.SingleResult {
	return db.Collection.FindOne(ctx, query, options.FindOne().SetProjection(project))
}

// return count of documents with custom query
func (db *mongoCollection) Count(query interface{}) (int64, error) {
	count, err := db.Collection.CountDocuments(ctx, query)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// get many documents with custom query, returns a mongo cursor
func (db *mongoCollection) GetMany(query interface{}) (*mongo.Cursor, error) {
	cursor, err := db.Collection.Find(ctx, query)
	if err != nil {
		return nil, err
	}
	return cursor, nil
}

// get many documents with custom query, returns a paginated list of queried documents
func (db *mongoCollection) GetPaginated(query interface{}, skip int64, size int64) (*mongo.Cursor, int64, error) {
	count, err := db.Count(query)
	if err != nil {
		return nil, 0, err
	}

	projectOption := options.Find().SetSkip(skip).SetLimit(size)
	cursor, err := db.Collection.Find(ctx, query, projectOption)
	if err != nil {
		return nil, 0, err
	}

	return cursor, count, nil
}

// returns a paginated list with options to sort
func (db *mongoCollection) GetManyWithSort(query interface{}, skip int64, size int64, field string, order int) (*mongo.Cursor, int64, error) {
	count, err := db.Count(query)
	if err != nil {
		return nil, 0, err
	}

	projectOption := options.Find().SetSort(bson.M{field: order}).SetSkip(skip).SetLimit(size)
	cursor, err := db.Collection.Find(ctx, query, projectOption)
	if err != nil {
		return nil, 0, err
	}

	return cursor, count, nil
}

// returns a paginated list of projected documents
func (db *mongoCollection) GetProjected(query interface{}, skip int64, size int64, fields bson.M) (*mongo.Cursor, int64, error) {
	count, err := db.Count(query)
	if err != nil {
		return nil, 0, err
	}

	projectOption := options.Find().SetSkip(skip).SetLimit(size).SetProjection(fields)
	cursor, err := db.Collection.Find(ctx, query, projectOption)
	if err != nil {
		return nil, 0, err
	}
	return cursor, count, nil
}

func (db *mongoCollection) Aggregate(query mongo.Pipeline) (*mongo.Cursor, error) {
	res, err := db.Collection.Aggregate(ctx, query)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (db *mongoCollection) UpdateById(id primitive.ObjectID, data interface{}) (*mongo.SingleResult, error) {
	updateOption := options.FindOneAndUpdate().SetReturnDocument(options.After)
	result := db.Collection.FindOneAndUpdate(
		ctx,
		bson.M{"_id": id},
		bson.M{"$push": data},
		updateOption,
	)

	return result, nil
}

func (db *mongoCollection) UpdateByQuery(query interface{}, data interface{}) (*mongo.UpdateResult, error) {
	result, err := db.Collection.UpdateOne(
		ctx,
		query,
		data,
	)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (db *mongoCollection) DeleteById(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	result, err := db.Collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (db *mongoCollection) DeleteByQuery(query interface{}) (bool, error) {
	_, err := db.Collection.DeleteOne(ctx, query)
	if err != nil {
		return false, err
	}

	return true, nil
}