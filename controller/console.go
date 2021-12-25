package controller

import (
	"DatabaseTask/model"
	"DatabaseTask/service"
	"fmt"
)

func errorHandler(err error) {
	fmt.Println("出现错误", err)
}

func stringScore(score int) string {
	if score != -1 {
		return fmt.Sprintf("%d", score)
	}
	return ""
}

func InsertStudentHandler() {
	var studentInfo model.StudentDetailInfo
	fmt.Println("请输入学生学号：")
	fmt.Scanln(&studentInfo.Number)
	fmt.Println("请输入学生姓名：")
	fmt.Scanln(&studentInfo.Name)
	fmt.Println("请输入学生年龄")
	fmt.Scanf("%d", &studentInfo.Age)
	fmt.Println("请输入学生所在系")
	fmt.Scanln(&studentInfo.Dept)
	for true {
		var code int
		fmt.Println("请输入学生性别：0 男生 1 女生")
		_, err := fmt.Scanf("%d", &code)
		if err == nil {
			switch code {
			case 0:
				studentInfo.Sex = "男"
			case 1:
				studentInfo.Sex = "女"
			default:
				fmt.Println("不合法的输入")
				continue
			}
			break
		}
		fmt.Println("不合法的输入")
	}
	err := service.InsertStudent(studentInfo)
	if err == nil {
		fmt.Println("成功插入")
	} else {
		errorHandler(err)
		return
	}
}

func printStudentBasic(studentNumber string) (info model.StudentBasicInfo, err error) {
	info, err = service.GetBasicStudentInfo(studentNumber)
	if err != nil {
		return info, err
	}
	fmt.Println("1-学号\t\t2-姓名\t3-性别\t4-年龄\t5-所在系")
	fmt.Printf("%s\t%s\t%s\t%d\t%s\n", info.Number, info.Name, info.Sex, info.Age, info.Dept)
	return info, nil
}

func UpdateStudentHandler() {
	var studentNumber string
	fmt.Println("请输入学生学号：")
	fmt.Scanln(&studentNumber)
	info, err := printStudentBasic(studentNumber)
	if err != nil {
		errorHandler(err)
		return
	}
	for true {
		fmt.Println("请输入需要修改的项目,0确认")
		var code int
		fmt.Scanf("%d", &code)
		if code < 0 && code > 5 {
			fmt.Println("错误输入")
			continue
		} else if code == 0 {
			err := service.UpdateStudent(studentNumber, info)
			if err != nil {
				errorHandler(err)
			} else {
				fmt.Println("修改成功")
			}
			break
		} else if code == 4 {
			fmt.Println("请输入需要修改的值")
			fmt.Scanf("%d", &info.Age)
		} else if code == 3 {
			for true {
				var code int
				fmt.Println("请输入学生性别：0 男生 1 女生")
				_, err := fmt.Scanf("%d", &code)
				if err == nil {
					switch code {
					case 0:
						info.Sex = "男"
					case 1:
						info.Sex = "女"
					default:
						fmt.Println("不合法的输入")
						continue
					}
					break
				}
				fmt.Println("不合法的输入")
			}
		} else {
			fmt.Println("请输入需要修改的值")
			var value string
			fmt.Scanf("%s", &value)
			switch code {
			case 1:
				info.Number = value
			case 2:
				info.Name = value
			case 5:
				info.Dept = value
			}
		}
	}
}

func InsertCourseHandler() {
	var courseInfo model.CourseInfo
	fmt.Println("请输入课程号：")
	fmt.Scanln(&courseInfo.Number)
	fmt.Println("请输入课程名：")
	fmt.Scanln(&courseInfo.Name)
	fmt.Println("请输入学分")
	fmt.Scanf("%d", &courseInfo.Credit)
	fmt.Println("请输入前置课程")
	fmt.Scanln(&courseInfo.PreviousCourseNumber)
	err := service.InsertCourse(courseInfo)
	if err == nil {
		fmt.Println("成功插入")
	} else {
		errorHandler(err)
		return
	}
}

func printCourseBasic(CourseNumber string) (info model.CourseInfo, err error) {
	info, err = service.GetBasicCourseInfo(CourseNumber)
	if err != nil {
		return info, err
	}
	fmt.Println("1-课程号\t2-课程名\t3-前置课程号\t4-学分")
	fmt.Printf("%-8s\t%-8s\t%-8s\t%d\n", info.Number, info.Name, info.PreviousCourseNumber, info.Credit)
	return info, nil
}

