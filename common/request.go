package common

//GenerateMasterKeyRequest ...
type GenerateMasterKeyRequest struct {
	Password string `json:"password,omitempty"`
}

//AddNewWalletRequest ...
type AddNewWalletRequest struct {
	MasterPublicKey string `json:"master_public_key,omitempty"`
}

//RetrieveAddressRequest ...
type RetrieveAddressRequest struct {
	MasterPublicKey string `json:"master_public_key,omitempty"`
	Currency        int    `json:"currency,omitempty"`
	WalletID        string `json:"wallet_id,omitempty"`
	Token           string `json:"token,omitempty"`
	//AccountID       string `json:"account_id,omitempty"`
}
