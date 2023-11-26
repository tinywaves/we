package dao

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	ErrorUserDuplicateEmail = errors.New("duplicate email address")
)

type UserDao struct {
	database *gorm.DB
}

// User mapping to database. entity/model/PO(persistent object).
type User struct {
	Id         int64  `gorm:"primaryKey,autoIncrement"`
	Email      string `gorm:"unique"`
	Password   string
	CreateTime int64 // millisecond UTC
	UpdateTime int64
	DeleteTime int64
}

func InitUsersDao(database *gorm.DB) *UserDao {
	return &UserDao{
		database: database,
	}
}

func (dao *UserDao) InsertNewUser(ctx context.Context, u User) error {
	now := time.Now().UnixMilli()
	u.CreateTime = now
	u.UpdateTime = now
	err := dao.database.WithContext(ctx).Create(&u).Error
	var mysqlError *mysql.MySQLError
	if errors.As(err, &mysqlError) {
		const uniqueConflictsErrorNumber uint16 = 1062 // Mysql error number for unique conflicts.
		if mysqlError.Number == uniqueConflictsErrorNumber {
			return ErrorUserDuplicateEmail
		}
	}
	return err
}
