package config

import (
	"github.com/spf13/viper"
)


func NewJwt(viper *viper.Viper)[]byte{
	Jwtkey := []byte(viper.GetString("jwt.secret"))
	return Jwtkey
}