package main

import (
	"log"

	"github.com/chippydip/go-sc2ai/api"
	"github.com/chippydip/go-sc2ai/botutil"
	"github.com/chippydip/go-sc2ai/client"
	"github.com/chippydip/go-sc2ai/runner"
)

type bot struct {
	*botutil.Bot
}

func main() {
	// Play a random map against a medium difficulty computer
	runner.SetComputer(api.Race_Random, api.Difficulty_Easy, api.AIBuild_RandomBuild)
	runner.SetGameVersion(92440, "A389D1F7DF9DD792FBE980533B7119FF")
	runner.SetMap("C:\\Program Files (x86)\\StarCraft II\\Maps\\Equilibrium513AIE.SC2Map")

	agent := client.AgentFunc(runAgent)
	runner.RunAgent(client.NewParticipant(api.Race_Protoss, agent, "StubBot"))
}

func runAgent(info client.AgentInfo) {
	bot := bot{Bot: botutil.NewBot(info)}
	bot.LogActionErrors()

	bot.init()
	for bot.IsInGame() {
		bot.doSmt()

		if err := bot.Step(1); err != nil {
			log.Print(err)
			break
		}
	}
}

func (bot *bot) init() {
	// Send a friendly hello
	bot.Chat("(glhf)")
}

func (bot *bot) doSmt() {

}
