package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var pathFile string
var buffer []string
var posisiQueens [][]int
var row int
var col int
var realBoard [][]string
var solFound int = 0
var iterationCount int = 0

var pilihanStrategi int
var freqVisualisasi int

func main() {
	fmt.Println("N-QUEENS SOLVER")

	fmt.Println("Masukkan Path Input File: ")
	fmt.Scanln(&pathFile)

	board := readInput(pathFile)
	if board == nil {
		fmt.Println("\nProgram dihentikan karena input tidak valid")
		return
	}

	if row != col {
		fmt.Println("Input tidak valid")
		return
	}

	// pilihan optimasi
	fmt.Println("Pilih Strategi Penyelesaian")
	fmt.Println("1. Pure Brute Force")
	fmt.Println("2. One Queen Per Column (optimasi kolom)")
	fmt.Println("3. Early Pruning + optimasi kolom")
	fmt.Println("Pilihan (1-3): ")
	fmt.Scanln(&pilihanStrategi)

	if pilihanStrategi < 1 || pilihanStrategi > 3 {
		fmt.Println("Pilihan tidak valid")
		return
	}

	// pilihan visualisasi
	fmt.Println("Visualisasi setiap berapa iterasi? (contoh: 1000000): ")
	fmt.Scanln(&freqVisualisasi)

	if freqVisualisasi <= 0 {
		freqVisualisasi = 1000000
	}

	// init working board
	tempBoard := make([][]string, row)
	for i := range tempBoard {
		tempBoard[i] = make([]string, col)
	}
	initQueens(tempBoard)

	namaStrategi := []string{"", "Pure Brute Force", "One Queen Per Column", "Early Pruning"}
	fmt.Printf("Strategi yang dipakai: %s\n", namaStrategi[pilihanStrategi])

	startTime := time.Now()
	var found bool

	switch pilihanStrategi {
	case 1:
		found = solveBruteForce(tempBoard, 0, 0)
	case 2:
		found = solveQueenKolom(tempBoard, 0)
	case 3:
		found = solvePruning(tempBoard, 0)
	}

	execTime := time.Since(startTime)

	if !found {
		fmt.Println("Tidak ada solusi ditemukan")
	} else {
		fmt.Println("Solusi ditemukan!")

		var save string
		fmt.Print("Apakah ingin menyimpan solusi? (y/n): ")
		fmt.Scanln(&save)
		if save == "y" || save == "Y" {
			saveToFile(tempBoard, execTime, iterationCount)
		}
	}
	fmt.Printf("Waktu Eksekusi: %.3f ms\n", float64(execTime.Microseconds())/1000.0)
	fmt.Printf("Total Iterasi: %d\n", iterationCount)
}

