package utils

import (
	"errors"
	"log"
	"os"
	"strconv"
)

type env struct {
	Port int
}

func logEnv(name string, env interface{}) {
	log.Printf("env found: %v=%v\n", name, env)
}

func getPort() int {
	env, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		Fatal(errors.New("env \"PORT\":  has to be valid integer tcp port"))
	}
	logEnv("PORT", env)
	return env
}

func Env() env {
	return env{Port: getPort()}
}
