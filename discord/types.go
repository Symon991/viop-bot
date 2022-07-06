package discord

import (
	"time"

	"github.com/symon991/pirate/sites"
)

type Payload struct {
	T  interface{} `json:"t"`
	S  interface{} `json:"s"`
	Op int         `json:"op"`
}

type HeartBeatPayload struct {
	Op int `json:"op"`
	D  int `json:"d"`
}

type HelloPayload struct {
	T  interface{} `json:"t"`
	S  interface{} `json:"s"`
	Op int         `json:"op"`
	D  struct {
		HeartbeatInterval int      `json:"heartbeat_interval"`
		Trace             []string `json:"_trace"`
	} `json:"d"`
}

type HeartBeatAckPayload struct {
	Op int `json:"op"`
}

type IdentifyPayload struct {
	Op int `json:"op"`
	D  struct {
		Token      string `json:"token"`
		Intents    int    `json:"intents"`
		Properties struct {
			Os      string `json:"os"`
			Browser string `json:"browser"`
			Device  string `json:"device"`
		} `json:"properties"`
	} `json:"d"`
}

type ResumePayload struct {
	Op int `json:"op"`
	D  struct {
		Token     string `json:"token"`
		SessionID string `json:"session_id"`
		Seq       int    `json:"seq"`
	} `json:"d"`
}

type GuildCreatePayload struct {
	T  string `json:"t"`
	S  int    `json:"s"`
	Op int    `json:"op"`
	D  struct {
		PreferredLocale           string        `json:"preferred_locale"`
		OwnerID                   string        `json:"owner_id"`
		Name                      string        `json:"name"`
		Emojis                    []interface{} `json:"emojis"`
		Large                     bool          `json:"large"`
		VoiceStates               []interface{} `json:"voice_states"`
		EmbeddedActivities        []interface{} `json:"embedded_activities"`
		AfkTimeout                int           `json:"afk_timeout"`
		Unavailable               bool          `json:"unavailable"`
		PremiumProgressBarEnabled bool          `json:"premium_progress_bar_enabled"`
		Members                   []struct {
			User struct {
				Username      string      `json:"username"`
				PublicFlags   int         `json:"public_flags"`
				ID            string      `json:"id"`
				Discriminator string      `json:"discriminator"`
				Bot           bool        `json:"bot"`
				Avatar        interface{} `json:"avatar"`
			} `json:"user"`
			Roles                      []string    `json:"roles"`
			PremiumSince               interface{} `json:"premium_since"`
			Pending                    bool        `json:"pending"`
			Nick                       interface{} `json:"nick"`
			Mute                       bool        `json:"mute"`
			JoinedAt                   time.Time   `json:"joined_at"`
			Flags                      int         `json:"flags"`
			Deaf                       bool        `json:"deaf"`
			CommunicationDisabledUntil interface{} `json:"communication_disabled_until"`
			Avatar                     interface{} `json:"avatar"`
		} `json:"members"`
		Nsfw                        bool          `json:"nsfw"`
		JoinedAt                    time.Time     `json:"joined_at"`
		Features                    []interface{} `json:"features"`
		Icon                        interface{}   `json:"icon"`
		Stickers                    []interface{} `json:"stickers"`
		PublicUpdatesChannelID      interface{}   `json:"public_updates_channel_id"`
		MfaLevel                    int           `json:"mfa_level"`
		MaxVideoChannelUsers        int           `json:"max_video_channel_users"`
		DefaultMessageNotifications int           `json:"default_message_notifications"`
		NsfwLevel                   int           `json:"nsfw_level"`
		GuildHashes                 struct {
			Version int `json:"version"`
			Roles   struct {
				Omitted bool   `json:"omitted"`
				Hash    string `json:"hash"`
			} `json:"roles"`
			Metadata struct {
				Omitted bool   `json:"omitted"`
				Hash    string `json:"hash"`
			} `json:"metadata"`
			Channels struct {
				Omitted bool   `json:"omitted"`
				Hash    string `json:"hash"`
			} `json:"channels"`
		} `json:"guild_hashes"`
		Threads            []interface{} `json:"threads"`
		SystemChannelFlags int           `json:"system_channel_flags"`
		DiscoverySplash    interface{}   `json:"discovery_splash"`
		AfkChannelID       interface{}   `json:"afk_channel_id"`
		MemberCount        int           `json:"member_count"`
		StageInstances     []interface{} `json:"stage_instances"`
		Channels           []struct {
			Type                 int           `json:"type"`
			Position             int           `json:"position"`
			PermissionOverwrites []interface{} `json:"permission_overwrites"`
			Name                 string        `json:"name"`
			ID                   string        `json:"id"`
			Flags                int           `json:"flags"`
			Topic                interface{}   `json:"topic,omitempty"`
			RateLimitPerUser     int           `json:"rate_limit_per_user,omitempty"`
			ParentID             string        `json:"parent_id,omitempty"`
			LastMessageID        string        `json:"last_message_id,omitempty"`
			UserLimit            int           `json:"user_limit,omitempty"`
			RtcRegion            interface{}   `json:"rtc_region,omitempty"`
			Bitrate              int           `json:"bitrate,omitempty"`
		} `json:"channels"`
		ExplicitContentFilter    int `json:"explicit_content_filter"`
		ApplicationCommandCounts struct {
		} `json:"application_command_counts"`
		SystemChannelID          string        `json:"system_channel_id"`
		VanityURLCode            interface{}   `json:"vanity_url_code"`
		Splash                   interface{}   `json:"splash"`
		ID                       string        `json:"id"`
		Lazy                     bool          `json:"lazy"`
		Description              interface{}   `json:"description"`
		PremiumTier              int           `json:"premium_tier"`
		ApplicationID            interface{}   `json:"application_id"`
		PremiumSubscriptionCount int           `json:"premium_subscription_count"`
		Region                   string        `json:"region"`
		RulesChannelID           interface{}   `json:"rules_channel_id"`
		GuildScheduledEvents     []interface{} `json:"guild_scheduled_events"`
		Roles                    []struct {
			UnicodeEmoji interface{} `json:"unicode_emoji"`
			Tags         struct {
			} `json:"tags,omitempty"`
			Position    int         `json:"position"`
			Permissions string      `json:"permissions"`
			Name        string      `json:"name"`
			Mentionable bool        `json:"mentionable"`
			Managed     bool        `json:"managed"`
			ID          string      `json:"id"`
			Icon        interface{} `json:"icon"`
			Hoist       bool        `json:"hoist"`
			Flags       int         `json:"flags"`
			Color       int         `json:"color"`
			Tags0       struct {
				BotID string `json:"bot_id"`
			} `json:"tags,omitempty"`
		} `json:"roles"`
		Presences         []interface{} `json:"presences"`
		MaxMembers        int           `json:"max_members"`
		HubType           interface{}   `json:"hub_type"`
		Banner            interface{}   `json:"banner"`
		VerificationLevel int           `json:"verification_level"`
	} `json:"d"`
}

