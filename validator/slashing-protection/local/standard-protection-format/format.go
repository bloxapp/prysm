// Package interchangeformat defines methods to parse, import, and export slashing protection data
// from a standard JSON file according to EIP-3076 https://eips.ethereum.org/EIPS/eip-3076. This format
// is critical to allow safe interoperability between eth2 clients.
package interchangeformat

import "github.com/sirupsen/logrus"

var log = logrus.WithField("prefix", "slashing-protection-format")

// INTERCHANGE_FORMAT_VERSION specified by https://eips.ethereum.org/EIPS/eip-3076.
// The version Prysm supports is version 5.
const INTERCHANGE_FORMAT_VERSION = "5"

// EIPSlashingProtectionFormat string representation of a standard
// format for representing validator slashing protection db data.
type EIPSlashingProtectionFormat struct {
	Metadata struct {
		InterchangeFormatVersion string `json:"interchange_format_version"`
		GenesisValidatorsRoot    string `json:"genesis_validators_root"`
	} `json:"metadata"`
	Data []*ProtectionData `json:"data"`
}

// ProtectionData field for the standard slashing protection format.
type ProtectionData struct {
	Pubkey             string               `json:"pubkey"`
	SignedBlocks       []*SignedBlock       `json:"signed_blocks"`
	SignedAttestations []*SignedAttestation `json:"signed_attestations"`
}

// SignedAttestation in the standard slashing protection format file, including
// a source epoch, target epoch, and an optional signing root.
type SignedAttestation struct {
	SourceEpoch string `json:"source_epoch"`
	TargetEpoch string `json:"target_epoch"`
	SigningRoot string `json:"signing_root,omitempty"`
}

// SignedBlock in the standard slashing protection format, including a slot
// and an optional signing root.
type SignedBlock struct {
	Slot        string `json:"slot"`
	SigningRoot string `json:"signing_root,omitempty"`
}
