package game

import (
	"fmt"
	"strings"
)

func checkNotInitialized(g *game) error {
	errMsgMap := map[string]func(g *game) bool{
		"board":               func(g *game) bool { return g.board == nil },
		"players":             func(g *game) bool { return g.playerSet == nil },
		"initialize function": func(g *game) bool { return g.initialize == nil },
		"mainphase function":  func(g *game) bool { return g.mainPhase == nil },
		"cleanUp function":    func(g *game) bool { return g.cleanUp == nil },
	}
	errMsgs := make([]string, 0)
	for msg, isNotInit := range errMsgMap {
		if isNotInit(g) {
			errMsgs = append(errMsgs, msg)
		}
	}
	if len(errMsgs) > 0 {
		return fmt.Errorf("[%s] not initialized", strings.Join(
			errMsgs, ", ",
		))
	}
	return nil
}

func checkMoreThanOnePlayer(g *game) error {
	if len(g.playerSet) <= 1 {
		return fmt.Errorf("there should be more than one player")
	}
	return nil
}
