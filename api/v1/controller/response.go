package controller

import "tier-up/internal/app/model"

type Response[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type LoginResponse struct {
	AccessToken string     `json:"access_token"`
	User        model.User `json:"user"`
}
