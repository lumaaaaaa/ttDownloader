package main

type SearchData struct {
	GlobalDoodleConfig struct {
		DisplayFilterBar int    `json:"display_filter_bar"`
		NewSource        string `json:"new_source"`
		SearchChannel    string `json:"search_channel"`
		HitDolphin       bool   `json:"hit_dolphin"`
		Keyword          string `json:"keyword"`
		TnsSearchResult  string `json:"tns_search_result"`
		HideResults      bool   `json:"hide_results"`
		FeedbackSurvey   []struct {
			FeedbackType    string `json:"feedback_type"`
			MultipleChoices []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"multiple_choices"`
		} `json:"feedback_survey"`
		HitShark bool `json:"hit_shark"`
	} `json:"global_doodle_config"`
	Extra struct {
		Now             int64         `json:"now"`
		Logid           string        `json:"logid"`
		FatalItemIds    []interface{} `json:"fatal_item_ids"`
		SearchRequestID string        `json:"search_request_id"`
	} `json:"extra"`
	Backtrace string `json:"backtrace"`
	HasMore   int    `json:"has_more"`
	Cursor    int    `json:"cursor"`
	LogPb     struct {
		ImprID string `json:"impr_id"`
	} `json:"log_pb"`
	FeedbackType string `json:"feedback_type"`
	StatusCode   int    `json:"status_code"`
	AwemeList    []struct {
		PreventDownload  bool          `json:"prevent_download"`
		ContentDescExtra []interface{} `json:"content_desc_extra"`
		Author           struct {
			Events         interface{} `json:"events"`
			SpecialAccount struct {
				SpecialAccountList interface{} `json:"special_account_list"`
				TtNow              struct {
					TtNowLogStatus int `json:"tt_now_log_status"`
				} `json:"tt_now"`
			} `json:"special_account"`
			ShieldFollowNotice     int         `json:"shield_follow_notice"`
			FollowersDetail        interface{} `json:"followers_detail"`
			LiveAgreement          int         `json:"live_agreement"`
			InsID                  string      `json:"ins_id"`
			SearchUserName         string      `json:"search_user_name"`
			SearchUserDesc         string      `json:"search_user_desc"`
			AdvancedFeatureInfo    interface{} `json:"advanced_feature_info"`
			AccountLabels          interface{} `json:"account_labels"`
			EnterpriseVerifyReason string      `json:"enterprise_verify_reason"`
			SearchHighlight        interface{} `json:"search_highlight"`
			NameField              string      `json:"name_field"`
			StoryStatus            int         `json:"story_status"`
			FollowStatus           int         `json:"follow_status"`
			FollowerCount          int         `json:"follower_count"`
			FbExpireTime           int         `json:"fb_expire_time"`
			ShieldDiggNotice       int         `json:"shield_digg_notice"`
			ShieldCommentNotice    int         `json:"shield_comment_notice"`
			CommentSetting         int         `json:"comment_setting"`
			AvatarMedium           struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
				Width   int      `json:"width"`
				Height  int      `json:"height"`
			} `json:"avatar_medium"`
			VerifyInfo      string      `json:"verify_info"`
			UserMode        int         `json:"user_mode"`
			ShortID         string      `json:"short_id"`
			ItemList        interface{} `json:"item_list"`
			TypeLabel       interface{} `json:"type_label"`
			BoldFields      interface{} `json:"bold_fields"`
			UserNowPackInfo struct {
				UserNowStatus int `json:"user_now_status"`
			} `json:"user_now_pack_info"`
			ShieldEditFieldInfo        interface{} `json:"shield_edit_field_info"`
			Nickname                   string      `json:"nickname"`
			IsAdFake                   bool        `json:"is_ad_fake"`
			CommerceUserLevel          int         `json:"commerce_user_level"`
			UserRate                   int         `json:"user_rate"`
			CommentFilterStatus        int         `json:"comment_filter_status"`
			MutualRelationAvatars      interface{} `json:"mutual_relation_avatars"`
			AwemeCount                 int         `json:"aweme_count"`
			SpecialLock                int         `json:"special_lock"`
			PlatformSyncInfo           interface{} `json:"platform_sync_info"`
			CanMessageFollowStatusList interface{} `json:"can_message_follow_status_list"`
			AuthorityStatus            int         `json:"authority_status"`
			CoverURL                   interface{} `json:"cover_url"`
			AdCoverURL                 interface{} `json:"ad_cover_url"`
			HomepageBottomToast        interface{} `json:"homepage_bottom_toast"`
			UserProfileGuide           interface{} `json:"user_profile_guide"`
			AvatarThumb                struct {
				URLList []string `json:"url_list"`
				Width   int      `json:"width"`
				Height  int      `json:"height"`
				URI     string   `json:"uri"`
			} `json:"avatar_thumb"`
			VerificationType int         `json:"verification_type"`
			Secret           int         `json:"secret"`
			NeedPoints       interface{} `json:"need_points"`
			StitchSetting    int         `json:"stitch_setting"`
			AvatarURI        string      `json:"avatar_uri"`
			FollowerStatus   int         `json:"follower_status"`
			CustomVerify     string      `json:"custom_verify"`
			LiveVerify       int         `json:"live_verify"`
			WithShopEntry    bool        `json:"with_shop_entry"`
			ShowImageBubble  bool        `json:"show_image_bubble"`
			VideoIcon        struct {
				Height  int           `json:"height"`
				URI     string        `json:"uri"`
				URLList []interface{} `json:"url_list"`
				Width   int           `json:"width"`
			} `json:"video_icon"`
			AcceptPrivatePolicy     bool        `json:"accept_private_policy"`
			IsStar                  bool        `json:"is_star"`
			MentionStatus           int         `json:"mention_status"`
			FriendsStatus           int         `json:"friends_status"`
			HideSearch              bool        `json:"hide_search"`
			ChaList                 interface{} `json:"cha_list"`
			WhiteCoverURL           interface{} `json:"white_cover_url"`
			AdvanceFeatureItemOrder interface{} `json:"advance_feature_item_order"`
			UID                     string      `json:"uid"`
			WithCommerceEntry       bool        `json:"with_commerce_entry"`
			Geofencing              interface{} `json:"geofencing"`
			UserPeriod              int         `json:"user_period"`
			UserCanceled            bool        `json:"user_canceled"`
			Status                  int         `json:"status"`
			DownloadPromptTs        int         `json:"download_prompt_ts"`
			Avatar168X168           struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
				Width   int      `json:"width"`
				Height  int      `json:"height"`
			} `json:"avatar_168x168"`
			RelativeUsers      interface{} `json:"relative_users"`
			IsBlock            bool        `json:"is_block"`
			UniqueID           string      `json:"unique_id"`
			NeedRecommend      int         `json:"need_recommend"`
			IsDisciplineMember bool        `json:"is_discipline_member"`
			LiveCommerce       bool        `json:"live_commerce"`
			Avatar300X300      struct {
				Height  int      `json:"height"`
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
				Width   int      `json:"width"`
			} `json:"avatar_300x300"`
			FollowingCount int         `json:"following_count"`
			TotalFavorited int         `json:"total_favorited"`
			CvLevel        string      `json:"cv_level"`
			SecUID         string      `json:"sec_uid"`
			UserTags       interface{} `json:"user_tags"`
			AvatarLarger   struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
				Width   int      `json:"width"`
				Height  int      `json:"height"`
			} `json:"avatar_larger"`
			RoomID           int         `json:"room_id"`
			PreventDownload  bool        `json:"prevent_download"`
			CanSetGeofencing interface{} `json:"can_set_geofencing"`
		} `json:"author"`
		UserDigged           int         `json:"user_digged"`
		Geofencing           interface{} `json:"geofencing"`
		IsHashTag            int         `json:"is_hash_tag"`
		GreenScreenMaterials interface{} `json:"green_screen_materials"`
		InteractPermission   struct {
			AllowAddingToStory   int `json:"allow_adding_to_story"`
			Duet                 int `json:"duet"`
			Stitch               int `json:"stitch"`
			DuetPrivacySetting   int `json:"duet_privacy_setting"`
			StitchPrivacySetting int `json:"stitch_privacy_setting"`
			Upvote               int `json:"upvote"`
		} `json:"interact_permission"`
		IsTextStickerTranslatable bool `json:"is_text_sticker_translatable"`
		Video                     struct {
			Cover struct {
				Width   int      `json:"width"`
				Height  int      `json:"height"`
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"cover"`
			Height       int `json:"height"`
			DynamicCover struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
				Width   int      `json:"width"`
				Height  int      `json:"height"`
			} `json:"dynamic_cover"`
			BitRate []struct {
				PlayAddr struct {
					FileHash string   `json:"file_hash"`
					FileCs   string   `json:"file_cs"`
					URI      string   `json:"uri"`
					URLList  []string `json:"url_list"`
					Width    int      `json:"width"`
					Height   int      `json:"height"`
					URLKey   string   `json:"url_key"`
					DataSize int      `json:"data_size"`
				} `json:"play_addr"`
				IsBytevc1   int         `json:"is_bytevc1"`
				DubInfos    interface{} `json:"dub_infos"`
				HDRType     string      `json:"HDR_type"`
				HDRBit      string      `json:"HDR_bit"`
				GearName    string      `json:"gear_name"`
				QualityType int         `json:"quality_type"`
				BitRate     int         `json:"bit_rate"`
			} `json:"bit_rate"`
			Tags      interface{}   `json:"tags"`
			BigThumbs []interface{} `json:"big_thumbs"`
			PlayAddr  struct {
				URLList  []string `json:"url_list"`
				Width    int      `json:"width"`
				Height   int      `json:"height"`
				URLKey   string   `json:"url_key"`
				DataSize int      `json:"data_size"`
				FileHash string   `json:"file_hash"`
				FileCs   string   `json:"file_cs"`
				URI      string   `json:"uri"`
			} `json:"play_addr"`
			OriginCover struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
				Width   int      `json:"width"`
				Height  int      `json:"height"`
			} `json:"origin_cover"`
			DownloadAddr struct {
				URLList  []string `json:"url_list"`
				Width    int      `json:"width"`
				Height   int      `json:"height"`
				DataSize int      `json:"data_size"`
				URI      string   `json:"uri"`
			} `json:"download_addr"`
			Duration       int  `json:"duration"`
			IsCallback     bool `json:"is_callback"`
			AiDynamicCover struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"ai_dynamic_cover"`
			SourceHDRType int `json:"source_HDR_type"`
			Width         int `json:"width"`
			CdnURLExpired int `json:"cdn_url_expired"`
			AnimatedCover struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"animated_cover"`
			CoverTsp          int    `json:"CoverTsp"`
			MiscDownloadAddrs string `json:"misc_download_addrs"`
			IsBytevc1         int    `json:"is_bytevc1"`
			Meta              string `json:"meta"`
			AiDynamicCoverBak struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"ai_dynamic_cover_bak"`
			Ratio        string `json:"ratio"`
			NeedSetToken bool   `json:"need_set_token"`
			HasWatermark bool   `json:"has_watermark"`
		} `json:"video"`
		CollectStat          int         `json:"collect_stat"`
		ItemReact            int         `json:"item_react"`
		Anchors              interface{} `json:"anchors"`
		LabelTopText         interface{} `json:"label_top_text"`
		SearchDesc           string      `json:"search_desc"`
		NoSelectedMusic      bool        `json:"no_selected_music"`
		FollowUpItemIDGroups string      `json:"follow_up_item_id_groups"`
		Music                struct {
			BindedChallengeID int    `json:"binded_challenge_id"`
			OwnerHandle       string `json:"owner_handle"`
			IsPgc             bool   `json:"is_pgc"`
			ID                int64  `json:"id"`
			Author            string `json:"author"`
			AvatarMedium      struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
				Width   int      `json:"width"`
				Height  int      `json:"height"`
			} `json:"avatar_medium"`
			PreviewEndTime      int    `json:"preview_end_time"`
			IsOriginalSound     bool   `json:"is_original_sound"`
			ShootDuration       int    `json:"shoot_duration"`
			IsPlayMusic         bool   `json:"is_play_music"`
			SecUID              string `json:"sec_uid"`
			CommercialRightType int    `json:"commercial_right_type"`
			IDStr               string `json:"id_str"`
			CoverThumb          struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
				Width   int      `json:"width"`
				Height  int      `json:"height"`
			} `json:"cover_thumb"`
			Artists           []interface{} `json:"artists"`
			IsMatchedMetadata bool          `json:"is_matched_metadata"`
			Duration          int           `json:"duration"`
			Status            int           `json:"status"`
			Title             string        `json:"title"`
			OfflineDesc       string        `json:"offline_desc"`
			IsOriginal        bool          `json:"is_original"`
			MuteShare         bool          `json:"mute_share"`
			IsAuthorArtist    bool          `json:"is_author_artist"`
			CoverMedium       struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
				Width   int      `json:"width"`
				Height  int      `json:"height"`
			} `json:"cover_medium"`
			PlayURL struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
				Width   int      `json:"width"`
				Height  int      `json:"height"`
			} `json:"play_url"`
			AuditionDuration int         `json:"audition_duration"`
			SearchHighlight  interface{} `json:"search_highlight"`
			AvatarThumb      struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
				Width   int      `json:"width"`
				Height  int      `json:"height"`
			} `json:"avatar_thumb"`
			VideoDuration        int         `json:"video_duration"`
			MultiBitRatePlayInfo interface{} `json:"multi_bit_rate_play_info"`
			SourcePlatform       int         `json:"source_platform"`
			Mid                  string      `json:"mid"`
			PreviewStartTime     int         `json:"preview_start_time"`
			Position             interface{} `json:"position"`
			OwnerID              string      `json:"owner_id"`
			IsCommerceMusic      bool        `json:"is_commerce_music"`
			TagList              interface{} `json:"tag_list"`
			IsAudioURLWithCookie bool        `json:"is_audio_url_with_cookie"`
			CoverLarge           struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
				Width   int      `json:"width"`
				Height  int      `json:"height"`
			} `json:"cover_large"`
			LyricShortPosition    interface{} `json:"lyric_short_position"`
			DurationHighPrecision struct {
				ShootDurationPrecision    int `json:"shoot_duration_precision"`
				AuditionDurationPrecision int `json:"audition_duration_precision"`
				VideoDurationPrecision    int `json:"video_duration_precision"`
				DurationPrecision         int `json:"duration_precision"`
			} `json:"duration_high_precision"`
			CanNotReuse   bool   `json:"can_not_reuse"`
			Album         string `json:"album"`
			Extra         string `json:"extra"`
			UserCount     int    `json:"user_count"`
			CollectStat   int    `json:"collect_stat"`
			OwnerNickname string `json:"owner_nickname"`
			AuthorDeleted bool   `json:"author_deleted"`
			StrongBeatURL struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
				Width   int      `json:"width"`
				Height  int      `json:"height"`
			} `json:"strong_beat_url"`
			ExternalSongInfo []interface{} `json:"external_song_info"`
			DmvAutoShow      bool          `json:"dmv_auto_show"`
			AuthorPosition   interface{}   `json:"author_position"`
			PreventDownload  bool          `json:"prevent_download"`
		} `json:"music"`
		Statistics struct {
			AwemeID            string `json:"aweme_id"`
			ForwardCount       int    `json:"forward_count"`
			LoseCommentCount   int    `json:"lose_comment_count"`
			WhatsappShareCount int    `json:"whatsapp_share_count"`
			CollectCount       int    `json:"collect_count"`
			CommentCount       int    `json:"comment_count"`
			DiggCount          int    `json:"digg_count"` // like count
			DownloadCount      int    `json:"download_count"`
			PlayCount          int    `json:"play_count"`
			ShareCount         int    `json:"share_count"`
			LoseCount          int    `json:"lose_count"`
		} `json:"statistics"`
		CmtSwt               bool          `json:"cmt_swt"`
		SortLabel            string        `json:"sort_label"`
		TextStickerMajorLang string        `json:"text_sticker_major_lang"`
		UniqidPosition       interface{}   `json:"uniqid_position"`
		CommerceConfigData   interface{}   `json:"commerce_config_data"`
		HybridLabel          interface{}   `json:"hybrid_label"`
		SearchHighlight      []interface{} `json:"search_highlight"`
		AwemeID              string        `json:"aweme_id"`
		Rate                 int           `json:"rate"`
		IsTop                int           `json:"is_top"`
		IsAds                bool          `json:"is_ads"`
		ContentDesc          string        `json:"content_desc"`
		CommerceInfo         struct {
			AuctionAdInvited       bool `json:"auction_ad_invited"`
			WithCommentFilterWords bool `json:"with_comment_filter_words"`
			AdvPromotable          bool `json:"adv_promotable"`
		} `json:"commerce_info"`
		ProductsInfo     interface{} `json:"products_info"`
		CreateTime       int         `json:"create_time"`
		Distance         string      `json:"distance"`
		NicknamePosition interface{} `json:"nickname_position"`
		VideoControl     struct {
			AllowMusic            bool `json:"allow_music"`
			AllowStitch           bool `json:"allow_stitch"`
			AllowDownload         bool `json:"allow_download"`
			ShowProgressBar       int  `json:"show_progress_bar"`
			AllowDuet             bool `json:"allow_duet"`
			PreventDownloadType   int  `json:"prevent_download_type"`
			TimerStatus           int  `json:"timer_status"`
			ShareType             int  `json:"share_type"`
			DraftProgressBar      int  `json:"draft_progress_bar"`
			AllowReact            bool `json:"allow_react"`
			AllowDynamicWallpaper bool `json:"allow_dynamic_wallpaper"`
		} `json:"video_control"`
		IsPgcshow      bool   `json:"is_pgcshow"`
		DistributeType int    `json:"distribute_type"`
		ShareURL       string `json:"share_url"`
		TextExtra      []struct {
			HashtagID   string `json:"hashtag_id"`
			IsCommerce  bool   `json:"is_commerce"`
			SecUID      string `json:"sec_uid"`
			Start       int    `json:"start"`
			End         int    `json:"end"`
			UserID      string `json:"user_id"`
			Type        int    `json:"type"`
			HashtagName string `json:"hashtag_name"`
		} `json:"text_extra"`
		LabelTop struct {
			URI     string   `json:"uri"`
			URLList []string `json:"url_list"`
			Width   int      `json:"width"`
			Height  int      `json:"height"`
		} `json:"label_top"`
		BodydanceScore int `json:"bodydance_score"`
		GroupIDList    struct {
			GroupdIDList0 interface{} `json:"GroupdIdList0"`
			GroupdIDList1 interface{} `json:"GroupdIdList1"`
		} `json:"group_id_list"`
		ImageInfos interface{} `json:"image_infos"`
		LongVideo  interface{} `json:"long_video"`
		AwemeAcl   struct {
			ShareListStatus int `json:"share_list_status"`
			ShareGeneral    struct {
				Transcode int  `json:"transcode"`
				Mute      bool `json:"mute"`
				Code      int  `json:"code"`
				ShowType  int  `json:"show_type"`
			} `json:"share_general"`
			PlatformList    interface{} `json:"platform_list"`
			DownloadGeneral struct {
				Code      int  `json:"code"`
				ShowType  int  `json:"show_type"`
				Transcode int  `json:"transcode"`
				Mute      bool `json:"mute"`
			} `json:"download_general"`
			DownloadMaskPanel struct {
				Code      int  `json:"code"`
				ShowType  int  `json:"show_type"`
				Transcode int  `json:"transcode"`
				Mute      bool `json:"mute"`
			} `json:"download_mask_panel"`
		} `json:"aweme_acl"`
		DisableSearchTrendingBar  bool          `json:"disable_search_trending_bar"`
		Position                  interface{}   `json:"position"`
		WithPromotionalMusic      bool          `json:"with_promotional_music"`
		ItemCommentSettings       int           `json:"item_comment_settings"`
		MaskInfos                 []interface{} `json:"mask_infos"`
		ItemStitch                int           `json:"item_stitch"`
		CoverLabels               interface{}   `json:"cover_labels"`
		PlaylistBlocked           bool          `json:"playlist_blocked"`
		IsDescriptionTranslatable bool          `json:"is_description_translatable"`
		Desc                      string        `json:"desc"`
		ChallengePosition         interface{}   `json:"challenge_position"`
		OriginCommentIds          interface{}   `json:"origin_comment_ids"`
		HasVsEntry                bool          `json:"has_vs_entry"`
		Status                    struct {
			AwemeID        string `json:"aweme_id"`
			AllowShare     bool   `json:"allow_share"`
			AllowComment   bool   `json:"allow_comment"`
			InReviewing    bool   `json:"in_reviewing"`
			DownloadStatus int    `json:"download_status"`
			IsDelete       bool   `json:"is_delete"`
			PrivateStatus  int    `json:"private_status"`
			Reviewed       int    `json:"reviewed"`
			SelfSee        bool   `json:"self_see"`
			IsProhibited   bool   `json:"is_prohibited"`
			ReviewResult   struct {
				ReviewStatus int `json:"review_status"`
			} `json:"review_result"`
		} `json:"status"`
		MiscInfo               string      `json:"misc_info"`
		BrandedContentAccounts interface{} `json:"branded_content_accounts"`
		NeedTrimStep           bool        `json:"need_trim_step"`
		MusicBeginTimeInMs     int         `json:"music_begin_time_in_ms"`
		RiskInfos              struct {
			Vote     bool   `json:"vote"`
			Warn     bool   `json:"warn"`
			RiskSink bool   `json:"risk_sink"`
			Type     int    `json:"type"`
			Content  string `json:"content"`
		} `json:"risk_infos"`
		GroupID             string        `json:"group_id"`
		InteractionStickers interface{}   `json:"interaction_stickers"`
		IsPreview           int           `json:"is_preview"`
		VideoLabels         []interface{} `json:"video_labels"`
		IsVr                bool          `json:"is_vr"`
		DescLanguage        string        `json:"desc_language"`
		MusicEndTimeInMs    int           `json:"music_end_time_in_ms,omitempty"`
		AwemeType           int           `json:"aweme_type"`
		Region              string        `json:"region"`
		WithoutWatermark    bool          `json:"without_watermark"`
		AnchorsExtras       string        `json:"anchors_extras"`
		ShareInfo           struct {
			NowInvitationCardImageUrls interface{} `json:"now_invitation_card_image_urls"`
			ShareURL                   string      `json:"share_url"`
			ShareDesc                  string      `json:"share_desc"`
			ShareTitle                 string      `json:"share_title"`
			ShareTitleMyself           string      `json:"share_title_myself"`
			ShareSignatureDesc         string      `json:"share_signature_desc"`
			WhatsappDesc               string      `json:"whatsapp_desc"`
			ShareDescInfo              string      `json:"share_desc_info"`
			BoolPersist                int         `json:"bool_persist"`
			ShareTitleOther            string      `json:"share_title_other"`
			ShareLinkDesc              string      `json:"share_link_desc"`
			ShareSignatureURL          string      `json:"share_signature_url"`
			ShareQuote                 string      `json:"share_quote"`
		} `json:"share_info"`
		VideoText             []interface{} `json:"video_text"`
		GeofencingRegions     interface{}   `json:"geofencing_regions"`
		AllowGift             bool          `json:"allow_gift,omitempty"`
		NeedVsEntry           bool          `json:"need_vs_entry"`
		HaveDashboard         bool          `json:"have_dashboard"`
		QuestionList          interface{}   `json:"question_list"`
		FollowUpPublishFromID int64         `json:"follow_up_publish_from_id"`
		ChaList               []struct {
			ExtraAttr struct {
				IsLive bool `json:"is_live"`
			} `json:"extra_attr"`
			UserCount       int           `json:"user_count"`
			ConnectMusic    []interface{} `json:"connect_music"`
			SubType         int           `json:"sub_type"`
			IsChallenge     int           `json:"is_challenge"`
			ChaName         string        `json:"cha_name"`
			Desc            string        `json:"desc"`
			ViewCount       int           `json:"view_count"`
			ShowItems       interface{}   `json:"show_items"`
			SearchHighlight interface{}   `json:"search_highlight"`
			Cid             string        `json:"cid"`
			Author          struct {
				HomepageBottomToast        interface{} `json:"homepage_bottom_toast"`
				UserTags                   interface{} `json:"user_tags"`
				AdvanceFeatureItemOrder    interface{} `json:"advance_feature_item_order"`
				AdvancedFeatureInfo        interface{} `json:"advanced_feature_info"`
				ShieldEditFieldInfo        interface{} `json:"shield_edit_field_info"`
				CanMessageFollowStatusList interface{} `json:"can_message_follow_status_list"`
				PlatformSyncInfo           interface{} `json:"platform_sync_info"`
				ItemList                   interface{} `json:"item_list"`
				ChaList                    interface{} `json:"cha_list"`
				NeedPoints                 interface{} `json:"need_points"`
				BoldFields                 interface{} `json:"bold_fields"`
				MutualRelationAvatars      interface{} `json:"mutual_relation_avatars"`
				UserProfileGuide           interface{} `json:"user_profile_guide"`
				Geofencing                 interface{} `json:"geofencing"`
				CoverURL                   interface{} `json:"cover_url"`
				TypeLabel                  interface{} `json:"type_label"`
				AdCoverURL                 interface{} `json:"ad_cover_url"`
				WhiteCoverURL              interface{} `json:"white_cover_url"`
				SearchHighlight            interface{} `json:"search_highlight"`
				AccountLabels              interface{} `json:"account_labels"`
				FollowersDetail            interface{} `json:"followers_detail"`
				RelativeUsers              interface{} `json:"relative_users"`
				CanSetGeofencing           interface{} `json:"can_set_geofencing"`
				Events                     interface{} `json:"events"`
			} `json:"author"`
			ShareInfo struct {
				NowInvitationCardImageUrls interface{} `json:"now_invitation_card_image_urls"`
				ShareDesc                  string      `json:"share_desc"`
				BoolPersist                int         `json:"bool_persist"`
				ShareTitleOther            string      `json:"share_title_other"`
				ShareSignatureDesc         string      `json:"share_signature_desc"`
				ShareQuote                 string      `json:"share_quote"`
				ShareDescInfo              string      `json:"share_desc_info"`
				ShareURL                   string      `json:"share_url"`
				ShareTitle                 string      `json:"share_title"`
				ShareTitleMyself           string      `json:"share_title_myself"`
				ShareSignatureURL          string      `json:"share_signature_url"`
			} `json:"share_info"`
			Type           int         `json:"type"`
			HashtagProfile string      `json:"hashtag_profile"`
			ChaAttrs       interface{} `json:"cha_attrs"`
			BannerList     interface{} `json:"banner_list"`
			Schema         string      `json:"schema"`
			IsPgcshow      bool        `json:"is_pgcshow"`
			CollectStat    int         `json:"collect_stat"`
			IsCommerce     bool        `json:"is_commerce"`
		} `json:"cha_list"`
		IsRelieve     bool  `json:"is_relieve"`
		AuthorUserID  int64 `json:"author_user_id"`
		ItemDuet      int   `json:"item_duet"`
		UpvotePreload struct {
			NeedPullUpvoteInfo bool `json:"need_pull_upvote_info"`
		} `json:"upvote_preload,omitempty"`
		RetryType int `json:"retry_type,omitempty"`
	} `json:"aweme_list"`
}
