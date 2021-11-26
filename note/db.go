package note

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

//101 redis数据库
//redis支持的数据类型: String/Hash(类似于golang的map[string]string)/List/Set(无序且不能重复)/zset是有序的(无序：set)

//101.1 常见key value操作
//101.1.1 添加和修改key value：set key value
//101.1.2 添加临时的key value：setex key seconds value
//101.1.3 批量添加修改key value：mset key1 value1 key2 value2 ...
//101.1.4 查看所有的key：keys *
//101.1.5 查看key对应的值：get key
//101.1.6 批量获取key的值：mget key1 key2 ...
//101.1.7 切换数据库（0~15 默认为0）：select index
//101.1.8 查看当前数据库key value数量：dbsize
//101.1.9 清空当前数据库所有key value：flushdb
//101.1.10 清空所有数据库所有key value：flushall
//101.1.11 删除指定的key value：del key

//101.2 hash 操作
//101.2.1 添加修改hash：hset key field value
//101.2.2 批量添加修改hash：hmset key field1 value1 field2 value2 ...
//101.2.3 查看key field对应的值：hget key field
//101.2.4 批量查看key的多个field的值：hmget key field1 field2 ...
//101.2.5 获取key对应的所有的field value: hgetall key
//101.2.6 删除对应的field和value：hdel key field1 field2 ...
//101.2.7 查看hash长度：hlen key
//101.2.8 查看字段是否存在：hexists key field

//101.3 list
//101.3.1 从左边插入元素：lpush key value2 value1 value0 ...
//101.3.2 从右边插入元素：rpush key value0 value1 value2 ...
//101.3.3 获取对应的元素：lrange key startIndex stopIndex (index可以为负数：-1为倒数第一个元素，-2为倒数第二个元素...)
//101.3.4 从左边推出元素（如果全部推出，则对应的key也会被删除）：lpop key
//101.3.5 从右边推出元素：rpop key
//101.3.6 统计list长度：llen key

//101.4 set
//101.4.1 添加元素：sadd key member1 member2 member0 ...
//101.4.2 获取所有的元素：smembers key
//101.4.3 判断元素是否存在：sismember key member
//101.4.4 删除元素：srem key member member ...

//101.5 通过redigo 操作redis
func RediGo() {
	//101.5.1 拨号连接
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis连接失败！", err)
		return
	}
	//101.5.2 关闭连接
	defer func() {
		err := conn.Close()
		if err != nil {
			fmt.Println("关闭与redis的连接失败！", err)
		}
	}()
	fmt.Println("redis连接成功！", conn)

	//101.5.3 操作redis
	conn.Do("set", "name", "fangpig")
	str, err := redis.String(conn.Do("get", "name"))
	if err != nil {
		fmt.Println("操作redis失败！", err)
		return
	}
	fmt.Println("redis返回：", str)
	conn.Do("hmset", "user1", "name", "fangpig", "age", "60")
	strs, err := redis.Strings(conn.Do("hgetall", "user1"))
	if err != nil {
		fmt.Println("操作redis失败！", err)
		return
	}
	for i, v := range strs {
		if i%2 == 0 {
			fmt.Printf("user1[\"%v\"]:", v)
		} else {
			fmt.Printf("\"%v\"\n", v)
		}
	}

}

//101.5.4 redis连接池（服务器端并发性能优化）
var redisPool *redis.Pool

func RedisPoolInit() { //需要在项目的init(初始化的意思）中调用
	redisPool = &redis.Pool{
		MaxIdle:     8,  //最大空闲连接数
		MaxActive:   0,  //最大连接数，0是没有限制
		IdleTimeout: 30, //最大空闲时间（秒）
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
}
func RedisPoolTest() { //通常在协程中使用
	//从redis中取出连接
	conn := redisPool.Get()
	defer func() {
		err := conn.Close()
		if err != nil {
			fmt.Println("关闭与redis的连接失败！", err)
		}
	}()
	fmt.Println("redis连接成功！", conn)
	conn.Do("hmset", "user1", "name", "fangpig", "age", "60")
	strs, err := redis.Strings(conn.Do("hgetall", "user1"))
	if err != nil {
		fmt.Println("操作redis失败！", err)
		return
	}
	for i, v := range strs {
		if i%2 == 0 {
			fmt.Printf("user1[\"%v\"]:", v)
		} else {
			fmt.Printf("\"%v\"\n", v)
		}
	}
}

//102 LevelDb google开发 c++keyValue嵌入式数据库 这里用第三方实现
func LevelDb() {
	//102.1 打开或新建，关闭数据库
	db, err := leveldb.OpenFile("pathToDb", nil)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	//102.2 常见操作
	//102.2.1 存入或修改数据
	err = db.Put([]byte("key1"), []byte("value1"), nil)
	if err != nil {
		fmt.Println("put失败 err=", err)
		return
	}
	//102.2.2 取出数据
	v, err := db.Get([]byte("key1"), nil) //！！！这边取出来的v不可以修改，修改的话复制一份再修改
	if err != nil {
		fmt.Println("get失败 err=", err)
		return
	}
	fmt.Println(string(v))
	//102.2.3 删除数据
	err = db.Delete([]byte("key1"), nil)
	if err != nil {
		fmt.Println("Delete失败 err=", err)
		return
	}
	err = db.Put([]byte("key0"), []byte("value0"), nil)
	if err != nil {
		fmt.Println("put失败 err=", err)
		return
	}
	//102.2.4 批量操作(短时间内操作大量数据性能更加)
	batch := new(leveldb.Batch)
	batch.Put([]byte("key1"), []byte("value1"))
	batch.Put([]byte("key_2"), []byte("value2"))
	batch.Put([]byte("key_3"), []byte("value3"))
	batch.Put([]byte("key_4"), []byte("value4"))
	batch.Put([]byte("key5"), []byte("value5"))
	batch.Delete([]byte("key0"))
	err = db.Write(batch, nil)
	if err != nil {
		fmt.Println("Write失败 err=", err)
		return
	}
	//102.2.5 判断数据是否存在
	ret, err := db.Has([]byte("key1"), nil)
	if err != nil {
		fmt.Println("Has失败 err=", err)
		return
	}
	if ret {
		fmt.Println("存在")
	} else {
		fmt.Println("不存在")
	}

	//102.3 完整遍历数据库
	fmt.Println("\n102.3 完整遍历数据库")
	//102.3.1 新建遍历器
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		//获取值，不能修改！
		fmt.Printf("[%v]=%v\n", string(iter.Key()), string(iter.Value()))
	}
	//102.3.2 释放遍历器
	iter.Release()
	err = iter.Error()
	if err != nil {
		fmt.Println("遍历失败 err=", err)
		return
	}

	//102.4 遍历部分数据库
	fmt.Println("\n//102.4 遍历部分数据库")
	iter = db.NewIterator(&util.Range{Start: []byte("key1"), Limit: []byte("key7")}, nil) //[key1:key5)
	for iter.Next() {
		fmt.Printf("[%v]=%v\n", string(iter.Key()), string(iter.Value()))
	}
	iter.Release()
	err = iter.Error()
	if err != nil {
		fmt.Println("遍历失败 err=", err)
		return
	}

	//102.5 遍历特定前缀的数据库
	fmt.Println("\n//102.5 遍历特定前缀的数据库")
	iter = db.NewIterator(util.BytesPrefix([]byte("key_")), nil)
	for iter.Next() {
		fmt.Printf("[%v]=%v\n", string(iter.Key()), string(iter.Value()))
	}
	iter.Release()
	err = iter.Error()
	if err != nil {
		fmt.Println("遍历失败 err=", err)
		return
	}
}
