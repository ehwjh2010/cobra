package types

import (
	"bytes"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/ehwjh2010/viper/constant"
	"github.com/ehwjh2010/viper/helper/basic/integer"
	"github.com/ehwjh2010/viper/helper/serialize"
)

//********************int64*****************************

// NullInt64 is an alias for sql.NullInt64 data type.
type NullInt64 struct {
	sql.NullInt64
}

// IsEmpty 判断为nil或0.
func (ni *NullInt64) IsEmpty() bool {
	if ni == nil || !ni.Valid || ni.Int64 == 0 {
		return true
	}

	return false
}

func (ni *NullInt64) String() string {
	if !ni.Valid {
		return constant.NullStr
	}

	return strconv.FormatInt(ni.Int64, 10)
}

// IsNull 是否是Nil.
func (ni *NullInt64) IsNull() bool {
	return ni == nil || !ni.Valid
}

// Equal 比较是否相等.
func (ni *NullInt64) Equal(v NullInt64) bool {
	return ni.Valid == v.Valid && (!ni.Valid || ni.Int64 == v.Int64)
}

// GetValue 获取值.
func (ni *NullInt64) GetValue() int64 {
	return ni.Int64
}

func NewInt64(v int64) NullInt64 {
	return NullInt64{NullInt64: sql.NullInt64{
		Int64: v,
		Valid: true,
	}}
}

func NewInt64Null() NullInt64 {
	return NullInt64{}
}

// MarshalJSON for NullInt64.
func (ni NullInt64) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return serialize.Marshal(ni.Int64)
}

// UnmarshalJSON for NullInt64.
func (ni *NullInt64) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, constant.NullBytes) {
		ni.Valid = false
		return nil
	}

	err := serialize.Unmarshal(b, &ni.Int64)

	if err != nil {
		ni.Valid = false
	} else {
		ni.Valid = true
	}

	return err
}

//********************int*****************************

// NullInt is an alias for sql.NullInt64 data type.
type NullInt struct {
	sql.NullInt64
}

// Empty 判断为nil或0.
func (ni *NullInt) Empty() bool {
	if ni == nil || !ni.Valid || ni.Int64 == 0 {
		return true
	}

	return false
}

func (ni NullInt) String() string {
	if !ni.Valid {
		return constant.NullStr
	}

	return strconv.FormatInt(ni.Int64, 10)
}

// IsNull 是否是Nil.
func (ni *NullInt) IsNull() bool {
	return ni == nil || !ni.Valid
}

// GetValue 获取值.
func (ni *NullInt) GetValue() int {
	return integer.Int64ToInt(ni.Int64)
}

func NewInt(v int) NullInt {
	return NullInt{NullInt64: sql.NullInt64{
		Int64: int64(v),
		Valid: true,
	}}
}

func NewIntNull() NullInt {
	return NullInt{}
}

// MarshalJSON for NullInt.
func (ni NullInt) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return serialize.Marshal(integer.Int64ToInt(ni.Int64))
}

// UnmarshalJSON for NullInt.
func (ni *NullInt) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, constant.NullBytes) {
		ni.Valid = false
		return nil
	}

	err := serialize.Unmarshal(b, &ni.Int64)
	if err != nil {
		ni.Valid = false
	} else {
		ni.Valid = true
	}

	return err
}

// Equal 比较是否相等.
func (ni *NullInt) Equal(v NullInt) bool {
	return ni.Valid == v.Valid && (!ni.Valid || ni.Int64 == v.Int64)
}

//********************int32*****************************

// NullInt32 is an alias for sql.NullInt64 data type.
type NullInt32 struct {
	sql.NullInt32
}

// Empty 判断为nil或0.
func (ni *NullInt32) Empty() bool {
	if ni == nil || !ni.Valid || ni.Int32 == 0 {
		return true
	}

	return false
}

func (ni NullInt32) String() string {
	if !ni.Valid {
		return constant.NullStr
	}

	return strconv.FormatInt(int64(ni.Int32), 10)
}

