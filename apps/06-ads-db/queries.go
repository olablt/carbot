package main

import "github.com/olablt/carbot"

var queries = []*carbot.Query{
	{
		// MB Vito from:2010
		ID:             1,
		Name:           "MB Vito from:2010",
		AdvertiserName: "autoplius.lt",
		FirstURL:       "https://autoplius.lt/skelbimai/naudoti-automobiliai?make_id_list=67&engine_capacity_from=&engine_capacity_to=&power_from=&power_to=&kilometrage_from=&kilometrage_to=&has_damaged_id=&condition_type_id=&make_date_from=2010&make_date_to=&sell_price_from=&sell_price_to=&co2_from=&co2_to=&euro_id=&fk_place_countries_id=&qt=&number_of_seats_id=&number_of_doors_id=&gearbox_id=&steering_wheel_id=&is_partner=&older_not=&save_search=1&slist=2237451705&category_id=2&order_by=&order_direction=&make_id%5B67%5D=674",
	},
	{
		// Lexus RX 450h from:2017
		ID:             2,
		Name:           "Lexus RX 450h from:2017",
		AdvertiserName: "autoplius.lt",
		FirstURL:       "https://autoplius.lt/skelbimai/naudoti-automobiliai?make_date_from=2017&make_date_to=&sell_price_from=&sell_price_to=&engine_capacity_from=&engine_capacity_to=&power_from=&power_to=&kilometrage_from=&kilometrage_to=&qt=&steering_wheel_id=10922&category_id=2&make_id=72&model_id=19859&slist=2241756133",
	},
}
