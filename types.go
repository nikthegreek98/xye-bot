package xyebot

// Response to Telegram
type Response struct {
	Chatid int64  `json:"chat_id"`
	Text   string `json:"text"`
	Method string `json:"method"`
}

// Chat Telegram structure
type Chat struct {
	ID int64 `json: "chat_id"`
}

// Message Telegram structure
type Message struct {
	Chat *Chat  `json:"chat"`
	Text string `json:"text"`
}

// Update - outer Telegram structure
type Update struct {
	Message *Message `json:"message"`
}

// DatastoreDelay type for DataStore
type DatastoreDelay struct {
	Delay int
}

// DatastoreGentle type for DataStore
type DatastoreGentle struct {
	Gentle bool
}