func UpdateCourseHandler() {
	var courseNumber string
	fmt.Println("请输入课程号：")
	fmt.Scanln(&courseNumber)
	info, err := printCourseBasic(courseNumber)
	if err != nil {
		errorHandler(err)
		return
	}
	for true {
		fmt.Println("请输入需要修改的项目,0确认")
		var code int
		fmt.Scanf("%d", &code)
		if code < 0 && code > 4 {
			fmt.Println("错误输入")
			continue
		} else if code == 0 {
			err := service.UpdateCourse(courseNumber, info)
			if err != nil {
				errorHandler(err)
				return
			} else {
				fmt.Println("修改成功")
			}
			break
		} else if code == 4 {
			fmt.Println("请输入需要修改的值")
			fmt.Scanf("%d", &info.Credit)
		} else {
			fmt.Println("请输入需要修改的值")
			var value string
			fmt.Scanf("%s", &value)
			switch code {
			case 1:
				info.Number = value
			case 2:
				info.Name = value
			case 3:
				info.PreviousCourseNumber = value
			}
		}
	}
}

func DeleteNotChoiceCourseHandler() {
	fmt.Println("删除未被选课未被依赖的课程,Y确认")
	var code string
	fmt.Scanf("%s", &code)
	if code == "Y" {
		err := service.DeleteNotChoiceCourse()
		if err != nil {
			errorHandler(err)
			return
		} else {
			fmt.Println("成功执行")
		}
	} else {
		fmt.Println("取消执行")
	}
}

func InsertStudentScoreHandler() {
	var info model.SCInfo
	fmt.Println("请输入学号：")
	fmt.Scanln(&info.StudentNumber)
	fmt.Println("请输入课程号：")
	fmt.Scanln(&info.CourseNumber)
	fmt.Println("请输入成绩[输入-1为未出]")
	fmt.Scanf("%d", &info.Score)
	err := service.InsertStudentCourseGrade(info)
	if err != nil {
		errorHandler(err)
		return
	} else {
		fmt.Println("成功执行")
	}
}

func UpdateStudentScoreHandler() {
	var info model.SCInfo
	fmt.Println("请输入学号：")
	fmt.Scanln(&info.StudentNumber)
	fmt.Println("请输入课程号：")
	fmt.Scanln(&info.CourseNumber)
	fmt.Println("请输入成绩[输入-1为未出]")
	fmt.Scanf("%d", &info.Score)
	err := service.UpdateStudentGrade(info)
	if err != nil {
		errorHandler(err)
		return
	} else {
		fmt.Println("成功执行")
	}
}

func PrintDeptScoreHandler() {
	var deptName, courseNumber string
	fmt.Println("请输入系名：")
	fmt.Scanln(&deptName)
	fmt.Println("请输入课程号：")
	fmt.Scanln(&courseNumber)
	info, err := service.GetDeptGradeInfo(deptName, courseNumber)
	if err != nil {
		errorHandler(err)
		return
	}
	fmt.Println("系名\t平均分\t最高分\t最低分\t优秀率\t不及格人数")
	fmt.Printf("%s\t%.1f\t%d\t%d\t%.1f%%\t%d\n", info.DeptName, info.Average, info.Best, info.Worst, info.GoodRate*100.0, info.FailedCount)
}

func PrintStudentRankHandler() {
	var deptName, courseNumber string
	fmt.Println("请输入系名：")
	fmt.Scanln(&deptName)
	fmt.Println("请输入课程号：")
	fmt.Scanln(&courseNumber)
	rankList, err := service.GetDeptStudentScoreList(deptName, courseNumber)
	if err != nil {
		errorHandler(err)
		return
	}
	fmt.Println("课程信息")
	_, err = printCourseBasic(courseNumber)
	if err != nil {
		errorHandler(err)
		return
	}
	fmt.Printf("\n%s系排名情况\n", deptName)
	fmt.Println("排名\t学号\t\t姓名\t性别\t年龄\t奖学金\t成绩")
	for i, item := range rankList {
		fmt.Printf("%d\t%s\t%s\t%s\t%d\t%s\t%s\n", i+1, item.Number, item.Name, item.Sex, item.Age, item.Scholarship, stringScore(item.Score))
	}
}

func PrintStudentHandler() {
	var studentNumber string
	fmt.Println("请输入学生学号：")
	fmt.Scanln(&studentNumber)
	info, err := service.GetStudentInfo(studentNumber)
	if err != nil {
		errorHandler(err)
		return
	}
	fmt.Println("基本信息")
	fmt.Println("学号\t\t姓名\t性别\t年龄\t所在系\t奖学金")
	fmt.Printf("%s\t%s\t%s\t%d\t%s\t%s\n", info.Number, info.Name, info.Sex, info.Age, info.Dept, info.Scholarship)
	fmt.Println("选课信息")
	fmt.Println("课程号\t课程名\t\t先修课\t学分\t成绩")
	for _, item := range info.CourseList {
		fmt.Printf("%s\t%-8s\t%s\t%d\t%s\n", item.Number, item.Name, item.PreviousCourseNumber, item.Credit, stringScore(item.Score))
	}
}
