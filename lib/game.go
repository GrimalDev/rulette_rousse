package game

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/gorilla/sessions"
)

type game struct {
	Code string
}

func RandomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}

func CreateGame(sess *sessions.Session) game {
	game := game{Code: ""}
	// game.Code, _ = randomHex(3)
	game.Code = sess.Values["code"].(string)

	return game
}
