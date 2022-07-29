package messages

import (
	"time"
)

type Base struct {
	T  string `json:"t"`
	Op int    `json:"op"`
	D  int    `json:"d"`
}

type HeartBeat struct {
	Op int `json:"op"`
	D  int `json:"d"`
}

type HeartBeatAck struct {
	Op int `json:"op"`
}

type Hello struct {
	T  interface{} `json:"t"`
	S  interface{} `json:"s"`
	Op int         `json:"op"`
	D  struct {
		HeartbeatInterval int      `json:"heartbeat_interval"`
		Trace             []string `json:"_trace"`
	} `json:"d"`
}

type Identify struct {
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

type Resume struct {
	Op int `json:"op"`
	D  struct {
		Token     string `json:"token"`
		SessionID string `json:"session_id"`
		Seq       int    `json:"seq"`
	} `json:"d"`
}

type InteractionCreate struct {
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
				User struct {
					Username         string      `json:"username"`
					PublicFlags      int         `json:"public_flags"`
					ID               string      `json:"id"`
					Discriminator    string      `json:"discriminator"`
					AvatarDecoration interface{} `json:"avatar_decoration"`
					Avatar           interface{} `json:"avatar"`
				} `json:"user"`
				Type int    `json:"type"`
				Name string `json:"name"`
				ID   string `json:"id"`
			} `json:"interaction"`
			Components []Components `json:"components"`
			Content    string       `json:"content"`
		} `json:"message"`
		Data struct {
			Options []struct {
				Name    string      `json:"name"`
				Value   interface{} `json:"value"`
				Options []struct {
					Name  string      `json:"name"`
					Value interface{} `json:"value"`
				} `json:"options"`
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

type InteractionCallback struct {
	Type            int `json:"type"`
	Data            `json:"data"`
	AllowedMentions AllowedMention `json:"allowedMentions"`
}

type Data struct {
	Content    string       `json:"content"`
	Embeds     []Embed      `json:"embeds"`
	Components []Components `json:"components"`
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
	Domain string `json:"domain"`
}

type ChannelMessage struct {
	Content string  `json:"content"`
	TTS     bool    `json:"tts"`
	Embeds  []Embed `json:"embeds"`
}
