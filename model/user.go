package model

type User struct {
	Id        string `bson:"_id,omitempty"`
	FirstName string `bson:"firstName,omitempty"`
	Code      string `bson:"Code"`
	LastName  string `bson:"lastName,omitempty"`
	BirthDate string `bson:"BirthDate,omitempty"`
	UserName  string `bson:"UserName,omitempty"`
	Password  string `bson:"Password,omitempty"`
}
