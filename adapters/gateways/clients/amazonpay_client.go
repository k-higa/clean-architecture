package clients

type AmazonPayClient interface {
	Entry(req AmazonPayReq, apiKey string) (*AmazonPayRes, error)
}
type amazonPayClient struct {
}

func NewAmazonPayClient() AmazonPayClient {
	return &amazonPayClient{}
}

func (a *amazonPayClient) Entry(req AmazonPayReq, apiKey string) (*AmazonPayRes, error) {
	return &AmazonPayRes{ID: 1, Amount: 100, Status: "authorize"}, nil
}

type AmazonPayReq struct {
	ID     int
	Amount int
}

type AmazonPayRes struct {
	ID     int
	Amount int
	Status string
	Log
}
type Log struct {
	Host        string
	ContentType string
	RequestLog  string
	ResponseLog string
}