type MessageCreatePayload struct {
	T  string `json:"t"`
	S  int    `json:"s"`
	Op int    `json:"op"`
	D  struct {
		Type              int           `json:"type"`
		Tts               bool          `json:"tts"`
		Timestamp         time.Time     `json:"timestamp"`
		ReferencedMessage interface{}   `json:"referenced_message"`
		Pinned            bool          `json:"pinned"`
		Nonce             string        `json:"nonce"`
		Mentions          []interface{} `json:"mentions"`
		MentionRoles      []interface{} `json:"mention_roles"`
		MentionEveryone   bool          `json:"mention_everyone"`
		Member            struct {
			Roles                      []interface{} `json:"roles"`
			PremiumSince               interface{}   `json:"premium_since"`
			Pending                    bool          `json:"pending"`
			Nick                       interface{}   `json:"nick"`
			Mute                       bool          `json:"mute"`
			JoinedAt                   time.Time     `json:"joined_at"`
			Flags                      int           `json:"flags"`
			Deaf                       bool          `json:"deaf"`
			CommunicationDisabledUntil interface{}   `json:"communication_disabled_until"`
			Avatar                     interface{}   `json:"avatar"`
		} `json:"member"`
		ID              string        `json:"id"`
		Flags           int           `json:"flags"`
		Embeds          []interface{} `json:"embeds"`
		EditedTimestamp interface{}   `json:"edited_timestamp"`
		Content         string        `json:"content"`
		Components      []interface{} `json:"components"`
		ChannelID       string        `json:"channel_id"`
		Author          struct {
			Username         string      `json:"username"`
			PublicFlags      int         `json:"public_flags"`
			ID               string      `json:"id"`
			Discriminator    string      `json:"discriminator"`
			AvatarDecoration interface{} `json:"avatar_decoration"`
			Avatar           interface{} `json:"avatar"`
		} `json:"author"`
		Attachments []interface{} `json:"attachments"`
		GuildID     string        `json:"guild_id"`
	} `json:"d"`
}

