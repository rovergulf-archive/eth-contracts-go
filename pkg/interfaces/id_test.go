package interfaces

import (
	"testing"
)

func TestHex(t *testing.T) {
	t.Run("test_interface_ids", func(t *testing.T) {
		if ERC165.Hex() != "0x01ffc9a7" {
			t.Fatalf("Invalid hex execution")
		}
		if ERC721.Hex() != "0x80ac58cd" {
			t.Fatalf("Invalid hex execution")
		}
		if ERC1155.Hex() != "0xd9b67a26" {
			t.Fatalf("Invalid hex execution")
		}
	})
}
