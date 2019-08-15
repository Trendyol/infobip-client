package infobip

type Message struct {
	Sender     string      `json:"sender"`
	Text       string      `json:"text"`
	Datacoding string      `json:"datacoding"`
	Nli        string      `json:"nli"`
	Type       string      `json:"type"`
	Recipients []Recipient `json:"recipients"`
}

func NewMessage(sender, text, datacoding, nli, typeValue string, recipients []Recipient) *Message {
	return &Message{Sender: sender, Text: text, Datacoding: datacoding, Nli: nli, Type: typeValue, Recipients: recipients}
}

type Recipient struct {
	Gsm string `json:"gsm"`
}

type Response struct {
	Result []Result `json:"results"`
}

type Result struct {
	Status      string `json:"status"`
	MessageID   string `json:"messageid"`
	Destination string `json:"destination"`
}

type Sender interface {
	Send(request []Message) (*Response, error)
}

