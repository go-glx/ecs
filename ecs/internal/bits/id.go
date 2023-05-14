package bits

// max components in world = 255(banks)*64(per_bank) = 16320
const maxBanks = 255  // hard-limit - uint8
const maxEntries = 64 // hard-limit - count bits in int64

type (
	bankID  uint8
	entryID uint8

	Id struct {
		bankID  bankID
		entryID entryID
	}
)

func (x Id) entryMask() int64 {
	mask := int64(1)
	mask = mask << int(x.entryID)
	return mask
}
