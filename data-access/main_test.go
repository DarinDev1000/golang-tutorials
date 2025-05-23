package main

import (
	"testing"
)

func TestAddAlbum(t *testing.T) {
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}

	newAlbum := Album{
		Title:  "Test Title",
		Artist: "Test Artist",
		Price:  49.99,
	}

	// Insert Album
	albID, err := addAlbum(tx, newAlbum)
	if err != nil {
		t.Fatal(err)
	}

	// Fetch Album
	_, err = albumByID(albID)
	if err != nil {
		t.Fatal(err)
	}

	// if album.Title != newAlbum.Title || album.Artist != newAlbum.Artist || album.Price != newAlbum.Price {
	// 	t.Errorf(`New Album does not match inserted album %+v, %+v`, album, newAlbum)
	// }

	t.Cleanup(func() {
		tx.Rollback()

	})
}
