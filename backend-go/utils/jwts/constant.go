package jwts

import "time"

const (
	/*
	   Token 验证结果
	   -1:验证失败
	   0:普通用户
	   1:管理员
	*/
	RoleError = -1
	RoleUser  = 0
	RoleAdmin = 1

	/*
		Token存储在redis的key
	*/
	TokenPrefix = "login:token:"

	/*
		Token存储的有效期: 1 hour
	*/
	TokenTimeout = time.Hour
)
