// (c) Jisin0

package utils

import "github.com/PaulSonOfLars/gotgbot/v2"

var TEXT map[string]string = map[string]string{

	"START": `
<b>Hᴇʏ %v ɪᴍ %v ᴀɴ Aᴡᴇsᴏᴍᴇ Filter bot with global filter support</b>

<i>I can save a custom reply for a word in any chat. Check my help menu for more details.</i>
	`,
	"ABOUT": `
<b>𝗗𝗲𝘃𝗲𝗹𝗼𝗽𝗲𝗿</b> : <a href='https://go.dev'>𝗔𝘁𝗵𝘂𝗹</a>
<b>𝗟𝗮𝗻𝗴𝘂𝗮𝗴𝗲</b> : <a href='https://go.dev/'>𝗚𝗼</a>
<b>𝗦𝗲𝗿𝘃𝗲𝗿</b> : <a href='heroku.com'>𝗛𝗲𝗿𝗼𝗸𝘂</a>
<b>𝗗𝗮𝘁𝗮𝗯𝗮𝘀𝗲</b> : <a href='mongodb.org'>𝗠𝗼𝗻𝗴𝗼𝗗𝗕</a>
<b>𝗖𝗵𝗮𝗻𝗻𝗲𝗹</b> : <a href='https://t.me/+L8SWfrF_7m04ODZl'>𝗛𝗲𝗿𝗲</a>
<b>𝗔𝘂𝘁𝗼 𝗙𝗶𝗹𝘁𝗲𝗿 𝗕𝗼𝘁</b> : <a href=''>𝗛𝗲𝗿𝗲</a>
	`,

	"MF": `
<b>Mᴀɴᴜᴀʟ ғɪʟᴛᴇʀs ᴀʟʟᴏᴡ ʏᴏᴜ ᴛᴏ sᴀᴠᴇ ᴄᴜsᴛᴏᴍ ғɪʟᴛᴇʀs ᴏᴛʜᴇʀ ᴛʜᴀɴ ᴛʜᴇ ᴀᴜᴛᴏᴍᴀᴛɪᴄ ᴏɴᴇs. Fɪʟᴛᴇʀs ᴄᴀɴ ʙᴇ ᴏғ ᴛᴇxᴛ/ᴘʜᴏᴛᴏ/ᴅᴏᴄᴜᴍᴇɴᴛ/ᴀᴜᴅɪᴏ/ᴀɴɪᴍᴀᴛɪᴏɴ/ᴠɪᴅᴇᴏ .</b>

<b><u>Nᴇᴡ ғɪʟᴛᴇʀ :</u></b>

<u>Fᴏʀᴍᴀᴛ</u>
<code>/filter "keyword" text</code> or
Rᴇᴘʟʏ ᴛᴏ ᴀ ᴍᴇssᴀɢᴇ -><code>/filter "keyword"</code>
<u>Usᴀɢᴇ</u>
<code>/filter "hi" hello</code>
[<code>hello</code>] -> Reply -> <code>/filter hi</code>

<b><u>Sᴛᴏᴘ ғɪʟᴛᴇʀ :</u></b>

<u>Fᴏʀᴍᴀᴛ</u>
<code>/stop "keyword"</code>
<u>Usᴀɢᴇ</u>
<code>/stop "hi"</code>

<b><u>Vɪᴇᴡ ғɪʟᴛᴇʀs :</u></b>
<code>/filters</code>
`,

	"GF": `
<b>Gʟᴏʙᴀʟ ғɪʟᴛᴇʀs ᴀʀᴇ ᴍᴀɴᴜᴀʟ ғɪʟᴛᴇʀs sᴀᴠᴇᴅ ʙʏ ʙᴏᴛ ᴀᴅᴍɪɴs ᴛʜᴀᴛ ᴡᴏʀᴋ ɪɴ ᴀʟʟ ᴄʜᴀᴛs. Tʜᴇʏ ᴘʀᴏᴠɪᴅᴇ ʟᴀᴛᴇsᴛ ᴍᴏᴠɪᴇs ɪɴ ᴀ ᴇᴀsʏ ᴛᴏ ᴜsᴇ ғᴏʀᴍᴀᴛ.</b>

<b><u>Sᴛᴏᴘ ғɪʟᴛᴇʀ :</u></b>

<u>Fᴏʀᴍᴀᴛ</u>
<code>/stop "keyword"</code>
<u>Usᴀɢᴇ</u>
<code>/stop "et"</code>

<b><u>Vɪᴇᴡ ғɪʟᴛᴇʀs :</u></b>
/gfilters
`,
	"CONNECT": `
<b>Cᴏɴɴᴇᴄᴛɪᴏɴs ᴀʟʟᴏᴡ ʏᴏᴜ ᴛᴏ ᴍᴀɴᴀɢᴇ ʏᴏᴜʀ ɢʀᴏᴜᴘ ʜᴇʀᴇ ɪɴ ᴘᴍ ɪɴsᴛᴇᴀᴅ ᴏғ sᴇɴᴅɪɴɢ ᴛʜᴏsᴇ ᴄᴏᴍᴍᴀɴᴅs ᴘᴜʙʟɪᴄʟʏ ɪɴ ᴛʜᴇ ɢʀᴏᴜᴘ ⠘⁾</b>

<b><u>Cᴏɴɴᴇᴄᴛ :</u></b>
-> Fɪʀsᴛ ɢᴇᴛ ʏᴏᴜʀ ɢʀᴏᴜᴘ's ɪᴅ ʙʏ sᴇɴᴅɪɴɢ /id ɪɴ ʏᴏᴜʀ ɢʀᴏᴜᴘ
-> <code>/connect [group_id]</code>

<b><u>Dɪsᴄᴏɴɴᴇᴄᴛ :</u></b>
<code>/disconnect</code>
`,

	"BROADCAST": `
<b>The broadcast feature allows bot admins to broadcast a message to all of the bot's users.</b>

<I>Broadcasts are of two types :</i>
 ~ Full Broadcast - Broadcast to all of the bot users <code>/broadcast</code>
 ~ Concast - Broadcast to only users who are connected to a chat <code>/concast</code>
`,

	"HELP": `
<b>To know how to use my modules use the buttons below to get all my commands with usage examples !</b>
`,
}

