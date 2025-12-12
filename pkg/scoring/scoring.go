package scoring

func CalculateScore(t, pfc, rc int) (score int) {
	score += proficiencyScore(pfc)
	score += timeScore(t, 20)
	score += timeScore(rc, 10)
	score /= 3

	return
}

func timeScore(t, limit int) int {
	if t > limit {
		return 0
	} else if t == limit {
		return 1
	}

	x := float64(10) / float64(limit)
	return 10 - int(float64(t)*x)
}

func proficiencyScore(pfc int) int {
	if pfc <= -16 {
		return 10
	} else if pfc >= 16 {
		return 0
	}

	return (-1 * pfc / 3) + 5
}