// IsNull 是否是Nil.
func (ni *NullInt32) IsNull() bool {
	return !ni.Valid
}

// GetValue 获取值.
func (ni *NullInt32) GetValue() int32 {
	return ni.Int32
}

func NewInt32(v int32) NullInt32 {
	return NullInt32{NullInt32: sql.NullInt32{
		Int32: v,
		Valid: true,
	}}
}

func NewInt32Null() NullInt32 {
	return NullInt32{}
}

// MarshalJSON for NullInt.
func (ni NullInt32) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return serialize.Marshal(ni.Int32)
}

// UnmarshalJSON for NullInt.
func (ni *NullInt32) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, constant.NullBytes) {
		ni.Valid = false
		return nil
	}

	err := serialize.Unmarshal(b, &ni.Int32)
	if err != nil {
		ni.Valid = false
	} else {
		ni.Valid = true
	}

	return err
}

// Equal 比较是否相等.
func (ni *NullInt32) Equal(v NullInt32) bool {
	return ni.Valid == v.Valid && (!ni.Valid || ni.Int32 == v.Int32)
}

//********************bool*****************************

// NullBool is an alias for sql.NullBool data type.
type NullBool struct {
	sql.NullBool
}

// Empty 判断为nil或false.
func (nb *NullBool) Empty() bool {
	if nb == nil || !nb.Valid || !nb.Bool {
		return true
	}

	return false
}

func (nb NullBool) String() string {
	if !nb.Valid {
		return constant.NullStr
	}

	return strconv.FormatBool(nb.Bool)
}

// IsNull 是否是Nil.
func (nb *NullBool) IsNull() bool {
	return !nb.NullBool.Valid
}

// GetValue 获取值.
func (nb *NullBool) GetValue() bool {
	return nb.NullBool.Bool
}

// MarshalJSON for NullBool.
func (nb NullBool) MarshalJSON() ([]byte, error) {
	if !nb.Valid {
		return []byte("null"), nil
	}
	return serialize.Marshal(nb.Bool)
}

func NewBool(v bool) NullBool {
	return NullBool{NullBool: sql.NullBool{
		Bool:  v,
		Valid: true,
	}}
}

func NewBoolNull() NullBool {
	return NullBool{}
}

// UnmarshalJSON for NullBool.
func (nb *NullBool) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, constant.NullBytes) {
		nb.Valid = false
		return nil
	}

	err := serialize.Unmarshal(b, &nb.Bool)
	if err != nil {
		nb.Valid = false
	} else {
		nb.Valid = true
	}

	return err
}

// Equal 比较是否相等.
func (nb *NullBool) Equal(v NullBool) bool {
	return nb.Valid == v.Valid && (!nb.Valid || nb.Bool == v.Bool)
}

//********************float64*****************************

// NullFloat64 is an alias for sql.NullFloat64 data type.
type NullFloat64 struct {
	sql.NullFloat64
}

// Empty 判断为nil或0.
func (nf NullFloat64) Empty() bool {
	if !nf.Valid || nf.Float64 == 0 {
		return true
	}

	return false
}

func (nf NullFloat64) String() string {
	if !nf.Valid {
		return constant.NullStr
	}

	return strconv.FormatFloat(nf.Float64, 'E', -1, 64)
}

// IsNull 是否是Nil.
func (nf *NullFloat64) IsNull() bool {
	return !nf.NullFloat64.Valid
}

// GetValue 获取值.
func (nf *NullFloat64) GetValue() float64 {
	return nf.NullFloat64.Float64
}

// MarshalJSON for NullFloat64.
func (nf NullFloat64) MarshalJSON() ([]byte, error) {
	if !nf.Valid {
		return []byte("null"), nil
	}
	return serialize.Marshal(nf.Float64)
}

func NewFloat64(v float64) NullFloat64 {
	return NullFloat64{NullFloat64: sql.NullFloat64{
		Float64: v,
		Valid:   true,
	}}
}

