// Code generated by fastssz. DO NOT EDIT.
package phase0

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the ETH1Data object
func (e *ETH1Data) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(e)
}

// MarshalSSZTo ssz marshals the ETH1Data object to a target array
func (e *ETH1Data) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'DepositRoot'
	dst = append(dst, e.DepositRoot[:]...)

	// Field (1) 'DepositCount'
	dst = ssz.MarshalUint64(dst, e.DepositCount)

	// Field (2) 'BlockHash'
	if len(e.BlockHash) != 32 {
		err = ssz.ErrBytesLength
		return
	}
	dst = append(dst, e.BlockHash...)

	return
}

// UnmarshalSSZ ssz unmarshals the ETH1Data object
func (e *ETH1Data) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 72 {
		return ssz.ErrSize
	}

	// Field (0) 'DepositRoot'
	copy(e.DepositRoot[:], buf[0:32])

	// Field (1) 'DepositCount'
	e.DepositCount = ssz.UnmarshallUint64(buf[32:40])

	// Field (2) 'BlockHash'
	if cap(e.BlockHash) == 0 {
		e.BlockHash = make([]byte, 0, len(buf[40:72]))
	}
	e.BlockHash = append(e.BlockHash, buf[40:72]...)

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the ETH1Data object
func (e *ETH1Data) SizeSSZ() (size int) {
	size = 72
	return
}

// HashTreeRoot ssz hashes the ETH1Data object
func (e *ETH1Data) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(e)
}

// HashTreeRootWith ssz hashes the ETH1Data object with a hasher
func (e *ETH1Data) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'DepositRoot'
	hh.PutBytes(e.DepositRoot[:])

	// Field (1) 'DepositCount'
	hh.PutUint64(e.DepositCount)

	// Field (2) 'BlockHash'
	if len(e.BlockHash) != 32 {
		err = ssz.ErrBytesLength
		return
	}
	hh.PutBytes(e.BlockHash)

	hh.Merkleize(indx)
	return
}
