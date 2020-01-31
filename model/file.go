package model

type File_uploader struct {
	 id int
	 UploaderId string
	 FileId int
	 Uploadtime string
}

type File_downloader struct {
	id int
	DownloaderId string
	FileId int
	Uploadtime string
}

type File_collecter struct {
	id int
	CollecterId string
	FileId int
	Uploadtime string
}

