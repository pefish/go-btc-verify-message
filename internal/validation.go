package internal

import (
	"fmt"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
)

// ValidateP2PKH ensures that the passed P2PKH address matches the address generated from the public key hash, recovery flag and network.
func ValidateP2PKH(pubKey *btcec.PublicKey, addr btcutil.Address, net *chaincfg.Params) (bool, error) {

	if p2pkhAddr, err := btcutil.NewAddressPubKeyHash(btcutil.Hash160(pubKey.SerializeCompressed()), net); err != nil {
		return false, err
	} else if addr.String() != p2pkhAddr.String() {
		return false, fmt.Errorf("generated address '%s' does not match expected address '%s'", p2pkhAddr.String(), addr.String())
	}

	return true, nil
}

// ValidateP2PKH ensures that the passed P2SH address matches the address generated from the public key hash, recovery flag and network.
func ValidateP2SH(pubKey *btcec.PublicKey, addr btcutil.Address, net *chaincfg.Params) (bool, error) {

	if scriptSig, err := txscript.NewScriptBuilder().AddOp(txscript.OP_0).AddData(btcutil.Hash160(pubKey.SerializeCompressed())).Script(); err != nil {
		return false, err
	} else if p2shAddr, err := btcutil.NewAddressScriptHash(scriptSig, net); err != nil {
		return false, err
	} else if addr.String() != p2shAddr.String() {
		return false, fmt.Errorf("generated address '%s' does not match expected address '%s'", p2shAddr.String(), addr.String())
	}

	return true, nil
}

// ValidateP2WPKH ensures that the passed P2WPKH address matches the address generated from the public key hash, recovery flag and network.
func ValidateP2WPKH(pubKey *btcec.PublicKey, addr btcutil.Address, net *chaincfg.Params) (bool, error) {

	if p2wkhAddr, err := btcutil.NewAddressWitnessPubKeyHash(btcutil.Hash160(pubKey.SerializeCompressed()), net); err != nil {
		return false, err
	} else if addr.String() != p2wkhAddr.String() {
		return false, fmt.Errorf("generated address '%s' does not match expected address '%s'", p2wkhAddr.String(), addr.String())
	}

	return true, nil
}

// ValidateP2TR ensures that the passed P2TR address matches the address generated from the public key hash, recovery flag and network.
func ValidateP2TR(pubKey *btcec.PublicKey, addr btcutil.Address, net *chaincfg.Params) (bool, error) {
	tapKey := txscript.ComputeTaprootKeyNoScript(pubKey)
	if p2trAddr, err := btcutil.NewAddressTaproot(schnorr.SerializePubKey(tapKey), net); err != nil {
		return false, err
	} else if addr.String() != p2trAddr.String() {
		return false, fmt.Errorf("generated address '%s' does not match expected address '%s'", p2trAddr.String(), addr.String())
	}

	return true, nil
}
