package player

type character struct {
	walkCycleOffset walkCycleOffset
	initialPower    int
	initialStepSize int
}

var BlackBoy = character{
	walkCycleOffset: walkCycleOffset{
		x: 0,
		y: spriteColHeight,
	},
	initialStepSize: 4,
	initialPower:    1,
}

var BlackGirl = character{
	walkCycleOffset: walkCycleOffset{
		x: 0,
		y: 0,
	},
	initialStepSize: 4,
	initialPower:    1,
}

var BrownBoy = character{
	walkCycleOffset: walkCycleOffset{
		x: spriteRowWidth,
		y: spriteColHeight,
	},
	initialStepSize: 6,
	initialPower:    1,
}

var BrownGirl = character{
	walkCycleOffset: walkCycleOffset{
		x: spriteRowWidth,
		y: 0,
	},
	initialStepSize: 6,
	initialPower:    1,
}

var YellowBoy = character{
	walkCycleOffset: walkCycleOffset{
		x: spriteRowWidth * 2,
		y: spriteColHeight,
	},
	initialStepSize: 6,
	initialPower:    1,
}

var YellowGirl = character{
	walkCycleOffset: walkCycleOffset{
		x: spriteRowWidth * 2,
		y: 0,
	},
	initialStepSize: 6,
	initialPower:    1,
}

var OrangeBoy = character{
	walkCycleOffset: walkCycleOffset{
		x: spriteRowWidth * 3,
		y: spriteColHeight,
	},
	initialStepSize: 6,
	initialPower:    1,
}

var OrangeGirl = character{
	walkCycleOffset: walkCycleOffset{
		x: spriteRowWidth * 3,
		y: 0,
	},
	initialStepSize: 6,
	initialPower:    1,
}
