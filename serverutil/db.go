package serverutil

//DB type of the database to use
type DB map[string][]byte

//CreateDatabase generates an empty database
func CreateDatabase() DB {
	var database = DB{}
	return database
}

//AddElement inserts a new element in the database
func AddElement(database DB, id string, value []byte) {
	database[id] = value
}

//GetElement gets the value for the provided key
func GetElement(database DB, id string) []byte {
	return database[id]
}

//RemoveElement eleminates an element in the database (Not required)
func RemoveElement(database DB, id string) {
	delete(database, id)
}
