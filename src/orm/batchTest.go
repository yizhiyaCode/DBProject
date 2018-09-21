package main

import (
	"resource"
	"skynet"
)

func main() {
	var jobs []resource.Job

	job := resource.Job{
		Name:    "name",
		Context: "context",
	}
	//const shortForm = "20060102"
	//time1, _ := time.Parse(shortForm, "20180920")
	//time2, _ := time.Parse(shortForm, "20180921")
	jobs = append(jobs, job)
	jobs = append(jobs, job)
	jobs = append(jobs, job)
	jobs = append(jobs, job)
	jobs = append(jobs, job)

	//skynet.BatchInsert(jobs)
	//skynet.BatchInsert(jobs)

	msg := []string{}
	msg = append(msg, "hello")
	msg = append(msg, "hello2")
	msg = append(msg, "hello3")
	msg = append(msg, "hello4")

	skynet.InsertTask("11", "job", msg)

	skynet.Finish()
}
