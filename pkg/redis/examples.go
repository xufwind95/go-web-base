package redis

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func redisCRUD() {

	conn := RedisPool.Get()
	defer conn.Close()

	// "HMSET myhash field1 "Hello" field2 "World""
	exists, err := redis.Bool(conn.Do("EXISTS", "user:1"))
	if err != nil {
		return
	}
	fmt.Println(exists)
	if !exists {
		success, err := conn.Do("HMSET", "user:1", "field1", "1", "field2", "2")
		if err != nil {
			fmt.Println("some error happened")
		}
		fmt.Println("=================", success)
	} else {
		// type User struct {
		// 	field1 string `redis:"field1"`
		// 	field2 string `redis:"field2"`
		// }

		// 获取map中的单个值
		data, err := redis.String(conn.Do("HGET", "user:1", "field1"))
		if err != nil {
			fmt.Println("get field1 failed")
			return
		}
		fmt.Println("get data is:", data)

		// 使用stringMap的形式获取hash中的所有值
		res, err := redis.StringMap(conn.Do("HGETALL", "user:1"))
		if err != nil {
			fmt.Println("get string map failed:", err)
		} else {
			fmt.Println("----------", res, "----------")
		}

		// 使用struct的形式直接获取hash中的所有值
		var user struct {
			Field1 string `redis:"field1"`
			Field2 string `redis:"field2"`
		}
		values, err := redis.Values(conn.Do("HGETALL", "user:1"))
		if err != nil {
			fmt.Println("get hash data failed:", err)
		} else {
			fmt.Println("data:", values)
			err := redis.ScanStruct(values, &user)
			if err != nil {
				fmt.Println("get map data failed:", err)
			} else {
				fmt.Println("get hash data success:", user)
			}
		}

		// 重新设置map中的值
		res_, err := conn.Do("HMSET", "user:1", "field1", "1", "field2", "2")
		if err != nil {
			fmt.Println("some error happened")
		}
		fmt.Println("set new data success=================", res_)

		// 直接修改map中的值(在原来值的基础上增加或减少)
		conn.Send("MULTI")
		conn.Send("HINCRBY", "user:1", "field1", "2")
		conn.Send("HINCRBY", "user:1", "field2", "3")
		conn.Send("HINCRBY", "user:1", "field3", "3")
		r, err := conn.Do("EXEC")
		if err != nil {
			fmt.Println("execute user update failed!")
		} else {
			fmt.Println(r)
		}

		values2, err := redis.Values(conn.Do("HGETALL", "user:1"))
		if err != nil {
			fmt.Println("get hash data failed:", err)
		} else {
			fmt.Println("data:", values2)
			err := redis.ScanStruct(values2, &user)
			if err != nil {
				fmt.Println("get map data failed:", err)
			} else {
				fmt.Println("get hash data success:", user)
			}
		}
	}
}
