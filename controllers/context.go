package controllers

import (
	"gopkg.in/mgo.v2"

	"github.com/ljcastro/taskmanager/common"
)

// Context struct used for maintaining HTTP Request Context
type Context struct {
	MongoSession *mgo.Session
}

// Close mgo.Session
func (c *Context) Close() {
	c.MongoSession.Close()
}

// DbCollection Returns collection for the given name
func (c *Context) DbCollection(name string) *mgo.Collection {
	return c.MongoSession.DB(common.AppConfig.Database).C(name)
}

// NewContext creates a new context object for each HTTP Request
func NewContext() *Context {
	session := common.GetSession().Copy()
	context := &Context{
		MongoSession: session,
	}
	return context
}
