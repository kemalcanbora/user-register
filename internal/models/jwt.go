package models

type JwtUserAuth struct {
	Authorized bool   `json:"authorized"`
	Email      string `json:"email"`
	Exp        int64  `json:"exp"`
}
