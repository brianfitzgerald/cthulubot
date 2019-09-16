package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	mentioned := false
	for _, user := range m.Mentions {
		if user.ID == botUserID {
			mentioned = true
		}
	}

	if !mentioned {
		return
	}

	message := m.Message.Content

	fmt.Println(message)
	if strings.Contains(message, "?") {
		response := askQuestion(message)
		translatedResponse := translate(response)
		s.ChannelMessageSend(m.ChannelID, translatedResponse)
	} else {
		translatedMessage := translate(message)
		s.ChannelMessageSend(m.ChannelID, translatedMessage)
	}
}

// Variables used for command line parameters
var (
	Token            = "NjIzMjMzOTY5NTU0NzE4NzIx.XX_fRg.RHT2V2LRl8nyhIr2M5Bnp3Gd-lw"
	useDiscord       = true
	botUserID        = "623233969554718721"
	generalChannelID = "559867622439518227"
)

func main() {
	discordInit()
}

func askQuestion(message string) string {

	responses := []string{
		"Sounds like a good idea",
		"I would not recommend it",
		"I’d rather just eat humans",
		"I’ll tell you after I’ve eaten my daily sacrifice",
		"Don’t count on it",
		"Seems sketchy",
		"I, Cthulhu, would not recommend it",
	}

	return responses[rand.Intn(len(responses))]

}

func discordInit() {
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()

}

func translate(inputStr string) string {
	rand.Seed(time.Now().Unix())

	var words string = "Ya vulgt shogg va goka y gotha Fm’latgh shuggoth syha’h Ya vulgt uh’e hupadgh shugg Gnaiih, ‘fhalmaa, ng gof’nn Mnahn’ hlirghh Lk’cvol ng mnulye mg jhn arw N’gha Ya Yvulgtm yron ooboshu rn shugg Tharanak yr’luhh grah’nn F’dkahnn, ya wgahn f’mnahnog tharanak li’hee syha’h fhtagn ch’ ebumna Ep uln skri lloigg ng ‘bthnkk Hrii gh Ronnyth Mgah'ehye Nilgh'nahyar  agll ah ah'gotha ah'kn'aa ah ahagl ahagl ahaimgr'luhh ahehyee ahephai ahf ahh ahhai ahlloig ahlloigehye ahna ahnah ahnythor ahog ahor ahornah ahuaaah ai azanahoth bthnk bthnknahor. bthnknahorr bugnah cahf ehye ehyee ehyeeog ehyeog epgoka ephai ephaiaglor ephaiah ephailease ephaimggokai ephaimgr'luh ephaimgsyha'h fhtagn gn'th'bthnkor gokahe hafh hnah hnahh k'yarnak kadishtu lllln'gha llllw'nafh lloig lloigehye lloighnahh lw'nafh lw'nafh'drn lw'nafhnah mgah mgah'ehye mgah'n'ghft mgehye mgep mgep'ai mgepah mgepah'kn'a mgepah. mgepahlloig mgepbug mgephai mgepmgah'n'ghft mgepmggokaog mgepmgr'luh mgepnogephaii mgepog mgepogg. mgepuaaah mgepyarog mgepzhro mgfm'latghnanah mggoka mggoka'ai mglagln mgleth mgn'ghftephai. mgng mgr'luh mgsyha'h mguh'e. mgvulgtnah n'ghft n'ghftor n'ghftorog n'ghftyar na'ah'ehye naIIII nafl nilgh'nahyar nilgh'ri nilgh'rinah nng nnnkadishtuor nog ph'lloig ph'nglui r'luhhor riuh'eor shuggoth shugnahoth uaaah uh'e uh'eaglnah; uh'eor vulgtmnah vulgtmnahog vulgtmnahor vulgtmog wgah'nagl y'or'nahyah'or'nahh yah'or'nanah yah'or'nanah yog yogagl yogfm'll yogor you’re zhro"
	cthuluWords := strings.Split(words, "")
	inputWords := strings.Split(inputStr, "")
	wordMap := make(map[string]string)
	var outputString string
	for _, word := range inputWords {
		var replacementWord string
		if wordMap[word] == "" {
			replacementWord = cthuluWords[rand.Intn(len(cthuluWords))]
			wordMap[word] = replacementWord
		} else {
			replacementWord = wordMap[word]
		}
		outputString += replacementWord
	}
	println(outputString)
	return outputString
}
