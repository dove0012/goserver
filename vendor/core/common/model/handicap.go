package model

type Handicap struct {
	Han_id  int64  `bson:"han_id"`
	Game_id int64  `bson:"game_id"`
	Name    string `bason:"name"`
}
