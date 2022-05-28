package repository

func Init() {
	go ConsumeComment()
	go ConsumeCommentsCache()
}
