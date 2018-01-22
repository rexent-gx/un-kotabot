package botcmds

import (
    "fmt"
    "time"
    "github.com/bwmarrin/discordgo"
    "log"
    "strings"
    "math/rand"
    "embed"
    //"encoding/json"
)

func GetCommands() []string {
    cmds := []string{
        "hi",
        "takustat",
        "status",
        "ullasummon",
        "help",
    }
    return cmds
}

func RunCommand(s *discordgo.Session, c *discordgo.Channel, CmdPrefix string, msg string) {
    guild, _ := s.Guild(c.GuildID)
    vs := guild.VoiceStates
    pr := guild.Presences
    member := guild.Members

    var isjoinvc bool
    var sendText string
    switch {
        case msg == fmt.Sprintf("%shi", CmdPrefix):
            reply := []string{
                "おれはボミ",
                "ラ",
                "陰キャだまれや",
                "ころすぞ",
                "口臭いぞ",
                "黙れ黙れ黙れ便器に向かって話しかけてろ",
                "ホホ＾～",
                "パァ＾～～～",
                "ガイジしんでくれ",
            }
            i := makeRand(len(reply))
            sendMessage(s, c, reply[i])

        case msg == fmt.Sprintf("%stakustat", CmdPrefix):
            status, playing, vcjoin := "", "", ""
            for _, p := range pr {
                if p.User.ID == "196502291812057089" {
                  status = fmt.Sprintf("%v", p.Status)
                  if p.Game != nil {
                      playing = fmt.Sprintf("%vを家のリビングでやってる", p.Game.Name)
                  } else {
                      playing = fmt.Sprintf("なんかやれや")
                  }
                }
            }
            for _, v := range vs {
                if v.UserID == "196502291812057089" {
                    isjoinvc = true
                }
            }
            if isjoinvc == true {
                vcjoin = fmt.Sprintf("たくたくは通話にいるよ・・・")
            } else {
                vcjoin = fmt.Sprintf("taku8730がいないぞ")
            }

            //なんかforの中で宣言したら怒られたやつ
            url := ""

            for _, mem := range member {
              user := mem.User
              if mem.User.ID == "196502291812057089" {
                url = user.AvatarURL("64")
              }
            }

            color := randColor()
            embed := embed.NewEmbed().
              SetTitle("タク・ステータス").
              SetDescription(strings.ToUpper(status)).
              AddField("タク・ゲーミング", playing, true).
              AddField("タク・ボイス", vcjoin, true).
              SetFooter("un-kotabot v0.0.3 α").
              SetThumbnail(url).
              SetTimestamp(time.Now().Format("2006-01-02T15:04:05-07:00")).
              SetColor(color).MessageEmbed

            sendEmbedMessage(s, c, embed)

        case strings.HasPrefix(msg, fmt.Sprintf("%sstatus", CmdPrefix)):
            if strings.Contains(msg, "status <@") {
                start := strings.Index(msg, "@")
                end := strings.Index(msg, ">")
                uid := msg[start+1:end]
                uid = strings.Trim(uid, "!")

                status, playing, vcjoin := "", "", ""

                for _, p := range pr {
                    if p.User.ID == uid {
                      status = fmt.Sprintf("%v", p.Status)
                      if p.Game != nil {
                          playing = fmt.Sprintf("%vやってる", p.Game.Name)
                      } else {
                          playing += fmt.Sprintf("なんもしてない")
                      }
                    }
                }
                for _, v := range vs {
                    if v.UserID == uid {
                        isjoinvc = true
                    }
                }
                if isjoinvc == true {
                    vcjoin = fmt.Sprintf("なんか喋ってる")
                } else {
                    vcjoin = fmt.Sprintf("いませんでした・・・")
                }

                url, uname := "", ""

                for _, mem := range member {
                  user := mem.User
                  if mem.User.ID == uid {
                    url = user.AvatarURL("64")
                    if mem.Nick != ""{
                      uname = mem.Nick
                    } else {
                      uname = mem.User.Username
                    }
                  }
                }

                color := randColor()
                embed := embed.NewEmbed().
                  SetTitle(fmt.Sprintf("%sのステータス", uname)).
                  SetDescription(strings.ToUpper(status)).
                  AddField("プレイ中", playing, true).
                  AddField("通話", vcjoin, true).
                  SetFooter("un-kotabot v0.0.3 α").
                  SetThumbnail(url).
                  SetTimestamp(time.Now().Format("2006-01-02T15:04:05-07:00")).
                  SetColor(color).MessageEmbed

                sendEmbedMessage(s, c, embed)

            } else {
              sendMessage(s, c, "コマンドがちがうぞボケ\n```?helpで調べてくれや```")
            }

          case msg == fmt.Sprintf("%sullasummon", CmdPrefix):
              msgnum := makeRand(5)
              for i := 0; i < msgnum; i++ {
                msglen := makeRand(10)
                for i := 0; i < msglen; i++ {
                  sendText += "<@245188551128514561>"
                }
                sendMessage(s, c, sendText)
              }

          case msg == fmt.Sprintf("%shelp", CmdPrefix):
            color := randColor()
            embed := embed.NewEmbed().
              SetTitle(":poop:ウンコタボットの使い方<:kotamori:384528257699151873>").
              SetDescription("（プログラミングのれんしゅうでつくっているあわれなBotだよ。やさしくしてね<:kotabi:281065031121108992>）").
              AddField("?hi", "こたびちゃんがあいさつをするよ！", false).
              AddField("?takustat", "サーバー名物\"タク\"について表示するよ！", false).
              AddField("?status [@ユーザー名]", "指定したユーザーについて表示するよ！", false).
              AddField("?ullasummon", "**SPAM!**", false).
              SetFooter("un-kotabot v0.0.3 α").
              SetTimestamp(time.Now().Format("2006-01-02T15:04:05-07:00")).
              SetColor(color).MessageEmbed

            sendEmbedMessage(s, c, embed)
      }
  return
}

func makeRand(l int) int {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(l)
}

func randColor() int {
    color := []int{
      0xba062f,
      0x00ff00,
      0x604ffc,
      0x30465e,//ao
      0xddefd3,//siro
      0x6427a4,//purple
      0x41b9a9,//bluegreen
      0x7d882d,//oleve
      0x42bd5d,//emerald
      0xccb043,//darkyellow
    }
    i := makeRand(len(color))
    return color[i]
}

func sendMessage(s *discordgo.Session, c *discordgo.Channel, msg string) {
    _, err := s.ChannelMessageSend(c.ID, msg)

    log.Println(">>> Message send")
    if err != nil {
        log.Println("Error sending message: ", err)
    }
}

func sendEmbedMessage(s *discordgo.Session, c *discordgo.Channel, embed *discordgo.MessageEmbed) {
    _, err := s.ChannelMessageSendEmbed(c.ID, embed)

    log.Println(">>> Embed Message send")
    if err != nil {
        log.Println("Error sending message: ", err)
    }
}
