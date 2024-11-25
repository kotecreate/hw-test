package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

type LogType struct {
	Time    string
	Level   string
	Message string
}

func parseLogFile(filePath string, level string) ([]LogType, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("open file: %s: %w", filePath, err)
	}
	defer file.Close()

	var entries []LogType
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) >= 3 && (level == "" || fields[1] == level) {
			entry := LogType{
				Time:    fields[0],
				Level:   fields[1],
				Message: strings.Join(fields[2:], " "),
			}
			entries = append(entries, entry)
		}
	}
	if err = scanner.Err(); err != nil {
		return nil, fmt.Errorf("сканирование файла: %s: %w", filePath, err)
	}
	return entries, nil
}

func analyze(logTypes []LogType) map[string]int {
	stats := make(map[string]int)
	for _, entry := range logTypes {
		stats[entry.Level]++
	}
	return stats
}

func printStats(stats map[string]int, outputPath string) error {
	var out *os.File
	var err error

	if outputPath == "" {
		out = os.Stdout
	} else {
		out, err = os.Create(outputPath)
		if err != nil {
			return fmt.Errorf("не удалось создать файл %s: %w", outputPath, err)
		}
		defer out.Close()
	}

	fmt.Fprintln(out, "Статистика:")
	for level, count := range stats {
		fmt.Fprintf(out, "%s: %d\n", level, count)
	}
	return nil
}

func main() {
	fileFlag := flag.String("file", "", "Path to the log file")
	levelFlag := flag.String("level", "INFO", "Filter by log level")
	outputFlag := flag.String("output", "", "Path to the output file")
	flag.Parse()

	logAnalyzerFile := os.Getenv("LOG_ANALYZER_FILE")
	logAnalyzerLevel := os.Getenv("LOG_ANALYZER_LEVEL")
	logAnalyzerOutput := os.Getenv("LOG_ANALYZER_OUTPUT")

	if *fileFlag == "" && logAnalyzerFile == "" {
		fmt.Println("Не задан путь к файлу логов")
	}

	filePath := *fileFlag
	if filePath == "" {
		filePath = logAnalyzerFile
	}

	level := *levelFlag
	if level == "" {
		level = logAnalyzerLevel
	}

	outputPath := *outputFlag
	if outputPath == "" {
		outputPath = logAnalyzerOutput
	}

	logTypes, err := parseLogFile(filePath, level)
	if err != nil {
		fmt.Println(err)
	}

	stats := analyze(logTypes)

	err = printStats(stats, outputPath)
	if err != nil {
		fmt.Println(err)
	}
}
