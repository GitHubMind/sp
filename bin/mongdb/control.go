package mongoContorl

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"sync"
	"time"
)

type Mongodb struct {
	client     *mongo.Client     //链接句柄
	name       string            // 数据库名
	maxTime    int64             // 链接超时时间
	num        uint64            // 链接数
	Db         *mongo.Database   // database 话柄
	collection *mongo.Collection // collection 话柄
	ctx        context.Context
}

//如果是多链接，要封装成mapÒ
var countGuard sync.Mutex
var mongodb *Mongodb

func connect(collection string) (*Mongodb, error) {
	var err error
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://admin:mindMonGoDBmima@45.253.64.106:9998/?authSource=admin"))
	//client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://tong:mind123@localhost:27017/?authSource=admin"))
	if err != nil {
		log.Println("ping", err)

	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	//关闭链接
	if err != nil {
		log.Println("Connect", err)
	}
	//检查错误
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Println("Ping", err)
	}
	Db := client.Database(collection)
	//初始化
	//mongodb.name = "text"
	//mongodb.maxTime = 5000
	//mongodb.num = 10000000000
	//
	mongodb = &Mongodb{client: client, name: collection, maxTime: 100000, num: 100, Db: Db, collection: nil, ctx: ctx}
	return mongodb, err
}

func (r *Mongodb) Insert(value interface{}, collectionName string) (string, error) {

	collection := r.Db.Collection(collectionName)
	insertOneResult, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		log.Println("Insert:", err)
	}

	if err != nil {
		//log.Println(err)
	}
	return fmt.Sprint(insertOneResult.InsertedID.(primitive.ObjectID).Hex()), err
}
func (r *Mongodb) InsertOne(value interface{}, collectionName string) (string, error) {
	r.collection = r.Db.Collection(collectionName)
	insertOneResult, err := r.collection.InsertOne(context.Background(), value)
	if err != nil {
		log.Println("InsertOne:", err)

	}
	return fmt.Sprint(insertOneResult.InsertedID.(primitive.ObjectID).Hex()), err
}
func (r *Mongodb) FindAll(collectionName string) ([]interface{}, error) {
	r.collection = r.Db.Collection(collectionName)

	filter := bson.M{}
	singleResult, err := r.collection.Find(context.Background(), filter)
	// 遍历游标允许我们一次解码一个文档
	var arr []interface{}
	for singleResult.Next(context.TODO()) {
		// 创建一个值，将单个文档解码为该值
		var elem interface{}
		err := singleResult.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		arr = append(arr, elem)
	}
	return arr, err
}

func (r *Mongodb) FindMul(collectionName, key, value string, limit, index int64) (count int64, arr []interface{}, err error) {
	//设置过期时间

	var wg sync.WaitGroup
	wg.Add(2)

	ctx, cannel := context.WithTimeout(context.Background(), time.Minute)
	defer cannel()
	var findoptions *options.FindOptions

	r.collection = r.Db.Collection(collectionName)
	filter := bson.M{key: value}
	go func() {
		count, err = r.collection.CountDocuments(ctx, filter)
	}()
	if err != nil {
		return
	}
	if limit > 0 && count > limit {
		findoptions = &options.FindOptions{}
		findoptions.SetLimit(limit)
		findoptions.SetSkip(limit * index)
	}

	singleResult, err := r.collection.Find(ctx, filter, findoptions)

	for singleResult.Next(context.TODO()) {
		// 创建一个值，将单个文档解码为该值é
		var elem interface{}
		err = singleResult.Decode(&elem)
		if err != nil {
			return
		}
		arr = append(arr, elem)
	}
	return
}

func (r *Mongodb) FindDRY(collectionName string, filter bson.D, limit, index int64) (count int64, arr []interface{}, err error) {
	//设置过期时间
	ctx, cannel := context.WithTimeout(context.Background(), time.Minute)
	defer cannel()
	var findoptions *options.FindOptions

	r.collection = r.Db.Collection(collectionName)
	if count, err = r.collection.CountDocuments(ctx, filter); err != nil {
		return
	}
	if limit > 0 && count > limit {
		findoptions = &options.FindOptions{}
		findoptions.SetLimit(limit)
		findoptions.SetSkip(limit * index)
	}
	singleResult, err := r.collection.Find(ctx, filter, findoptions)
	for singleResult.Next(context.TODO()) {
		// 创建一个值，将单个文档解码为该值
		var elem interface{}
		err = singleResult.Decode(&elem)
		if err != nil {
			return
		}
		arr = append(arr, elem)
	}
	return
}
func (r *Mongodb) CollectionCount(collectionName string) (string, int64) {
	r.collection = r.Db.Collection(collectionName)
	name := r.collection.Name()
	size, _ := r.collection.EstimatedDocumentCount(context.Background())
	return name, size
}
func GetSingleInstanceMongoDB(collection string) (*Mongodb, error) {

	var err error

	if err != nil {
		log.Println("Ping", err)
	}

	if mongodb == nil {
		mongodb, err = connect(collection)
	}
	//保证n能ping的通
	if err != nil {
		ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
		err = mongodb.client.Ping(ctx, readpref.Primary())
		for err != nil {
			countGuard.Lock()
			mongodb, err = connect(collection)
			countGuard.Unlock()
		}
	}
	return mongodb, err
}

func (r *Mongodb) Close() {
	r.client.Disconnect(r.ctx)
}
