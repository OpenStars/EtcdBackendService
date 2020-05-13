package checksig

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/crypto"
)

func Hash256(s string) []byte {
	h := sha256.New()
	h.Write([]byte(s))
	var out []byte
	out = h.Sum(nil)
	return out
}

func CheckSignature(pubKey string, message string, sig string) bool {
	aHash := Hash256(message)
	sigBin, err := hex.DecodeString(sig)
	if err != nil {
		return false
	}

	aRecoveredPub, errSig := crypto.SigToPub(aHash, sigBin)
	if errSig != nil {
		return false
	}

	compressedPub := crypto.CompressPubkey(aRecoveredPub)
	return pubKey == hex.EncodeToString(compressedPub)
}
