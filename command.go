package main

import "errors"

type Command struct {
	command string
	args    []string
}

type Commands struct {
	commands map[string]func(*State, Command) error
}

func (c *Commands) run(s *State, cmd Command) error {
	command, ok := c.commands[cmd.command]

	if !ok {
		return errors.New("Command not found: " + cmd.command)
	}

	return command(s, cmd)
}

func (c *Commands) register(name string, f func(*State, Command) error) error {
	_, ok := c.commands[name]

	if ok {
		return errors.New("Command already has a handler")
	}

	c.commands[name] = f

	return nil
}
