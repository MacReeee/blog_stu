package models

type Category struct {
	Cid      int
	Name     string
	CreateAt string
	UpdateAt string
}

type CategoryData struct {
	HomeData
	CategoryName string
}
