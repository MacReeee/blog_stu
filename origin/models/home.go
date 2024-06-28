package models

import "goblog/config"

type HomeData struct {
	config.Viewer
	Categorys []Category
	Posts []PostMore
	Total int
	Page int
	Pages []int
	PageEnd bool
}
