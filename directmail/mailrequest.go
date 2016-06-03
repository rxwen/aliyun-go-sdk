package directmail

// MailRequest struct represents a mail to be sent.
type MailRequest struct {
	Action         string
	AccountName    string
	ReplyToAddress bool
	AddressType    int
	ToAddress      string
	FromAlias      string
	Subject        string
	HtmlBody       string
	TextBody       string
}
