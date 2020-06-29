// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/skv2/api/multicluster/v1alpha1/cluster.proto

package v1alpha1

import (
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"hash/fnv"

	"github.com/mitchellh/hashstructure"
	safe_hasher "github.com/solo-io/protoc-gen-ext/pkg/hasher"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = new(hash.Hash64)
	_ = fnv.New64
	_ = hashstructure.Hash
	_ = new(safe_hasher.SafeHasher)
)

// Hash function
func (m *KubernetesClusterSpec) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("multicluster.solo.io.github.com/solo-io/skv2/pkg/api/multicluster.solo.io/v1alpha1.KubernetesClusterSpec")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetSecretName())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetClusterDomain())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *KubernetesClusterStatus) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("multicluster.solo.io.github.com/solo-io/skv2/pkg/api/multicluster.solo.io/v1alpha1.KubernetesClusterStatus")); err != nil {
		return 0, err
	}

	for _, v := range m.GetStatus() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}