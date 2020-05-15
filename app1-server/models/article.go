package models

import (
	"time"
)

type (
	Article struct {
		Id            int32     `db:"id"` // id
		Aid           int64     `db:"aid"`
		Uid           int64     `db:"uid"` // id
		Title         string    `db:"title"`
		Mdtext        string    `db:"mdtext"` // markdown
		Content       string    `db:"content"`
		ViewCnt       int32     `db:"view_cnt"`
		CoverImageUrl string    `db:"cover_image_url"`
		Extra         string    `db:"extra"` // extra
		Status        int8      `db:"status"`
		CreateTime    time.Time `db:"create_time"`
		UpdateTime    time.Time `db:"update_time"`
	}
)
