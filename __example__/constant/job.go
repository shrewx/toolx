package constant

//go:generate toolx gen enum Job

type Job int

const (
	JOB__DOCTOR  Job = iota // 医生
	JOB__TEACHER            // 教师
	JOB__DRIVER             // 司机
)
