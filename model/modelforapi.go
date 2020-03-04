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
