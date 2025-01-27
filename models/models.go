package models

type room struct {
	name string
	x string
	y string
}
type AntFarm struct {
	ants int
	rooms []room
	From string
	To string
}