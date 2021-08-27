package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	//"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	//ExampleClient()
	//Lua_script()
	Remove_keys_by_lua_script()
}

///
func Remove_keys_by_lua_script() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.52.128:46379",
		Password: "sparksubmit666", // no password set
		DB:       2,                // use default DB
	})

	Remove_Keys := redis.NewScript(`
		local ks = redis.call('KEYS', KEYS[1])     	 
		for i=1,#ks,5000 do             
			redis.call('del', unpack(ks, i, math.min(i+4999, #ks)) )       
		end 
		return true 
	`)

	//key_name:= "monitor:atsinspect:0083f73c-cb1d-428b-b788-7ff7008d8d96:Porccess_NeedToCalcCount"
	n, err := Remove_Keys.Run(ctx, rdb, []string{"monitor:atsinspect*"}).Result()
	fmt.Println(n, err)

}

/// 示例： 模糊删除指定的keys
func ExampleClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.52.128:46379",
		Password: "sparksubmit666", // no password set
		DB:       2,                // use default DB
	})

	//opt, err := redis.ParseURL("redis://<user>:<pass>@localhost:6379/<db>")
	//if err != nil {
	//	panic(err)
	//}
	//

	key_name := "monitor:atsinspect:0083f73c-cb1d-428b-b788-7ff7008d8d96:Porccess_NeedToCalcCount"
	err := rdb.Set(ctx, key_name, "68823", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, key_name).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(key_name, val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}

func Lua_script() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.52.128:46379",
		Password: "sparksubmit666", // no password set
		DB:       2,                // use default DB
	})

	//key_name:= "monitor:atsinspect:0083f73c-cb1d-428b-b788-7ff7008d8d96:Porccess_NeedToCalcCount"

	IncrByXX := redis.NewScript(`
		if redis.call("GET", KEYS[1]) ~= false then
			return redis.call("INCRBY", KEYS[1], ARGV[1])
		end
		return false
	`)

	n, err := IncrByXX.Run(ctx, rdb, []string{"xx_counter"}, 2).Result()
	fmt.Println(n, err)

	err = rdb.Set(ctx, "xx_counter", "40", 0).Err()
	if err != nil {
		panic(err)
	}

	n, err = IncrByXX.Run(ctx, rdb, []string{"xx_counter"}, 2).Result()
	fmt.Println(n, err)

}
