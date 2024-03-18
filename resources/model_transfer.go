/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Transfer struct {
	Key
	Attributes *TransferAttributes `json:"attributes,omitempty"`
}
type TransferResponse struct {
	Data     Transfer `json:"data"`
	Included Included `json:"included"`
}

type TransferListResponse struct {
	Data     []Transfer `json:"data"`
	Included Included   `json:"included"`
	Links    *Links     `json:"links"`
}

// MustTransfer - returns Transfer from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustTransfer(key Key) *Transfer {
	var transfer Transfer
	if c.tryFindEntry(key, &transfer) {
		return &transfer
	}
	return nil
}
