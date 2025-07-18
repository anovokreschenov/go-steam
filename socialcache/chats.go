package socialcache

import (
	"errors"
	. "github.com/anovokreschenov/go-steam/protocol/steamlang"
	"github.com/anovokreschenov/go-steam/steamid"
	"sync"
)

// Chats list is a thread safe map
// They can be iterated over like so:
// 	for id, chat := range client.Social.Chats.GetCopy() {
// 		log.Println(id, chat.Name)
// 	}
type ChatsList struct {
	mutex sync.RWMutex
	byId  map[steamid.SteamId]*Chat
}

// Returns a new chats list
func NewChatsList() *ChatsList {
	return &ChatsList{byId: make(map[steamid.SteamId]*Chat)}
}

// Adds a chat to the chat list
func (list *ChatsList) Add(chat Chat) {
	list.mutex.Lock()
	defer list.mutex.Unlock()
	_, exists := list.byId[chat.SteamId]
	if !exists { //make sure this doesnt already exist
		list.byId[chat.SteamId] = &chat
	}
}

// Removes a chat from the chat list
func (list *ChatsList) Remove(id steamid.SteamId) {
	list.mutex.Lock()
	defer list.mutex.Unlock()
	delete(list.byId, id)
}

// Adds a chat member to a given chat
func (list *ChatsList) AddChatMember(id steamid.SteamId, member ChatMember) {
	list.mutex.Lock()
	defer list.mutex.Unlock()
	chat := list.byId[id]
	if chat == nil { //Chat doesn't exist
		chat = &Chat{SteamId: id}
		list.byId[id] = chat
	}
	if chat.ChatMembers == nil { //New chat
		chat.ChatMembers = make(map[steamid.SteamId]ChatMember)
	}
	chat.ChatMembers[member.SteamId] = member
}

// Removes a chat member from a given chat
func (list *ChatsList) RemoveChatMember(id steamid.SteamId, member steamid.SteamId) {
	list.mutex.Lock()
	defer list.mutex.Unlock()
	chat := list.byId[id]
	if chat == nil { //Chat doesn't exist
		return
	}
	if chat.ChatMembers == nil { //New chat
		return
	}
	delete(chat.ChatMembers, member)
}

// Returns a copy of the chats map
func (list *ChatsList) GetCopy() map[steamid.SteamId]Chat {
	list.mutex.RLock()
	defer list.mutex.RUnlock()
	glist := make(map[steamid.SteamId]Chat)
	for key, chat := range list.byId {
		glist[key] = *chat
	}
	return glist
}

// Returns a copy of the chat of a given steamid.SteamId
func (list *ChatsList) ById(id steamid.SteamId) (Chat, error) {
	list.mutex.RLock()
	defer list.mutex.RUnlock()
	if val, ok := list.byId[id]; ok {
		return *val, nil
	}
	return Chat{}, errors.New("Chat not found")
}

// Returns the number of chats
func (list *ChatsList) Count() int {
	list.mutex.RLock()
	defer list.mutex.RUnlock()
	return len(list.byId)
}

// A Chat
type Chat struct {
	SteamId     steamid.SteamId `json:",string"`
	GroupId     steamid.SteamId `json:",string"`
	ChatMembers map[steamid.SteamId]ChatMember
}

// A Chat Member
type ChatMember struct {
	SteamId         steamid.SteamId `json:",string"`
	ChatPermissions EChatPermission
	ClanPermissions EClanPermission
}
