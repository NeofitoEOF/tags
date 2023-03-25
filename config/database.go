package config

import (
	"fmt"
	"ginFramework/helper"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	StringConexaoBanco = ""
	Porta              = 0
)

func DatabaseConnection() *gorm.DB {
	StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"),
	)
	db, erro := gorm.Open(mysql.Open(StringConexaoBanco), &gorm.Config{})
	helper.ErrorPanic(erro)

	return db
}