func NewFloat64Null() NullFloat64 {
	return NullFloat64{}
}

// UnmarshalJSON for NullFloat64.
func (nf *NullFloat64) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, constant.NullBytes) {
		nf.Valid = false
		return nil
	}

	err := serialize.Unmarshal(b, &nf.Float64)
	if err != nil {
		nf.Valid = false
	} else {
		nf.Valid = true
	}

	return err
}

// Equal 比较是否相等.
func (nf *NullFloat64) Equal(v NullFloat64) bool {
	return nf.Valid == v.Valid && (!nf.Valid || nf.Float64 == v.Float64)
}

//********************string*****************************

// NullString is an alias for sql.NullString data type.
type NullString struct {
	sql.NullString
}

// Empty 判断为nil或"".
func (ns NullString) Empty() bool {
	if !ns.Valid || ns.NullString.String == "" {
		return true
	}

	return false
}

func (ns NullString) String() string {
	if !ns.Valid {
		return constant.NullStr
	}

	return ns.NullString.String
}

// IsNull 是否是Nil.
func (ns *NullString) IsNull() bool {
	return !ns.NullString.Valid
}

// GetValue 获取值.
func (ns *NullString) GetValue() string {
	return ns.NullString.String
}

func NewStr(str string) NullString {
	return NullString{NullString: sql.NullString{
		String: str,
		Valid:  true,
	}}
}

func NewEmptyStr() NullString {
	return NullString{NullString: sql.NullString{
		String: "",
		Valid:  true,
	}}
}

func NewStrNull() NullString {
	return NullString{}
}

// MarshalJSON for NullString.
func (ns NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return serialize.Marshal(ns.NullString.String)
}

// UnmarshalJSON for NullString.
func (ns *NullString) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, constant.NullBytes) {
		ns.Valid = false
		return nil
	}

	err := serialize.Unmarshal(b, &ns.NullString.String)
	if err != nil {
		ns.Valid = false
	} else {
		ns.Valid = true
	}

	return err
}

// Equal 比较是否相等.
func (ns *NullString) Equal(v NullString) bool {
	return ns.Valid == v.Valid && (!ns.Valid || ns.NullString.String == v.NullString.String)
}

//********************time*****************************

// NullTime is an alias for mysql.NullTime data type.
type NullTime struct {
	sql.NullTime
}

func (nt NullTime) String() string {
	if !nt.Valid {
		return constant.NullStr
	}

	return nt.Time.Format(constant.DefaultTimePattern)
}

// IsNull 是否是Nil.
func (nt *NullTime) IsNull() bool {
	return !nt.Valid
}

// Empty 判断为nil或0.
func (nt NullTime) Empty() bool {
	if !nt.Valid || nt.Time.Equal(time.Unix(0, 0)) {
		return true
	}

	return false
}

// GetValue 获取值.
func (nt *NullTime) GetValue() time.Time {
	return nt.Time
}

// TimeStamp 获取时间戳, 单位: s.
func (nt *NullTime) TimeStamp() int64 {
	return nt.Time.Unix()
}

// MarshalJSON for NullTime.
func (nt NullTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return constant.NullBytes, nil
	}
	val := fmt.Sprintf("\"%s\"", nt.Time.Format(constant.DefaultTimePattern))
	return []byte(val), nil
}

func NewTime(t time.Time) NullTime {
	return NullTime{NullTime: sql.NullTime{
		Time:  t,
		Valid: true,
	}}
}

func NewTimeNull() NullTime {
	return NullTime{}
}

// UnmarshalJSON for NullTime.
func (nt *NullTime) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, constant.NullBytes) {
		nt.Valid = false
		return nil
	}

	err := serialize.Unmarshal(b, &nt.Time)
	if err != nil {
		nt.Valid = false
	} else {
		nt.Valid = true
	}

	return err
}

// Equal 比较是否相等.
func (nt *NullTime) Equal(v NullTime) bool {
	return nt.Valid == v.Valid && (!nt.Valid || nt.Time == v.Time)
}
