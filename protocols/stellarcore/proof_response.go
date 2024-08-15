package stellarcore

type ProofResponse struct {
	Ledger uint32 `json:"ledger"`
	Proof  string `json:"proof,omitempty"`
}
