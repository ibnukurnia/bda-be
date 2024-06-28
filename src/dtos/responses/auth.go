package responses

type Auth struct {
	Token string `json:"token"`
}

type LdapResponse struct {
	StatusCode      int
	ResponseCode    string   `json:"responseCode"`
	ResponseMessage string   `json:"responseMessage"`
	ResponseData    LdapData `json:"responseData"`
}

type LdapData struct {
	Pernr string `json:"pernr"`
	Nip   string `json:"nip"`
	Name  string `json:"sname"`
}
