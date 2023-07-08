package global

func setup() {
	setupPgsql()
	//setupRedis()
	setupSSO()
}

func init() {
	setup()
}
