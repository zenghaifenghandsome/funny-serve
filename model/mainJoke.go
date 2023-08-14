package model

type Joke struct {
	Content    string
	HashId     string
	Unixtime   string
	Updatetime string
}

type Results struct {
	Data []Joke
}

type JokeInfo struct {
	Error_code int
	Reason     string
	Result     Results
}
