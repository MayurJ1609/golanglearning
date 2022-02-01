package controller

const connectionString = "mongodb://localhost:27017"
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection
