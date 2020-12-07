package greetings

import (
    "errors"
    "fmt"
    "math/rand"
    "time"
)

func Init() {
    rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
    messages := []string {
        "Hi, %v!",
        "Welcome, %v!",
        "Great to see you, %v!",
    }

    return messages[rand.Intn(len(messages))]
}

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    if name == "" {
        return "", errors.New("empty new")
    }
    Init()
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}
