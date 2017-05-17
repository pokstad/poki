package poki

type MetaData struct {
	Title string
}

type Post struct {
	Path string
	Raw  []byte
	Meta MetaData
}

type PostRev struct {
	Post
	RevisionID string
}
