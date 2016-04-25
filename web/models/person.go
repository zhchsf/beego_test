
package models

type Person struct {
  Username string `bson:"username"`
  Email string `bson:"email"`
}