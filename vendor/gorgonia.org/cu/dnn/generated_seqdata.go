package cudnn

/* WAS Generated by gencudnn. DO NOT EDIT */

// #include <cudnn.h>
import "C"
import (
	"runtime"
	"unsafe"
)

// SeqData is a representation of cudnnSeqDataDescriptor_t.
type SeqData struct {
	internal C.cudnnSeqDataDescriptor_t

	dataType           DataType
	nbDims             int
	dimA               []int
	axes               []SeqDataAxis
	seqLengthArraySize uintptr
	seqLengthArray     []int
	paddingFill        Memory
}

// NewSeqData creates a new SeqData.
func NewSeqData(dataType DataType, nbDims int, dimA []int, axes []SeqDataAxis, seqLengthArraySize uintptr, seqLengthArray []int, paddingFill Memory) (retVal *SeqData, err error) {
	var internal C.cudnnSeqDataDescriptor_t
	if err := result(C.cudnnCreateSeqDataDescriptor(&internal)); err != nil {
		return nil, err
	}

	dimAC, dimACManaged := ints2CIntPtr(dimA)
	defer returnManaged(dimACManaged)

	seqLengthArrayC, seqLengthArrayCManaged := ints2CIntPtr(seqLengthArray)
	defer returnManaged(seqLengthArrayCManaged)

	if err := result(C.cudnnSetSeqDataDescriptor(internal, dataType.C(), C.int(nbDims), dimAC, axes2Ptr(axes), C.size_t(seqLengthArraySize), seqLengthArrayC, unsafe.Pointer(paddingFill.Uintptr()))); err != nil {
		return nil, err
	}

	retVal = &SeqData{
		internal:           internal,
		dataType:           dataType,
		nbDims:             nbDims,
		dimA:               dimA,
		axes:               axes,
		seqLengthArraySize: seqLengthArraySize,
		seqLengthArray:     seqLengthArray,
		paddingFill:        paddingFill,
	}
	runtime.SetFinalizer(retVal, destroySeqData)
	return retVal, nil
}

// C() returns the internal cgo representation
func (s *SeqData) C() C.cudnnSeqDataDescriptor_t { return s.internal }

// DataType returns the internal dataType.
func (s *SeqData) DataType() DataType { return s.dataType }

// NbDims returns the internal nbDims.
func (s *SeqData) NbDims() int { return s.nbDims }

// DimA returns the internal dimA slice.
func (s *SeqData) DimA() []int { return s.dimA }

// Axes returns the internal axes.
func (s *SeqData) Axes() []SeqDataAxis { return s.axes }

// SeqLengthArraySize returns the internal seqLengthArraySize.
func (s *SeqData) SeqLengthArraySize() uintptr { return s.seqLengthArraySize }

// SeqLengthArray returns the internal `seqLengthArray` slice.
func (s *SeqData) SeqLengthArray() []int { return s.seqLengthArray }

// PaddingFill returns the internal paddingFill.
func (s *SeqData) PaddingFill() Memory { return s.paddingFill }

func destroySeqData(obj *SeqData) { C.cudnnDestroySeqDataDescriptor(obj.internal) }

/* UTIL */

func axes2Ptr(a []SeqDataAxis) (ptr *C.cudnnSeqDataAxis_t) {
	return (*C.cudnnSeqDataAxis_t)(unsafe.Pointer(&a[0]))
}
