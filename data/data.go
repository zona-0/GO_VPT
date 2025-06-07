package data

// This is global function
// TODO: Structs data sec
// TODO: Skill struct sec
type Skill struct {
	Name string
}

// TODO: exp struct sec
type Experience struct {
	Title, Company, Description string
}

// TODO: educations struct sec
type Education struct {
	School, Degree, Major string
	Year                  int
}

// TODO: Skill sec
var Skills [100]Skill
var SkillCount int

// TODO: Exp sec
var Experiences [100]Experience
var ExperienceCount int

// TODO: Educations sec
var Educations [100]Education
var EducationCount int
