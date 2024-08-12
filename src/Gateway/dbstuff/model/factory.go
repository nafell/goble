package model

func Int32Ptr(int int) *int32 {
	i := int32(int)
	return &i
}
