package datasource

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Categories struct {
	MenuUpdate       int `json:"menu_update"`
	GlobalCategoryID int `json:"global_category_id"`
	Data             []struct {
		ID         int `json:"id"`
		ExternalID any `json:"external_id"`
		Status     int `json:"status"`
		Name       struct {
			En string `json:"en"`
			Ar string `json:"ar"`
		} `json:"name"`
		CategoryMetaDescription struct {
			En any `json:"en"`
			Ar any `json:"ar"`
		} `json:"category_meta_description"`
		Items struct {
			CurrentPage int `json:"current_page"`
			Data        []struct {
				ID          int    `json:"id"`
				ExternalID  any    `json:"external_id"`
				CategoryID  int    `json:"category_id"`
				OrderType   int    `json:"order_type"`
				Price       int    `json:"price"`
				Weight      int    `json:"weight"`
				Unit        string `json:"unit"`
				Calories    any    `json:"calories"`
				Image       string `json:"image"`
				IsSize      int    `json:"is_size"`
				Status      int    `json:"status"`
				Type        int    `json:"type"`
				MinQuantity int    `json:"min_quantity"`
				OutOfStock  int    `json:"out_of_stock"`
				IsDonation  int    `json:"is_donation"`
				Name        struct {
					En string `json:"en"`
					Ar string `json:"ar"`
				} `json:"name"`
				Description struct {
					En string `json:"en"`
					Ar string `json:"ar"`
				} `json:"description"`
				Sort                int   `json:"sort"`
				Earn                int   `json:"earn"`
				Redeem              int   `json:"redeem"`
				IsModifiers         int   `json:"is_modifiers"`
				IsRequiredModifiers int   `json:"is_required_modifiers"`
				IsExclusions        int   `json:"is_exclusions"`
				Modifiers           []any `json:"modifiers"`
				Exclusions          []any `json:"exclusions"`
				Sizes               []struct {
					ID         int    `json:"id"`
					ExternalID any    `json:"external_id"`
					Price      int    `json:"price"`
					Weight     int    `json:"weight"`
					Unit       string `json:"unit"`
					Calories   any    `json:"calories"`
					Name       struct {
						En string `json:"en"`
						Ar string `json:"ar"`
					} `json:"name"`
				} `json:"sizes"`
				Images             []any `json:"images"`
				UpsalesItems       []any `json:"upsales_items"`
				CategoryExternalID any   `json:"category_external_id"`
			} `json:"data"`
			FirstPageURL string `json:"first_page_url"`
			From         int    `json:"from"`
			LastPage     int    `json:"last_page"`
			LastPageURL  string `json:"last_page_url"`
			NextPageURL  any    `json:"next_page_url"`
			Path         string `json:"path"`
			PerPage      int    `json:"per_page"`
			PrevPageURL  any    `json:"prev_page_url"`
			To           int    `json:"to"`
			Total        int    `json:"total"`
		} `json:"items"`
		Description struct {
			En string `json:"en"`
			Ar string `json:"ar"`
		} `json:"description"`
	} `json:"data"`
}

var categories *Categories

func init() {
	callApi()
	periodicallyCleanCache()
}

func periodicallyCleanCache() {
	go func() {
		for {
			categories = nil
			time.Sleep(48 * time.Hour)
		}
	}()
}

func callApi() {
	fmt.Println("calling Api")

	url := "https://shawarma-house.my.taker.io/api/v4/categories/products?page=1&per_page=1000"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}
	req.Header.Add("accept-language", "ar")

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&categories)
	if err != nil {
		panic(err)
	}

	count := 0
	for _, cat := range categories.Data {
		for _ = range cat.Items.Data {
			count++
		}
	}

	fmt.Printf("%d item loaded successfully\n", count)
}

func GetCategories() *Categories {
	if categories == nil {
		callApi()
	}
	return categories
}
