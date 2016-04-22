package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Users_20160422_165951 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Users_20160422_165951{}
	m.Created = "20160422_165951"
	migration.Register("Users_20160422_165951", m)
}

// Run the migrations
func (m *Users_20160422_165951) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE users(`id` int(11) NOT NULL AUTO_INCREMENT,`username` varchar(128) NOT NULL,`password` varchar(128) NOT NULL,`email` varchar(128) NOT NULL,`phone` varchar(128) NOT NULL,`created_at` datetime NOT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *Users_20160422_165951) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `users`")
}
