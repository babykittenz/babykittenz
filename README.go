package developer

// Skill represents a technical skill
type Skill string

// Define skill constants
const (
	SkillPhp             Skill = "PHP"
	SkillTypescript      Skill = "TypeScript"
	SkillLaravel         Skill = "Laravel"
	SkillReactNative     Skill = "React Native"
	SkillReact           Skill = "React"
	SkillNextjs          Skill = "Next.js"
	SkillScssCss         Skill = "SCSS/CSS"
	SkillResponsiveDesign Skill = "Responsive Design"
	SkillWordPress       Skill = "WordPress"
	SkillMicroservices   Skill = "Microservices"
)

// Workplace describes employment information
type Workplace struct {
	Company  string
	Position string
}

// ProfileInfo defines the interface for developer profile information
type ProfileInfo interface {
	CurrentWorkplace() Workplace
	DailySkills() []Skill
	FutureGoal() string
}

// Developer implements the ProfileInfo interface
type Developer struct {
	name string
}

// NewDeveloper creates a new Developer instance
func NewDeveloper(name string) *Developer {
	return &Developer{
		name: name,
	}
}

// CurrentWorkplace returns current employment information
func (d *Developer) CurrentWorkplace() Workplace {
	return Workplace{
		Company:  "millMountainDigital",
		Position: "owner",
	}
}

// DailySkills returns a slice of technical skills used daily
func (d *Developer) DailySkills() []Skill {
	return []Skill{
		SkillPhp,
		SkillTypescript,
		SkillLaravel,
		SkillReact,
		SkillNextjs,
		SkillReactNative,
		SkillScssCss,
		SkillResponsiveDesign,
		SkillWordPress,
		SkillMicroservices,
	}
}

// FutureGoal returns aspirational information
func (d *Developer) FutureGoal() string {
	return "To contribute to open source, marketing, and digital art projects."
}

// Name returns the developer's name
func (d *Developer) Name() string {
	return d.name
}
