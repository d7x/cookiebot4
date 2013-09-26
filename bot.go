package main

import (
	"log"
//	"runtime"
)

type BotConfig struct {
	Username string
	Password string
	Cdkey string
	Server string
}

type Bot struct {
	ProfileName string
	Config BotConfig

	CdkeyData *CdkeyData
	ExeInfo *ExeInfo

	Bncs *BncsSocket
	Bnls *BnlsSocket
}

type CdkeyData struct {
	ClientToken int
	KeyLength int
	ProductValue int
	PublicValue int
	CdkeyHash []int
}

type ExeInfo struct {
	Version int
	Checksum int
	StatString []byte
	Cookie int
	VersionByte int
}

func NewBot(profileName string, config BotConfig) *Bot {
	b := new(Bot)
	b.ProfileName = profileName
	b.Config = config

	return b
}

func (b *Bot) Connect(bnls *BnlsSocket) bool {
	b.Bnls = bnls
	b.Bncs = NewBncsSocket(b)
	log.Printf("[bncs] %s connecting to %s...", b.ProfileName, b.Config.Server)
	err := b.Bncs.Connect(b.Config.Server)
	if err != nil {
		log.Printf("[bncs] %s failed to connect to %s [%s]", b.ProfileName, b.Config.Server, err.Error())
		return false
	}

	log.Printf("[bncs] %s successfully connected to %s", b.ProfileName, b.Bncs.Server)
	log.Printf("[bncs] (%s) sending protocol byte (0x01)", b.ProfileName)
	err = b.Bncs.SendProtocolByte()
	if err != nil {
		log.Printf("[bncs] (%s) failed to send protocol byte [%s]", b.ProfileName, err.Error())
		return false
	}
	b.Bncs.SendSid_AuthInfo()
	
	return true
}