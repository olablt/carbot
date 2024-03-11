package carbot

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestDelete(t *testing.T) {
	tests := []struct {
		want Ad
	}{
		{
			want: Ad{
				ID:    1,
				Title: "get bmw",
				// Project: "groceries",
				// Status:  "todo",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.want.Title, func(t *testing.T) {
			adsDB := setup()
			defer teardown(adsDB)

			if _, err := adsDB.Insert(&tc.want); err != nil {
				t.Fatalf("unable to insert tasks: %v", err)
			}
			ads, err := adsDB.GetAds()
			if err != nil {
				t.Fatalf("unable to get tasks: %v", err)
			}
			tc.want.DelistedTime = ads[0].DelistedTime
			if !reflect.DeepEqual(tc.want, ads[0]) {
				t.Fatalf("got %v, want %v", tc.want, ads)
			}
			if err := adsDB.Delete(1); err != nil {
				t.Fatalf("unable to delete tasks: %v", err)
			}
			ads, err = adsDB.GetAds()
			if err != nil {
				t.Fatalf("unable to get ads: %v", err)
			}
			if len(ads) != 0 {
				t.Fatalf("expected ads to be empty, got: %v", ads)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	tests := []struct {
		new  *Ad
		old  *Ad
		want Ad
	}{
		{
			old: &Ad{
				ID:       1,
				Title:    "get bmw",
				Subtitle: "get fast",
			},
			new: &Ad{
				ID:       1,
				Title:    "get mercedes",
				Subtitle: "",
			},
			want: Ad{
				ID:       1,
				Title:    "get mercedes",
				Subtitle: "get fast",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.new.Title, func(t *testing.T) {
			adsDB := setup()
			defer teardown(adsDB)
			if _, err := adsDB.Insert(tc.old); err != nil {
				t.Fatalf("we ran into an unexpected error: %v", err)
			}
			if err := adsDB.Update(*tc.new); err != nil {
				t.Fatalf("we ran into an unexpected error: %v", err)
			}
			ad, err := adsDB.GetAd(tc.want.ID)
			if err != nil {
				t.Fatalf("we ran into an unexpected error: %v", err)
			}
			if ad.Title != tc.want.Title {
				t.Fatalf("got: %v, want: %v", ad.Title, tc.want.Title)
			}
			if ad.Subtitle != tc.want.Subtitle {
				t.Fatalf("got: %v, want: %v", ad.Subtitle, tc.want.Subtitle)
			}
			// if ad.DelistedTime.IsZero() {
			// 	t.Fatalf("expected ad to be delisted, got: %v", ad)
			// }
			// tc.want.Created = Ad.Created
			// if !reflect.DeepEqual(Ad, tc.want) {
			// 	t.Fatalf("got: %#v, want: %#v", Ad, tc.want)
			// }
		})
	}
}

func setup() *AdsDB {
	path := filepath.Join(os.TempDir(), "test.db")
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err)
	}
	t := AdsDB{DB: db, TableName: "tasks", FilePath: path}
	if !t.TableExists() {
		err := t.CreateTable()
		if err != nil {
			log.Fatal(err)
		}
	}
	return &t
}

func teardown(adsDB *AdsDB) {
	adsDB.DB.Close()
	os.Remove(adsDB.FilePath)
}
