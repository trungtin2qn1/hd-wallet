package display

//Account ...
type Account struct {
	ID         string `json:"id,omitempty"`
	PrivateKey string `json:"private_key,omitempty"`
	PublicKey  string `json:"public_key,omitempty"`
	Mnemonic   string `json:"mnemonic,omitempty"`
}
