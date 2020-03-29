package display

//Address ...
type Address struct {
	ID       string `json:"id,omitempty"`
	Index    int    `json:"index"`
	Currency int    `json:"currency"`
	Address  string `json:"address,omitempty"`
}
