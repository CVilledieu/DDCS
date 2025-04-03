package char

//Character info
type Character struct {
	Name        string
	Class       string
	Level       int
	Health      []int
	Stats       []Stat
	Skills      []Skill
	Items       []Skill
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

type Effect struct {
	Name         string
	BoostedStats Stat
	Duration     int
	EffectType   string // Buff or Debuff
}

// CLASSLIST [12]string = {"Barbarian", "Bard", "Cleric", "Druid", "Fighter", "Monk", "Paladin", "Ranger", "Rogue", "Sorcerer", "Warlock", "Wizard"}

func CreateCharacter(name, class string) *Character {
	return &Character{
		Name:        name,
		Class:       class,
		Level:       1,
		Health:      []int{0},
		Stats:       []Stat{},
		Skills:      []Skill{},
		Items:       []Skill{},
		TempEffects: []Effect{},
	}
}

func (c *Character) GetName() string {
	return c.Name
}

func (c *Character) AddSkill(sk Skill) {
	c.Skills = append(c.Skills, sk)
}

func (c *Character) IncreaseMaxHealth(i int) {
	if i < 0 {
		panic("Health cannot be negative")
	}
	c.Health[0] += i
}
