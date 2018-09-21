package skynet

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"resource"
	"strings"
	"time"
)

var dbs *gorm.DB

func init() {
	fmt.Println("-------------------------开始")
	var err error
	//path := strings.Join([]string{"root", ":", "123456", "@tcp(", "192.168.26.19", ":", "30306", ")/",
	//	"k8s_skynet", "?charset=utf8&parseTime=true&loc=Local"}, "")
	path := strings.Join([]string{resource.UserName, ":", resource.Password, "@tcp(", resource.Ip, ":", resource.Port, ")/",
		resource.DbName, "?charset=utf8&parseTime=true&loc=Local"}, "")
	dbs, err = gorm.Open("mysql", path)
	fmt.Printf("%v  \n", path)
	dbs.DB().SetConnMaxLifetime(100)
	dbs.DB().SetMaxIdleConns(10)

	//验证连接
	if err != nil {
		fmt.Println("open database fail   \n")
		return
	}
	fmt.Printf("连接成功 \n")

	create(resource.Job{})
	create(resource.Task{})
	create(resource.Mail{})
}

//创建表
func create(i interface{}) bool {
	if !dbs.HasTable(i) {
		if err := dbs.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(i).
			Error; err != nil {
			return false
		}
	}
	return true
}
func BatchInsert(jobList []resource.Job) bool {

	tx := dbs.Begin()
	sql := "INSERT INTO jobs (id,created_at,updated_at,deleted_at,name,context)  VALUES (?,?,?,?,?,?)"

	//sqlStr := "INSERT INTO jobs (id,created_at,updated_at,deleted_at,name,context)  VALUES"
	//vals := []interface{}{}
	//const rowSQL = "(?,?,?,?,?,?)"
	//var inserts []string
	//for _, elem := range jobList {
	//	inserts = append(inserts, rowSQL)
	//	vals = append(vals, nil, nil, nil, nil, elem.Name, elem.Context)
	//}
	//sqlStr = sqlStr + strings.Join(inserts, ",")
	//fmt.Printf("sql      %s \n", sqlStr)
	//fmt.Printf("vals     %s\n", vals)
	var err interface{}
	for _, job := range jobList {
		err = tx.Exec(sql, 0, time.Now(), nil, nil, job.Name, job.Context).Error
		fmt.Println("--------------------------------------------v")
	}

	if err != nil {
		tx.Rollback()
		fmt.Printf("插入失败     %v \n", err)
		return false
	}
	fmt.Printf("插入成功\n")
	tx.Commit()
	return true
}

//插入task信息
func InsertTask(taskName string, jobName string, list []string) bool {
	//插入sql
	sql := "INSERT INTO tasks (id,created_at,updated_at,deleted_at,job_name,name,message)  VALUES (?,?,?,?,?,?,?)"
	//开启事务
	tx := dbs.Begin()
	var err interface{}
	for _, msg := range list {
		//插入数据
		err = tx.Exec(sql, 0, time.Now(), nil, nil, jobName, taskName, msg).Error
	}
	if err != nil {
		tx.Rollback()
		fmt.Printf("插入失败   %v \n", err)
		return false
	}
	fmt.Printf("插入成功\n")
	//关闭事务
	tx.Commit()
	return true
}

func JobInsert(job *resource.Job) bool {
	fmt.Println("-------------------------insert开始")
	tx := dbs.Begin()

	fmt.Printf("%v \n", job)
	if err := tx.Create(job).Error; err != nil {
		tx.Rollback()
		fmt.Errorf("%v\n", err)
		fmt.Printf("插入失败 \n")
		return false
	}
	tx.Commit()
	fmt.Printf("插入成功\n")
	return true
}

//
func SelectByName(name string, begin time.Time, end time.Time) (job []resource.Job) {
	tx := dbs.Begin()
	//error := tx.Where("name LIKE ? AND (created_at BETWEEN ? AND ? )",name,begin,end).Find(&job).Error
	error := tx.Where("created_at BETWEEN ? AND ? ", begin, end).Find(&job).Error
	if error != nil {
		tx.Rollback()
		return nil
	}

	tx.Commit()
	fmt.Printf("查询成功\n")
	return job

}

//func main() {
//
//	job := &resource.Job{
//		Name:       "name",
//		Context:    "context",
//	}
//	jobInsert(job)
//
//	fmt.Println("hello world")
//	const shortForm = "20060102"
//	time1,err := time.Parse(shortForm,"20180920")
//	time2,err := time.Parse(shortForm,"20180921")
//
//	if err!=nil {
//		return
//	}
//
//	//SelectById("name",time)
//	fmt.Printf("job------------------------%d\n",len(SelectByName("name",time1,time2)))
//	defer dbs.Close()
//}
func Finish() {
	dbs.Close()
}
