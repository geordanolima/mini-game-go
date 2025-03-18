package engine

import "mini-game-go/domain/entitie"

func verifyConflictInObstacles(object entitie.Object, listObjects []entitie.Obstacle) bool {
	results := make([]bool, len(listObjects))
	for _, item := range listObjects {
		hasConflict := verifyConflict(object, item.Object, true)
		results = append(results, hasConflict)
	}
	for _, item := range results {
		if item {
			return item
		}
	}
	return false
}
