package model

import (
	"encoding/json"
	"errors"

	"github.com/garyburd/redigo/redis"
)

// 服务器启动后，就初始化一个userDao
// 全局的，当需要和redis操作时，就调用它
var (
	MyUserDao *UserDao
)

type UserDao struct {
	pool *redis.Pool
}

// 使用工厂模式，创建UserDao实例
func NewUserDao(pool *redis.Pool) *UserDao {
	return &UserDao{pool: pool}
}

// 1. 根据用户ID，返回实例
func (ud *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	// 通过给定ID，去redis查询
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		if err == redis.ErrNil {
			// 表示users哈希中，没有找到对应的id
			err = ERROR_USER_NOTEXISTS
		}
		return nil, err
	}

	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		return nil, errors.New("json.Unmarshal err")
	}
	return user, err
}

// 2. 登录验证
func (ud *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	conn := ud.pool.Get()
	defer conn.Close()
	user, err = ud.getUserById(conn, userId)
	if err != nil {
		return nil, err
	}

	if user.UserPwd != userPwd {
		err = ERROR_USER_PWDINVALID
		return nil, err
	}
	return user, err

}
