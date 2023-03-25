package database

import (
	"context"
	"log"
	


	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	
)

var DB *mongo.Collection
var ctx = context.TODO();

func Setup() *mongo.Collection{
	   clientOptions := options.Client().ApplyURI("mongodb+srv://Kd_singh:kuldeep@cluster0.yjd3k8n.mongodb.net/?retryWrites=true&w=majority");

	   client,err := mongo.Connect(ctx,clientOptions);

	   if err!= nil{
            log.Fatal(err);
	   }

	   err = client.Ping(ctx,nil)
	   if err!= nil{
		log.Fatal(err);
	   }

	   DB = client.Database("Auth").Collection("credentials");
       log.Print("DB connected");
	   return DB;
}