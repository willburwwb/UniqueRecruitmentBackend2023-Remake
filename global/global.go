package global

func setup() {
	setupPgsql()
	setupCOS()
	setupSess()
}

func init() {
	setup()
}
