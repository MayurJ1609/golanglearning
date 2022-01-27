package controller

// import ("go.mongodb.org/mongo-dirver/mongo")

const connectionString = "mongodb://localhost:27017"
const dbName = "netflix"
const colName = "watchlist"

// MOST IMPORTANT
var collection *mongo.collection

// Connect with mongoDB
func init() {
	// client option
	clientOption := options.Client()
}
