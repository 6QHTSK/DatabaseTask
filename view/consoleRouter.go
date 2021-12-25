package view

import (
	"DatabaseTask/controller"
	"fmt"
)

func ConsoleRouter() {
	fmt.Println("欢迎使用学生管理系统")
	for true {
		fmt.Println("菜单")
		fmt.Println("1-插入学生\t2-修改学生\t3-新增课程\t4-修改课程")
		fmt.Println("5-录入成绩\t6-修改成绩\t7-院系成绩\t8-课程成绩")
		fmt.Println("9-学生成绩\t10-删除未选课未被依赖课程\t0-退出")
		var code int
		_, err := fmt.Scanf("%d", &code)
		if err != nil {
			fmt.Printf("输入错误：%s\t", err)
			continue
		}
		if code == 0 {
			break
		}
		switch code {
		case 1:
			controller.InsertStudentHandler()
		case 2:
			controller.UpdateStudentHandler()
		case 3:
			controller.InsertCourseHandler()
		case 4:
			controller.UpdateCourseHandler()
		case 5:
			controller.InsertStudentScoreHandler()
		case 6:
			controller.UpdateStudentScoreHandler()
		case 7:
			controller.PrintDeptScoreHandler()
		case 8:
			controller.PrintStudentRankHandler()
		case 9:
			controller.PrintStudentHandler()
		case 10:
			controller.DeleteNotChoiceCourseHandler()
		default:
			fmt.Println("未找到输入命令")
		}
	}
	fmt.Println("再见")
}
