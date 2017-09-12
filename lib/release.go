package lib

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"strings"
)

// GetSystemReleaseName retrieves information about the distribution release as given by the vendor
func GetSystemReleaseName() string {
	f, err := os.Open("/etc/os-release")
	if err != nil {
		log.Fatalln(err)
	}
	scanner := bufio.NewScanner(f)
	var buffer bytes.Buffer
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "=")
		configMap := make(map[string]string)
		for _ = range parts {
			configMap[parts[0]] = parts[1]
		}
		for key, value := range configMap {
			if key == "NAME" || key == "VERSION" {
				buffer.WriteString(value[1:len(value)-1] + " ")
			}
		}
	}
	release := strings.Replace(buffer.String()[:len(buffer.String())-1], " ", "_", -1)
	release = strings.Replace(release, "-", "_", -1)
	return release
}