var BUTTONS map[string][][]gotgbot.InlineKeyboardButton = map[string][][]gotgbot.InlineKeyboardButton{
	"START": {
		{
			{Text: "ᴀʙᴏᴜᴛ", CallbackData: "edit(ABOUT)"},
			{Text: "ʜᴇʟᴘ", CallbackData: "edit(HELP)"},
			{Text: "ꜱᴜᴘᴘᴏʀᴛ", Url: "https://t.me/+L8SWfrF_7m04ODZl"},
		},
	},
	"ABOUT": {
		{
			{Text: "ʜᴏᴍᴇ", CallbackData: "edit(START)"},
			{Text: "ꜱᴛᴀᴛꜱ", CallbackData: "stats"},
		},
	},
	"STATS": {
		{
			{Text: "ʙᴀᴄᴋ", CallbackData: "edit(ABOUT)"},
			{Text: "ʀᴇꜰʀᴇꜱʜ", CallbackData: "stats"},
		},
	},
	"HELP": {
		{{Text: "ꜰɪʟᴛᴇʀ", CallbackData: "edit(MF)"},
			{Text: "ɢʟᴏʙᴀʟ", CallbackData: "edit(GF)"},
		}, {
			{Text: "ᴄᴏɴɴᴇᴄᴛ", CallbackData: "edit(CONNECT)"}, {Text: "Broadcast", CallbackData: "edit(BROADCAST)"},
		},
		{{Text: "ʙᴀᴄᴋ", CallbackData: "edit(START)"}},
	},
}
