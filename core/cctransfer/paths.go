package cctransfer

import (
	"path"
)

// Data store pathes.
const (
	pathCrossChannelTransfer = "/transfer/"                       // transfer - cross channel transfer
	pathTransferFrom         = pathCrossChannelTransfer + "from/" // f - From + ID
	pathTransferTo           = pathCrossChannelTransfer + "to/"   // t - To + ID
)

const (
	pathCrossChannelMultiTransfer = "/multi_transfer/"
	pathMultiTransferFrom         = pathCrossChannelMultiTransfer + "from/"
	pathMultiTransferTo           = pathCrossChannelMultiTransfer + "to/"
)

// Base returns the last element of path.
// Trailing slashes are removed before extracting the last element.
func Base(fullPath string) string {
	return path.Base(fullPath)
}

// CCFromTransfers returns path to store key.
func CCFromTransfers() string {
	return pathTransferFrom
}

// CCFromTransfer returns path to store key.
func CCFromTransfer(id string) string {
	return path.Join(CCFromTransfers(), id)
}

// CCToTransfers returns path to store key.
func CCToTransfers() string {
	return pathTransferTo
}

// CCToTransfer returns path to store key.
func CCToTransfer(id string) string {
	return path.Join(CCToTransfers(), id)
}

// CCFromMultiTransfers returns a path to a store key.
func CCFromMultiTransfers() string {
	return pathMultiTransferFrom
}

// CCFromMultiTransfer returns a path to a store key.
func CCFromMultiTransfer(id string) string {
	return path.Join(CCFromMultiTransfers(), id)
}

// CCToMultiTransfers returns a path to a store key.
func CCToMultiTransfers() string {
	return pathMultiTransferTo
}

// CCToMultiTransfer returns a path to a store key.
func CCToMultiTransfer(id string) string {
	return path.Join(CCToMultiTransfers(), id)
}
