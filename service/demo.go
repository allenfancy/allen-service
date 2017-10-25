package main

import (
	"bytes"
	"container/ring"
	"crypto/aes"
	"crypto/cipher"
	"database/sql"
	"encoding/base64"
	"fmt"
	"runtime"
	"time"

	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	zk "github.com/samuel/go-zookeeper/zk"
)

const sql1 = ""

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

		//
		//zkOp()
		//redisOp()
		who := "World"
		if len(os.Args) > 1 {
			who = strings.Join(os.Args[1:], ",")
		}
		fmt.Println(who)
		testAes()

		// md5
		h := md5.New()
		io.WriteString(h, "I am the X")
		io.WriteString(h, "I am the Y")
		fmt.Printf("%x", h.Sum(nil))
		fmt.Println()
		data := []byte("These pretzels are making me thirsty.")
		fmt.Printf("%x", md5.Sum(data))
		fmt.Println()
		testContainerRing()
		mysql()
		fmt.Println(" .... ")
		go fun1()
		// 新建一个Channel
		channel := make(chan string, 10)
		go funChann(channel)
		//读取Channel
		fmt.Println(<-channel)

		// 多个Channel
		channel1 := make(chan string)
		channel2 := make(chan string)
		go funcMutil(channel1)
		go funcMutil(channel2)
		fmt.Println("channel one sends ", <-channel1)
		fmt.Println("channel two sends ", <-channel2)

	channel3 := make(chan string)
	channel4 := make(chan string)
	go funcMutil(channel3)
	go funcMutil(channel4)
	count := 0
	for count < 3 {
		select {
		case v := <-channel3:
			fmt.Println(v)
		case v2 := <-channel4:
			fmt.Println(v2)
		default:
			fmt.Println("没有channl准备好!")
		}
		time.Sleep(time.Second)
		count++
	}
	time.Sleep(1 * time.Hour)

	a := 1
	mas := make(map[int]string)
	fmt.Println(&a)
	mas[1] = "allen"
	fmt.Println(&mas)
	**/
	type demo struct {
		info string
	}

	dd := &demo{}
	var info1 string = ""
	fmt.Println(dd.info)
	fmt.Println(len(info1))
	/**
	info := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	var slat string
	for i := 0; i < 8; i++ {
		slat += info[rand.Intn(len(info))]
	}
	fmt.Println(slat)
	input := []byte("hello wu" + "biliCould")
	e := base64.StdEncoding.EncodeToString(input)
	fmt.Println(e)
	d, _ := base64.StdEncoding.DecodeString(e)
	fmt.Println(string(d))
	xxx()
	t1 := time.NewTimer(time.Second * 1)
	for {
		select {
		case <-t1.C:
			t := time.Now()
			h := t.Hour()
			if h >= 23 || h <= 5 {
				println("5s timer")
				t1.Reset(time.Second * 1)
			} else {
				println("开始睡觉了")
				//time.Sleep(26 * time.Hour)
				t1.Reset(time.Second * 10)
			}

		}
	}
	**/
}

func xxx() {
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
}

func funcMutil(channel chan string) {
	fmt.Println("haha")
	channel <- "boom"
}

func funChann(channel chan string) {
	fmt.Println("hhhh")
	// 往channal中写数据
	channel <- "boom!!!!"
}
func fun1() {
	fmt.Println("demo ...")
	time.Sleep(1 * time.Second)
}

func testContainerRing() {
	ring := ring.New(5)
	for i := 1; i <= 5; i++ {
		ring.Next().Value = i
		ring.Prev().Value = i - 1
	}

	next := ring.Next()
	fmt.Println(next.Value)
	pre := ring.Prev()
	fmt.Println(pre.Value)
}

func testAes() {
	// AES-128。key长度：16, 24, 32 bytes 对应 AES-128, AES-192, AES-256
	key := []byte("sfe023f_9fd&fwfl")
	result, err := AesEncrypt([]byte("allenfancy@163.com"), key)
	if err != nil {
		panic(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(result))
	origData, err := AesDecrypt(result, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(origData))
}

// AesEncrypt aes
func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// AesDecrypt aes
func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

// PKCS5Padding add
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS5UnPadding add
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
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

	db, err := sql.Open("mysql", "root:123@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	for i := 10000; i < 20000; i++ {
		stmt, _ := db.Prepare("INSERT INTO `test` ( `mid`, `timestamp`, `loginip`) VALUES(?, ?, ?);")
		res, _ := stmt.Exec(i, i, i)
		fmt.Println(res)
		time.Sleep(200 * time.Microsecond)
	}

	for i := 1000; i < 10000; i++ {
		stmt, _ := db.Prepare("INSERT INTO `test1` ( `mid`, `timestamp`, `loginip`) VALUES(?, ?, ?);")
		res, _ := stmt.Exec(i, i, i)
		fmt.Println(res)
		time.Sleep(200 * time.Microsecond)
	}

}
