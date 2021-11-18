package types

import (
	"bytes"
	"database/sql"
	"fmt"
	"ginLearn/enum"
	"ginLearn/util/intutils"
	"ginLearn/util/jsonutils"
	"time"
)

var nullBytes = []byte("null")

//********************int64*****************************

// NullInt64 is an alias for sql.NullInt64 data type
type NullInt64 struct {
	sql.NullInt64
}

//IsNil 是否是Nil
func (ni *NullInt64) IsNil() bool {
	return !ni.NullInt64.Valid
}

//Equal 比较是否相等
func (ni *NullInt64) Equal(v NullInt64) bool {
	return ni.Valid == v.Valid && (!ni.Valid || ni.Int64 == v.Int64)
}

//GetValue 获取值
func (ni *NullInt64) GetValue() int64 {
	return ni.Int64
}

type nullInt64Opt func(nullInt64 *NullInt64)

func newInt64WithInt64(v int64) nullInt64Opt {
	return func(nullInt64 *NullInt64) {
		nullInt64.Int64 = v
	}
}

func newInt64WithValid(valid bool) nullInt64Opt {
	return func(nullInt64 *NullInt64) {
		nullInt64.NullInt64.Valid = valid
	}
}

func newNullInt64(args ...nullInt64Opt) *NullInt64 {
	nt := &NullInt64{}

	for _, arg := range args {
		arg(nt)
	}

	return nt
}

func NewInt64(v int64) NullInt64 {
	return *newNullInt64(newInt64WithInt64(v), newInt64WithValid(true))
}

func NewInt64Null() NullInt64 {
	return *newNullInt64()
}

// MarshalJSON for NullInt64
func (ni NullInt64) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return jsonutils.Marshal(ni.Int64)
}

//UnmarshalJSON for NullInt64
func (ni *NullInt64) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, nullBytes) {
		ni.Valid = false
		return nil
	}

	err := jsonutils.Unmarshal(b, &ni.Int64)

	if err != nil {
		ni.Valid = false
	} else {
		ni.Valid = true
	}

	return err
}

//********************int64*****************************

// NullInt is an alias for sql.NullInt64 data type
type NullInt struct {
	sql.NullInt64
}

//IsNil 是否是Nil
func (ni *NullInt) IsNil() bool {
	return !ni.Valid
}

//GetValue 获取值
func (ni *NullInt) GetValue() int {
	return intutils.Int64ToInt(ni.Int64)
}

type nullIntOpt func(nullInt *NullInt)

func newIntWithInt(v int) nullIntOpt {
	return func(nullInt *NullInt) {
		nullInt.Int64 = int64(v)
	}
}

func newIntWithValid(valid bool) nullIntOpt {
	return func(nullInt *NullInt) {
		nullInt.Valid = valid
	}
}

func newNullInt(args ...nullIntOpt) *NullInt {
	nt := &NullInt{}

	for _, arg := range args {
		arg(nt)
	}

	return nt
}

func NewInt(v int) NullInt {
	return *newNullInt(newIntWithInt(v), newIntWithValid(true))
}

func NewIntNull() NullInt {
	return *newNullInt()
}

// MarshalJSON for NullInt
func (ni NullInt) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return jsonutils.Marshal(intutils.Int64ToInt(ni.Int64))
}

//UnmarshalJSON for NullInt
func (ni *NullInt) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, nullBytes) {
		ni.Valid = false
		return nil
	}

	err := jsonutils.Unmarshal(b, &ni.Int64)
	if err != nil {
		ni.Valid = false
	} else {
		ni.Valid = true
	}

	return err
}

//Equal 比较是否相等
func (ni *NullInt) Equal(v NullInt) bool {
	return ni.Valid == v.Valid && (!ni.Valid || ni.Int64 == v.Int64)
}

//********************bool*****************************

// NullBool is an alias for sql.NullBool data type
type NullBool struct {
	sql.NullBool
}

//IsNil 是否是Nil
func (nb *NullBool) IsNil() bool {
	return !nb.NullBool.Valid
}

//GetValue 获取值
func (nb *NullBool) GetValue() bool {
	return nb.NullBool.Bool
}

// MarshalJSON for NullBool
func (nb NullBool) MarshalJSON() ([]byte, error) {
	if !nb.Valid {
		return []byte("null"), nil
	}
	return jsonutils.Marshal(nb.Bool)
}

type nullBoolOpt func(nullBool *NullBool)

func newBoolWithBool(v bool) nullBoolOpt {
	return func(nullBool *NullBool) {
		nullBool.NullBool.Bool = v
	}
}

func newBoolWithValid(valid bool) nullBoolOpt {
	return func(nullBool *NullBool) {
		nullBool.NullBool.Valid = valid
	}
}

func newNullBool(args ...nullBoolOpt) *NullBool {
	nt := &NullBool{}

	for _, arg := range args {
		arg(nt)
	}

	return nt
}

func NewBool(v bool) NullBool {
	return *newNullBool(newBoolWithBool(v), newBoolWithValid(true))
}

func NewBoolNull() NullBool {
	return *newNullBool()
}

//UnmarshalJSON for NullBool
func (nb *NullBool) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, nullBytes) {
		nb.Valid = false
		return nil
	}

	err := jsonutils.Unmarshal(b, &nb.Bool)
	if err != nil {
		nb.Valid = false
	} else {
		nb.Valid = true
	}

	return err
}

//Equal 比较是否相等
func (nb *NullBool) Equal(v NullBool) bool {
	return nb.Valid == v.Valid && (!nb.Valid || nb.Bool == v.Bool)
}

