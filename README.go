package main

import "fmt"

// Skills type represents various technical skills
type Skills string

// Define skill constants
const (
	Php            Skills = "Php"
	Javascript     Skills = "Javascript"
	Laravel        Skills = "Laravel"
	ReactNative    Skills = "ReactNative"
	ScssCss        Skills = "ScssCss"
	ResponsiveDesign Skills = "ResponsiveDesign"
	WordPress      Skills = "WordPress"
)

// Workplace represents employment information
type Workplace struct {
	Company  string
	Position string
}

// Me is a base struct that About extends
type Me struct{}

// About extends the Me struct to provide personal information
type About struct {
	Me
}

// GetCurrentWorkplace returns information about current employment
func (a *About) GetCurrentWorkplace() map[string]Workplace {
	return map[string]Workplace{
		"workplace": {
			Company:  "millMountainDigital",
			Position: "owner",
		},
	}
}

// GetDailyKnowledge returns an array of technical skills
func (a *About) GetDailyKnowledge() []Skills {
	return []Skills{
		Php,
		Javascript,
		Laravel,
		ReactNative,
		ScssCss,
		ResponsiveDesign,
		WordPress,
	}
}

// GetFutureGoal returns a string describing future aspirations
func (a *About) GetFutureGoal() string {
	return "To contribute to open source, marketing, and digital art projects."
}

func main() {
	about := &About{}
	
	fmt.Println("Current Workplace:")
	for _, workplace := range about.GetCurrentWorkplace() {
		fmt.Printf("Company: %s, Position: %s\n", workplace.Company, workplace.Position)
	}
	
	fmt.Println("\nDaily Knowledge:")
	for _, skill := range about.GetDailyKnowledge() {
		fmt.Println("-", skill)
	}
	
	fmt.Println("\nFuture Goal:")
	fmt.Println(about.GetFutureGoal())
}
