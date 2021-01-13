package sync

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"

	"github.com/gogo/protobuf/proto"
	"github.com/pkg/errors"
	eth "github.com/prysmaticlabs/ethereumapis/eth/v1alpha1"
	"github.com/prysmaticlabs/prysm/beacon-chain/cache"
	"github.com/prysmaticlabs/prysm/beacon-chain/core/feed"
	"github.com/prysmaticlabs/prysm/beacon-chain/core/feed/operation"
	"github.com/prysmaticlabs/prysm/beacon-chain/core/helpers"
	"github.com/prysmaticlabs/prysm/shared/attestationutil"
	"github.com/prysmaticlabs/prysm/shared/params"
	"github.com/prysmaticlabs/prysm/shared/sliceutil"
)

func (s *Service) committeeIndexBeaconAttestationSubscriber(ctx context.Context, msg proto.Message) error {
	a, ok := msg.(*eth.Attestation)
	if !ok {
		return fmt.Errorf("message was not type *eth.Attestation, type=%T", msg)
	}

	if a.Data == nil {
		return errors.New("nil attestation")
	}
	s.setSeenCommitteeIndicesSlot(a.Data.Slot, a.Data.CommitteeIndex, a.AggregationBits)

	exists, err := s.attPool.HasAggregatedAttestation(a)
	if err != nil {
		return errors.Wrap(err, "Could not determine if attestation pool has this atttestation")
	}
	if exists {
		return nil
	}

	// Broadcast the unaggregated attestation on a feed to notify other services in the beacon node
	// of a received unaggregated attestation.
	s.attestationNotifier.OperationFeed().Send(&feed.Event{
		Type: operation.UnaggregatedAttReceived,
		Data: &operation.UnAggregatedAttReceivedData{
			Attestation: a,
		},
	})

	// for testing only
	s.logWireAttestationsForTesting(ctx, a)

	return s.attPool.SaveUnaggregatedAttestation(a)
}

func (s *Service) logWireAttestationsForTesting(ctx context.Context, att *eth.Attestation) {
	state, err := s.chain.AttestationPreState(ctx, att)
	if err != nil {
		log.WithError(err).Error("committeeIndexBeaconAttestationSubscriber: could not fetch pre attestation state")
	}
	committee, err := helpers.BeaconCommitteeFromState(state, att.Data.Slot, att.Data.CommitteeIndex)
	if err != nil {
		log.WithError(err).Error("committeeIndexBeaconAttestationSubscriber: could not fetch committee")
	}
	indexedAtt := attestationutil.ConvertToIndexed(ctx, att, committee)
	attRaw, err := json.Marshal(indexedAtt)
	if err != nil {
		log.WithError(err).Error("committeeIndexBeaconAttestationSubscriber: failed to marshal attestation")
	}
	logrus.WithField("indexedAttRaw", string(attRaw)).Info("committeeIndexBeaconAttestationSubscriber: got attestation object")
}

func (s *Service) persistentSubnetIndices() []uint64 {
	return cache.SubnetIDs.GetAllSubnets()
}

func (s *Service) aggregatorSubnetIndices(currentSlot uint64) []uint64 {
	endEpoch := helpers.SlotToEpoch(currentSlot) + 1
	endSlot := endEpoch * params.BeaconConfig().SlotsPerEpoch
	var commIds []uint64
	for i := currentSlot; i <= endSlot; i++ {
		commIds = append(commIds, cache.SubnetIDs.GetAggregatorSubnetIDs(i)...)
	}
	return sliceutil.SetUint64(commIds)
}

func (s *Service) attesterSubnetIndices(currentSlot uint64) []uint64 {
	endEpoch := helpers.SlotToEpoch(currentSlot) + 1
	endSlot := endEpoch * params.BeaconConfig().SlotsPerEpoch
	var commIds []uint64
	for i := currentSlot; i <= endSlot; i++ {
		commIds = append(commIds, cache.SubnetIDs.GetAttesterSubnetIDs(i)...)
	}
	return sliceutil.SetUint64(commIds)
}
