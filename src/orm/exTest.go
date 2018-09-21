package main

import (
	"resource"
	"skynet"
	"time"
)

func main() {

	//var jobs []resource.Job

	job := &resource.Job{
		Name:    "name",
		Context: "context",
	}
	const shortForm = "20060102"
	time1, _ := time.Parse(shortForm, "20180920")
	time2, _ := time.Parse(shortForm, "20180921")
	//jobs = append(jobs,job)
	//go func() {skynet.SelectByName("name", time1, time2)}()
	//go func() {skynet.JobInsert(job)}()
	//死循环，目的不让主goroutine结束
	//for{
	//	time.Sleep(time.Second)
	//	fmt.Println("job------------------------", )
	//}
	//skynet.JobInsert(jobs)

	skynet.JobInsert(job)
	skynet.SelectByName("name", time1, time2)
	skynet.Finish()
	time.Sleep(10 * time.Second)
	skynet.JobInsert(job)
	//skynet.Finish()
}
