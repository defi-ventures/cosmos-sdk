package types

import (
	fmt "fmt"

	"github.com/gogo/protobuf/proto"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
)

type (
	// Msg defines the interface a transaction message must fulfill.
	Msg interface {
		proto.Message

		// ValidateBasic does a simple validation check that
		// doesn't require access to any other information.
		ValidateBasic() error

		// Signers returns the addrs of signers that must sign.
		// CONTRACT: All signatures must be present to be valid.
		// CONTRACT: Returns addrs in some deterministic order.
		GetSigners() []AccAddress
	}

	// LegacyMsg defines the old interface a message must fulfill, containing
	// Amino signing method and legacy router info.
	// Deprecated: Please use `Msg` instead.
	LegacyMsg interface {
		Msg

		// Get the canonical byte representation of the Msg.
		GetSignBytes() []byte

		// Return the message type.
		// Must be alphanumeric or empty.
		Route() string

		// Returns a human-readable string for the message, intended for utilization
		// within tags
		Type() string
	}

	// Fee defines an interface for an application application-defined concrete
	// transaction type to be able to set and return the transaction fee.
	Fee interface {
		GetGas() uint64
		GetAmount() Coins
	}

	// Signature defines an interface for an application application-defined
	// concrete transaction type to be able to set and return transaction signatures.
	Signature interface {
		GetPubKey() cryptotypes.PubKey
		GetSignature() []byte
	}

	// Tx defines the interface a transaction must fulfill.
	Tx interface {
		// Gets the all the transaction's messages.
		GetMsgs() []Msg

		// ValidateBasic does a simple and lightweight validation check that doesn't
		// require access to any other information.
		ValidateBasic() error
	}

	// FeeTx defines the interface to be implemented by Tx to use the FeeDecorators
	FeeTx interface {
		Tx
		GetGas() uint64
		GetFee() Coins
		FeePayer() AccAddress
		FeeGranter() AccAddress
	}

	// Tx must have GetMemo() method to use ValidateMemoDecorator
	TxWithMemo interface {
		Tx
		GetMemo() string
	}

	// TxWithTimeoutHeight extends the Tx interface by allowing a transaction to
	// set a height timeout.
	TxWithTimeoutHeight interface {
		Tx

		GetTimeoutHeight() uint64
	}
)

// TxDecoder unmarshals transaction bytes
type TxDecoder func(txBytes []byte) (Tx, error)

// TxEncoder marshals transaction to bytes
type TxEncoder func(tx Tx) ([]byte, error)

// MsgName returns the protobuf MessageName of a sdk.Msg.
func MsgName(msg Msg) string {
	return proto.MessageName(msg)
}

// MsgTypeURL returns the TypeURL of a `sdk.Msg`.
func MsgTypeURL(msg Msg) string {
	return fmt.Sprintf("/%s", MsgName(msg))
}
