package main

import (
	"flag"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"os/exec"
)

func loadYamlFile(filename string) (map[string]string, error) {
	var data map[string]string

	// Read file
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Printf("Error while reading YAML file %v.", err)
		return data, err
	}
	// Unmarshal YAML content into a map
	err = yaml.Unmarshal(content, &data)
	if err != nil {
		log.Printf("Error processing YAML data: %v.", err)
		return data, err
	}

	return data, nil
}

// loadEnvironmentVars loads environment variables from the passed map
func loadEnvironmentVars(vars map[string]string) {
	for key, value := range vars {
		err := os.Setenv(key, value)
		if err != nil {
			log.Printf("Error setting environment variable %s: %v", key, err)
			continue
		}
		log.Printf("Loaded environment variable %s", key)
	}
}

func launchProgram(programCommand []string) error {
	// Check if the programCommand is not empty
	if len(programCommand) == 0 {
		log.Print("No program specified")
		return nil
	}

	// The first element is program name, rest elements are arguments.
	programName := programCommand[0]
	programArgs := programCommand[1:]

	// Create an *exec.Cmd
	cmd := exec.Command(programName, programArgs...)

	// Connect the stdout and stderr to os.Stdout and os.Stderr,
	// this forwards the output of the child program to the parent.
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Launch the program
	err := cmd.Run()
	if err != nil {
		// Log the error and return it
		log.Printf("Error launching program: %v", err)
		return err
	}

	// Log a success message
	log.Printf("Program %s launched successfully", programName)

	return nil
}

func main() {
	// Define a new command-line flag
	yamlFileName := flag.String("f", "", "Path to the YAML file")

	// Parse all command-line flags
	flag.Parse()

	if *yamlFileName == "" {
		log.Fatalf("No YAML file specified!")
	}

	// Load YAML file
	envVars, err := loadYamlFile(*yamlFileName)
	if err != nil {
		log.Fatalf("Failed to load YAML file: %v", err)
	}

	// Load environment variables
	loadEnvironmentVars(envVars)

	// Check if there are non-flag command-line arguments
	if flag.NArg() == 0 {
		log.Fatalf("No command specified!")
	}

	// Everything after the "--" is considered as the command to execute
	programCommand := flag.Args()

	// Launch the specified program
	err = launchProgram(programCommand)
	if err != nil {
		log.Fatalf("Failed to launch program: %v", err)
	}
}
