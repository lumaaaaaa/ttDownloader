package main

type SearchData struct {
	StatusCode int `json:"status_code"`
	AwemeList  []struct {
		AwemeId    string `json:"aweme_id"`
		Desc       string `json:"desc"`
		CreateTime int    `json:"create_time"`
		Author     struct {
			Uid          string `json:"uid"`
			ShortId      string `json:"short_id"`
			Nickname     string `json:"nickname"`
			Gender       int    `json:"gender"`
			Signature    string `json:"signature"`
			AvatarLarger struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"avatar_larger"`
			AvatarThumb struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"avatar_thumb"`
			AvatarMedium struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"avatar_medium"`
			Birthday            string `json:"birthday"`
			FollowStatus        int    `json:"follow_status"`
			AwemeCount          int    `json:"aweme_count"`
			FollowingCount      int    `json:"following_count"`
			FollowerCount       int    `json:"follower_count"`
			FavoritingCount     int    `json:"favoriting_count"`
			TotalFavorited      int    `json:"total_favorited"`
			IsBlock             bool   `json:"is_block"`
			HideSearch          bool   `json:"hide_search"`
			CustomVerify        string `json:"custom_verify"`
			UniqueId            string `json:"unique_id"`
			BindPhone           string `json:"bind_phone"`
			SpecialLock         int    `json:"special_lock"`
			NeedRecommend       int    `json:"need_recommend"`
			HasFacebookToken    bool   `json:"has_facebook_token"`
			HasTwitterToken     bool   `json:"has_twitter_token"`
			FbExpireTime        int    `json:"fb_expire_time"`
			TwExpireTime        int    `json:"tw_expire_time"`
			HasYoutubeToken     bool   `json:"has_youtube_token"`
			YoutubeExpireTime   int    `json:"youtube_expire_time"`
			RoomId              int    `json:"room_id"`
			LiveVerify          int    `json:"live_verify"`
			AuthorityStatus     int    `json:"authority_status"`
			VerifyInfo          string `json:"verify_info"`
			ShieldFollowNotice  int    `json:"shield_follow_notice"`
			ShieldDiggNotice    int    `json:"shield_digg_notice"`
			ShieldCommentNotice int    `json:"shield_comment_notice"`
			ShareInfo           struct {
				ShareUrl       string `json:"share_url"`
				ShareWeiboDesc string `json:"share_weibo_desc"`
				ShareDesc      string `json:"share_desc"`
				ShareTitle     string `json:"share_title"`
				ShareQrcodeUrl struct {
					Uri       string        `json:"uri"`
					UrlList   []interface{} `json:"url_list"`
					Width     int           `json:"width"`
					Height    int           `json:"height"`
					UrlPrefix interface{}   `json:"url_prefix"`
				} `json:"share_qrcode_url"`
				ShareTitleMyself           string      `json:"share_title_myself"`
				ShareTitleOther            string      `json:"share_title_other"`
				ShareDescInfo              string      `json:"share_desc_info"`
				NowInvitationCardImageUrls interface{} `json:"now_invitation_card_image_urls"`
			} `json:"share_info"`
			WithCommerceEntry      bool          `json:"with_commerce_entry"`
			VerificationType       int           `json:"verification_type"`
			EnterpriseVerifyReason string        `json:"enterprise_verify_reason"`
			IsAdFake               bool          `json:"is_ad_fake"`
			FollowersDetail        interface{}   `json:"followers_detail"`
			Region                 string        `json:"region"`
			AccountRegion          string        `json:"account_region"`
			CommerceUserLevel      int           `json:"commerce_user_level"`
			LiveAgreement          int           `json:"live_agreement"`
			PlatformSyncInfo       interface{}   `json:"platform_sync_info"`
			WithShopEntry          bool          `json:"with_shop_entry"`
			IsDisciplineMember     bool          `json:"is_discipline_member"`
			Secret                 int           `json:"secret"`
			HasOrders              bool          `json:"has_orders"`
			PreventDownload        bool          `json:"prevent_download"`
			ShowImageBubble        bool          `json:"show_image_bubble"`
			Geofencing             []interface{} `json:"geofencing"`
			UniqueIdModifyTime     int           `json:"unique_id_modify_time"`
			VideoIcon              struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"video_icon"`
			InsId               string `json:"ins_id"`
			GoogleAccount       string `json:"google_account"`
			YoutubeChannelId    string `json:"youtube_channel_id"`
			YoutubeChannelTitle string `json:"youtube_channel_title"`
			AppleAccount        int    `json:"apple_account"`
			WithFusionShopEntry bool   `json:"with_fusion_shop_entry"`
			IsPhoneBinded       bool   `json:"is_phone_binded"`
			AcceptPrivatePolicy bool   `json:"accept_private_policy"`
			TwitterId           string `json:"twitter_id"`
			TwitterName         string `json:"twitter_name"`
			UserCanceled        bool   `json:"user_canceled"`
			HasEmail            bool   `json:"has_email"`
			LiveAgreementTime   int    `json:"live_agreement_time"`
			Status              int    `json:"status"`
			CreateTime          int    `json:"create_time"`
			AvatarUri           string `json:"avatar_uri"`
			FollowerStatus      int    `json:"follower_status"`
			CommentSetting      int    `json:"comment_setting"`
			DuetSetting         int    `json:"duet_setting"`
			ReflowPageGid       int    `json:"reflow_page_gid"`
			ReflowPageUid       int    `json:"reflow_page_uid"`
			UserRate            int    `json:"user_rate"`
			DownloadSetting     int    `json:"download_setting"`
			DownloadPromptTs    int    `json:"download_prompt_ts"`
			ReactSetting        int    `json:"react_setting"`
			LiveCommerce        bool   `json:"live_commerce"`
			CoverUrl            []struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"cover_url"`
			Language            string      `json:"language"`
			HasInsights         bool        `json:"has_insights"`
			ShareQrcodeUri      string      `json:"share_qrcode_uri"`
			ItemList            interface{} `json:"item_list"`
			UserMode            int         `json:"user_mode"`
			UserPeriod          int         `json:"user_period"`
			IsStar              bool        `json:"is_star"`
			CvLevel             string      `json:"cv_level"`
			TypeLabel           interface{} `json:"type_label"`
			AdCoverUrl          interface{} `json:"ad_cover_url"`
			CommentFilterStatus int         `json:"comment_filter_status"`
			Avatar168X168       struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"avatar_168x168"`
			Avatar300X300 struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"avatar_300x300"`
			RelativeUsers           interface{} `json:"relative_users"`
			ChaList                 interface{} `json:"cha_list"`
			SecUid                  string      `json:"sec_uid"`
			NeedPoints              interface{} `json:"need_points"`
			NameField               string      `json:"name_field"`
			HomepageBottomToast     interface{} `json:"homepage_bottom_toast"`
			CanSetGeofencing        interface{} `json:"can_set_geofencing"`
			WhiteCoverUrl           interface{} `json:"white_cover_url"`
			UserTags                interface{} `json:"user_tags"`
			StitchSetting           int         `json:"stitch_setting"`
			BoldFields              interface{} `json:"bold_fields"`
			SearchUserName          string      `json:"search_user_name"`
			SearchUserDesc          string      `json:"search_user_desc"`
			SearchHighlight         interface{} `json:"search_highlight"`
			MutualRelationAvatars   interface{} `json:"mutual_relation_avatars"`
			Events                  interface{} `json:"events"`
			AdvanceFeatureItemOrder interface{} `json:"advance_feature_item_order"`
			MentionStatus           int         `json:"mention_status"`
			SpecialAccount          struct {
				SpecialAccountList interface{} `json:"special_account_list"`
			} `json:"special_account"`
			AdvancedFeatureInfo        interface{} `json:"advanced_feature_info"`
			UserProfileGuide           interface{} `json:"user_profile_guide"`
			ShieldEditFieldInfo        interface{} `json:"shield_edit_field_info"`
			FriendsStatus              int         `json:"friends_status"`
			CanMessageFollowStatusList interface{} `json:"can_message_follow_status_list"`
			AccountLabels              interface{} `json:"account_labels"`
			FakeDataInfo               struct {
			} `json:"fake_data_info"`
			IsMute   int `json:"is_mute"`
			QaStatus int `json:"qa_status,omitempty"`
		} `json:"author"`
		Music struct {
			Id      int64  `json:"id"`
			IdStr   string `json:"id_str"`
			Title   string `json:"title"`
			Author  string `json:"author"`
			Album   string `json:"album"`
			CoverHd struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"cover_hd"`
			CoverLarge struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"cover_large"`
			CoverMedium struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"cover_medium"`
			CoverThumb struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"cover_thumb"`
			PlayUrl struct {
				Uri       string        `json:"uri"`
				UrlList   []interface{} `json:"url_list"`
				Width     int           `json:"width"`
				Height    int           `json:"height"`
				UrlPrefix interface{}   `json:"url_prefix"`
			} `json:"play_url"`
			SchemaUrl                 string        `json:"schema_url"`
			SourcePlatform            int           `json:"source_platform"`
			StartTime                 int           `json:"start_time"`
			EndTime                   int           `json:"end_time"`
			Duration                  int           `json:"duration"`
			Extra                     string        `json:"extra"`
			UserCount                 int           `json:"user_count"`
			Position                  interface{}   `json:"position"`
			CollectStat               int           `json:"collect_stat"`
			Status                    int           `json:"status"`
			OfflineDesc               string        `json:"offline_desc"`
			OwnerId                   string        `json:"owner_id,omitempty"`
			OwnerNickname             string        `json:"owner_nickname"`
			IsOriginal                bool          `json:"is_original"`
			Mid                       string        `json:"mid"`
			BindedChallengeId         int           `json:"binded_challenge_id"`
			Redirect                  bool          `json:"redirect"`
			IsRestricted              bool          `json:"is_restricted"`
			AuthorDeleted             bool          `json:"author_deleted"`
			IsDelVideo                bool          `json:"is_del_video"`
			IsVideoSelfSee            bool          `json:"is_video_self_see"`
			OwnerHandle               string        `json:"owner_handle"`
			AuthorPosition            interface{}   `json:"author_position"`
			PreventDownload           bool          `json:"prevent_download"`
			PreventItemDownloadStatus int           `json:"prevent_item_download_status"`
			ExternalSongInfo          []interface{} `json:"external_song_info"`
			SecUid                    string        `json:"sec_uid,omitempty"`
			AvatarThumb               struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"avatar_thumb,omitempty"`
			AvatarMedium struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"avatar_medium,omitempty"`
			AvatarLarge struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"avatar_large,omitempty"`
			PreviewStartTime      float64       `json:"preview_start_time"`
			PreviewEndTime        int           `json:"preview_end_time"`
			IsCommerceMusic       bool          `json:"is_commerce_music"`
			IsOriginalSound       bool          `json:"is_original_sound"`
			AuditionDuration      int           `json:"audition_duration"`
			ShootDuration         int           `json:"shoot_duration"`
			ReasonType            int           `json:"reason_type"`
			Artists               []interface{} `json:"artists"`
			LyricShortPosition    interface{}   `json:"lyric_short_position"`
			MuteShare             bool          `json:"mute_share"`
			TagList               interface{}   `json:"tag_list"`
			DmvAutoShow           bool          `json:"dmv_auto_show"`
			IsAuthorArtist        bool          `json:"is_author_artist"`
			IsPgc                 bool          `json:"is_pgc"`
			IsMatchedMetadata     bool          `json:"is_matched_metadata"`
			IsAudioUrlWithCookie  bool          `json:"is_audio_url_with_cookie"`
			VideoDuration         int           `json:"video_duration"`
			SearchHighlight       interface{}   `json:"search_highlight"`
			MultiBitRatePlayInfo  interface{}   `json:"multi_bit_rate_play_info"`
			DurationHighPrecision struct {
				DurationPrecision         float64 `json:"duration_precision"`
				ShootDurationPrecision    float64 `json:"shoot_duration_precision"`
				AuditionDurationPrecision float64 `json:"audition_duration_precision"`
				VideoDurationPrecision    float64 `json:"video_duration_precision"`
			} `json:"duration_high_precision"`
			CanNotReuse         bool        `json:"can_not_reuse"`
			IsPlayMusic         bool        `json:"is_play_music"`
			CommercialRightType int         `json:"commercial_right_type"`
			TtToDspSongInfos    interface{} `json:"tt_to_dsp_song_infos"`
			IsShootingAllow     bool        `json:"is_shooting_allow"`
			RecommendStatus     int         `json:"recommend_status"`
			HasCommerceRight    bool        `json:"has_commerce_right"`
			MemeSongInfo        struct {
			} `json:"meme_song_info"`
			UncertArtists interface{} `json:"uncert_artists"`
			StrongBeatUrl struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"strong_beat_url,omitempty"`
			MatchedSong struct {
				Id          string `json:"id"`
				Author      string `json:"author"`
				Title       string `json:"title"`
				H5Url       string `json:"h5_url"`
				CoverMedium struct {
					Uri       string      `json:"uri"`
					UrlList   []string    `json:"url_list"`
					Width     int         `json:"width"`
					Height    int         `json:"height"`
					UrlPrefix interface{} `json:"url_prefix"`
				} `json:"cover_medium"`
				Performers interface{} `json:"performers"`
				ChorusInfo struct {
					StartMs    int `json:"start_ms"`
					DurationMs int `json:"duration_ms"`
				} `json:"chorus_info"`
				FullDuration int `json:"full_duration"`
			} `json:"matched_song,omitempty"`
			MusicReleaseInfo struct {
				IsNewReleaseSong bool `json:"is_new_release_song"`
				GroupReleaseDate int  `json:"group_release_date"`
			} `json:"music_release_info,omitempty"`
		} `json:"music"`
		ChaList []struct {
			Cid     string `json:"cid"`
			ChaName string `json:"cha_name"`
			Desc    string `json:"desc"`
			Schema  string `json:"schema"`
			Author  struct {
				FollowersDetail            interface{} `json:"followers_detail"`
				PlatformSyncInfo           interface{} `json:"platform_sync_info"`
				Geofencing                 interface{} `json:"geofencing"`
				CoverUrl                   interface{} `json:"cover_url"`
				ItemList                   interface{} `json:"item_list"`
				TypeLabel                  interface{} `json:"type_label"`
				AdCoverUrl                 interface{} `json:"ad_cover_url"`
				RelativeUsers              interface{} `json:"relative_users"`
				ChaList                    interface{} `json:"cha_list"`
				NeedPoints                 interface{} `json:"need_points"`
				HomepageBottomToast        interface{} `json:"homepage_bottom_toast"`
				CanSetGeofencing           interface{} `json:"can_set_geofencing"`
				WhiteCoverUrl              interface{} `json:"white_cover_url"`
				UserTags                   interface{} `json:"user_tags"`
				BoldFields                 interface{} `json:"bold_fields"`
				SearchHighlight            interface{} `json:"search_highlight"`
				MutualRelationAvatars      interface{} `json:"mutual_relation_avatars"`
				Events                     interface{} `json:"events"`
				AdvanceFeatureItemOrder    interface{} `json:"advance_feature_item_order"`
				AdvancedFeatureInfo        interface{} `json:"advanced_feature_info"`
				UserProfileGuide           interface{} `json:"user_profile_guide"`
				ShieldEditFieldInfo        interface{} `json:"shield_edit_field_info"`
				CanMessageFollowStatusList interface{} `json:"can_message_follow_status_list"`
				AccountLabels              interface{} `json:"account_labels"`
			} `json:"author"`
			UserCount int `json:"user_count"`
			ShareInfo struct {
				ShareUrl                   string      `json:"share_url"`
				ShareWeiboDesc             string      `json:"share_weibo_desc"`
				ShareDesc                  string      `json:"share_desc"`
				ShareTitle                 string      `json:"share_title"`
				BoolPersist                int         `json:"bool_persist"`
				ShareTitleMyself           string      `json:"share_title_myself"`
				ShareTitleOther            string      `json:"share_title_other"`
				ShareSignatureUrl          string      `json:"share_signature_url"`
				ShareSignatureDesc         string      `json:"share_signature_desc"`
				ShareQuote                 string      `json:"share_quote"`
				ShareDescInfo              string      `json:"share_desc_info"`
				NowInvitationCardImageUrls interface{} `json:"now_invitation_card_image_urls"`
			} `json:"share_info"`
			ConnectMusic   []interface{} `json:"connect_music"`
			Type           int           `json:"type"`
			SubType        int           `json:"sub_type"`
			IsPgcshow      bool          `json:"is_pgcshow"`
			CollectStat    int           `json:"collect_stat"`
			IsChallenge    int           `json:"is_challenge"`
			ViewCount      int           `json:"view_count"`
			IsCommerce     bool          `json:"is_commerce"`
			HashtagProfile string        `json:"hashtag_profile"`
			ChaAttrs       interface{}   `json:"cha_attrs"`
			BannerList     interface{}   `json:"banner_list"`
			ExtraAttr      struct {
				IsLive bool `json:"is_live"`
			} `json:"extra_attr"`
			ShowItems       interface{} `json:"show_items"`
			SearchHighlight interface{} `json:"search_highlight"`
			UseCount        int         `json:"use_count"`
		} `json:"cha_list"`
		Video struct {
			PlayAddr struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlKey    string      `json:"url_key"`
				DataSize  int         `json:"data_size"`
				FileHash  string      `json:"file_hash"`
				UrlPrefix interface{} `json:"url_prefix"`
				FileCs    string      `json:"file_cs,omitempty"`
			} `json:"play_addr"`
			Cover struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"cover"`
			Height       int `json:"height"`
			Width        int `json:"width"`
			DynamicCover struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"dynamic_cover"`
			OriginCover struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"origin_cover"`
			Ratio        string `json:"ratio"`
			DownloadAddr struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				DataSize  int         `json:"data_size"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"download_addr"`
			HasWatermark bool `json:"has_watermark"`
			BitRate      []struct {
				GearName    string `json:"gear_name"`
				QualityType int    `json:"quality_type"`
				BitRate     int    `json:"bit_rate"`
				PlayAddr    struct {
					Uri       string      `json:"uri"`
					UrlList   []string    `json:"url_list"`
					Width     int         `json:"width"`
					Height    int         `json:"height"`
					UrlKey    string      `json:"url_key"`
					DataSize  int         `json:"data_size"`
					FileHash  string      `json:"file_hash"`
					UrlPrefix interface{} `json:"url_prefix"`
					FileCs    string      `json:"file_cs,omitempty"`
				} `json:"play_addr"`
				IsH265      int         `json:"is_h265"`
				IsBytevc1   int         `json:"is_bytevc1"`
				DubInfos    interface{} `json:"dub_infos"`
				HDRType     string      `json:"HDR_type"`
				HDRBit      string      `json:"HDR_bit"`
				VideoExtra  string      `json:"video_extra"`
				PlayAddr265 struct {
					Uri       string      `json:"uri"`
					UrlList   []string    `json:"url_list"`
					Width     int         `json:"width"`
					Height    int         `json:"height"`
					UrlKey    string      `json:"url_key"`
					DataSize  int         `json:"data_size"`
					FileHash  string      `json:"file_hash"`
					FileCs    string      `json:"file_cs"`
					UrlPrefix interface{} `json:"url_prefix"`
				} `json:"play_addr_265,omitempty"`
			} `json:"bit_rate"`
			Duration               int `json:"duration"`
			DownloadSuffixLogoAddr struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				DataSize  int         `json:"data_size"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"download_suffix_logo_addr,omitempty"`
			HasDownloadSuffixLogoAddr bool        `json:"has_download_suffix_logo_addr,omitempty"`
			IsH265                    int         `json:"is_h265"`
			CdnUrlExpired             int         `json:"cdn_url_expired"`
			NeedSetToken              bool        `json:"need_set_token"`
			CoverTsp                  float64     `json:"CoverTsp"`
			MiscDownloadAddrs         string      `json:"misc_download_addrs,omitempty"`
			IsCallback                bool        `json:"is_callback"`
			Tags                      interface{} `json:"tags"`
			BigThumbs                 []struct {
				ImgNum   int      `json:"img_num"`
				Uri      string   `json:"uri"`
				ImgUrl   string   `json:"img_url"`
				ImgXSize int      `json:"img_x_size"`
				ImgYSize int      `json:"img_y_size"`
				ImgXLen  int      `json:"img_x_len"`
				ImgYLen  int      `json:"img_y_len"`
				Duration float64  `json:"duration"`
				Interval int      `json:"interval"`
				Fext     string   `json:"fext"`
				ImgUris  []string `json:"img_uris"`
				ImgUrls  []string `json:"img_urls"`
			} `json:"big_thumbs"`
			IsBytevc1      int    `json:"is_bytevc1"`
			Meta           string `json:"meta"`
			AiDynamicCover struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"ai_dynamic_cover"`
			AiDynamicCoverBak struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"ai_dynamic_cover_bak"`
			SourceHDRType int           `json:"source_HDR_type"`
			BitRateAudio  []interface{} `json:"bit_rate_audio"`
			PlayAddr265   struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlKey    string      `json:"url_key"`
				DataSize  int         `json:"data_size"`
				FileHash  string      `json:"file_hash"`
				FileCs    string      `json:"file_cs"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"play_addr_265,omitempty"`
			PlayAddrH264 struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlKey    string      `json:"url_key"`
				DataSize  int         `json:"data_size"`
				FileHash  string      `json:"file_hash"`
				FileCs    string      `json:"file_cs"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"play_addr_h264,omitempty"`
			CoverIsCustom bool   `json:"cover_is_custom,omitempty"`
			IsLongVideo   int    `json:"is_long_video,omitempty"`
			VideoModel    string `json:"video_model,omitempty"`
		} `json:"video"`
		ShareUrl   string `json:"share_url"`
		UserDigged int    `json:"user_digged"`
		Statistics struct {
			AwemeId            string `json:"aweme_id"`
			CommentCount       int    `json:"comment_count"`
			DiggCount          int    `json:"digg_count"`
			DownloadCount      int    `json:"download_count"`
			PlayCount          int    `json:"play_count"`
			ShareCount         int    `json:"share_count"`
			ForwardCount       int    `json:"forward_count"`
			LoseCount          int    `json:"lose_count"`
			LoseCommentCount   int    `json:"lose_comment_count"`
			WhatsappShareCount int    `json:"whatsapp_share_count"`
			CollectCount       int    `json:"collect_count"`
			RepostCount        int    `json:"repost_count"`
		} `json:"statistics"`
		Status struct {
			AwemeId        string `json:"aweme_id"`
			IsDelete       bool   `json:"is_delete"`
			AllowShare     bool   `json:"allow_share"`
			AllowComment   bool   `json:"allow_comment"`
			IsPrivate      bool   `json:"is_private"`
			WithGoods      bool   `json:"with_goods"`
			PrivateStatus  int    `json:"private_status"`
			InReviewing    bool   `json:"in_reviewing"`
			Reviewed       int    `json:"reviewed"`
			SelfSee        bool   `json:"self_see"`
			IsProhibited   bool   `json:"is_prohibited"`
			DownloadStatus int    `json:"download_status"`
			ReviewResult   struct {
				ReviewStatus int `json:"review_status"`
			} `json:"review_result"`
		} `json:"status"`
		Rate      int `json:"rate"`
		TextExtra []struct {
			Start       int    `json:"start"`
			End         int    `json:"end"`
			UserId      string `json:"user_id"`
			Type        int    `json:"type"`
			HashtagName string `json:"hashtag_name"`
			HashtagId   string `json:"hashtag_id,omitempty"`
			IsCommerce  bool   `json:"is_commerce,omitempty"`
			SecUid      string `json:"sec_uid"`
		} `json:"text_extra"`
		IsTop    int `json:"is_top"`
		LabelTop struct {
			Uri       string      `json:"uri"`
			UrlList   []string    `json:"url_list"`
			Width     int         `json:"width"`
			Height    int         `json:"height"`
			UrlPrefix interface{} `json:"url_prefix"`
		} `json:"label_top"`
		ShareInfo struct {
			ShareUrl                   string      `json:"share_url"`
			ShareWeiboDesc             string      `json:"share_weibo_desc"`
			ShareDesc                  string      `json:"share_desc"`
			ShareTitle                 string      `json:"share_title"`
			BoolPersist                int         `json:"bool_persist"`
			ShareTitleMyself           string      `json:"share_title_myself"`
			ShareTitleOther            string      `json:"share_title_other"`
			ShareLinkDesc              string      `json:"share_link_desc"`
			ShareSignatureUrl          string      `json:"share_signature_url"`
			ShareSignatureDesc         string      `json:"share_signature_desc"`
			ShareQuote                 string      `json:"share_quote"`
			ShareDescInfo              string      `json:"share_desc_info"`
			NowInvitationCardImageUrls interface{} `json:"now_invitation_card_image_urls"`
		} `json:"share_info"`
		Distance    string        `json:"distance"`
		VideoLabels []interface{} `json:"video_labels"`
		IsVr        bool          `json:"is_vr"`
		IsAds       bool          `json:"is_ads"`
		Duration    int           `json:"duration"`
		AwemeType   int           `json:"aweme_type"`
		CmtSwt      bool          `json:"cmt_swt"`
		ImageInfos  interface{}   `json:"image_infos"`
		RiskInfos   struct {
			Vote     bool   `json:"vote"`
			Warn     bool   `json:"warn"`
			RiskSink bool   `json:"risk_sink"`
			Type     int    `json:"type"`
			Content  string `json:"content"`
		} `json:"risk_infos"`
		IsRelieve            bool          `json:"is_relieve"`
		SortLabel            string        `json:"sort_label"`
		Position             interface{}   `json:"position"`
		UniqidPosition       interface{}   `json:"uniqid_position"`
		AuthorUserId         int64         `json:"author_user_id"`
		BodydanceScore       int           `json:"bodydance_score"`
		Geofencing           []interface{} `json:"geofencing"`
		IsHashTag            int           `json:"is_hash_tag"`
		IsPgcshow            bool          `json:"is_pgcshow"`
		Region               string        `json:"region"`
		VideoText            []interface{} `json:"video_text"`
		CollectStat          int           `json:"collect_stat"`
		LabelTopText         interface{}   `json:"label_top_text"`
		GroupId              string        `json:"group_id"`
		PreventDownload      bool          `json:"prevent_download"`
		NicknamePosition     interface{}   `json:"nickname_position"`
		ChallengePosition    interface{}   `json:"challenge_position"`
		ItemCommentSettings  int           `json:"item_comment_settings"`
		WithPromotionalMusic bool          `json:"with_promotional_music"`
		LongVideo            interface{}   `json:"long_video"`
		ItemDuet             int           `json:"item_duet"`
		ItemReact            int           `json:"item_react"`
		WithoutWatermark     bool          `json:"without_watermark"`
		DescLanguage         string        `json:"desc_language"`
		InteractionStickers  []struct {
			Type            int    `json:"type"`
			Index           int    `json:"index"`
			IsNonGlobal     bool   `json:"is_non_global"`
			TrackInfo       string `json:"track_info"`
			Attr            string `json:"attr"`
			TextInfo        string `json:"text_info,omitempty"`
			TextStickerInfo struct {
				TextSize     int     `json:"text_size"`
				TextColor    string  `json:"text_color"`
				BgColor      string  `json:"bg_color"`
				TextLanguage string  `json:"text_language"`
				Alignment    int     `json:"alignment"`
				SourceWidth  float64 `json:"source_width"`
				SourceHeight float64 `json:"source_height"`
			} `json:"text_sticker_info,omitempty"`
			MaterialIndex        int  `json:"material_index"`
			IsNonGlobalV2        bool `json:"is_non_global_v2"`
			AutoVideoCaptionInfo struct {
				Location     int         `json:"location"`
				Utterances   interface{} `json:"utterances"`
				AutoCaptions []struct {
					Language string `json:"language"`
					Url      struct {
						UrlList   []string    `json:"url_list"`
						UrlPrefix interface{} `json:"url_prefix"`
					} `json:"url"`
				} `json:"auto_captions"`
				IsTranslatable bool `json:"is_translatable"`
			} `json:"auto_video_caption_info,omitempty"`
		} `json:"interaction_stickers"`
		MiscInfo           string      `json:"misc_info"`
		OriginCommentIds   interface{} `json:"origin_comment_ids"`
		CommerceConfigData interface{} `json:"commerce_config_data"`
		DistributeType     int         `json:"distribute_type"`
		VideoControl       struct {
			AllowDownload         bool `json:"allow_download"`
			ShareType             int  `json:"share_type"`
			ShowProgressBar       int  `json:"show_progress_bar"`
			DraftProgressBar      int  `json:"draft_progress_bar"`
			AllowDuet             bool `json:"allow_duet"`
			AllowReact            bool `json:"allow_react"`
			PreventDownloadType   int  `json:"prevent_download_type"`
			AllowDynamicWallpaper bool `json:"allow_dynamic_wallpaper"`
			TimerStatus           int  `json:"timer_status"`
			AllowMusic            bool `json:"allow_music"`
			AllowStitch           bool `json:"allow_stitch"`
		} `json:"video_control"`
		CommerceInfo struct {
			AuctionAdInvited       bool `json:"auction_ad_invited"`
			WithCommentFilterWords bool `json:"with_comment_filter_words"`
			AdvPromotable          bool `json:"adv_promotable"`
			BrandedContentType     int  `json:"branded_content_type"`
		} `json:"commerce_info"`
		IsPreview         int         `json:"is_preview"`
		Anchors           interface{} `json:"anchors"`
		HybridLabel       interface{} `json:"hybrid_label"`
		GeofencingRegions interface{} `json:"geofencing_regions"`
		HaveDashboard     bool        `json:"have_dashboard"`
		AwemeAcl          struct {
			DownloadGeneral struct {
				Code      int    `json:"code"`
				ShowType  int    `json:"show_type"`
				Transcode int    `json:"transcode"`
				Mute      bool   `json:"mute"`
				Extra     string `json:"extra,omitempty"`
			} `json:"download_general"`
			DownloadMaskPanel struct {
				Code      int    `json:"code"`
				ShowType  int    `json:"show_type"`
				Transcode int    `json:"transcode"`
				Mute      bool   `json:"mute"`
				Extra     string `json:"extra,omitempty"`
			} `json:"download_mask_panel"`
			ShareListStatus int `json:"share_list_status"`
			ShareGeneral    struct {
				Code      int    `json:"code"`
				ShowType  int    `json:"show_type"`
				Transcode int    `json:"transcode"`
				Mute      bool   `json:"mute"`
				ToastMsg  string `json:"toast_msg,omitempty"`
				Extra     string `json:"extra,omitempty"`
			} `json:"share_general"`
			PlatformList    interface{} `json:"platform_list"`
			ShareActionList interface{} `json:"share_action_list"`
			PressActionList interface{} `json:"press_action_list"`
		} `json:"aweme_acl"`
		IsStory              int           `json:"is_story"`
		ItemStitch           int           `json:"item_stitch"`
		CoverLabels          interface{}   `json:"cover_labels"`
		MaskInfos            []interface{} `json:"mask_infos"`
		SearchDesc           string        `json:"search_desc"`
		SearchHighlight      []interface{} `json:"search_highlight"`
		PlaylistBlocked      bool          `json:"playlist_blocked"`
		GreenScreenMaterials interface{}   `json:"green_screen_materials"`
		NeedTrimStep         bool          `json:"need_trim_step"`
		InteractPermission   struct {
			Duet                 int `json:"duet"`
			Stitch               int `json:"stitch"`
			DuetPrivacySetting   int `json:"duet_privacy_setting"`
			StitchPrivacySetting int `json:"stitch_privacy_setting"`
			Upvote               int `json:"upvote"`
			AllowAddingToStory   int `json:"allow_adding_to_story"`
			AllowCreateSticker   struct {
				Status int `json:"status"`
			} `json:"allow_create_sticker"`
			AllowStorySwitchToPost struct {
			} `json:"allow_story_switch_to_post"`
			AllowAddingAsPost struct {
				Status int `json:"status"`
			} `json:"allow_adding_as_post"`
		} `json:"interact_permission"`
		QuestionList     interface{} `json:"question_list"`
		ContentDesc      string      `json:"content_desc"`
		ContentDescExtra []struct {
			Start       int    `json:"start"`
			End         int    `json:"end"`
			Type        int    `json:"type"`
			HashtagName string `json:"hashtag_name"`
			HashtagId   string `json:"hashtag_id"`
			IsCommerce  bool   `json:"is_commerce"`
			LineIdx     int    `json:"line_idx"`
		} `json:"content_desc_extra"`
		ProductsInfo             interface{} `json:"products_info"`
		FollowUpPublishFromId    int64       `json:"follow_up_publish_from_id"`
		DisableSearchTrendingBar bool        `json:"disable_search_trending_bar"`
		GroupIdList              struct {
			GroupdIdList0 interface{} `json:"GroupdIdList0"`
			GroupdIdList1 []int64     `json:"GroupdIdList1"`
		} `json:"group_id_list"`
		MusicBeginTimeInMs        int         `json:"music_begin_time_in_ms"`
		BrandedContentAccounts    interface{} `json:"branded_content_accounts"`
		IsDescriptionTranslatable bool        `json:"is_description_translatable"`
		PoiReTagSignal            int         `json:"poi_re_tag_signal"`
		NoSelectedMusic           bool        `json:"no_selected_music"`
		FollowUpItemIdGroups      string      `json:"follow_up_item_id_groups"`
		IsTextStickerTranslatable bool        `json:"is_text_sticker_translatable"`
		TextStickerMajorLang      string      `json:"text_sticker_major_lang"`
		CcTemplateInfo            struct {
			TemplateId           string `json:"template_id"`
			Desc                 string `json:"desc"`
			AuthorName           string `json:"author_name"`
			ClipCount            int    `json:"clip_count"`
			DurationMilliseconds int    `json:"duration_milliseconds"`
			RelatedMusicId       string `json:"related_music_id"`
		} `json:"cc_template_info"`
		MusicSelectedFrom       string      `json:"music_selected_from"`
		TtsVoiceIds             []string    `json:"tts_voice_ids"`
		ReferenceTtsVoiceIds    interface{} `json:"reference_tts_voice_ids"`
		VoiceFilterIds          interface{} `json:"voice_filter_ids"`
		ReferenceVoiceFilterIds interface{} `json:"reference_voice_filter_ids"`
		MusicTitleStyle         int         `json:"music_title_style"`
		AnimatedImageInfo       struct {
			Type   int `json:"type"`
			Effect int `json:"effect"`
		} `json:"animated_image_info"`
		CommentConfig struct {
			EmojiRecommendList interface{} `json:"emoji_recommend_list"`
			Preload            struct {
				Preds string `json:"preds"`
			} `json:"preload"`
			QuickComment struct {
				Enabled bool `json:"enabled"`
			} `json:"quick_comment"`
		} `json:"comment_config"`
		IsOnThisDay         int `json:"is_on_this_day"`
		AddedSoundMusicInfo struct {
			Id      int64  `json:"id"`
			IdStr   string `json:"id_str"`
			Title   string `json:"title"`
			Author  string `json:"author"`
			Album   string `json:"album"`
			CoverHd struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"cover_hd"`
			CoverLarge struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"cover_large"`
			CoverMedium struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"cover_medium"`
			CoverThumb struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"cover_thumb"`
			PlayUrl struct {
				Uri       string        `json:"uri"`
				UrlList   []interface{} `json:"url_list"`
				Width     int           `json:"width"`
				Height    int           `json:"height"`
				UrlPrefix interface{}   `json:"url_prefix"`
			} `json:"play_url"`
			SchemaUrl                 string        `json:"schema_url"`
			SourcePlatform            int           `json:"source_platform"`
			StartTime                 int           `json:"start_time"`
			EndTime                   int           `json:"end_time"`
			Duration                  int           `json:"duration"`
			Extra                     string        `json:"extra"`
			UserCount                 int           `json:"user_count"`
			Position                  interface{}   `json:"position"`
			CollectStat               int           `json:"collect_stat"`
			Status                    int           `json:"status"`
			OfflineDesc               string        `json:"offline_desc"`
			OwnerId                   string        `json:"owner_id,omitempty"`
			OwnerNickname             string        `json:"owner_nickname"`
			IsOriginal                bool          `json:"is_original"`
			Mid                       string        `json:"mid"`
			BindedChallengeId         int           `json:"binded_challenge_id"`
			Redirect                  bool          `json:"redirect"`
			IsRestricted              bool          `json:"is_restricted"`
			AuthorDeleted             bool          `json:"author_deleted"`
			IsDelVideo                bool          `json:"is_del_video"`
			IsVideoSelfSee            bool          `json:"is_video_self_see"`
			OwnerHandle               string        `json:"owner_handle"`
			AuthorPosition            interface{}   `json:"author_position"`
			PreventDownload           bool          `json:"prevent_download"`
			PreventItemDownloadStatus int           `json:"prevent_item_download_status"`
			ExternalSongInfo          []interface{} `json:"external_song_info"`
			SecUid                    string        `json:"sec_uid,omitempty"`
			AvatarThumb               struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"avatar_thumb,omitempty"`
			AvatarMedium struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"avatar_medium,omitempty"`
			AvatarLarge struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"avatar_large,omitempty"`
			PreviewStartTime      float64       `json:"preview_start_time"`
			PreviewEndTime        int           `json:"preview_end_time"`
			IsCommerceMusic       bool          `json:"is_commerce_music"`
			IsOriginalSound       bool          `json:"is_original_sound"`
			AuditionDuration      int           `json:"audition_duration"`
			ShootDuration         int           `json:"shoot_duration"`
			ReasonType            int           `json:"reason_type"`
			Artists               []interface{} `json:"artists"`
			LyricShortPosition    interface{}   `json:"lyric_short_position"`
			MuteShare             bool          `json:"mute_share"`
			TagList               interface{}   `json:"tag_list"`
			DmvAutoShow           bool          `json:"dmv_auto_show"`
			IsAuthorArtist        bool          `json:"is_author_artist"`
			IsPgc                 bool          `json:"is_pgc"`
			IsMatchedMetadata     bool          `json:"is_matched_metadata"`
			IsAudioUrlWithCookie  bool          `json:"is_audio_url_with_cookie"`
			VideoDuration         int           `json:"video_duration"`
			SearchHighlight       interface{}   `json:"search_highlight"`
			MultiBitRatePlayInfo  interface{}   `json:"multi_bit_rate_play_info"`
			DurationHighPrecision struct {
				DurationPrecision         float64 `json:"duration_precision"`
				ShootDurationPrecision    float64 `json:"shoot_duration_precision"`
				AuditionDurationPrecision float64 `json:"audition_duration_precision"`
				VideoDurationPrecision    float64 `json:"video_duration_precision"`
			} `json:"duration_high_precision"`
			CanNotReuse         bool        `json:"can_not_reuse"`
			IsPlayMusic         bool        `json:"is_play_music"`
			CommercialRightType int         `json:"commercial_right_type"`
			TtToDspSongInfos    interface{} `json:"tt_to_dsp_song_infos"`
			IsShootingAllow     bool        `json:"is_shooting_allow"`
			RecommendStatus     int         `json:"recommend_status"`
			HasCommerceRight    bool        `json:"has_commerce_right"`
			MemeSongInfo        struct {
			} `json:"meme_song_info"`
			UncertArtists interface{} `json:"uncert_artists"`
			StrongBeatUrl struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"strong_beat_url,omitempty"`
			MatchedSong struct {
				Id          string `json:"id"`
				Author      string `json:"author"`
				Title       string `json:"title"`
				H5Url       string `json:"h5_url"`
				CoverMedium struct {
					Uri       string      `json:"uri"`
					UrlList   []string    `json:"url_list"`
					Width     int         `json:"width"`
					Height    int         `json:"height"`
					UrlPrefix interface{} `json:"url_prefix"`
				} `json:"cover_medium"`
				Performers interface{} `json:"performers"`
				ChorusInfo struct {
					StartMs    int `json:"start_ms"`
					DurationMs int `json:"duration_ms"`
				} `json:"chorus_info"`
				FullDuration int `json:"full_duration"`
			} `json:"matched_song,omitempty"`
			MusicReleaseInfo struct {
				IsNewReleaseSong bool `json:"is_new_release_song"`
				GroupReleaseDate int  `json:"group_release_date"`
			} `json:"music_release_info,omitempty"`
		} `json:"added_sound_music_info"`
		OriginVolume               string      `json:"origin_volume"`
		MusicVolume                string      `json:"music_volume"`
		SupportDanmaku             bool        `json:"support_danmaku"`
		HasDanmaku                 bool        `json:"has_danmaku"`
		MufCommentInfoV2           interface{} `json:"muf_comment_info_v2"`
		BehindTheSongMusicIds      interface{} `json:"behind_the_song_music_ids"`
		BehindTheSongVideoMusicIds interface{} `json:"behind_the_song_video_music_ids"`
		IsNffOrNr                  bool        `json:"is_nff_or_nr"`
		OperatorBoostInfo          interface{} `json:"operator_boost_info"`
		MainArchCommon             string      `json:"main_arch_common"`
		AigcInfo                   struct {
			AigcLabelType int  `json:"aigc_label_type"`
			CreatedByAi   bool `json:"created_by_ai"`
		} `json:"aigc_info"`
		Banners             interface{}   `json:"banners"`
		PickedUsers         []interface{} `json:"picked_users"`
		IsTitleTranslatable bool          `json:"is_title_translatable"`
		TitleLanguage       string        `json:"title_language"`
		CommentTopbarInfo   interface{}   `json:"comment_topbar_info"`
		ContentModel        struct {
			StandardBiz struct {
				TtsVoiceInfo struct {
					TtsVoiceAttr        string `json:"tts_voice_attr"`
					TtsVoiceReuseParams string `json:"tts_voice_reuse_params"`
				} `json:"tts_voice_info,omitempty"`
				VcFilterInfo struct {
					VcFilterAttr string `json:"vc_filter_attr"`
				} `json:"vc_filter_info,omitempty"`
			} `json:"standard_biz"`
			CustomBiz struct {
				AwemeTrace string `json:"aweme_trace"`
			} `json:"custom_biz"`
		} `json:"content_model"`
		CreationInfo struct {
			CreationUsedFunctions []string `json:"creation_used_functions"`
		} `json:"creation_info"`
		UsedFullSong       bool `json:"used_full_song"`
		MusicEndTimeInMs   int  `json:"music_end_time_in_ms,omitempty"`
		OriginalClientText struct {
			MarkupText string `json:"markup_text"`
			TextExtra  []struct {
				Start       int    `json:"start,omitempty"`
				End         int    `json:"end,omitempty"`
				Type        int    `json:"type"`
				HashtagName string `json:"hashtag_name"`
				IsCommerce  bool   `json:"is_commerce,omitempty"`
				SubType     int    `json:"sub_type,omitempty"`
				LineIdx     int    `json:"line_idx,omitempty"`
				TagId       string `json:"tag_id"`
			} `json:"text_extra"`
		} `json:"original_client_text,omitempty"`
		ContentOriginalType  int    `json:"content_original_type,omitempty"`
		TttProductRecallType int    `json:"ttt_product_recall_type,omitempty"`
		RetryType            int    `json:"retry_type,omitempty"`
		ShootTabName         string `json:"shoot_tab_name,omitempty"`
		ContentType          string `json:"content_type,omitempty"`
		ContentSizeType      int    `json:"content_size_type,omitempty"`
		AddYoursInfo         struct {
			VideoSource int `json:"video_source"`
		} `json:"add_yours_info,omitempty"`
		Stickers      string `json:"stickers,omitempty"`
		StickerDetail struct {
			Id       string      `json:"id"`
			Name     string      `json:"name"`
			Children interface{} `json:"children"`
			IconUrl  struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"icon_url"`
			OwnerId       string      `json:"owner_id"`
			Tags          interface{} `json:"tags"`
			SecUid        string      `json:"sec_uid"`
			LinkedAnchors interface{} `json:"linked_anchors"`
			Attributions  interface{} `json:"attributions"`
		} `json:"sticker_detail,omitempty"`
		FollowUpFirstItemId string `json:"follow_up_first_item_id,omitempty"`
		MomentsModeInfo     struct {
			IsMoments int `json:"is_moments"`
		} `json:"moments_mode_info,omitempty"`
	} `json:"aweme_list"`
	HasMore int `json:"has_more"`
	Cursor  int `json:"cursor"`
	Extra   struct {
		Now              int64         `json:"now"`
		Logid            string        `json:"logid"`
		FatalItemIds     []interface{} `json:"fatal_item_ids"`
		SearchRequestId  string        `json:"search_request_id"`
		ServerStreamTime int           `json:"server_stream_time"`
		ApiDebugInfo     interface{}   `json:"api_debug_info"`
	} `json:"extra"`
	LogPb struct {
		ImprId string `json:"impr_id"`
	} `json:"log_pb"`
	Backtrace          string `json:"backtrace"`
	GlobalDoodleConfig struct {
		Keyword          string      `json:"keyword"`
		DisplayFilterBar int         `json:"display_filter_bar"`
		NewSource        string      `json:"new_source"`
		SearchChannel    string      `json:"search_channel"`
		TnsSearchResult  string      `json:"tns_search_result"`
		HideResults      bool        `json:"hide_results"`
		FeedbackSurvey   interface{} `json:"feedback_survey"`
		HitShark         bool        `json:"hit_shark"`
		HitDolphin       bool        `json:"hit_dolphin"`
		SearchIntent     struct {
		} `json:"search_intent"`
	} `json:"global_doodle_config"`
	FeedbackType   string      `json:"feedback_type"`
	SearchItemList interface{} `json:"search_item_list"`
	Components     interface{} `json:"components"`
}
