package carbot

import (
	"database/sql"
	"log"
	"reflect"
	"time"
)

type AdsDB struct {
	DB        *sql.DB
	TableName string
	// dataDir string
}

func (adsDB *AdsDB) TableExists() bool {
	var count int
	query := `SELECT COUNT(*) FROM sqlite_master WHERE type='table' AND name=?;`
	err := adsDB.DB.QueryRow(query, adsDB.TableName).Scan(&count)
	if err != nil {
		log.Printf("Error checking if table exists: %v", err)
		return false
	}
	return count > 0
}

// CreateTable creates the ads table in the database.
func (adsDB *AdsDB) CreateTable() error {
	_, err := adsDB.DB.Exec(`
		CREATE TABLE "ads" ( 
			"id" INTEGER, 
			"query_id" INTEGER NOT NULL,
			"title" TEXT NOT NULL,
			"subtitle" TEXT,
			"price" TEXT,
			"other" TEXT,
			"description" TEXT,
			"image_url" TEXT,
			"listed_time" DATETIME,
			"delisted_time" DATETIME,
			"scrape_time" DATETIME,
			"advertiser_ad_id" TEXT,
			"is_new" BOOLEAN DEFAULT TRUE,
			"deleted_time" DATETIME,
			PRIMARY KEY("id" AUTOINCREMENT)
		); 
		CREATE INDEX idx_ads_query_id_scrape_time ON ads (query_id, scrape_time);
	`)
	return err
}

