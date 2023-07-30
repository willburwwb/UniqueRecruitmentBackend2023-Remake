package global

func setup() {
	setupPgsql()
	//setupRedis()
	setupSSO()
	setupCOS()
	setupSess()
}

func init() {
	setup()
}
