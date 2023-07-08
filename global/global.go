package global

func setup() {
	setupPgsql()
	//setupRedis()
	setupSSO()
	setupCOS()
}

func init() {
	setup()
}
