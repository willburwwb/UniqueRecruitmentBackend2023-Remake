package models

import "google.golang.org/protobuf/types/known/timestamppb"

type User struct {
	Uid         string
	Phone       string
	Email       string
	Name        string
	JoinTime    *timestamppb.Timestamp
	AvatarUrl   string
	Gender      string
	Groups      []string
	LarkUnionId string
}
