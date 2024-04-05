package db

func InitDatabases() {
	InitMysql()
	InitRedis()
}

func Close() {
	db, _ := Conn.DB()
	_ = Rdb.Close()
	_ = db.Close()

}
