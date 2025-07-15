package types

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

// NullString 自定义的可空字符串类型，序列化时只返回值
type NullString struct {
	sql.NullString
}

// MarshalJSON 实现 JSON 序列化，只返回字符串值
func (ns NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(ns.String)
}

// UnmarshalJSON 实现 JSON 反序列化
func (ns *NullString) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s == "" {
		ns.Valid = false
		ns.String = ""
	} else {
		ns.Valid = true
		ns.String = s
	}
	return nil
}

// Scan 实现 sql.Scanner 接口
func (ns *NullString) Scan(value interface{}) error {
	return ns.NullString.Scan(value)
}

// Value 实现 driver.Valuer 接口
func (ns NullString) Value() (driver.Value, error) {
	return ns.NullString.Value()
}

// NullInt64 自定义的可空整数类型，序列化时只返回值
type NullInt64 struct {
	sql.NullInt64
}

// MarshalJSON 实现 JSON 序列化，只返回整数值
func (ni NullInt64) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(ni.Int64)
}

// UnmarshalJSON 实现 JSON 反序列化
func (ni *NullInt64) UnmarshalJSON(data []byte) error {
	var i int64
	if err := json.Unmarshal(data, &i); err != nil {
		return err
	}
	ni.Valid = true
	ni.Int64 = i
	return nil
}

// Scan 实现 sql.Scanner 接口
func (ni *NullInt64) Scan(value interface{}) error {
	return ni.NullInt64.Scan(value)
}

// Value 实现 driver.Valuer 接口
func (ni NullInt64) Value() (driver.Value, error) {
	return ni.NullInt64.Value()
}

// NullFloat64 自定义的可空浮点数类型，序列化时只返回值
type NullFloat64 struct {
	sql.NullFloat64
}

// MarshalJSON 实现 JSON 序列化，只返回浮点数值
func (nf NullFloat64) MarshalJSON() ([]byte, error) {
	if !nf.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(nf.Float64)
}

// UnmarshalJSON 实现 JSON 反序列化
func (nf *NullFloat64) UnmarshalJSON(data []byte) error {
	var f float64
	if err := json.Unmarshal(data, &f); err != nil {
		return err
	}
	nf.Valid = true
	nf.Float64 = f
	return nil
}

// Scan 实现 sql.Scanner 接口
func (nf *NullFloat64) Scan(value interface{}) error {
	return nf.NullFloat64.Scan(value)
}

// Value 实现 driver.Valuer 接口
func (nf NullFloat64) Value() (driver.Value, error) {
	return nf.NullFloat64.Value()
}
