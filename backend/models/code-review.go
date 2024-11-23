package models

type File struct{
	Name string `form:"name"`
	FileContent string `file:"file_upload"`
}