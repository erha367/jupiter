package form

type CastInfo struct {
	ClassId       int64 `form:"classId" binding:"required"`
	ClientClassId int64 `form:"clientClassId"`
	EndTimestamp  int64 `form:"endTimestamp" binding:"required"`
}
