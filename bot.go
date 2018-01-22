package main

import (
    "fmt"
    "time"
    "github.com/bwmarrin/discordgo"
    "log"
    "strings"
    "botcmds"
    "os"
)

var(
    Token = "Bot Token" //"Bot"という接頭辞がないと401 unauthorizedエラーが起きます
    BotName = "<@ID>"
    stopBot = make(chan bool)
    CmdPrefix = "?"
    cmds = botcmds.GetCommands()
)

func main() {
    log.Println("Starting new discord session")
    discord, err := discordgo.New()
    discord.Token = Token
    if err != nil {
        log.Println("Error logging in")
        log.Println(err)
    }

    discord.AddHandler(onReady)
    discord.AddHandler(onPresenceUpdate)
    discord.AddHandler(onMessageCreate)

    log.Println("Connecting to the Server")
    err = discord.Open()
    if err != nil {
        log.Println(err)
    }

    log.Println("KOTABOT Ready to BORN")
    <-stopBot //プログラムが終了しないようロック
    return
}

func onReady(s *discordgo.Session, event *discordgo.Ready) {
    err := s.UpdateStatus(0, "召喚の儀")
    if err != nil {
        log.Println("Error updating status: %v", err)
    }
    log.Println("Recieved READY payload")
}

func onPresenceUpdate(s *discordgo.Session, m *discordgo.PresenceUpdate) {
  guild, _ := s.Guild(m.GuildID)
  pr := guild.Presences
  memcnt := guild.MemberCount
  onlinecnt := 0
  for _, p := range pr {
        if p.Status == "online" || p.Status == "idle" || p.Status == "dnd" {
          onlinecnt++
        }
  }
  err := s.UpdateStatus(0, fmt.Sprintf("Online: %d/%d", onlinecnt, memcnt))
  if err != nil {
      log.Println("Error updating status: %v", err)
  }
  log.Println("Online: %d/%d", onlinecnt, memcnt)
}

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
    c, err := s.State.Channel(m.ChannelID) //チャンネル取得
    if err != nil {
        log.Println("Error getting channel: ", err)
        return
    }
    if m.Author.Username != "コタボット" {
        fmt.Printf("%20s %20s %20s > %s\n", m.ChannelID, time.Now().Format(time.Stamp), m.Author.Username, m.Content)
    }

    isexecute := false
    for _, cmd := range cmds {
        if strings.HasPrefix(m.Content, fmt.Sprintf("%s%s", CmdPrefix, cmd)) {
            botcmds.RunCommand(s, c, CmdPrefix, m.Content)
            isexecute =true
        }
    }
    if strings.HasPrefix(m.Content, fmt.Sprintf(CmdPrefix)) && len(m.Content) > 4 && isexecute != true && m.Content != fmt.Sprintf("%sshutdown", CmdPrefix) {
      sendMessage(s, c, "存在しないコマンドをうつな\n```?helpで調べてくれや```")
    }

    if m.Content == fmt.Sprintf("%sshutdown", CmdPrefix) && m.Author.ID == "202435215375728641" {
      sendMessage(s, c, ":wave:<:kotabi:281065031121108992>")
      err := s.UpdateStatus(3, "")
      if err != nil {
          log.Println("Error updating status: %v", err)
      }
      log.Println("Set status to Offline")
      err = s.Close()
      if err != nil {
          log.Println("Failed to close websocket: %v", err)
      }
      log.Println("Socket closed")
      log.Println("Exit")
      os.Exit(0)
    }
}

//メッセージを送信する関数
func sendMessage(s *discordgo.Session, c *discordgo.Channel, msg string) {
    _, err := s.ChannelMessageSend(c.ID, msg)

    log.Println(">>> " + msg)
    if err != nil {
        log.Println("Error sending message: ", err)
    }
}
