package main

type Response struct {
	Item struct {
		ID         int    `json:"id"`
		AssetID    string `json:"asset_id"`
		VdpID      int    `json:"vdp_id"`
		Video360   int    `json:"video360"`
		SerieTitle string `json:"serie_title"`
		Title      string `json:"title"`
		Lead       string `json:"lead"`
		Content    string `json:"content"`
		Thumbnail  []struct {
			URL    string `json:"url"`
			Srcx   string `json:"srcx"`
			Srcy   string `json:"srcy"`
			Srcw   string `json:"srcw"`
			Srch   string `json:"srch"`
			RatioX int    `json:"ratioX"`
			RatioY int    `json:"ratioY"`
			Type   int    `json:"type"`
			Text   string `json:"text"`
		} `json:"thumbnail"`
		ThumbnailUrls []interface{} `json:"thumbnail_urls"`
		ChannelLogo   []struct {
			URL    string `json:"url"`
			Srcx   string `json:"srcx"`
			Srcy   string `json:"srcy"`
			Srcw   string `json:"srcw"`
			Srch   string `json:"srch"`
			RatioX int    `json:"ratioX"`
			RatioY int    `json:"ratioY"`
			Type   int    `json:"type"`
			Text   string `json:"text"`
		} `json:"channel_logo"`
		BrandingZonePhoto []struct {
			URL    string `json:"url"`
			Srcx   string `json:"srcx"`
			Srcy   string `json:"srcy"`
			Srcw   string `json:"srcw"`
			Srch   string `json:"srch"`
			RatioX int    `json:"ratioX"`
			RatioY int    `json:"ratioY"`
			Type   int    `json:"type"`
			Text   string `json:"text"`
		} `json:"branding_zone_photo"`
		Episode              int           `json:"episode"`
		Season               int           `json:"season"`
		NextItemID           int           `json:"next_item_id"`
		RunTime              string        `json:"run_time"`
		EndCreditsStart      interface{}   `json:"end_credits_start"`
		HasProductPlacement  int           `json:"has_product_placement"`
		Rating               string        `json:"rating"`
		ShowRating           int           `json:"show_rating"`
		ShowProductPlacement int           `json:"show_product_placement"`
		ShowAdds             int           `json:"show_adds"`
		StartDate            string        `json:"start_date"`
		EndDate              string        `json:"end_date"`
		ReleaseYear          interface{}   `json:"release_year"`
		Credits              []interface{} `json:"credits"`
		Keywords             []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
			Slug string `json:"slug"`
		} `json:"keywords"`
		TypeEpisode string `json:"type_episode"`
		Countries   []struct {
			ID    string `json:"id"`
			Alias string `json:"alias"`
			Name  string `json:"name"`
		} `json:"countries"`
		Payable                                int         `json:"payable"`
		Tmobile                                string      `json:"tmobile"`
		PackageContractors                     []string    `json:"package_contractors"`
		LimitPremiere                          int         `json:"limit_premiere"`
		Protected                              int         `json:"protected"`
		DevicesLimit                           int         `json:"devices_limit"`
		NumberingEpisodes                      int         `json:"numbering_episodes"`
		EpisodeOne                             int         `json:"episode_one"`
		BrandingzonePhotoApplicationHorizontal interface{} `json:"brandingzone_photo_application_horizontal"`
		BrandingzonePhotoApplicationVertical   interface{} `json:"brandingzone_photo_application_vertical"`
		Videos                                 struct {
			Main struct {
				SpritesPath  string   `json:"sprites_path"`
				AspectRatio  string   `json:"aspect_ratio"`
				Gstream      string   `json:"gstream"`
				GstreamID    string   `json:"gstream_id"`
				Adserver     string   `json:"adserver"`
				AdBreakes    []string `json:"ad_breakes"`
				VideoContent []struct {
					ProfileName string `json:"profile_name"`
					URL         string `json:"url"`
				} `json:"video_content"`
				VideoContentLicenseType     interface{} `json:"video_content_license_type"`
				VideoContentError           interface{} `json:"video_content_error"`
				HardwareDecoding            int         `json:"hardware_decoding"`
				VideoContentLicenserenewURL interface{} `json:"video_content_licenserenew_url"`
			} `json:"main"`
		} `json:"videos"`
		TrailerURL            interface{} `json:"trailer_url"`
		UserTmobileMsisdnHash interface{} `json:"user_tmobile_msisdn_hash"`
		OriginalDistribution  string      `json:"original_distribution"`
		OriginalInTv          string      `json:"original_in_tv"`
		Firstaired            string      `json:"firstaired"`
		IsPlayerplus          int         `json:"is_playerplus"`
		IsNcp                 int         `json:"is_ncp"`
		Logo                  string      `json:"logo"`
		TrailerAdserver       string      `json:"trailer_adserver"`
	} `json:"item"`
	Vistatus string `json:"vistatus"`
}
