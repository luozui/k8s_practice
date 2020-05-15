package models

import (
	"time"
)

type (
	User struct {
		Id            int32     `db:"id"`  // id
		Uid           int64     `db:"uid"` // id
		Pass          string    `db:"pass"`
		Name          string    `db:"name"`
		Email         string    `db:"email"` // email
		About         string    `db:"about"`
		CoverImageUrl string    `db:"cover_image_url"`
		Extra         string    `db:"extra"` // extra
		Status        int8      `db:"status"`
		CreateTime    time.Time `db:"create_time"`
		UpdateTime    time.Time `db:"update_time"`
	}
)
