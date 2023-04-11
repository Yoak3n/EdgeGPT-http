package edgegpt

import (
	"log"
)

// Combines everything to make it seamless
type ChatBot struct {
	cookiePath string
	cookies    []map[string]interface{}
	proxy      string
	chatHub    *ChatHub
	addr       string
	path       string
}

func NewChatBot(cookiePath string, cookies []map[string]interface{}, proxy string) *ChatBot {
	addr := "sydney.bing.com"
	path := "/sydney/ChatHub"
	bot := &ChatBot{
		cookiePath: cookiePath,
		cookies:    cookies,
		proxy:      proxy,
		addr:       addr,
		path:       path,
		chatHub:    nil,
	}
	return bot
}

func (bot *ChatBot) Init() error {
	conversation := NewConversation(bot.cookiePath, bot.cookies, bot.proxy)
	err := conversation.Init()
	if err != nil {
		return err
	}
	bot.chatHub = NewChatHub(bot.addr, bot.path, conversation)
	err = bot.chatHub.newConnect()
	if err != nil {
		return err
	}
	log.Println("init success")
	return nil
}

/*
// Ask a question to the bot
The callback function is streaming,
it will be called every time data is received,
if you only want to get the final result,
you can use `answer.IsDone()` to judge whether it is finished
*/

func (bot *ChatBot) Ask(prompt string, conversationStyle ConversationStyle) (answer *Answer, err error) {
	// defer bot.chatHub.Close()

	return bot.chatHub.askStream(prompt, conversationStyle)
}

func (bot *ChatBot) Close() error {
	return bot.chatHub.Close()
}

// Reset the conversation
func (bot *ChatBot) Reset() error {
	bot.chatHub.Close()
	err := bot.Init()
	if err != nil {
		return err
	}
	return nil
}
