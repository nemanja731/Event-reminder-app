// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"time"
)

type Event struct {
	ID        int64     `json:"id"`
	IDUser    int64     `json:"id_user"`
	Title     string    `json:"title"`
	EventTime time.Time `json:"event_time"`
}

type Session struct {
	ID           string    `json:"id"`
	IDUser       int64     `json:"id_user"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Password string `json:"password"`
}