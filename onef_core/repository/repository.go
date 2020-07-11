package repository

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/hoaxoan/onef-api/onef_core/setting"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Connection1() (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(setting.Config.Mongo.Uri))
	if err != nil {
		log.Fatalf("connect error :%v", err)
		return nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	return client, nil
}

func Connection() (*gorm.DB, error) {
	gorm.NowFunc = func() time.Time {
		return time.Now().UTC().Truncate(1000 * time.Nanosecond)
	}

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", setting.Config.Mysql.Username, setting.Config.Mysql.Password, setting.Config.Mysql.Host, setting.Config.Mysql.Port, setting.Config.Mysql.DbName)
	val := url.Values{}
	val.Add("charset", "utf8")
	val.Add("parseTime", "True")
	val.Add("loc", "Local")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	db, err := gorm.Open("mysql", dsn)

	if err != nil {
		log.Fatalf("Connect error :%v", err)
		//log.Fatalf("Connect not success to mysql database at host:%s with user:%s and db:%s", setting.Config.Mysql.Host, setting.Config.Mysql.Username, setting.Config.Mysql.DbName)
		return nil, err
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().Ping()
	db.LogMode(true)
	db.SingularTable(true)

	// defer func() {
	// 	err := db.Close()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()

	return db, nil
}
