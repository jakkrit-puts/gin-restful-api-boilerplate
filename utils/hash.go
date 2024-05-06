package utils

import "github.com/matthewhartstonge/argon2"

func HashPassword(password string) string {
	argon := argon2.DefaultConfig()
	encoded, _ := argon.HashEncoded([]byte(password))

	return string(encoded)
}
