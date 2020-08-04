package storage

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type Mgo struct {
	Uri        string //数据库网络地址
	Database   string //要连接的数据库
	Collection string //要连接的集合
}

type MongoTemplate struct {
	collection *mongo.Collection
}

func (m *Mgo) Connect() (*mongo.Collection, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.Uri).SetMaxPoolSize(20))
	if err != nil {
		log.Print(err)
	}
	collection := client.Database(m.Database).Collection(m.Collection)
	return collection, nil
}

/**
 * 插入数据
 */
func (t *MongoTemplate) Insert(document interface{}) (interface{}, error) {
	insertResult, err := t.collection.InsertOne(context.TODO(), document)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return insertResult.InsertedID, nil
}

/**
 * 查询数据
 */
func (t *MongoTemplate) Query(filter bson.D) (interface{}, error) {
	var result interface{}
	err := t.collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	fmt.Printf("Found a single document: %+v\n", result)
	return result, nil
}

/**
 * 更新数据
 */
func (t *MongoTemplate) Update(filter bson.D, update bson.D) (interface{}, error) {
	updateResult, err := t.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	return updateResult.ModifiedCount, nil
}

/**
 * 删除数据
 */
func (t *MongoTemplate) Delete(filter bson.D) (interface{}, error) {
	deleteResult, err := t.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
	return deleteResult, nil
}
