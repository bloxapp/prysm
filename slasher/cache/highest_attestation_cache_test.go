package cache

import (
	"testing"

	ethereum_slashing "github.com/prysmaticlabs/prysm/proto/slashing"

	"github.com/prysmaticlabs/prysm/shared/testutil/require"
)

func TestStoringAndFetching(t *testing.T) {
	cache, err := NewHighestAttestationCache(10, nil)
	require.NoError(t, err)

	// cache
	cache.Set(1, &ethereum_slashing.HighestAttestation{
		ValidatorId:        1,
		HighestSourceEpoch: 2,
		HighestTargetEpoch: 3,
	})

	// has
	require.Equal(t, true, cache.Has(1))

	// fetch
	res, b := cache.Get(1)
	require.Equal(t, true, b)
	require.Equal(t, uint64(1), res[1].ValidatorId)
	require.Equal(t, uint64(2), res[1].HighestSourceEpoch)
	require.Equal(t, uint64(3), res[1].HighestTargetEpoch)

	// delete
	require.Equal(t, true, cache.Delete(1))
	// confirm
	res2, b2 := cache.Get(1)
	require.Equal(t, false, b2)
	require.Equal(t, true, res2 == nil)
}

func TestPurge(t *testing.T) {
	wasEvicted := false
	onEvicted := func(key interface{}, value interface{}) {
		wasEvicted = true
	}
	cache, err := NewHighestAttestationCache(10, onEvicted)
	require.NoError(t, err)

	// cache
	cache.Set(1, &ethereum_slashing.HighestAttestation{
		ValidatorId:        1,
		HighestSourceEpoch: 2,
		HighestTargetEpoch: 3,
	})
	cache.Set(2, &ethereum_slashing.HighestAttestation{
		ValidatorId:        4,
		HighestSourceEpoch: 5,
		HighestTargetEpoch: 6,
	})
	cache.Set(3, &ethereum_slashing.HighestAttestation{
		ValidatorId:        7,
		HighestSourceEpoch: 8,
		HighestTargetEpoch: 9,
	})

	// purge
	cache.Purge()

	require.Equal(t, false, cache.Has(1))
	require.Equal(t, false, cache.Has(2))
	require.Equal(t, false, cache.Has(3))

	require.Equal(t, true, wasEvicted)
}

func TestClear(t *testing.T) {
	wasEvicted := false
	onEvicted := func(key interface{}, value interface{}) {
		wasEvicted = true
	}
	cache, err := NewHighestAttestationCache(10, onEvicted)
	require.NoError(t, err)

	// cache
	cache.Set(1, &ethereum_slashing.HighestAttestation{
		ValidatorId:        1,
		HighestSourceEpoch: 2,
		HighestTargetEpoch: 3,
	})
	cache.Set(2, &ethereum_slashing.HighestAttestation{
		ValidatorId:        4,
		HighestSourceEpoch: 5,
		HighestTargetEpoch: 6,
	})
	cache.Set(3, &ethereum_slashing.HighestAttestation{
		ValidatorId:        7,
		HighestSourceEpoch: 8,
		HighestTargetEpoch: 9,
	})

	// purge
	cache.Clear()

	require.Equal(t, false, cache.Has(1))
	require.Equal(t, false, cache.Has(2))
	require.Equal(t, false, cache.Has(3))

	require.Equal(t, true, wasEvicted)
}
