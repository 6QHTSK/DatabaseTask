package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type DbCourse struct {
	Cno     string         `db:"Cno"`
	Cname   string         `db:"Cname"`
	Cpno    sql.NullString `db:"Cpno"`
	Ccredit int            `db:"Ccredit"`
}

type DbSC struct {
	Sno   string       `db:"Sno"`
	Cno   string       `db:"Cno"`
	Grade sql.NullByte `db:"Grade"`
}

type DbStudent struct {
	Sno         string `db:"Sno"`
	Sname       string `db:"Sname"`
	Ssex        string `db:"Ssex"`
	Sage        byte   `db:"Sage"`
	Sdept       string `db:"Sdept"`
	Scholarship string `db:"Scholarship"`
}

type DbDeptGradeInfo struct {
	Avr  float64 `db:"avg"`
	Max  byte    `db:"max"`
	Min  byte    `db:"min"`
	Rate float32 `db:"rate"`
	Fail int     `db:"fail"`
}

type DbStudentGradeInfo struct {
	Sno         string       `db:"Sno"`
	Sname       string       `db:"Sname"`
	Ssex        string       `db:"Ssex"`
	Sage        byte         `db:"Sage"`
	Scholarship string       `db:"Scholarship"`
	Grade       sql.NullByte `db:"Grade"`
}

type DbCourseGradeInfo struct {
	Cno     string         `db:"Cno"`
	Cname   string         `db:"Cname"`
	Cpno    sql.NullString `db:"Cpno"`
	Ccredit int            `db:"Ccredit"`
	Grade   sql.NullByte   `db:"Grade"`
}

type StudentBasicInfo struct {
	Number string `json:"number"`
	Name   string `json:"name"`
	Sex    string `json:"sex"`
	Age    byte   `json:"age"`
	Dept   string `json:"dept"`
}

type SCInfo struct {
	StudentNumber string `json:"student_number"`
	CourseNumber  string `json:"course_number"`
	Score         int    `json:"score"`
}

type CourseInfo struct {
	Number               string `json:"number"`
	Name                 string `json:"name"`
	PreviousCourseNumber string `json:"previous_number"`
	Credit               int    `json:"credit"`
}

type StudentDetailInfo struct {
	StudentBasicInfo
	Scholarship string        `json:"scholarship"`
	CourseList  []CourseScore `json:"course_list,omitempty"`
}

type StudentScore struct {
	StudentBasicInfo
	Scholarship string `json:"scholarship"`
	Score       int    `json:"score"`
}

type CourseScore struct {
	CourseInfo
	Score int `json:"score"`
}

type DeptGradeInfo struct {
	DeptName    string  `json:"dept_name"`
	Average     float64 `json:"average"`
	Best        byte    `json:"best"`
	Worst       byte    `json:"worst"`
	GoodRate    float32 `json:"good_rate"`
	FailedCount int     `json:"failed_count"`
}
