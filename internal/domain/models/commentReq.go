package models

type CommentReq struct {
	FilmID  int32
	Title   string
	Content string
}

type Comment struct {
	Title   string
	Content string
}
