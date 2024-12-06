package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	table := make([]string, 0)
	for {
		var s string
		_, err := fmt.Fscanf(f, "%s\n", &s)
		if err != nil {
			break
		}
		table = append(table, s)
	}

	sum := 0

	for i := 0; i < len(table); i++ {
		for z := 0; z < len(table[i]); z++ {
			if len(table[i]) > z+3 {
				if table[i][z] == 'X' && table[i][z+1] == 'M' && table[i][z+2] == 'A' && table[i][z+3] == 'S' {
					sum++
				}
			}

			if i-3 >= 0 {
				if table[i][z] == 'X' && table[i-1][z] == 'M' && table[i-2][z] == 'A' && table[i-3][z] == 'S' {
					sum++
				}
			}

			if z-3 >= 0 {
				if table[i][z] == 'X' && table[i][z-1] == 'M' && table[i][z-2] == 'A' && table[i][z-3] == 'S' {
					sum++
				}
			}

			if i-3 >= 0 && z-3 >= 0 {
				if table[i][z] == 'X' && table[i-1][z-1] == 'M' && table[i-2][z-2] == 'A' && table[i-3][z-3] == 'S' {
					sum++
				}
			}

			if i < len(table)-3 && z < len(table[i])-3 {
				if table[i][z] == 'X' && table[i+1][z+1] == 'M' && table[i+2][z+2] == 'A' && table[i+3][z+3] == 'S' {
					sum++
				}
			}

			if i-3 >= 0 && len(table[i]) > z+3 {
				if table[i][z] == 'X' && table[i-1][z+1] == 'M' && table[i-2][z+2] == 'A' && table[i-3][z+3] == 'S' {
					sum++
				}
			}

			if len(table) > i+3 && z-3 >= 0 {
				if table[i][z] == 'X' && table[i+1][z-1] == 'M' && table[i+2][z-2] == 'A' && table[i+3][z-3] == 'S' {
					sum++
				}
			}

			if len(table) > i+3 {
				if table[i][z] == 'X' && table[i+1][z] == 'M' && table[i+2][z] == 'A' && table[i+3][z] == 'S' {
					sum++
				}
			}
		}
	}

	fmt.Println("Answer to part one is", sum)

	sum = 0

	for i := 0; i < len(table); i++ {
		if i == 0 {
			continue
		}
		if i == len(table)-1 {
			continue
		}
		for z := 0; z < len(table[i]); z++ {
			if table[i][z] != 'A' {
				continue
			}

			if z == 0 {
				continue
			}

			if z == len(table[i])-1 {
				continue
			}

			if table[i+1][z+1] == 'M' {
				if table[i-1][z-1] != 'S' {
					continue
				}
			} else if table[i+1][z+1] == 'S' {
				if table[i-1][z-1] != 'M' {
					continue
				}
			} else {
				continue
			}

			if table[i+1][z-1] == 'M' {
				if table[i-1][z+1] == 'S' {
					sum++
				}
			} else if table[i+1][z-1] == 'S' {
				if table[i-1][z+1] == 'M' {
					sum++
				}
			}
		}
	}

	fmt.Println("Answer to part two is", sum)
}
