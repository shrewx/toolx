package constant

//go:generate toolx gen enum City

type City string

const (
	CITY__BEIJING  City = "beijing"  // 医生
	CITY__SHANGHAI City = "shanghai" // 教师
	CITY__CHENGDU  City = "chengdu"  // 司机
)
