package db

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"v1/src/errors"
)

type MySQL struct {
	Dsn      string `yaml:"dsn"`
	Name     string `yaml:"name"`
	Password string `yaml:"password"`
}

func GetMySQLProperties() *MySQL {

	yamlFile, err := ioutil.ReadFile("src/resources/mysql.yaml")
	errors.CheckErr(err)

	data := &MySQL{}

	err = yaml.Unmarshal(yamlFile, data)
	errors.CheckErr(err)

	return data
}

type Postgres struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbName"`
	SSLMode  string `yaml:"sslmode"`
}

func GetPostgresData() *Postgres {
	yamlFile, err := ioutil.ReadFile("src/resources/postgresql.yaml")
	errors.CheckErr(err)

	data := &Postgres{}

	err = yaml.Unmarshal(yamlFile, data)
	errors.CheckErr(err)

	return data
}
