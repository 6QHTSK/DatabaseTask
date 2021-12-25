package service

import (
	"DatabaseTask/model"
	"database/sql"
	"fmt"
)

// 定义空字符串和null等价
func stringToNullString(s string) sql.NullString {
	if s != "" {
		return sql.NullString{String: s, Valid: true}
	} else {
		return sql.NullString{}
	}
}

func nullStringToString(nullStr sql.NullString) string {
	if nullStr.Valid {
		return nullStr.String
	} else {
		return ""
	}
}

// 定义-1为无成绩
func intToNullByte(n int) sql.NullByte {
	if n != -1 {
		return sql.NullByte{
			Byte:  byte(n),
			Valid: true,
		}
	}
	return sql.NullByte{}
}

func nullByteToInt(n sql.NullByte) int {
	if !n.Valid {
		return -1
	}
	return int(n.Byte)
}

func InsertStudent(student model.StudentDetailInfo) (err error) {
	dbStudent := model.DbStudent{
		Sno:   student.Number,
		Sname: student.Name,
		Ssex:  student.Sex,
		Sage:  student.Age,
		Sdept: student.Dept,
	}
	_, err = SqlDB.Exec("INSERT INTO Student(Sno, Sname, Ssex, Sage, Sdept) VALUES(?,?,?,?,?)", dbStudent.Sno, dbStudent.Sname, dbStudent.Ssex, dbStudent.Sage, dbStudent.Sdept)
	return err
}

func UpdateStudent(oldStudentNumber string, student model.StudentBasicInfo) (err error) {
	found, err := findStudent(oldStudentNumber)
	if err != nil {
		return err
	}
	if !found {
		return fmt.Errorf("student not found")
	}
	dbStudent := model.DbStudent{
		Sno:   student.Number,
		Sname: student.Name,
		Ssex:  student.Sex,
		Sage:  student.Age,
		Sdept: student.Dept,
	}
	_, err = SqlDB.Exec("UPDATE Student SET Sno = ?,Sname = ?,Ssex = ?,Sage = ?,Sdept = ? where Sno = ?", dbStudent.Sno, dbStudent.Sname, dbStudent.Ssex, dbStudent.Sage, dbStudent.Sdept, oldStudentNumber)
	return err
}

func GetBasicStudentInfo(studentNumber string) (info model.StudentBasicInfo, err error) {
	found, err := findStudent(studentNumber)
	if err != nil {
		return info, err
	}
	if !found {
		return info, fmt.Errorf("不存在该学生")
	}
	var dbStudent model.DbStudent
	err = SqlDB.Get(&dbStudent, "SELECT Sno,Sname,Ssex,Sage,Sdept,Scholarship from Student where Sno = ?", studentNumber)
	return model.StudentBasicInfo{
		Number: dbStudent.Sno,
		Name:   dbStudent.Sname,
		Sex:    dbStudent.Ssex,
		Age:    dbStudent.Sage,
		Dept:   dbStudent.Sdept,
	}, err
}

func GetBasicCourseInfo(courseNumber string) (info model.CourseInfo, err error) {
	found, err := findCourse(courseNumber)
	if err != nil {
		return info, err
	}
	if !found {
		return info, fmt.Errorf("student not found")
	}
	var dbCourse model.DbCourse
	err = SqlDB.Get(&dbCourse, "SELECT Cno,Cname,Cpno,Ccredit from Course where Cno = ?", courseNumber)
	return model.CourseInfo{
		Number:               dbCourse.Cno,
		Name:                 dbCourse.Cname,
		PreviousCourseNumber: nullStringToString(dbCourse.Cpno),
		Credit:               dbCourse.Ccredit,
	}, nil
}

func findCourse(courseNumber string) (exist bool, err error) {
	var count int
	err = SqlDB.Get(&count, "SELECT count(*) from Course where Cno=?", courseNumber)
	return count > 0, err
}

func findStudent(studentNumber string) (exist bool, err error) {
	var count int
	err = SqlDB.Get(&count, "SELECT count(*) from Student where Sno=?", studentNumber)
	return count > 0, err
}

func findDept(deptName string) (exist bool, err error) {
	var count int
	err = SqlDB.Get(&count, "SELECT count(*) from Student where Sdept=?", deptName)
	return count > 0, err
}

func validDeptGrade(deptName string, courseNumber string) (valid bool, err error) {
	var count int
	err = SqlDB.Get(&count, "SELECT count(*) from Student,SC where Sdept=? and Student.Sno = SC.Sno and SC.Cno = ? and Grade is not null", deptName, courseNumber)
	return count > 0, err
}

func InsertCourse(course model.CourseInfo) (err error) {
	dbCourse := model.DbCourse{
		Cno:     course.Number,
		Cname:   course.Name,
		Ccredit: course.Credit,
		Cpno:    stringToNullString(course.PreviousCourseNumber),
	}
	if dbCourse.Cpno.Valid {
		exist, err := findCourse(course.PreviousCourseNumber)
		if err != nil {
			return err
		}
		if !exist {
			return fmt.Errorf("不存在前置课程")
		}
	}
	_, err = SqlDB.Exec("INSERT INTO Course VALUES(?,?,?,?)", dbCourse.Cno, dbCourse.Cname, dbCourse.Cpno, dbCourse.Ccredit)
	return err
}

