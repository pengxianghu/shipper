// consignment-service/datastore.go
package main

import (
	"gopkg.in/mgo.v2"
)

// CreateSession() 创建了连接到 mongodb 的主 session
func CreateSession(host string) (*mgo.Session, error) {
	
	session, err := mgo.Dial(host)
	if err != nil {
		return nil, err
	}

	session.SetMode(mgo.Monotonic, true)

	return session, nil
}
