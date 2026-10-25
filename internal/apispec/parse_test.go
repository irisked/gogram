package apispec

import (
	"reflect"
	"strings"
	"testing"
)

// page builds a reference page around the given sections, as Telegram
// lays it out.
func page(sections string) string {
	return `<html><body><div id="menu"><div>nav</div></div>
<div id="dev_page_content"><!-- scroll_nav -->
<h3><a class="anchor" name="recent-changes" href="#recent-changes"><i class="anchor-icon"></i></a>Recent changes</h3>
<h4><a class="anchor" name="august-24-2026" href="#august-24-2026"><i class="anchor-icon"></i></a>August 24, 2026</h4>
<p><strong>Bot API 10.3</strong></p>
<hr>
` + sections + `
</div>
<script>if (a < b) {}</script>
<!-- page generated in 82.33ms -->
</body></html>`
}

const sections = `
<h3><a class="anchor" name="available-types" href="#available-types"><i class="anchor-icon"></i></a>Available types</h3>
<h4><a class="anchor" name="user" href="#user"><i class="anchor-icon"></i></a>User</h4>
<p>This object represents a Telegram user or bot.</p>
<table class="table">
<thead><tr><th>Field</th><th>Type</th><th>Description</th></tr></thead>
<tbody>
<tr><td>id</td><td>Integer</td><td>Unique identifier for this user or bot.</td></tr>
<tr><td>username</td><td>String</td><td><em>Optional</em>. User&#39;s or bot&#39;s username</td></tr>
</tbody>
</table>
<h4><a class="anchor" name="message" href="#message"><i class="anchor-icon"></i></a>Message</h4>
<p>This object represents a message.</p>
<table class="table"><tbody>
<tr><td>message_id</td><td>Integer</td><td>Unique message identifier</td></tr>
<tr><td>photo</td><td>Array of <a href="#photosize">PhotoSize</a></td><td><em>Optional</em>. Message is a photo, available sizes of the photo</td></tr>
</tbody></table>
<h4><a class="anchor" name="photosize" href="#photosize"><i class="anchor-icon"></i></a>PhotoSize</h4>
<p>This object represents one size of a photo.</p>
<table class="table"><tbody>
<tr><td>file_id</td><td>String</td><td>Identifier for this file</td></tr>
</tbody></table>
<h4><a class="anchor" name="reactiontype" href="#reactiontype"><i class="anchor-icon"></i></a>ReactionType</h4>
<p>This object describes the type of a reaction. Currently, it can be one of</p>
<ul>
<li><a href="#reactiontypeemoji">ReactionTypeEmoji</a></li>
</ul>
<h4><a class="anchor" name="reactiontypeemoji" href="#reactiontypeemoji"><i class="anchor-icon"></i></a>ReactionTypeEmoji</h4>
<p>The reaction is based on an emoji.</p>
<table class="table"><tbody>
<tr><td>type</td><td>String</td><td>Type of the reaction, always “emoji”</td></tr>
</tbody></table>
<h3><a class="anchor" name="available-methods" href="#available-methods"><i class="anchor-icon"></i></a>Available methods</h3>
<h4><a class="anchor" name="sendmessage" href="#sendmessage"><i class="anchor-icon"></i></a>sendMessage</h4>
<p>Use this method to send text messages. On success, the sent <a href="#message">Message</a> is returned.</p>
<table class="table"><tbody>
<tr><td>chat_id</td><td>Integer or String</td><td>Yes</td><td>Unique identifier for the target chat</td></tr>
<tr><td>text</td><td>String</td><td>Yes</td><td>Text of the message to be sent, 1-4096 characters</td></tr>
<tr><td>disable_notification</td><td>Boolean</td><td>Optional</td><td>Sends the message <a href="https://telegram.org/blog/channels-2-0#silent-messages">silently</a>.</td></tr>
</tbody></table>
<h4><a class="anchor" name="getuserprofilephotos" href="#getuserprofilephotos"><i class="anchor-icon"></i></a>getUserProfilePhotos</h4>
<p>Use this method to get a list of profile pictures. Returns an Array of <a href="#photosize">PhotoSize</a> objects.</p>
<table class="table"><tbody>
<tr><td>user_id</td><td>Integer</td><td>Yes</td><td>Unique identifier of the target user</td></tr>
</tbody></table>
<h4><a class="anchor" name="logout" href="#logout"><i class="anchor-icon"></i></a>logOut</h4>
<p>Use this method to log out. Returns <em>True</em> on success. See <a href="/bots/api#logout">the docs »</a></p>
`

