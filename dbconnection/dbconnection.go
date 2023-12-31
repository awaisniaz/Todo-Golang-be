package dbconnection

import (
    "context"
    "fmt"
    "time"
 
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
)
 
var Connection *mongo.Client = nil 
func Close(client *mongo.Client, ctx context.Context,
	cancel context.CancelFunc){
	 
defer cancel()

defer func(){

 if err := client.Disconnect(ctx); err != nil{
	 panic(err)
 }
}()
}
func Connect(uri string)(*mongo.Client, context.Context,
	context.CancelFunc, error) {
	 
// ctx will be used to set deadline for process, here
// deadline will of 30 seconds.
ctx, cancel := context.WithTimeout(context.Background(),
				 30 * time.Second)

// mongo.Connect return mongo.Client method
client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
Connection = client
return client, ctx, cancel, err
}

func Ping(client *mongo.Client, ctx context.Context) error{
 
    // mongo.Client has Ping to ping mongoDB, deadline of
    // the Ping method will be determined by cxt
    // Ping method return error if any occurred, then
    // the error can be handled.
    if err := client.Ping(ctx, readpref.Primary()); err != nil {
        return err
    }
    fmt.Println("connected successfully")
    return nil
}

func ConnectionClient() *mongo.Client {
    return Connection
}