package global

func setup() {
	setupPgsql()
	setupCOS()
	setupSess()
	setupRedis()
}

func init() {
	setup()
}
