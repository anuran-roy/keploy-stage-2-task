package models

type JokeResponseFlags struct {
	Nsfw      bool `json:"nsfw"`
	Religious bool `json:"religious"`
	Political bool `json:"political"`
	Racist    bool `json:"racist"`
	Sexist    bool `json:"sexist"`
	Explicit  bool `json:"explicit"`
}

type JokeResponse struct {
	Errors    bool              `json:"error"`
	Category  string            `json:"category"`
	Joke_type string            `json:"type"`
	Joke      string            `json:"joke"`
	Flags     JokeResponseFlags `json:"flags"`
	Id        uint              `json:"id,uint"` // `json:"id,uint"`
	Safe      bool              `json:"safe"`
	Lang      string            `json:"lang"`
}

type JokeNotFound struct {
	Errors  bool   `json:"error"`
	Code    uint   `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
}
