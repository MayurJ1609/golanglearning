package controller

const connectionString = "mongodb://localhost:27017"
const dbName = "netflix"
const colName = "watchlist"

// MOST IMPORTANT
var collection *mongo.Collection

// Connect with mongoDB
func init() {
	// client option
	clientOption := options.Client()
}
