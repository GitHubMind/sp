package mongoContorl

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"reflect"
	"time"
)

var (
	client     *mongo.Client                        //链接句柄
	name                         = "test"           // 数据库名
	maxTime                      = time.Duration(2) // 链接超时时间
	num        uint64            = 50               // 链接数
	table                        = "test"           // 表名
	db         *mongo.Database                      // database 话柄
	collection *mongo.Collection                    // collection 话柄
)

// You will be using this Trainer type later in the program
type Trainer struct {
	Name string
	Age  int
	City string
}

func main() {
	//config()
	//insert()
	initConnect()
}
func initConnect() {
	err := fmt.Errorf("test")
	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://tong:mind123@localhost:27017/?authSource=admin"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	//关闭链接
	defer client.Disconnect(ctx)
	//检查错误
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connect is successful")
	database, err := client.ListDatabaseNames(ctx, bson.M{})
	fmt.Println(database)
	//删除
	collection = client.Database("text").Collection("trainers")
	insertOne()
}
func findOne() {
	var one Trainer
	err := collection.FindOne(context.Background(), bson.M{"city": "Pallet Town"}).Decode(&one)
	if err != nil {

	}
	log.Println(one)
}
func findMutl() {
	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Println(err)
	}
	if err = cur.Err(); err != nil {

	}
	var many []*Trainer
	err = cur.All(context.Background(), &many)
	if err != nil {
		log.Println(err)
	}
	cur.Close(context.Background())
	for _, value := range many {
		log.Println(&value)
	}

}

func DocumentSum() {
	count, err := collection.CountDocuments(context.Background(), bson.D{})
	if err != nil {

	}
	log.Println("Document.len:", count)
}
func clearTable(data, collections string) {
	collection = client.Database(data).Collection(collections)
	err := collection.Drop(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
func insertOne() interface{} {
	ash := Trainer{"Ash", 10, "Pallet Town"}
	insertOneResult, err := collection.InsertOne(context.Background(), ash)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(reflect.TypeOf(insertOneResult.InsertedID))
	//log.Println(string(insertOneResult.InsertedID))
	test := make(map[string]interface{})
	test["test"] = insertOneResult.InsertedID
	insertOneResult, err = collection.InsertOne(context.Background(), test)
	if err != nil {
		log.Fatal(err)
	}
	return test["text"]
}
func config() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://tong:mind123@localhost:27017/?authSource=admin")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	db = client.Database("test")
}
func insert() {
	//collection := client.Database("test").Collection("test")
	//log.Println(collection)
	ash := Trainer{"Ash", 10, "Pallet Town"}
	misty := Trainer{"Misty", 10, "Cerulean City"}
	brock := Trainer{"Brock", 15, "Pewter City"}
	insertResult, err := collection.InsertOne(context.TODO(), ash)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	trainers := []interface{}{misty, brock}

	insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}

//func connect(cName string) (*mgo.Session, *mgo.Collection) {
//	//session, err := mgo.Dial("$mongoHost") //Mongodb's connection
//	//if err != nil {
//	//	panic(err)
//	//}
//	//session.SetMode(mgo.Monotonic, true)
//	////return a instantiated collect
//	//return session, session.DB(mongoDB).C(cName)
//}
