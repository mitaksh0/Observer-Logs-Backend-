package pkg

type Data struct {
	Time int64 `json:"time"`
	Log  Info  `json:"log"`
}

type Info struct {
	Body     string `json:"body"`
	Service  string `json:"service"`
	Severity string `json:"severity"`
}

type Response struct {
	Message interface{} `json:"message"`
}

var data []Data
