package solutions

import (
	utils "adventofcode/utils"
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	fields map[string]string
}

func (p *Passport) validByr() bool {
	value, ok := p.fields["byr"]
	if !ok {
		return false
	}

	asInt, err := strconv.Atoi(value)
	if err != nil || asInt < 1920 || asInt > 2002 {
		return false
	}

	return true
}

func (p *Passport) validIyr() bool {
	value, ok := p.fields["iyr"]
	if !ok {
		return false
	}

	asInt, err := strconv.Atoi(value)
	if err != nil || asInt < 2010 || asInt > 2020 {
		return false
	}

	return true
}

func (p *Passport) validEyr() bool {
	value, ok := p.fields["eyr"]
	if !ok || len(value) != 4 {
		return false
	}

	asInt, err := strconv.Atoi(value)
	if err != nil || asInt < 2020 || asInt > 2030 {
		return false
	}

	return true
}

func (p *Passport) validHgt() bool {
	value, ok := p.fields["hgt"]
	if !ok || len(value) < 3 {
		return false
	}

	asInt, err := strconv.Atoi(value[:len(value)-2])
	if err != nil {
		return false
	}

	metric := value[len(value)-2:]
	if metric != "in" && metric != "cm" {
		return false
	}

	if metric == "in" && (asInt < 59 || asInt > 76) {
		return false
	}

	if metric == "cm" && (asInt < 150 || asInt > 193) {
		return false
	}

	return true
}

func (p *Passport) validHcl() bool {
	value, ok := p.fields["hcl"]
	if !ok {
		return false
	}

	if value[0] != '#' || len(value) != 7 {
		return false
	}

	var isStringAlphabetic = regexp.MustCompile(`^[a-f0-9_]*$`).MatchString
	return isStringAlphabetic(value[1:])
}

func (p *Passport) validEcl() bool {
	value, ok := p.fields["ecl"]
	if !ok {
		return false
	}

	var eyes = map[string]bool{
		"amb": true,
		"blu": true,
		"brn": true,
		"gry": true,
		"grn": true,
		"hzl": true,
		"oth": true,
	}

	_, ok = eyes[value]
	return ok
}

func (p *Passport) validPid() bool {
	value, ok := p.fields["pid"]
	if !ok || len(value) != 9 {
		return false
	}

	var isStringAlphabetic = regexp.MustCompile(`^[0-9_]*$`).MatchString
	return isStringAlphabetic(value[1:])
}

func (p *Passport) valid() bool {

	//neededKeys := [7]string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	if !p.validByr() {
		return false
	}
	if !p.validIyr() {
		return false
	}
	if !p.validEyr() {
		return false
	}
	if !p.validHgt() {
		return false
	}
	if !p.validHcl() {
		return false
	}
	if !p.validEcl() {
		return false
	}
	if !p.validPid() {
		return false
	}

	//for _, key := range neededKeys {
	//	_, ok := p.fields[key]
	//	if !ok {
	//			return false
	//		}
	//	}

	return true
}

func NewPassport(data string) (passport Passport) {

	passport.fields = make(map[string]string)

	fields := strings.Fields(data)

	for _, field := range fields {

		keyValue := strings.Split(field, ":")
		passport.fields[keyValue[0]] = keyValue[1]

	}

	return
}

func fileAsPassports() (passports []Passport) {

	inputs := utils.FileAsLines("input/day4.txt")

	data := ""
	for _, input := range inputs {

		if input == "" {
			passports = append(passports, NewPassport(data))
			data = ""
		}

		data += " "
		data += input
	}
	passports = append(passports, NewPassport(data))

	return
}

func SolveDay4() (string, string) {

	part1 := 0
	part2 := 0

	passports := fileAsPassports()
	for _, passport := range passports {
		if passport.valid() {
			part1 += 1
		}
	}

	return utils.ResultStrings(part1, part2)
}
