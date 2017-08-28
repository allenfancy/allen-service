package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	zk "github.com/samuel/go-zookeeper/zk"
)

type person struct {
	Name string
	Age  int32
}

/**
1. & * 的用法
2. new make 用法
*/
func main() {

	/**var pN = new(person)
	fmt.Println(pN)
	var pA = &person{}
	fmt.Println(pA)
	pN = pA
	fmt.Println(pN)
	var pP *person //不复制，就会是nil
	if pP == nil {
		//fmt.Println("is nil", pP)
	} else {
		fmt.Println(pP)
	}
	**/
	//mysql()
	//zkOp()
	redisOp()
}

// zk operator github.com/samuel/go-zookeeper/zk
func zkOp() {
	var (
		path  = "/allen2"
		value = "this is test zk use golang"
		stat  *zk.Stat
		err   error
		str   []string
		//acl   = []zk.ACL{zk.ACL{Perms: zk.PermAll, Scheme: path, ID: "1"}} //ACL 权限控制
	)
	c, _, err := zk.Connect([]string{"127.0.0.1"}, time.Second)

	if err != nil {
		fmt.Println("connect error ", err)
		return
	}
	cStr, err := c.Create(path, []byte("1"), zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	if err != nil {
		fmt.Println("create node error", err)
		return
	}
	str, stat, err = c.Children(cStr)
	if err != nil {
		fmt.Println("get child error ", err)
		return
	}
	fmt.Println(stat)
	stat, err = c.Set(path, []byte(value), stat.Version)
	if err != nil {
		fmt.Println("set error", err)
		return
	}
	str, stat, err = c.Children(cStr)
	if err != nil {
		fmt.Println("get child error ", err)
		return
	}
	fmt.Println(str)
	var by []byte
	by, stat, err = c.Get(cStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(by), stat)
}

// redis operator redigo/redis
func redisOp() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("dail redis error is ", err)
		return
	}
	redisOpSet(conn)
	//redisOpString(conn)
}

func redisOpString(conn redis.Conn) (err error) {
	// set  get operator
	v, err := conn.Do("SET", "name", "red")
	if err != nil {
		fmt.Println("set is error ", err)
		return
	}
	fmt.Println(v)
	v, err = redis.String(conn.Do("GET", "name"))
	if err != nil {
		fmt.Println("get is error", err)
		return
	}
	fmt.Println(v)
	// inrc decr
	v, err = conn.Do("SET", "incrdecr", 1)
	if err != nil {
		fmt.Println("set is error", err)
		return
	}
	// inrc decr
	v, err = conn.Do("incr", "incrdecr")
	if err != nil {
		fmt.Println("incr is error", err)
		return
	}
	fmt.Println(v)
	// decr
	v, err = conn.Do("decr", "incrdecr")
	if err != nil {
		fmt.Println("decr is error", err)
		return
	}
	fmt.Println(v)

	return
}
func redisOpSet(conn redis.Conn) (err error) {
	v, err := conn.Do("sadd", "setkey", "1", "2", "3")
	if err != nil {
		fmt.Println("sadd is error", err)
		return
	}
	fmt.Println(v)
	v, err = conn.Do("smembers", "setkey")
	if err != nil {
		fmt.Println("smembers is error", err)
		return
	}
	fmt.Println(v)
	v, err = conn.Do("scard", "setkey")
	if err != nil {
		fmt.Println("scard is error", err)
		return
	}
	fmt.Println(v)
	return
}

// kafka operator  github.com/wvanbergen/kafka/consumergroup
func kafkaOp() {

}

// HBase operator    gohbase/hbase
func hbaseOp() {

}

// MC operator
func mcOp() {

}

func mysql() {
	db, err := sql.Open("mysql", "root:123@tcp(127.0.0.1:3306)/aso?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	fmt.Println("db is ", db)
	mid := 0
	uname := ""
	rows, _ := db.Query("select mid,uname from aso_account")
	for rows.Next() {
		rows.Scan(&mid, &uname)
		fmt.Println(mid, uname)
	}

}
