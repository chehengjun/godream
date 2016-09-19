package postgresql

import (
	"godream/db/global"
	"log"
)

type RDB global.RDB

func (prdb *RDB) Add(args string, reply *string) error {
	log.Println("\033[32m[ADD TEST]\033[0mThis is a add data function of postgresql.")
	*reply = args + ": RDB Add Success."
	return nil
}

func (prdb *RDB) Delete(args string, reply *string) error {
	log.Println("\033[32m[DELETE TEST]\033[0mThis is a delete data function of postgresql.")
	*reply = args + ": RDB Delete Success."
	return nil
}

func (prdb *RDB) Modify(args string, reply *string) error {
	log.Println("\033[32m[MODIFY TEST]\033[0mThis is a modify data function of postgresql.")
	*reply = args + ": RDB Modify Success."
	return nil
}

func (prdb *RDB) Query(args string, reply *string) error {
	log.Println("\033[32m[Query TEST]\033[0mThis is a modify data function of postgresql.")
	*reply = args + ": RDB Query Success."
	return nil
}
