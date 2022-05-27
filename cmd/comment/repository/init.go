package repository

func Init() {
	go ConsumeCreateComment()
	go ConsumeDeleteComment()
}
