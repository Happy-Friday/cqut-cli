package main

import (
	"flag"
	"github.com/Happy-Friday/zf-cqut/cqut"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"log"
)

const (
	Grades   = "grades"
	CTable   = "ctable"
	UserInfo = "userinfo"
	GPoint   = "gpoint"
	Photo    = "photo"
)

var (
	op = flag.String("tp", "help", `
	查询类型:
		grades 成绩
		ctable 课表
		userinfo 学生信息
		gpoint 绩点和学分总和
		photo 照片
	`)
	username  = flag.String("u", "", "登陆账号")
	password  = flag.String("p", "", "登陆密码")
	qusername = flag.String("qu", "", "要查询的用户名")
	year      = flag.String("y", "", "学年")
	term      = flag.String("t", "", "学期")
	fname     = flag.String("f", "", "文件名")
	plog      = flag.Bool("l", false, "是否打印log")
	//cache     = flag.Bool("c", false, "是否缓存链接")
)

var cq *cqut.Cqut

func stringJson(m interface{}) string {
	buf, err := json.MarshalIndent(m, "", "	");
	if err != nil {
		return ""
	}
	return string(buf)
}

func getGrades() string {
	var result []map[string]string
	if *year != "" {
		if *term != "" {
			result = cq.GetGrades(true, *year, *term)
		} else {
			result = cq.GetGrades(true, *year)
		}
	} else {
		result = cq.GetGrades(true)
	}
	return stringJson(result)
}

func getGradesPoint() string {
	var result map[string]string
	if *year != "" {
		if *term != "" {
			result = cq.GetGradesPoint(true, *year, *term)
		} else {
			result = cq.GetGradesPoint(true, *year)
		}
	} else {
		result = cq.GetGradesPoint(true)
	}
	return stringJson(result)
}
func getUserInfo() string {
	result := cq.GetUserInfo()
	return stringJson(result)
}

func getCoursesTable() string {
	var result map[string]interface{}
	if *year != "" {
		if *term != "" {
			result = cq.GetCoursesTable(true, *year, *term)
		} else {
			result = cq.GetCoursesTable(true, *year)
		}
	} else {
		result = cq.GetCoursesTable(true)
	}
	return stringJson(result)
}

func savePhoto() {
	log.Println(stringJson(cq))
	if *fname != "" {
		if *qusername != "" {
			ioutil.WriteFile(*fname, cq.GetPhoto(*qusername), os.ModePerm)
			return
		}
		ioutil.WriteFile(*fname, cq.GetPhoto(), os.ModePerm)
	}
	fmt.Println("无效的文件名")
}

func createObject() {
	cq = cqut.NewCqut(*username, *password)
	err := cq.Initialize()
	if err != nil {
		panic(err)
	}
}

//func saveObject(fname string) {
//	file, err := os.Create(fname)
//	if err != nil {
//		return
//	}
//	enc := gob.NewEncoder(file)
//	enc.Encode(cq)
//	file.Close()
//}
//
//func readObject(fname string) error {
//	file, err := os.Open(fname)
//	if err != nil {
//		return err
//	}
//	dec := gob.NewDecoder(file)
//	dec.Decode(&cq)
//	file.Close()
//	return nil
//}

func main() {
	flag.Parse()
	if len(os.Args) == 1 {
		flag.Usage()
		return
	}
	if *plog {
		cqut.SetLogOk(true)
	}

	createObject()

	switch(*op) {
	case Grades:
		fmt.Println(getGrades())
	case UserInfo:
		fmt.Println(getUserInfo())
	case GPoint:
		fmt.Println(getGradesPoint())
	case CTable:
		fmt.Println(getCoursesTable())
	case Photo:
		savePhoto()
	}

}
