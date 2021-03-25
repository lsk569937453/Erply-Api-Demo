package dao

import (
	"database/sql/driver"
	"fmt"
	"time"
)

const DB_TIME_FORMAT = "2006-01-02 15:04:05.000"

type SelfFormatTime struct {
	time.Time
}

func (t SelfFormatTime) String() string {
	return t.Format(DB_TIME_FORMAT)
}

//重写 MarshaJSON 方法，在此方法中实现自定义格式的转换；程序中解析到JSON
func (t SelfFormatTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format(DB_TIME_FORMAT))
	return []byte(formatted), nil
}

//JSON中解析到程序中
func (t *SelfFormatTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+DB_TIME_FORMAT+`"`, string(data), time.Local)
	*t = SelfFormatTime{Time: now}
	return
}

//写入数据库时会调用该方法将自定义时间类型转换并写入数据库
func (t SelfFormatTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

//读取数据库时会调用该方法将时间数据转换成自定义时间类型
func (t *SelfFormatTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = SelfFormatTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
