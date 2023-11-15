package global

func setup() {
	setupPgsql()
	//setupRedis()
	setupCOS()
	setupSess()
}

func init() {
	setup()
}
