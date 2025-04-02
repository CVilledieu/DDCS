package char

//Character info

type Character struct {
	Name        string
	Class       string
	Level       int
	Health      int
	Stats       Stat
	Skills      []Skill
	Items       []Item
	TempEffects []Effect
}

type Stat struct {
	Strength     int
	Intelligence int
	Dexterity    int
	Constitution int
	Wisdom       int
	Charisma     int
	ArmorClass   int
	Speed        int
}

type Skill struct {
	Name          string
	Description   string
	RequiredLevel int
	Damage        int
	Range         int
}
