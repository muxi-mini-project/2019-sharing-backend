package model

//type Res struct {
//	Msg  string `json:"message"`
//}

type Error struct {
	Msg string `json:"message"`
}

type Note struct {
	Hostid  string `json:"host_id"`
	Content string `json:"message"`
}

type Getmessage struct {
	Msg          string        `json:"message"`
	Collectarray []Messagelist `json:"message_list"`
}

type Messagelist struct {
	Writerid  string `json:"writer_id"`
	Content   string `json:"message"`
	Image_url string `json:"image_url"`
	Time      string `json:"write_time"`
}

type Getfile struct {
	Message  string `json:"message"`
	Fileinfo File   `json:"file"`
}

type Downloadfile struct {
	Message string `json:"message"`
	Fileurl string `json:"file_url"`
}

type Getfiles struct {
	Message string `json:"message"`
	Files   []File `json:"files"`
}
