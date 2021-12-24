package main

func checkCollision(x1, y1, x2, y2 float64, box1w, box1h, box2w, box2h float64) bool {
	return x1 < x2+box2w &&
		x1+box1w > x2 &&
		y1 < y2+box2h &&
		box1h+y1 > y2
}

func (g *Game) resolveCollisions() {
	for _, ghost := range g.ghosts {
		for _, shot := range g.shots {
			if checkCollision(ghost.x, ghost.y, shot.x, shot.y, 16.0, 16.0, 16.0*ShotScalingFactor, 16.0*ShotScalingFactor) {
				g.points++
				ghost.Kill()
				shot.Destroy()
			}
		}
	}

	for _, pac := range g.pacs {
		for _, ghost := range g.ghosts {
			if checkCollision(pac.x, pac.y, ghost.x, ghost.y, 32.0*PacScalingFactor, 32.0*PacScalingFactor, 16.0, 16.0) {
				g.points -= 10
				pac.Kill()
			}
		}
	}

	for _, pac := range g.pacs {
		for _, shot := range g.shots {
			if checkCollision(pac.x, pac.y, shot.x, shot.y, 32.0*PacScalingFactor, 32.0*PacScalingFactor, 16.0*ShotScalingFactor, 16.0*ShotScalingFactor) {
				g.points += 10
				pac.Kill()
				shot.Destroy()
			}
		}
	}
}