func TestExtract(t *testing.T) {
	tests := []struct {
		name    string
		page    string
		want    string // prefix and suffix of the snapshot
		wantErr string
	}{
		{"keeps only the documentation", page(""), `<div id="dev_page_content">`, ""},
		{"no documentation", "<html></html>", "", "no <div"},
		{"unclosed documentation", `<div id="dev_page_content"><div></div>`, "", "never closed"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Extract([]byte(tt.page))
			if tt.wantErr != "" {
				if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
					t.Fatalf("error = %v, want one containing %q", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			s := string(got)
			if !strings.HasPrefix(s, tt.want) || !strings.HasSuffix(s, "</div>\n") {
				t.Errorf("snapshot = %q", s)
			}
			for _, dropped := range []string{"menu", "script", "page generated"} {
				if strings.Contains(s, dropped) {
					t.Errorf("snapshot keeps %q", dropped)
				}
			}
		})
	}
}

func TestParse(t *testing.T) {
	snapshot, err := Extract([]byte(page(sections)))
	if err != nil {
		t.Fatal(err)
	}
	s, err := Parse(snapshot)
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}

	tests := []struct {
		name string
		got  any
		want any
	}{
		{"version", s.Version, "Bot API 10.3"},
		{"release date", s.ReleaseDate, "August 24, 2026"},
		{"changelog", s.Changelog, URL + "#august-24-2026"},
		{"types", len(s.Types), 5},
		{"methods", len(s.Methods), 3},
		{"type description", s.Types["User"].Description, []string{"This object represents a Telegram user or bot."}},
		{"href", s.Types["User"].Href, URL + "#user"},
		{"required field", s.Types["User"].Fields[0], Field{Name: "id", Types: []string{"Integer"}, Required: true, Description: "Unique identifier for this user or bot."}},
		{"optional field", s.Types["User"].Fields[1], Field{Name: "username", Types: []string{"String"}, Description: "Optional. User's or bot's username"}},
		{"array field", s.Types["Message"].Fields[1].Types, []string{"Array of PhotoSize"}},
		{"union members", s.Types["ReactionType"].Subtypes, []string{"ReactionTypeEmoji"}},
		{"union member of", s.Types["ReactionTypeEmoji"].SubtypeOf, []string{"ReactionType"}},
		{"union description lists members", s.Types["ReactionType"].Description[1], "- ReactionTypeEmoji"},
		{"typographic quotes", s.Types["ReactionTypeEmoji"].Fields[0].Description, `Type of the reaction, always "emoji"`},
		{"method parameters", len(s.Methods["sendMessage"].Fields), 3},
		{"several types", s.Methods["sendMessage"].Fields[0].Types, []string{"Integer", "String"}},
		{"optional parameter", s.Methods["sendMessage"].Fields[2].Required, false},
		{"returns an object", s.Methods["sendMessage"].Returns, []string{"Message"}},
		{"returns an array", s.Methods["getUserProfilePhotos"].Returns, []string{"Array of PhotoSize"}},
		{"returns True", s.Methods["logOut"].Returns, []string{"Boolean"}},
		{"» links are spelled out", s.Methods["logOut"].Description[0], "Use this method to log out. Returns True on success. See the docs: " + URL + "#logout"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(tt.got, tt.want) {
				t.Errorf("got %#v, want %#v", tt.got, tt.want)
			}
		})
	}
}

func TestParseErrors(t *testing.T) {
	tests := []struct {
		name     string
		sections string
		wantErr  string
	}{
		{
			name:     "unknown field type",
			sections: `<h4><a class="anchor" name="user" href="#user"></a>User</h4><table><tbody><tr><td>id</td><td>Wibble</td><td>x</td></tr></tbody></table>`,
			wantErr:  "User.id has an unknown type Wibble",
		},
		{
			name:     "method without a result",
			sections: `<h4><a class="anchor" name="ping" href="#ping"></a>ping</h4><p>Use this method to ping.</p>`,
			wantErr:  "method ping: no result type",
		},
		{
			name:     "table layout changed",
			sections: `<h4><a class="anchor" name="user" href="#user"></a>User</h4><table><tbody><tr><td>id</td><td>Integer</td></tr></tbody></table>`,
			wantErr:  "has 2 cells",
		},
		{
			name:     "type without fields",
			sections: `<h4><a class="anchor" name="empty" href="#empty"></a>Empty</h4><p>Nothing here.</p>`,
			wantErr:  "type Empty has no fields or subtypes",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			snapshot, err := Extract([]byte(page(tt.sections)))
			if err != nil {
				t.Fatal(err)
			}
			_, err = Parse(snapshot)
			if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
				t.Fatalf("error = %v, want one containing %q", err, tt.wantErr)
			}
		})
	}
}

func TestCleanType(t *testing.T) {
	tests := []struct {
		in   string
		want []string
	}{
		{"Integer", []string{"Integer"}},
		{"Integer or String", []string{"Integer", "String"}},
		{"Array of PhotoSize", []string{"Array of PhotoSize"}},
		{"Array of Array of KeyboardButton", []string{"Array of Array of KeyboardButton"}},
		{"Array of InputMediaAudio, InputMediaPhoto and InputMediaVideo", []string{"Array of InputMediaAudio", "Array of InputMediaPhoto", "Array of InputMediaVideo"}},
		{"True", []string{"Boolean"}},
		{"Float number", []string{"Float"}},
		{"Messages", []string{"Message"}},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			if got := cleanType(tt.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cleanType(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestReturnTypes(t *testing.T) {
	tests := []struct {
		description string
		want        []string
	}{
		{"On success, the sent Message is returned.", []string{"Message"}},
		{"Returns True on success.", []string{"Boolean"}},
		{"On success, an array of Messages that were sent is returned.", []string{"Array of Message"}},
		{"On success, if the message is not an inline message, the edited Message is returned, otherwise True is returned.", []string{"Message", "Boolean"}},
		{"Returns the number of members in a chat. Returns Int on success.", []string{"Integer"}},
		{"Use this method to do nothing.", nil},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			if got := returnTypes(tt.description); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("returnTypes = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestCollapseSpace(t *testing.T) {
	tests := []struct{ in, want string }{
		{"a b", "a b"},
		{"a  b", "a b"},
		{"a \n b", "a b"},
		{"a \n", "a\n"},
		{"a  b", "a b"},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			if got := collapseSpace(tt.in); got != tt.want {
				t.Errorf("collapseSpace(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
