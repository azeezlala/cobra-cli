package customer

type CustomerDetail struct {
	Id            int    `json:"id"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	Gender        string `json:"gender"`
	AccountNumber string `json:"accountNumber"`
	BvnNumber     string `json:"bvnNumber"`
	Pin           int    `json:"pin"`
}
