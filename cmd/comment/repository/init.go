package repository

import "context"

func Init() {
	go ConsumeComments(context.Background())
	go ConsumeCommentsCache(context.Background())
}
