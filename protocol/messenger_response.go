package protocol

import (
	"encoding/json"

	"github.com/status-im/status-go/protocol/common"
	"github.com/status-im/status-go/protocol/communities"
	"github.com/status-im/status-go/protocol/encryption/multidevice"
	"github.com/status-im/status-go/protocol/transport"
	"github.com/status-im/status-go/services/mailservers"
)

type MessengerResponse struct {
	Messages                []*common.Message
	Contacts                []*Contact
	Installations           []*multidevice.Installation
	EmojiReactions          []*EmojiReaction
	Invitations             []*GroupChatInvitation
	CommunityChanges        []*communities.CommunityChanges
	RequestsToJoinCommunity []*communities.RequestToJoin
	Filters                 []*transport.Filter
	RemovedFilters          []*transport.Filter
	Mailservers             []mailservers.Mailserver
	MailserverTopics        []mailservers.MailserverTopic
	MailserverRanges        []mailservers.ChatRequestRange
	// Notifications a list of MessageNotificationBody derived from received messages that are useful to notify the user about
	Notifications []MessageNotificationBody

	chats        map[string]*Chat
	removedChats map[string]bool
	communities  map[string]*communities.Community
}

func (r *MessengerResponse) MarshalJSON() ([]byte, error) {
	responseItem := struct {
		Chats                   []*Chat                         `json:"chats,omitempty"`
		RemovedChats            []string                        `json:"removedChats,omitempty"`
		Messages                []*common.Message               `json:"messages,omitempty"`
		Contacts                []*Contact                      `json:"contacts,omitempty"`
		Installations           []*multidevice.Installation     `json:"installations,omitempty"`
		EmojiReactions          []*EmojiReaction                `json:"emojiReactions,omitempty"`
		Invitations             []*GroupChatInvitation          `json:"invitations,omitempty"`
		CommunityChanges        []*communities.CommunityChanges `json:"communityChanges,omitempty"`
		RequestsToJoinCommunity []*communities.RequestToJoin    `json:"requestsToJoinCommunity,omitempty"`
		Filters                 []*transport.Filter             `json:"filters,omitempty"`
		RemovedFilters          []*transport.Filter             `json:"removedFilters,omitempty"`
		Mailservers             []mailservers.Mailserver        `json:"mailservers,omitempty"`
		MailserverTopics        []mailservers.MailserverTopic   `json:"mailserverTopics,omitempty"`
		MailserverRanges        []mailservers.ChatRequestRange  `json:"mailserverRanges,omitempty"`
		// Notifications a list of MessageNotificationBody derived from received messages that are useful to notify the user about
		Notifications []MessageNotificationBody `json:"notifications"`
		Communities   []*communities.Community  `json:"communities,omitempty"`
	}{
		Messages:                r.Messages,
		Contacts:                r.Contacts,
		Installations:           r.Installations,
		EmojiReactions:          r.EmojiReactions,
		Invitations:             r.Invitations,
		CommunityChanges:        r.CommunityChanges,
		RequestsToJoinCommunity: r.RequestsToJoinCommunity,
		Filters:                 r.Filters,
		RemovedFilters:          r.RemovedFilters,
		Mailservers:             r.Mailservers,
		MailserverTopics:        r.MailserverTopics,
		MailserverRanges:        r.MailserverRanges,
		Notifications:           r.Notifications,
	}

	responseItem.Chats = r.Chats()
	responseItem.Communities = r.Communities()
	responseItem.RemovedChats = r.RemovedChats()

	return json.Marshal(responseItem)
}

func (r *MessengerResponse) Chats() []*Chat {
	var chats []*Chat
	for _, chat := range r.chats {
		chats = append(chats, chat)
	}
	return chats
}

func (r *MessengerResponse) RemovedChats() []string {
	var chats []string
	for chatID := range r.removedChats {
		chats = append(chats, chatID)
	}
	return chats
}

