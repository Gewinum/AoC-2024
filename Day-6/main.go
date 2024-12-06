package main

import (
    "bufio"
    "errors"
    "fmt"
    "os"
    "strconv"
)

type Direction int

func (d Direction) IsNorth() bool {
    return d == 1
}

func (d Direction) IsEast() bool {
    return d == 2
}

func (d Direction) IsSouth() bool {
    return d == 3
}

func (d Direction) IsWest() bool {
    return d == 4
}

func (d Direction) GetSymbolByFacing() byte {
    if d.IsNorth() {
        return '^'
    } else if d.IsEast() {
        return '>'
    } else if d.IsSouth() {
        return 'v'
    } else if d.IsWest() {
        return '<'
    }

    panic(errors.New("unknown direction - " + strconv.Itoa(int(d))))
}

func getDirectionByFacing(bt byte) Direction {
    if bt == '^' {
        return 1
    } else if bt == '>' {
        return 2
    } else if bt == 'v' {
        return 3
    } else if bt == '<' {
        return 4
    }

    panic(errors.New("unknown direction " + string(bt)))
}

func main() {
    f, err := os.Open("input.txt")
    if err != nil {
        panic(err)
    }
    scanner := bufio.NewScanner(f)
    table := make([][]byte, 0)
    for scanner.Scan() {
        line := scanner.Text()
        table = append(table, []byte(line))
    }
    alreadyWas := make(map[int][]int)
    count := 0
    for {
        newX, newY := makeGuardMove(table)
        if newX == -1 {
            break
        }
        if _, ok := alreadyWas[newY]; !ok {
            alreadyWas[newY] = make([]int, 0)
        }
        isExisting := false
        for j := 0; j < len(alreadyWas[newY]); j++ {
            if alreadyWas[newY][j] == newX {
                isExisting = true
                break
            }
        }
        if !isExisting {
            alreadyWas[newY] = append(alreadyWas[newY], newX)
            count++
        }
    }
    fmt.Println("Answer to part one is", count)

    f.Close()
    f = nil
    // part 2
    f, err = os.Open("input.txt")
    if err != nil {
        panic(err)
    }
    scanner = bufio.NewScanner(f)
    initialTable := make([][]byte, 0)
    for scanner.Scan() {
        line := scanner.Text()
        initialTable = append(initialTable, []byte(line))
    }
    table = make([][]byte, 0)
    for j := 0; j < len(initialTable); j++ {
        hm := make([]byte, 0)
        for i := 0; i < len(initialTable[j]); i++ {
            hm = append(hm, initialTable[j][i])
        }
        table = append(table, hm)
    }
    startX, startY := determineGuardPosition(table)
    count = 0
    for checkY, checkXList := range alreadyWas {
        for _, checkX := range checkXList {
            table = make([][]byte, 0)
            for j := 0; j < len(initialTable); j++ {
                hm := make([]byte, 0)
                for i := 0; i < len(initialTable[j]); i++ {
                    hm = append(hm, initialTable[j][i])
                }
                table = append(table, hm)
            }
            if checkX == startX && checkY == startY {
                continue
            }
            if initialTable[checkY][checkX] == '#' {
                continue
            }
            table[checkY][checkX] = '#'

            iHateMyself := make(map[int]map[int][]Direction)

            for {
                newX, newY := makeGuardMove(table)
                if newX == -1 {
                    break
                }
                if _, ok := iHateMyself[newY]; !ok {
                    iHateMyself[newY] = make(map[int][]Direction)
                }
                if _, ok := iHateMyself[newY][newX]; !ok {
                    iHateMyself[newY][newX] = make([]Direction, 0)
                }
                whatIsLove := false
                for j := 0; j < len(iHateMyself[newY][newX]); j++ {
                    if iHateMyself[newY][newX][j] == getDirectionByFacing(table[newY][newX]) {
                        count++
                        whatIsLove = true
                    }
                }
                if whatIsLove {
                    break
                } else {
                    iHateMyself[newY][newX] = append(iHateMyself[newY][newX], getDirectionByFacing(table[newY][newX]))
                }
            }
        }
    }

    fmt.Println("Answer to part two", count)
}

// determineGuardPosition will return -1, -1 in case guard is no longer on map anymore
func determineGuardPosition(table [][]byte) (int, int) {
    for y := 0; y < len(table); y++ {
        for x := 0; x < len(table[y]); x++ {
            if table[y][x] == '^' || table[y][x] == '>' || table[y][x] == 'v' || table[y][x] == '<' {
                return x, y
            }
        }
    }
    return -1, -1
}

func makeGuardMove(table [][]byte) (int, int) {
    x, y := determineGuardPosition(table)
    if x == -1 {
        return -1, -1
    }
    dir := getDirectionByFacing(table[y][x])

    newX, newY := x, y

    if dir.IsNorth() {
        newY -= 1
    } else if dir.IsEast() {
        newX += 1
    } else if dir.IsSouth() {
        newY += 1
    } else if dir.IsWest() {
        newX -= 1
    }

    if newX < 0 || newY < 0 {
        return -1, -1
    }

    if newY >= len(table) || newX >= len(table[newY]) {
        return -1, -1
    }

    if table[newY][newX] == '#' {
        dir += 1
        if dir > 4 {
            dir -= 4
        }
        table[y][x] = dir.GetSymbolByFacing()
        return x, y
    }

    table[y][x] = '.'
    table[newY][newX] = dir.GetSymbolByFacing()
    return newX, newY
}
