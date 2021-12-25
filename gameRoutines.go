package main

import "time"

func (g *Game) spawnGhostsRoutine() {
	for {
		s := RandIntInRange(GhostSpawnIntervalMin, GhostSpawnIntervalMax)
		time.Sleep(time.Duration(s * int(time.Millisecond)))
		g.ghosts = append(g.ghosts, g.spawnGhosts(g.win, 1)...)
	}
}

func (g *Game) spawnPacsRoutine() {
	for {
		s := RandIntInRange(PacSpawnIntervalMin, PacSpawnIntervalMax)
		time.Sleep(time.Duration(s * int(time.Millisecond)))
		g.pacs = append(g.pacs, g.spawnPacs(g.win, 1)...)
	}
}
