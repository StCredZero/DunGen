package main

func main() {
	dungeon := &DunGen{
		xsize:      80,
		ysize:      24,
		targetObj:  20,
		chanceRoom: 60,
	}
	dungeon.createDungeon(GridCoord{0, 0}, DunGenEntropy{6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6})

	println(dungeon.debugPrint())
}
