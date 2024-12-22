package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const modulo = 16777216

type Sequence struct {
	x1 int
	x2 int
	x3 int
	x4 int
}

type BananaBank struct {
	gaps   [2000]int
	banana [2000]int
}

func (s *Sequence) add(x int) {
	s.x1, s.x2, s.x3 = s.x2, s.x3, s.x4
	s.x4 = x
}

func removeEmptyStrings(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func readInput() []int {
	fileName := "input.txt"
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	dataString := string(data)
	secretsStr := strings.Split(dataString, "\n")
	secretsStr = removeEmptyStrings(secretsStr)
	secrets := make([]int, len(secretsStr))
	for i, secretStr := range secretsStr {
		secret, _ := strconv.Atoi(secretStr)
		secrets[i] = secret
	}
	return secrets
}

func generateNewNumber(x int) int {
	initialSecret := x
	x *= 64
	x ^= initialSecret
	x %= modulo
	initialSecret = x
	x /= 32
	x ^= initialSecret
	x %= modulo
	initialSecret = x
	x *= 2048
	x ^= initialSecret
	x %= modulo
	return x
}

func checksum(secrets []int) int {
	sum := 0
	for _, secret := range secrets {
		for i := 0; i < 2000; i++ {
			secret = generateNewNumber(secret)
		}
		sum += secret
	}
	return sum
}

func (b *BananaBank) fillBank(secret int) {
	var oldSecret int
	b.gaps[0] = 98
	for i := 1; i < 2000; i++ {
		oldSecret = secret
		secret = generateNewNumber(secret)
		b.gaps[i] = (secret % 10) - (oldSecret % 10)
		b.banana[i] = secret % 10
	}
}

func fillBananaBanks(secrets []int) []BananaBank {
	bananaBanks := make([]BananaBank, len(secrets))
	for i, secret := range secrets {
		b := BananaBank{}
		b.fillBank(secret)
		bananaBanks[i] = b
	}
	return bananaBanks
}

func getNumberFromSequence(s Sequence, b BananaBank) int {
	currentSequence := Sequence{99, 99, 99, 99}
	for i := 1; true; i++ {
		if i == 2000 {
			return 0
		}
		currentSequence.add(b.gaps[i])
		if currentSequence == s {
			return b.banana[i]
		}
	}
	return 0
}

func getBananaCountFromSequence(s Sequence, bananaBanks []BananaBank) int {
	sum := 0
	for _, bananaBank := range bananaBanks {
		sum += getNumberFromSequence(s, bananaBank)
	}
	return sum
}

func brutforceBananaCount(bananaBanks []BananaBank) int {
	var bananaCount int
	max := 0
	for i := -5; i < 6; i++ { // Gap above 5 are rare and will probably generate a lot of 0
		for j := -5; j < 6; j++ {
			for k := -5; k < 6; k++ {
				for l := 0; l < 6; l++ { // We want a positive gap for the final one
					bananaCount = getBananaCountFromSequence(Sequence{i, j, k, l}, bananaBanks)
					if bananaCount > max {
						max = bananaCount
					}
				}
			}
		}
	}
	return max
}

func main() {
	secrets := readInput()
	fmt.Printf("Part 1 : %d\n", checksum(secrets))

	bananaBanks := fillBananaBanks(secrets)
	fmt.Printf("Part 2 : %d\n", brutforceBananaCount(bananaBanks))
}