func UpdateCourse(courseNumber string, course model.CourseInfo) (err error) {
	exist, err := findCourse(courseNumber)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("不存在该课程")
	}
	dbCourse := model.DbCourse{
		Cno:     course.Number,
		Cname:   course.Name,
		Ccredit: course.Credit,
		Cpno:    stringToNullString(course.PreviousCourseNumber),
	}
	if dbCourse.Cpno.Valid {
		exist, err := findCourse(course.PreviousCourseNumber)
		if err != nil {
			return err
		}
		if !exist {
			return fmt.Errorf("不存在前置课程")
		}
	}
	_, err = SqlDB.Exec("UPDATE Course SET Cname = ?, Cpno = ?, Ccredit = ?,Cno = ? where Cno = ?", dbCourse.Cname, dbCourse.Cpno, dbCourse.Ccredit, dbCourse.Cno, courseNumber)
	return err
}

func DeleteNotChoiceCourse() (err error) {
	_, err = SqlDB.Exec("CALL DeleteCourseStrict()")
	return err
}

func InsertStudentCourseGrade(info model.SCInfo) (err error) {
	exist, err := findCourse(info.CourseNumber)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("不存在该课程")
	}
	exist, err = findStudent(info.StudentNumber)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("不存在该学生")
	}
	dbSc := model.DbSC{
		Sno:   info.StudentNumber,
		Cno:   info.CourseNumber,
		Grade: intToNullByte(info.Score),
	}
	_, err = SqlDB.Exec("INSERT INTO SC VALUES(?,?,?)", dbSc.Sno, dbSc.Cno, dbSc.Grade)
	return err
}

func UpdateStudentGrade(info model.SCInfo) (err error) {
	exist, err := findCourse(info.CourseNumber)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("不存在该课程")
	}
	exist, err = findStudent(info.StudentNumber)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("不存在该学生")
	}
	dbSc := model.DbSC{
		Sno:   info.StudentNumber,
		Cno:   info.CourseNumber,
		Grade: intToNullByte(info.Score),
	}
	_, err = SqlDB.Exec("UPDATE SC SET Grade = ? WHERE Sno = ? and Cno = ?", dbSc.Grade, dbSc.Sno, dbSc.Cno)
	return err
}

func GetDeptGradeInfo(deptName string, courseNumber string) (info model.DeptGradeInfo, err error) {
	exist, err := findDept(deptName)
	if err != nil {
		return info, err
	}
	if !exist {
		return info, fmt.Errorf("该系不存在")
	}
	exist, err = findCourse(courseNumber)
	if err != nil {
		return info, err
	}
	if !exist {
		return info, fmt.Errorf("不存在该课程")
	}
	valid, err := validDeptGrade(deptName, courseNumber)
	if err != nil {
		return info, err
	}
	if !valid {
		return info, fmt.Errorf("该系该课程所有学生的成绩未出")
	}
	var dbInfo model.DbDeptGradeInfo
	err = SqlDB.Get(&dbInfo, "CALL ShowDeptGrade(?,?)", deptName, courseNumber)
	info = model.DeptGradeInfo{
		DeptName:    deptName,
		Average:     dbInfo.Avr,
		Best:        dbInfo.Max,
		Worst:       dbInfo.Min,
		GoodRate:    dbInfo.Rate,
		FailedCount: dbInfo.Fail,
	}
	return info, err
}

func GetDeptStudentScoreList(deptName string, courseNumber string) (list []model.StudentScore, err error) {
	exist, err := findDept(deptName)
	if err != nil {
		return list, err
	}
	if !exist {
		return list, fmt.Errorf("不存在该系")
	}
	exist, err = findCourse(courseNumber)
	if err != nil {
		return list, err
	}
	if !exist {
		return list, fmt.Errorf("不存在该课程")
	}
	var dbInfo []model.DbStudentGradeInfo
	err = SqlDB.Select(&dbInfo, "SELECT Sno,Sname,Ssex,Sage,Scholarship,Grade from StudentGrade where Cno = ? and Sdept = ? order by Grade desc", courseNumber, deptName)
	if err != nil {
		return list, err
	}
	for _, item := range dbInfo {
		list = append(list, model.StudentScore{
			StudentBasicInfo: model.StudentBasicInfo{
				Number: item.Sno,
				Name:   item.Sname,
				Sex:    item.Ssex,
				Age:    item.Sage,
			},
			Scholarship: item.Scholarship,
			Score:       nullByteToInt(item.Grade),
		})
	}
	return list, err
}

func GetStudentInfo(studentNumber string) (info model.StudentDetailInfo, err error) {
	exist, err := findStudent(studentNumber)
	if err != nil {
		return info, err
	}
	if !exist {
		return info, fmt.Errorf("不存在该学生")
	}
	var dbInfo model.DbStudent
	err = SqlDB.Get(&dbInfo, "SELECT * from Student where Sno = ?", studentNumber)
	var dbList []model.DbCourseGradeInfo
	err = SqlDB.Select(&dbList, "SELECT Cno,Cname,Cpno,Ccredit,Grade from CourseGrade where Sno = ?", studentNumber)
	info = model.StudentDetailInfo{
		StudentBasicInfo: model.StudentBasicInfo{
			Number: dbInfo.Sno,
			Name:   dbInfo.Sname,
			Sex:    dbInfo.Ssex,
			Age:    dbInfo.Sage,
			Dept:   dbInfo.Sdept,
		},
		Scholarship: dbInfo.Scholarship,
		CourseList:  []model.CourseScore{},
	}
	for _, item := range dbList {
		info.CourseList = append(info.CourseList, model.CourseScore{
			CourseInfo: model.CourseInfo{
				Number:               item.Cno,
				Name:                 item.Cname,
				PreviousCourseNumber: nullStringToString(item.Cpno),
				Credit:               item.Ccredit,
			},
			Score: nullByteToInt(item.Grade),
		})
	}
	return info, err
}
