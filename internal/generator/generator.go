package generator

import (
	"embed"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

//go:embed templates/*
var content embed.FS

func GenerateGoMain(path string) error {
	outputFile := filepath.Join(path, "main.go")

	// Create a new file or overwrite an existing one
	file, err := os.OpenFile(outputFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	sourceFile, err := content.Open("templates/main.go")
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	_, err = io.Copy(file, sourceFile)
	if err != nil {
		// Write the content to the file
		return err
	}

	log.Println("File 'main.go' created successfully!")

	return nil
}

func GenerateeBPFProgram(path string) error {
	outputFile := filepath.Join(path, "program.bpf.c")

	// Create a new eBPF file
	file, err := os.OpenFile(outputFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	sourceFile, err := content.Open("templates/program.bpf.c")
	if err != nil {
		return err
	}
	defer sourceFile.Close()
	_, err = io.Copy(file, sourceFile)
	// Write the content to the file
	if err != nil {
		return err
	}

	log.Println("File 'program.bpf.c' created successfully!")

	return nil
}

func DumpBTF(path string) error {
	// DumpBTF runs the bpftool command to dump BTF information into vmlinux.h
	outputFile := filepath.Join(path, "vmlinux.h")

	// Create vmlinux.h for writing
	file, err := os.OpenFile(outputFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	cmd := exec.Command("bpftool", "btf", "dump", "file", "/sys/kernel/btf/vmlinux", "format", "c")
	cmd.Stdout = file
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err == nil {
		log.Println("File 'vmlinux.h' created successfully!")
	}

	// Run the command and return any error
	return err
}

// DumpMake creates a Makefile for convenience
func DumpMake(path string) error {
	outputFile := filepath.Join(path, "Makefile")

	// Create vmlinux.h for writing
	file, err := os.OpenFile(outputFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	sourceFile, err := content.Open("templates/Makefile")
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	// Write the content to the file
	_, err = io.Copy(file, sourceFile)
	if err != nil {
		return err
	}

	log.Println("File 'Makefile' created successfully!")

	// Run the command and return any error
	return err
}
