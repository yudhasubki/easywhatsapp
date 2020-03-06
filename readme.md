## EasyWhatsapp

easywhatsapp is whatsapp library for Go to make easy implement package from [Rhymen/go-whatsapp](https://github.com/Rhymen/go-whatsapp).

Big thanks to [sigalor/whatsapp-web-reveng](https://github.com/sigalor/whatsapp-web-reveng) and all contributors who made it.

## Installation
```bash
go get github.com/yudhasubki/easywhatsapp
```

## Task
- [x] Login
- [x] Send Message
- [x] Retrieve All Message
- [x] Retrieve Message By JID
- [x] Get All JID
- [x] Create Group
- [x] Search Message
- [x] Get Group ID

## Usage

init struct
```go
wa, err := easywhatsapp.New(5, true)
if err != nil {
    log.Printf("err : %v", err)
}
// then you need to send qrcode with Pusher. 

client := make(map[string]string)
client["APP_ID"] = "YOUR_APP_ID"
client["APP_KEY"] = "YOUR_APP_KEY"
client["APP_SECRET"] = "YOUR_APP_SECRET"
client["CLUSTER"] = "YOUR_CLIENT_CLUSTER"

// Add handler to retrieve message, jids, or related with messages
wa.AddHandler()

pusherClient := wa.Client(client)
wa.Streamer.Pusher = pusherClient

err = wa.Login()
if err != nil {
    log.Printf("err : %v", err)
}
```

Send Message
```go
msgId, err := wa.SendText(jid@s.whatsapp.net, "hello world!")
if err != nil {
    log.Printf("err : %v", err)
}
log.Printf("Message ID : %v", msgId)
```

Search Message
```go
exist, messages, err := wa.SearchMessage("Input your message key")
if err != nil {
    log.Printf("err : %v", err)
}

log.Printf("Message available : %v", exist)
for _, m := range messages {
    log.Printf("%v \n", m)
}
```

Create Group
```go
group, err := wa.CreateGroup("Group Name", []string{"participants@s.whatsapp.net, ..."})
if err != nil {
    log.Printf("err : %v", err)
}
log.Printf("%v", group)
```

Get All JID
```go
for _, j := range wa.Message.RemoteJID {
    log.Printf("jid : %v \n", j)
}
```

Get Group JID
```go
groupJids := wa.GetGroupJID()
for _, g := range groupJids {
    log.Printf("%v \n", g)
}
```

Get Chat by JID
```go
chats := wa.GetChatByJID("62812000000@s.whatsapp.net", 100)
for _, c := range chats {
    log.Printf("chats : %v", c)
}
```

Get All Chats
```go
chats := wa.GetChats(100)
for _, c := range chats {
    log.Printf("chats : %v", c)
}
```

Get All only Group Chats
```go
chats := wa.GetGroupChats(100)
for _, c := range chats {
    log.Printf("chats : %v", c)
}
```

## Status
This project is still under maintenance.