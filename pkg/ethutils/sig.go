package ethutils

import (
	"bytes"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var ErrInvalidSignatureLength = fmt.Errorf("signature must be 65 bytes long")

func VerifySign(payload, sig []byte, address common.Address) error {
	if len(sig) != 65 {
		return ErrInvalidSignatureLength
	}

	if sig[64] == 27 || sig[64] == 28 {
		sig[64] -= 27
	}

	sigPublicKeyECDSA, err := crypto.SigToPub(payload, sig)
	if err != nil {
		return err
	}

	sigPublicKeyBytes := crypto.FromECDSAPub(sigPublicKeyECDSA)
	recoveredAccount := common.BytesToAddress(sigPublicKeyBytes)

	if !bytes.Equal(recoveredAccount.Bytes(), address.Bytes()) {
		return fmt.Errorf("invalid recovered key: %s vs %s", recoveredAccount.Hex(), address.Hex())
	}

	return nil
}

func VerifyEcRecover(payload, sig []byte, address common.Address) error {
	if len(sig) != 65 {
		return ErrInvalidSignatureLength
	}

	sigPublicKey, err := crypto.Ecrecover(payload, sig)
	if err != nil {
		return err
	}

	recoveredAccount := common.BytesToAddress(sigPublicKey)
	if !bytes.Equal(recoveredAccount.Bytes(), address.Bytes()) {
		return fmt.Errorf("invalid recovered key: %s vs %s", recoveredAccount.Hex(), address.Hex())
	}

	return nil
}

//func VerifySignNoRecoverId(payload, sig []byte, address common.Address) error {
//	if len(sig) != 65 {
//		return ErrInvalidSignatureLength
//	}
//
//	signatureNoRecoverID := sig[:len(sig)-1] // remove recovery id
//	if !crypto.VerifySignature(address.Bytes(), payload, signatureNoRecoverID) {
//		return fmt.Errorf("invalid signature")
//	}
//
//	return nil
//}