//********************float64*****************************

// NullFloat64 is an alias for sql.NullFloat64 data type
type NullFloat64 struct {
	sql.NullFloat64
}

//IsNil 是否是Nil
func (nf *NullFloat64) IsNil() bool {
	return !nf.NullFloat64.Valid
}

//GetValue 获取值
func (nf *NullFloat64) GetValue() float64 {
	return nf.NullFloat64.Float64
}

// MarshalJSON for NullFloat64
func (nf NullFloat64) MarshalJSON() ([]byte, error) {
	if !nf.Valid {
		return []byte("null"), nil
	}
	return jsonutils.Marshal(nf.Float64)
}

type nullFloat64Opt func(nullFloat64 *NullFloat64)

func newFloat64WithFloat64(v float64) nullFloat64Opt {
	return func(nullFloat64 *NullFloat64) {
		nullFloat64.NullFloat64.Float64 = v
	}
}

func newFloat64WithValid(valid bool) nullFloat64Opt {
	return func(nullFloat64 *NullFloat64) {
		nullFloat64.NullFloat64.Valid = valid
	}
}

func newNullFloat64(args ...nullFloat64Opt) *NullFloat64 {
	nt := &NullFloat64{}

	for _, arg := range args {
		arg(nt)
	}

	return nt
}

func NewFloat64(v float64) *NullFloat64 {
	return newNullFloat64(newFloat64WithFloat64(v), newFloat64WithValid(true))
}

func NewFloat64Null() *NullFloat64 {
	return newNullFloat64()
}

// UnmarshalJSON for NullFloat64
func (nf *NullFloat64) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, nullBytes) {
		nf.Valid = false
		return nil
	}

	err := jsonutils.Unmarshal(b, &nf.Float64)
	if err != nil {
		nf.Valid = false
	} else {
		nf.Valid = true
	}

	return err
}

//Equal 比较是否相等
func (nf *NullFloat64) Equal(v NullFloat64) bool {
	return nf.Valid == v.Valid && (!nf.Valid || nf.Float64 == v.Float64)
}

//********************string*****************************

// NullString is an alias for sql.NullString data type
type NullString struct {
	sql.NullString
}

//IsNil 是否是Nil
func (ns *NullString) IsNil() bool {
	return !ns.NullString.Valid
}

//GetValue 获取值
func (ns *NullString) GetValue() string {
	return ns.NullString.String
}

type nullStrOpt func(ns *NullString)

func nullStrWithStr(s string) nullStrOpt {
	return func(ns *NullString) {
		ns.NullString.String = s
	}
}

func nullStrWithValid(valid bool) nullStrOpt {
	return func(ns *NullString) {
		ns.NullString.Valid = valid
	}
}

func newNullString(args ...nullStrOpt) *NullString {
	ns := &NullString{}

	for _, arg := range args {
		arg(ns)
	}

	return ns
}

func NewStr(str string) NullString {
	return *newNullString(nullStrWithStr(str), nullStrWithValid(true))
}

func NewEmptyStr() NullString {
	return *newNullString(nullStrWithValid(true))
}

func NewStrNull() NullString {
	return *newNullString()
}

// MarshalJSON for NullString
func (ns NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return jsonutils.Marshal(ns.String)
}

// UnmarshalJSON for NullString
func (ns *NullString) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, nullBytes) {
		ns.Valid = false
		return nil
	}

	err := jsonutils.Unmarshal(b, &ns.String)
	if err != nil {
		ns.Valid = false
	} else {
		ns.Valid = true
	}

	return err
}

//Equal 比较是否相等
func (ns *NullString) Equal(v NullString) bool {
	return ns.Valid == v.Valid && (!ns.Valid || ns.String == v.String)
}

//********************time*****************************

//NullTime is an alias for mysql.NullTime data type
type NullTime struct {
	sql.NullTime
}

//type NullTime sql.NullTime

//IsNil 是否是Nil
func (nt *NullTime) IsNil() bool {
	return !nt.Valid
}

//GetValue 获取值
func (nt *NullTime) GetValue() time.Time {
	return nt.Time
}

// MarshalJSON for NullTime
func (nt NullTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return []byte("null"), nil
	}
	val := fmt.Sprintf("\"%s\"", nt.Time.Format(enum.DefaultTimePattern))
	return []byte(val), nil
}

type nullTimeOpt func(nullTime *NullTime)

func newTimeWithTime(v time.Time) nullTimeOpt {
	return func(nullTime *NullTime) {
		nullTime.Time = v
	}
}

func newTimeWithValid(valid bool) nullTimeOpt {
	return func(nullTime *NullTime) {
		nullTime.Valid = valid
	}
}

func newNullTime(args ...nullTimeOpt) *NullTime {
	nt := &NullTime{}

	for _, arg := range args {
		arg(nt)
	}

	return nt
}

func NewTime(t time.Time) *NullTime {
	return newNullTime(newTimeWithTime(t), newTimeWithValid(true))
}

func NewTimeNull() *NullTime {
	return newNullTime()
}

// UnmarshalJSON for NullTime
func (nt *NullTime) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, nullBytes) {
		nt.Valid = false
		return nil
	}

	err := jsonutils.Unmarshal(b, &nt.Time)
	if err != nil {
		nt.Valid = false
	} else {
		nt.Valid = true
	}

	return err
}

//Equal 比较是否相等
func (nt *NullTime) Equal(v NullTime) bool {
	return nt.Valid == v.Valid && (!nt.Valid || nt.Time == v.Time)
}