type InteractionCreatePayload struct {
	T  string `json:"t"`
	S  int    `json:"s"`
	Op int    `json:"op"`
	D  struct {
		Version int `json:"version"`
		User    struct {
			Username         string      `json:"username"`
			PublicFlags      int         `json:"public_flags"`
			ID               string      `json:"id"`
			Discriminator    string      `json:"discriminator"`
			AvatarDecoration interface{} `json:"avatar_decoration"`
			Avatar           interface{} `json:"avatar"`
		} `json:"user"`
		Type    int    `json:"type"`
		Token   string `json:"token"`
		Locale  string `json:"locale"`
		ID      string `json:"id"`
		Message struct {
			ID          string `json:"id"`
			Interaction struct {
				User interface{} `json:"user"`
				Type int         `json:"type"`
				Name string      `json:"name"`
				ID   string      `json:"id"`
			} `json:"interaction"`
			Components []Components `json:"components"`
			Content    string       `json:"content"`
		} `json:"message"`
		Data struct {
			Options []struct {
				Name  string      `json:"name"`
				Value interface{} `json:"value"`
			} `json:"options"`
			Type   int      `json:"type"`
			Name   string   `json:"name"`
			ID     string   `json:"id"`
			Values []string `json:"values"`
		} `json:"data"`
		Member struct {
			User struct {
				Username         string      `json:"username"`
				PublicFlags      int         `json:"public_flags"`
				ID               string      `json:"id"`
				Discriminator    string      `json:"discriminator"`
				AvatarDecoration interface{} `json:"avatar_decoration"`
				Avatar           interface{} `json:"avatar"`
			} `json:"user"`
			Roles                      []interface{} `json:"roles"`
			PremiumSince               interface{}   `json:"premium_since"`
			Permissions                string        `json:"permissions"`
			Pending                    bool          `json:"pending"`
			Nick                       interface{}   `json:"nick"`
			Mute                       bool          `json:"mute"`
			JoinedAt                   time.Time     `json:"joined_at"`
			IsPending                  bool          `json:"is_pending"`
			Flags                      int           `json:"flags"`
			Deaf                       bool          `json:"deaf"`
			CommunicationDisabledUntil interface{}   `json:"communication_disabled_until"`
			Avatar                     interface{}   `json:"avatar"`
		} `json:"member"`
		ChannelID     string `json:"channel_id"`
		ApplicationID string `json:"application_id"`
	} `json:"d"`
}

type InteractionCallbackPayload struct {
	Type int `json:"type"`
	Data struct {
		Content    string       `json:"content"`
		Embeds     []Embed      `json:"embeds"`
		Components []Components `json:"components"`
	} `json:"data"`
	AllowedMentions AllowedMention `json:"allowedMentions"`
}

type Embed struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Type        string  `json:"type"`
	Fields      []Field `json:"fields"`
	Thumbnail   struct {
		URL string `json:"url"`
	} `json:"thumbnail"`
	Image struct {
		URL string `json:"url"`
	} `json:"image"`
}

type Field struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type AllowedMention struct {
	Parse []string `json:"parse"`
}

type Components struct {
	Type       int         `json:"type"`
	Components []Component `json:"components"`
}

type Component struct {
	Type     int      `json:"type"`
	Label    string   `json:"label"`
	Style    int      `json:"style"`
	CustomID string   `json:"custom_id"`
	Options  []Option `json:"options"`
	Url      string   `json:"url"`
}

type Option struct {
	Label       string `json:"label"`
	Value       string `json:"value"`
	Description string `json:"description"`
	Emoji       Emoji  `json:"emoji"`
}

type Emoji struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type TinyUrlRequest struct {
	Url    string `json:"url"`
	Domain string `json:domain`
}

type TinyUrlResponse struct {
	Data struct {
		Domain    string `json:"domain"`
		Alias     string `json:"alias"`
		Deleted   bool   `json:"deleted"`
		Archived  bool   `json:"archived"`
		Analytics struct {
			Enabled bool `json:"enabled"`
			Public  bool `json:"public"`
		} `json:"analytics"`
		Tags      []interface{} `json:"tags"`
		CreatedAt time.Time     `json:"created_at"`
		ExpiresAt interface{}   `json:"expires_at"`
		TinyURL   string        `json:"tiny_url"`
		URL       string        `json:"url"`
	} `json:"data"`
	Code   int           `json:"code"`
	Errors []interface{} `json:"errors"`
}

type CachePirateEntry struct {
	Metadata []sites.Metadata
	Site     string
}
