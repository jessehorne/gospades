package game

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"go.deanishe.net/env"
)

type Config struct {
	ServerName  string `validate:"required"`
	MaxPlayers  uint8  `validate:"required,gt=0,lt=255"`
	Team1Name   string `validate:"required,gt=0,lt=11"`
	Team2Name   string `validate:"required,gt=0,lt=11"`
	DamageTorso uint8  `validate:"required,gte=0,lte=100"`
	DamageHead  uint8  `validate:"required,gte=0,lte=100"`
	DamageArms  uint8  `validate:"required,gte=0,lte=100"`
	DamageLegs  uint8  `validate:"required,gte=0,lte=100"`
	DamageMelee uint8  `validate:"required,gte=0,lte=100"`
}

func GetConfigFromEnv() (Config, error) {
	var newConfig Config

	// load .env
	err := godotenv.Load()
	if err != nil {
		return newConfig, errors.New("error reading .env")
	}

	newConfig = Config{
		ServerName:  env.Get("SERVER_NAME", "gospades server"),
		MaxPlayers:  uint8(env.GetInt("MAX_PLAYERS", 64)),
		Team1Name:   env.Get("TEAM_1_NAME", "Red"),
		Team2Name:   env.Get("TEAM_1_NAME", "Blue"),
		DamageTorso: uint8(env.GetInt("DMG_TORSO", 35)),
		DamageHead:  uint8(env.GetInt("DMG_HEAD", 100)),
		DamageArms:  uint8(env.GetInt("DMG_ARMS", 25)),
		DamageLegs:  uint8(env.GetInt("DMG_LEGS", 25)),
		DamageMelee: uint8(env.GetInt("DMG_MELEE", 100)),
	}

	validate := validator.New()
	if err := validate.Struct(newConfig); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return newConfig, validationErrors
	}

	return newConfig, nil
}
