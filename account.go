package anaconda

import (
	"net/url"
)

// Verify the credentials by making a very small request
func (a TwitterApi) VerifyCredentials() (ok bool, err error) {
	v := cleanValues(nil)
	v.Set("include_entities", "false")
	v.Set("skip_status", "true")

	_, err = a.GetSelf(v)
	return err == nil, err
}

// Get the user object for the authenticated user. Requests /account/verify_credentials
func (a TwitterApi) GetSelf(v url.Values) (u User, err error) {
	v = cleanValues(v)
	response_ch := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/account/verify_credentials.json", v, &u, _GET, response_ch}
	return u, (<-response_ch).err
}

func (a TwitterApi) UpdateScreenName(v url.Values) (p Profile, err error) {
	v = cleanValues(v)
	response_ch := make(chan response)
	a.queryQueue <- query{a.baseUrl + "//account/update_profile.json", v, &p, _POST, response_ch}
	return p, (<-response_ch).err
}

type Profile struct {
	ContributorsEnabled            bool   `json:"contributors_enabled"`
	CreatedAt                      string `json:"created_at"`
	DefaultProfile                 bool   `json:"default_profile"`
	DefaultProfileImage            bool   `json:"default_profile_image"`
	Description                    string `json:"description"`
	FavouritesCount                int    `json:"favourites_count"`
	FollowRequestSent              bool   `json:"follow_request_sent"`
	FollowersCount                 int    `json:"followers_count"`
	Following                      bool   `json:"following"`
	FriendsCount                   int    `json:"friends_count"`
	GeoEnabled                     bool   `json:"geo_enabled"`
	ID                             int    `json:"id"`
	IDStr                          string `json:"id_str"`
	IsTranslator                   bool   `json:"is_translator"`
	Lang                           string `json:"lang"`
	ListedCount                    int    `json:"listed_count"`
	Location                       string `json:"location"`
	Name                           string `json:"name"`
	Notifications                  bool   `json:"notifications"`
	ProfileBackgroundColor         string `json:"profile_background_color"`
	ProfileBackgroundImageURL      string `json:"profile_background_image_url"`
	ProfileBackgroundImageURLHTTPS string `json:"profile_background_image_url_https"`
	ProfileBackgroundTile          bool   `json:"profile_background_tile"`
	ProfileImageURL                string `json:"profile_image_url"`
	ProfileImageURLHTTPS           string `json:"profile_image_url_https"`
	ProfileLinkColor               string `json:"profile_link_color"`
	ProfileSidebarBorderColor      string `json:"profile_sidebar_border_color"`
	ProfileSidebarFillColor        string `json:"profile_sidebar_fill_color"`
	ProfileTextColor               string `json:"profile_text_color"`
	ProfileUseBackgroundImage      bool   `json:"profile_use_background_image"`
	Protected                      bool   `json:"protected"`
	ScreenName                     string `json:"screen_name"`
	ShowAllInlineMedia             bool   `json:"show_all_inline_media"`
	StatusesCount                  int    `json:"statuses_count"`
	TimeZone                       string `json:"time_zone"`
	URL                            string `json:"url"`
	UtcOffset                      int    `json:"utc_offset"`
	Verified                       bool   `json:"verified"`
}
