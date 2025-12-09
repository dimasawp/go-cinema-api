package models

import "time"

type User struct {
    ID           int64     `db:"id" json:"id"`
    FullName     string    `db:"full_name" json:"full_name"`
    Email        string    `db:"email" json:"email"`
    PasswordHash string    `db:"password_hash" json:"-"`
    CreatedAt    time.Time `db:"created_at" json:"created_at"`
}

type Showtime struct {
    ID           int64     `db:"id" json:"id"`
    MovieID      int64     `db:"movie_id" json:"movie_id"`
    AuditoriumID int64     `db:"auditorium_id" json:"auditorium_id"`
    StartTime    time.Time `db:"start_time" json:"start_time"`
    EndTime      time.Time `db:"end_time" json:"end_time"`
    Status       string    `db:"status" json:"status"`
}