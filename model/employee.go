package model

type Employee struct {
    ID       int    `bson:"_id"`
    Name     string `bson:"name"`
    Position string `bson:"position"`
    Salary   int    `bson:"salary"`
}