// InsertOne inserts a single ad into the db.
func (adsDB *AdsDB) InsertOne(ad *Ad) (int, error) {
	if len(ad.ImageURLs) == 0 {
		ad.ImageURLs = append(ad.ImageURLs, "-")
	}
	res, err := adsDB.DB.Exec(
		`INSERT INTO ads(query_id, title, subtitle, price, other, description, image_url, listed_time, delisted_time, scrape_time, advertiser_ad_id, is_new, deleted_time)
		VALUES(? , ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		ad.QueryID, ad.Title, ad.Subtitle, ad.Price, ad.Other, ad.Description, ad.ImageURLs[0], ad.ListedTime, ad.DelistedTime, ad.ScrapeTime, ad.AdvertiserAdID, ad.IsNew, ad.DeletedTime)

	if err != nil {
		return 0, err
	}
	// get the id of the newly inserted row
	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return 0, err
	}
	return int(id), nil
}

// InsertMany inserts multiple ads into the db.
func (adsDB *AdsDB) InsertMany(ads []Ad) error {
	tx, err := adsDB.DB.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(
		`INSERT INTO ads(query_id, title, subtitle, price, other, description, image_url, listed_time, delisted_time, scrape_time, advertiser_ad_id, is_new, deleted_time)
		VALUES(? , ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	for _, ad := range ads {
		if len(ad.ImageURLs) == 0 {
			ad.ImageURLs = append(ad.ImageURLs, "-")
		}
		_, err = stmt.Exec(ad.QueryID, ad.Title, ad.Subtitle, ad.Price, ad.Other, ad.Description, ad.ImageURLs[0], ad.ListedTime, ad.DelistedTime, ad.ScrapeTime, ad.AdvertiserAdID, ad.IsNew, ad.DeletedTime)
		if err != nil {
			return err
		}
	}
	return tx.Commit()
}

// Delete the ad from the db.
func (adsDB *AdsDB) Delete(id int) error {
	_, err := adsDB.DB.Exec("DELETE FROM ads WHERE id = ?", id)
	return err
}

// DeleteByQueryID deletes all ads with the given query ID from the db.
func (adsDB *AdsDB) DeleteByQueryID(queryID int) error {
	_, err := adsDB.DB.Exec("DELETE FROM ads WHERE query_id = ?", queryID)
	return err
}

func (orig *Ad) merge(t Ad) {
	uValues := reflect.ValueOf(&t).Elem()
	oValues := reflect.ValueOf(orig).Elem()
	for i := 0; i < uValues.NumField(); i++ {
		uField := uValues.Field(i)
		oField := oValues.Field(i)

		// Ensure the original field can be set
		if oField.CanSet() {
			switch uFieldValue := uField.Interface().(type) {
			case int64:
				if uFieldValue != 0 {
					oField.SetInt(uFieldValue)
				}
			case string:
				if uFieldValue != "" {
					oField.SetString(uFieldValue)
				}
			case time.Time:
				// Check if the time is not the zero value
				if !uFieldValue.IsZero() {
					oField.Set(uField)
				}
				// Add more cases as necessary for other types
			}
		}
	}
}

// Update the ad by merging the changed fields to the original ad (like task update and merge)
func (adsDB *AdsDB) Update(ad Ad) error {
	// Get the existing state of the ad we want to update.
	orig, err := adsDB.GetAd(ad.ID)
	if err != nil {
		return err
	}
	orig.merge(ad)
	imageURL := orig.ImageURLs[0]
	if len(ad.ImageURLs) > 0 {
		imageURL = ad.ImageURLs[0]
	}
	_, err = adsDB.DB.Exec(
		`UPDATE ads 
		SET query_id = ?, title = ?, subtitle = ?, price = ?, other = ?, description = ?, image_url = ?, listed_time = ?, delisted_time = ?, scrape_time = ?, advertiser_ad_id = ?, is_new = ?, deleted_time = ?
		WHERE id = ?`,
		orig.QueryID, orig.Title, orig.Subtitle, orig.Price, orig.Other, orig.Description, imageURL, orig.ListedTime, orig.DelistedTime, orig.ScrapeTime, orig.AdvertiserAdID, orig.ID)
	return err
}

// GetAd returns the ad from the db.
func (adsDB *AdsDB) GetAd(id int) (*Ad, error) {
	ad := Ad{}
	imageURL := ""
	err := adsDB.DB.QueryRow("SELECT * FROM ads WHERE id = ?", id).
		Scan(
			&ad.ID,
			&ad.QueryID,
			&ad.Title,
			&ad.Subtitle,
			&ad.Price,
			&ad.Other,
			&ad.Description,
			&imageURL,
			&ad.ListedTime,
			&ad.DelistedTime,
			&ad.ScrapeTime,
			&ad.AdvertiserAdID,
			&ad.IsNew,
			&ad.DeletedTime,
		)
	ad.ImageURLs = append(ad.ImageURLs, imageURL)
	return &ad, err
}

// GetAds returns all ads from the db.
func (adsDB *AdsDB) GetAds() ([]Ad, error) {
	rows, err := adsDB.DB.Query("SELECT * FROM ads")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var ads []Ad
	for rows.Next() {
		ad := Ad{}
		imageURL := ""
		err = rows.Scan(
			&ad.ID,
			&ad.QueryID,
			&ad.Title,
			&ad.Subtitle,
			&ad.Price,
			&ad.Other,
			&ad.Description,
			&imageURL,
			&ad.ListedTime,
			&ad.DelistedTime,
			&ad.ScrapeTime,
			&ad.AdvertiserAdID,
			&ad.IsNew,
			&ad.DeletedTime,
		)
		ad.ImageURLs = append(ad.ImageURLs, imageURL)
		ads = append(ads, ad)
	}
	return ads, err
}

// GetAdsByQueryID returns all ads with the given query ID from the db.
func (adsDB *AdsDB) GetAdsByQueryID(queryID int) ([]Ad, error) {
	rows, err := adsDB.DB.Query("SELECT * FROM ads WHERE query_id = ?", queryID)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var ads []Ad
	for rows.Next() {
		ad := Ad{}
		imageURL := ""
		err = rows.Scan(
			&ad.ID,
			&ad.QueryID,
			&ad.Title,
			&ad.Subtitle,
			&ad.Price,
			&ad.Other,
			&ad.Description,
			&imageURL,
			&ad.ListedTime,
			&ad.DelistedTime,
			&ad.ScrapeTime,
			&ad.AdvertiserAdID,
			&ad.IsNew,
			&ad.DeletedTime,
		)
		ad.ImageURLs = append(ad.ImageURLs, imageURL)
		ads = append(ads, ad)
	}
	return ads, err
}

// func (t *AdsDB) getTasks() ([]task, error) {
// 	var tasks []task
// 	rows, err := t.db.Query("SELECT * FROM tasks")
// 	if err != nil {
// 		return tasks, fmt.Errorf("unable to get values: %w", err)
// 	}
// 	for rows.Next() {
// 		var task task
// 		err = rows.Scan(
// 			&task.ID,
// 			&task.Name,
// 			&task.Project,
// 			&task.Status,
// 			&task.Created,
// 		)
// 		if err != nil {
// 			return tasks, err
// 		}
// 		tasks = append(tasks, task)
// 	}
// 	return tasks, err
// }

// func (t *AdsDB) getTasksByStatus(status string) ([]task, error) {
// 	rows, err := t.db.Query("SELECT * FROM tasks WHERE status = ?", status)
// 	defer rows.Close()
// 	var tasks []task
// 	if err != nil {
// 		return tasks, fmt.Errorf("unable to get values: %w", err)
// 	}
// 	for rows.Next() {
// 		var task task
// 		err = rows.Scan(
// 			&task.ID,
// 			&task.Name,
// 			&task.Project,
// 			&task.Status,
// 			&task.Created,
// 		)
// 		if err != nil {
// 			return tasks, err
// 		}
// 		tasks = append(tasks, task)
// 	}
// 	return tasks, err
// }

// func (t *AdsDB) getTask(id uint) (task, error) {
// 	var task task
// 	err := t.db.QueryRow("SELECT * FROM tasks WHERE id = ?", id).
// 		Scan(
// 			&task.ID,
// 			&task.Name,
// 			&task.Project,
// 			&task.Status,
// 			&task.Created,
// 		)
// 	return task, err
// }
