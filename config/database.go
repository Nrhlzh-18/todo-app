package config

import (
	"fmt"

	"github.com/Nrhlzh-18/todo-app/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type credential struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Name     string
	SSLMode  string
	Timezone string
}

func NewDatabase(env *Environment, driver string) *gorm.DB {
	credential := &credential{
		Host:     env.Get("DB_HOST"),
		Port:     env.Get("DB_PORT"),
		User:     env.Get("DB_USER"),
		Pass:     env.Get("DB_PASS"),
		Name:     env.Get("DB_NAME"),
		SSLMode:  env.Get("DB_SSLMODE"),
		Timezone: env.Get("DB_TZ"),
	}

	switch driver {
	case "mysql":
		return credential.getMysql()
	default:
		return credential.getMysql()
	}
}

func (c *credential) getMysql() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=%s", c.User, c.Pass, c.Host, c.Port, c.Name, c.Timezone)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,   // add default size for string fields, by default, will use db type `longtext` for fields without size, not a primary key, no index defined and don't have default values
		DisableDatetimePrecision:  true,  // disable datetime precision support, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create index when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // use change when rename column, rename rename not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // smart configure based on used version
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}

	migrate(db)

	return db
}

func migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.MTasks{},
		&models.MSchedule{},
		&models.MUser{},
		&models.MUserRole{},
	)
	if err != nil {
		fmt.Println("Error while migrating tables:", err)
	} else {
		fmt.Println("Tables migrated successfully!")
	}
}
