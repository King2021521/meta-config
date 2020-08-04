package storage

import (
	"configservice"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type mgo struct {
	uri        string //数据库网络地址
	database   string //要连接的数据库
	collection string //要连接的集合
}

func (m *mgo) Connect() (*mongo.Collection, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.uri).SetMaxPoolSize(20))
	if err != nil {
		log.Print(err)
	}
	collection := client.Database(m.database).Collection(m.collection)
	return collection, nil
}

/**
 * 插入数据
 */
func (m *mgo) Insert(collection *mongo.Collection, user configservice.Userspace) (interface{}, error) {
	insertResult, err := collection.InsertOne(context.TODO(), user)
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
func (m *mgo) Query(collection *mongo.Collection, filter bson.D) (interface{}, error) {
	var result configservice.Userspace
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
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
func (m *mgo) Update(collection *mongo.Collection, filter bson.D) (interface{}, error) {
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
	return deleteResult, nil
}
