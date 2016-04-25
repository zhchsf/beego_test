
package mongopool

import (
  "gopkg.in/mgo.v2"
)

type MongoPool struct {
  session *mgo.Session
}

var (
  MongoClient *MongoPool
)

func init(){
  session, err := mgo.Dial("127.0.0.1")
  if err != nil {
    panic(err)
  }
  MongoClient = &MongoPool{session}
}

func Get() *mgo.Session {
  return MongoClient.session.Copy()
}