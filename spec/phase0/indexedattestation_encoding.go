// Code generated by fastssz. DO NOT EDIT.
package phase0

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the IndexedAttestation object
func (i *IndexedAttestation) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(i)
}

// MarshalSSZTo ssz marshals the IndexedAttestation object to a target array
func (i *IndexedAttestation) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(228)

	// Offset (0) 'AttestingIndices'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(i.AttestingIndices) * 8

	// Field (1) 'Data'
	if i.Data == nil {
		i.Data = new(AttestationData)
	}
	if dst, err = i.Data.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (2) 'Signature'
	dst = append(dst, i.Signature[:]...)

	// Field (0) 'AttestingIndices'
	if len(i.AttestingIndices) > 2048 {
		err = ssz.ErrListTooBig
		return
	}
	for ii := 0; ii < len(i.AttestingIndices); ii++ {
		dst = ssz.MarshalUint64(dst, i.AttestingIndices[ii])
	}

	return
}

// UnmarshalSSZ ssz unmarshals the IndexedAttestation object
func (i *IndexedAttestation) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 228 {
		return ssz.ErrSize
	}

	tail := buf
	var o0 uint64

	// Offset (0) 'AttestingIndices'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	// Field (1) 'Data'
	if i.Data == nil {
		i.Data = new(AttestationData)
	}
	if err = i.Data.UnmarshalSSZ(buf[4:132]); err != nil {
		return err
	}

	// Field (2) 'Signature'
	copy(i.Signature[:], buf[132:228])

	// Field (0) 'AttestingIndices'
	{
		buf = tail[o0:]
		num, err := ssz.DivideInt2(len(buf), 8, 2048)
		if err != nil {
			return err
		}
		i.AttestingIndices = ssz.ExtendUint64(i.AttestingIndices, num)
		for ii := 0; ii < num; ii++ {
			i.AttestingIndices[ii] = ssz.UnmarshallUint64(buf[ii*8 : (ii+1)*8])
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the IndexedAttestation object
func (i *IndexedAttestation) SizeSSZ() (size int) {
	size = 228

	// Field (0) 'AttestingIndices'
	size += len(i.AttestingIndices) * 8

	return
}

// HashTreeRoot ssz hashes the IndexedAttestation object
func (i *IndexedAttestation) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(i)
}

// HashTreeRootWith ssz hashes the IndexedAttestation object with a hasher
func (i *IndexedAttestation) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'AttestingIndices'
	{
		if len(i.AttestingIndices) > 2048 {
			err = ssz.ErrListTooBig
			return
		}
		subIndx := hh.Index()
		for _, i := range i.AttestingIndices {
			hh.AppendUint64(i)
		}
		hh.FillUpTo32()
		numItems := uint64(len(i.AttestingIndices))
		hh.MerkleizeWithMixin(subIndx, numItems, ssz.CalculateLimit(2048, numItems, 8))
	}

	// Field (1) 'Data'
	if err = i.Data.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (2) 'Signature'
	hh.PutBytes(i.Signature[:])

	hh.Merkleize(indx)
	return
}