func (r *MessengerResponse) Communities() []*communities.Community {
	var communities []*communities.Community
	for _, c := range r.communities {
		communities = append(communities, c)
	}
	return communities
}

func (m *MessengerResponse) IsEmpty() bool {
	return len(m.chats)+len(m.Messages)+len(m.Contacts)+len(m.Installations)+len(m.Invitations)+len(m.EmojiReactions)+len(m.communities)+len(m.CommunityChanges)+len(m.Filters)+len(m.RemovedFilters)+len(m.removedChats)+len(m.MailserverTopics)+len(m.Mailservers)+len(m.MailserverRanges)+len(m.Notifications)+len(m.RequestsToJoinCommunity) == 0
}

// Merge takes another response and appends the new Chats & new Messages and replaces
// the existing Messages & Chats if they have the same ID
func (m *MessengerResponse) Merge(response *MessengerResponse) error {
	if len(response.Contacts)+len(response.Installations)+len(response.EmojiReactions)+len(response.Invitations)+len(response.RequestsToJoinCommunity)+len(response.Mailservers)+len(response.MailserverTopics)+len(response.MailserverRanges)+len(response.Notifications)+len(response.EmojiReactions)+len(response.CommunityChanges) != 0 {
		return ErrNotImplemented
	}

	m.AddChats(response.Chats())
	m.AddRemovedChats(response.RemovedChats())
	m.overrideMessages(response.Messages)
	m.overrideFilters(response.Filters)
	m.overrideRemovedFilters(response.Filters)
	m.AddCommunities(response.Communities())

	return nil
}

// overrideMessages append new messages and override existing ones in response.Messages
func (m *MessengerResponse) overrideMessages(messages []*common.Message) {
	for _, overrideMessage := range messages {
		var found = false
		for idx, chat := range m.Messages {
			if chat.ID == overrideMessage.ID {
				m.Messages[idx] = overrideMessage
				found = true
			}
		}
		if !found {
			m.Messages = append(m.Messages, overrideMessage)
		}
	}
}

// overrideFilters append new filters and override existing ones in response.Filters
func (m *MessengerResponse) overrideFilters(filters []*transport.Filter) {
	for _, overrideFilter := range filters {
		var found = false
		for idx, filter := range m.Filters {
			if filter.FilterID == overrideFilter.FilterID {
				m.Filters[idx] = overrideFilter
				found = true
			}
		}
		if !found {
			m.Filters = append(m.Filters, overrideFilter)
		}
	}
}

// overrideRemovedFilters append removed filters and override existing ones in response.Filters
func (m *MessengerResponse) overrideRemovedFilters(filters []*transport.Filter) {
	for _, overrideFilter := range filters {
		var found = false
		for idx, filter := range m.RemovedFilters {
			if filter.FilterID == overrideFilter.FilterID {
				m.RemovedFilters[idx] = overrideFilter
				found = true
			}
		}
		if !found {
			m.RemovedFilters = append(m.RemovedFilters, overrideFilter)
		}
	}
}

func (m *MessengerResponse) AddCommunities(communities []*communities.Community) {
	for _, overrideCommunity := range communities {
		m.AddCommunity(overrideCommunity)
	}
}

func (m *MessengerResponse) AddCommunity(c *communities.Community) {
	if m.communities == nil {
		m.communities = make(map[string]*communities.Community)
	}

	m.communities[c.IDString()] = c
}

func (m *MessengerResponse) AddChat(c *Chat) {
	if m.chats == nil {
		m.chats = make(map[string]*Chat)
	}

	m.chats[c.ID] = c
}

func (m *MessengerResponse) AddChats(chats []*Chat) {
	for _, c := range chats {
		m.AddChat(c)
	}
}

func (m *MessengerResponse) AddRemovedChats(chats []string) {
	for _, c := range chats {
		m.AddRemovedChat(c)
	}
}

func (m *MessengerResponse) AddRemovedChat(chatID string) {
	if m.removedChats == nil {
		m.removedChats = make(map[string]bool)
	}

	m.removedChats[chatID] = true
}
