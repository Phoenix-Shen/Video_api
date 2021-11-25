package model

type RcvObj struct {
	Project  string
	Timeline []Timeline
}

type Timeline struct {
	Index    int
	OnLineId string
}
