# TODO

- add separate caching for queries ads
    - save ads to cache file and append new ads when new ads are fetched
    - save new ads to ads file and append new ads when new ads are fetched, user can freely delete ads file or some ads from it
- API
    - get_ads - get stored ads from all advertisers
        - response is grouped by query:
        ```json
        {
            "query": "query",
            "ads": [
                {
                    "id": "ad_id",
                    "title": "ad_title",
                    "description": "ad_description",
                    "price": "ad_price",
                    "url": "ad_url",
                    "image": "ad_image"
                }
            ]
        }
        ```
    - get_queries - get stored queries
    - get_advertisers - get stored advertisers
    - get_queries_by_advertiser - get stored queries from specific advertiser
    - get_ads_by_advertiser - get stored ads from specific advertiser
    - get_ads_by_query - get stored ads from specific query
    - update_query - update query by id, if id is not found, create new query
    - delete_query - delete query by id
    - delete_ad - delete ad by id
    - delete_ads_by_query - delete all ads related to query
- Store ads images locally
- Add new advertiser "skelbiu.lt"
- docker image with go alpine

