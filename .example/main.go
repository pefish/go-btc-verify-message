package main

import (
	"fmt"

	"github.com/btcsuite/btcd/chaincfg"

	"github.com/bitonicnl/verify-signed-message/pkg"
)

func main() {
	// P2PKH
	fmt.Println(verifier.VerifyWithChain(verifier.SignedMessage{
		Address:   "156JQcYPQ1pLyUQNuQtcwoZ5J8FNZSAq4N",
		Message:   "Test123",
		Signature: "G2BB8Dsc5/zRhiHSFHbCrPZUIrgPOkJ8j17wNIK5EdTLas0RBy/IXWRjrkvYDm9b7jjEvvZ2YXrTE3U6tyC0eEs=",
	}, &chaincfg.MainNetParams))

	// P2WPKH
	fmt.Println(verifier.VerifyWithChain(verifier.SignedMessage{
		Address:   "bc1qq5w0x7hhvnfm509nvrddng2ufgxjsepvmcaryx",
		Message:   "Test123",
		Signature: "HExN/Z3c0Pvqj0Qqeryq77gLhyjs7nS5rdnFNqF65O7iHFoDawGJeWMDpbPUeih+T1LXFSHsEEyTvdeemYrBf8w=",
	}, &chaincfg.MainNetParams))

	// P2SH
	fmt.Println(verifier.VerifyWithChain(verifier.SignedMessage{
		Address:   "39fBhi58x4ExEx7WBpRHVVZs2H86jub63J",
		Message:   "Test123",
		Signature: "G5/ngG7QFsLe4vNdX5U7EarHt7ZnHENX8z8OzFJ5bblyHrTgHHI3I5ijToR6/gcVxm6NqSPywX6nU2ry5pV/0Ng=",
	}, &chaincfg.MainNetParams))

	// P2TR
	fmt.Println(verifier.VerifyWithChain(verifier.SignedMessage{
		Address:   "bc1pda09fk3x5fgg0rqjpklandzu8cpwmx3k8ypt4y5m5yy2c3rv837shjvhh6",
		Message:   "Test123",
		Signature: "GwIfz7CSX+UkoZd9QcfNTnFDUQCUgWvOnK+sdrep/8pLRF9N6juh6C3UfkO/NVJ1XAfI9ld7GnbKRMiI//MlUeY=",
	}, &chaincfg.MainNetParams))

	// Bitcoin Testnet3
	//fmt.Println(verifier.VerifyWithChain(verifier.SignedMessage{
	//	Address:   "tb1qr97cuq4kvq7plfetmxnl6kls46xaka78n2288z",
	//	Message:   "The outage comes at a time when bitcoin has been fast approaching new highs not seen since June 26, 2019.",
	//	Signature: "H/bSByRH7BW1YydfZlEx9x/nt4EAx/4A691CFlK1URbPEU5tJnTIu4emuzkgZFwC0ptvKuCnyBThnyLDCqPqT10=",
	//}, &chaincfg.TestNet3Params))
}