func readInput(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error membuka file:", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var tempBuffer []string

	// trim whitespace
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		line = strings.ReplaceAll(line, " ", "")
		line = strings.ReplaceAll(line, "\t", "")

		if line != "" {
			tempBuffer = append(tempBuffer, line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error membaca file:", err)
		return nil
	}

	// cek valid or invalid board
	if !isEmpty(tempBuffer) {
		return nil
	}

	tempRow := len(tempBuffer)
	tempCol := len(tempBuffer[0])

	if !isSquare(tempRow, tempCol) {
		return nil
	}

	if !isConsistentRows(tempBuffer, tempCol) {
		return nil
	}

	if !isAlphabetic(tempBuffer) {
		return nil
	}

	uniqueColors := countUniqueColors(tempBuffer)
	if !hasValidColorCount(uniqueColors, tempRow, tempCol) {
		return nil
	}

	// make board
	buffer = tempBuffer
	row = tempRow
	col = tempCol
	realBoard = make([][]string, row)

	for i := range realBoard {
		realBoard[i] = make([]string, col)
	}

	for i := range buffer {
		for j := range buffer[i] {
			realBoard[i][j] = string(buffer[i][j])
		}
	}
	return realBoard
}

// validity cekkk
func isEmpty(buffer []string) bool {
	if len(buffer) == 0 {
		fmt.Println("Error: File kosong atau tidak berisi data valid")
		return false
	}
	return true
}

func isSquare(rows, cols int) bool {
	if rows != cols {
		fmt.Printf("Error: Papan harus berbentuk persegi (N x N)\n")
		fmt.Printf("Papan saat ini: %d baris x %d kolom\n", rows, cols)
		return false
	}
	return true
}

func isConsistentRows(buffer []string, expectedLength int) bool {
	for i, line := range buffer {
		if len(line) != expectedLength {
			fmt.Printf("Error: Baris %d memiliki panjang %d, seharusnya %d\n", i+1, len(line), expectedLength)
			return false
		}
	}
	return true
}

func isAlphabetic(buffer []string) bool {
	for i, line := range buffer {
		for j, char := range line {
			if !((char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z')) {
				fmt.Printf("Error: Karakter tidak valid '%c' di baris %d kolom %d\n", char, i+1, j+1)
				fmt.Println("Hanya karakter alphabet (A-Z, a-z) yang diperbolehkan")
				return false
			}
		}
	}
	return true
}

func countUniqueColors(buffer []string) int {
	colorSet := make(map[rune]bool)
	for _, line := range buffer {
		for _, char := range line {
			upperChar := char
			if char >= 'a' && char <= 'z' {
				upperChar = char - 32
			}
			colorSet[upperChar] = true
		}
	}
	return len(colorSet)
}

func hasValidColorCount(uniqueColors, rows, cols int) bool {
	if uniqueColors != rows {
		fmt.Printf("Error: Jumlah warna unik (%d) tidak sesuai dengan ukuran papan (%d x %d)\n", uniqueColors, rows, cols)
		fmt.Printf("Seharusnya ada tepat %d warna yang berbeda\n", rows)
		return false
	}
	return true
}

func printBoard(board [][]string) {
	for i := range board {
		for j := range board[i] {
			fmt.Print(board[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func initQueens(tempBoard [][]string) {
	posisiQueens = make([][]int, row)
	for i := range posisiQueens {
		posisiQueens[i] = []int{0, 0}
	}

	for i := range tempBoard {
		for j := range tempBoard[i] {
			tempBoard[i][j] = realBoard[i][j]
		}
	}
}

// pilihan 1: pure brute force
func solveBruteForce(tempBoard [][]string, idxQueen int, mulaiDari int) bool {
	totalCell := row * col
	// base
	if idxQueen == row {
		iterationCount++

		visualize(iterationCount, tempBoard)

		if isValid(row) {
			solFound++
			fmt.Printf("\nSOLUSI DITEMUKAN DI ITERASI KE-%d\n", iterationCount)
			printBoard(tempBoard)
			return true
		}
		return false
	}

	// iterate
	for idx := 0; idx < totalCell; idx++ {
		r := idx / col
		c := idx % col

		// Skip if this cell already has a queen from a previous recursive call
		if tempBoard[r][c] == "#" {
			continue
		}

		if idxQueen > 0 || idx > 0 {
			barisLama := posisiQueens[idxQueen][0]
			kolomLama := posisiQueens[idxQueen][1]
			if tempBoard[barisLama][kolomLama] == "#" {
				tempBoard[barisLama][kolomLama] = realBoard[barisLama][kolomLama]
			}
		}

		// place new queen
		tempBoard[r][c] = "#"
		posisiQueens[idxQueen] = []int{r, c}

		// rekursi
		if solveBruteForce(tempBoard, idxQueen+1, 0) {
			return true
		}
	}

	// backtrack
	barisLama := posisiQueens[idxQueen][0]
	kolomLama := posisiQueens[idxQueen][1]
	if tempBoard[barisLama][kolomLama] == "#" {
		tempBoard[barisLama][kolomLama] = realBoard[barisLama][kolomLama]
	}
	posisiQueens[idxQueen] = []int{0, 0}
	tempBoard[0][0] = "#"

	return false
}

// pilihan 2: one queen per column
func solveQueenKolom(tempBoard [][]string, idxQueen int) bool {
	// base
	if idxQueen == row {
		iterationCount++

		visualize(iterationCount, tempBoard)

		if isValid(row) {
			solFound++
			fmt.Printf("\nSOLUSI DITEMUKAN DI ITERASI KE-%d\n", iterationCount)
			printBoard(tempBoard)
			return true
		}
		return false
	}

	// iterate
	idxKolom := idxQueen
	for r := 0; r < row; r++ {
		if r > 0 {
			barisLama := posisiQueens[idxQueen][0]
			kolomLama := posisiQueens[idxQueen][1]
			if tempBoard[barisLama][kolomLama] == "#" {
				tempBoard[barisLama][kolomLama] = realBoard[barisLama][kolomLama]
			}
		}

		// place queen
		tempBoard[r][idxKolom] = "#"
		posisiQueens[idxQueen] = []int{r, idxKolom}

		//recursif
		if solveQueenKolom(tempBoard, idxQueen+1) {
			return true
		}
	}

	//backtrack
	barisLama := posisiQueens[idxQueen][0]
	kolomLama := posisiQueens[idxQueen][1]
	if tempBoard[barisLama][kolomLama] == "#" {
		tempBoard[barisLama][kolomLama] = realBoard[barisLama][kolomLama]
	}
	posisiQueens[idxQueen] = []int{0, idxKolom}
	tempBoard[0][idxKolom] = "#"

	return false
}

// pilihan 3: one queen per column + early pruning
func solvePruning(tempBoard [][]string, idxQueen int) bool {
	// base
	if idxQueen == row {
		iterationCount++
		visualize(iterationCount, tempBoard)

		// double check validation
		if isValid(row) {
			solFound++
			fmt.Printf("\nSOLUSI DITEMUKAN DI ITERASI KE-%d\n", iterationCount)
			printBoard(tempBoard)
			return true
		}
		return false
	}

	idxKolom := idxQueen

	// iterate
	for r := 0; r < row; r++ {
		iterationCount++

		tempBoard[r][idxKolom] = "#"
		posisiQueens[idxQueen] = []int{r, idxKolom}

		// early pruning
		if isSquareSafe(r, idxKolom, idxQueen) {
			visualize(iterationCount, tempBoard)
			//recursif
			if solvePruning(tempBoard, idxQueen+1) {
				return true
			}
		}

		tempBoard[r][idxKolom] = realBoard[r][idxKolom]
	}

	return false
}

// hhelper func
func isSquareSafe(r, c, idxQueen int) bool {
	// cek queen sebelumnya
	for i := 0; i < idxQueen; i++ {
		rowQ := posisiQueens[i][0]
		colQ := posisiQueens[i][1]

		// cek vertikal
		if rowQ == r {
			return false
		}

		// cek horizontal
		if colQ == c {
			return false
		}

		// cek tetangga
		if isTetangga(rowQ, r, colQ, c) {
			return false
		}

		// cek warna sama
		if isWarnaSama(rowQ, colQ, r, c) {
			return false
		}
	}

	return true
}

func visualize(iterSekarang int, board [][]string) {
	if iterSekarang%freqVisualisasi == 0 {
		fmt.Printf("\nIterasi: %d\n", iterSekarang)
		printBoard(board)
	}
}

func isValid(jumlahQueen int) bool {
	for i := 0; i < jumlahQueen; i++ {
		for j := i + 1; j < jumlahQueen; j++ {
			r1 := posisiQueens[i][0]
			c1 := posisiQueens[i][1]
			r2 := posisiQueens[j][0]
			c2 := posisiQueens[j][1]

			if r1 == r2 || c1 == c2 || isTetangga(r1, r2, c1, c2) || isWarnaSama(r1, c1, r2, c2) {
				return false
			}
		}
	}
	return true
}

func isTetangga(row1, row2, col1, col2 int) bool {
	return math.Abs(float64(row1-row2)) <= 1 && math.Abs(float64(col1-col2)) <= 1
}

func isWarnaSama(row1, col1, row2, col2 int) bool {
	return realBoard[row1][col1] == realBoard[row2][col2]
}

func saveToFile(board [][]string, execTime time.Duration, iterations int) {
	// gabungin nama inputfile ke output file path
	inputBaseName := filepath.Base(pathFile)
	inputNameOnly := strings.TrimSuffix(inputBaseName, filepath.Ext(inputBaseName))
	outputFileName := "solusi_" + inputNameOnly + ".txt"
	outputPath := "test"
	outputFullPath := filepath.Join(outputPath, outputFileName)
	file, err := os.Create(outputFullPath)
	if err != nil {
		fmt.Println("Error membuat file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	namaStrategi := []string{"", "Pure Brute Force", "One Queen Per Column", "Early Pruning"}
	writer.WriteString(fmt.Sprintf("Menggunakan optimasi: %s\n", namaStrategi[pilihanStrategi]))
	for i := range board {
		for j := range board[i] {
			writer.WriteString(board[i][j])
			writer.WriteString(" ")
		}
		writer.WriteString("\n")
	}

	writer.WriteString(fmt.Sprintf("\nWaktu Eksekusi: %.3f ms\n", float64(execTime.Microseconds())/1000.0))
	writer.WriteString(fmt.Sprintf("Total Iterasi: %d\n", iterations))

	writer.Flush()
	fmt.Println("Solusi berhasil disave ke:  ", outputFullPath)
}
