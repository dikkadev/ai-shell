# for ai-shell

Welcome to ai-shell! It's a command-line tool that uses OpenAI's GPT-3.5-turbo to generate precise console commands based on your instructions. It takes in your commands, gathers some context, and lets an AI model figure out the best course of action.

## What You'll Need

1. An OpenAI API key: This is non-negotiable. You'll need this to connect to OpenAI's API and make magic happen. Make sure to set it as an environment variable `OPENAI_API_KEY`.

## Getting ai-shell on Your Machine

### Option 1: Use Go

If you have Go set up on your machine, it's a breeze. Just use the following command:


```bash
go install github.com/sett17/ai-shell` 
```

### Option 2: Grab the Binary

I've also got pre-compiled binaries for you. Just head over to the [Releases](https://github.com/sett17/ai-shell/releases) page, download the version you need, and you're good to go.

## Configuration

ai-shell uses a configuration file (`config.toml`) to customize its behavior. By default, this file is located in the local configuration directory for your system:

- On Linux and Unix systems: `~/.config/ai-shell/config.toml`
- On Windows: `%APPDATA%/ai-shell/config.toml` or `C:\Users\%USER%\AppData\Roaming\ai-shell\config.toml`

Don't worry if you don't find the config file there, ai-shell will create one with default values for you on its first run.

### Configuration Structure

Each context item has two common settings:

- `Enabled`: A Boolean that controls whether this context is enabled.
- `Priority`: An integer that determines the order in which this context is considered by the AI. Higher numbers mean higher priority.

Moreover, some contexts may have additional configuration settings.

### Context-Specific Configurations

| Context | Description | Additional Settings |
| --- | --- | --- |
| FileListing | Provides file listing context to the AI. | / |
| Shell | Provides shell context to the AI. | / |

This structure will be maintained as more context types are added in the future. Each new context will come with its own `Enabled` and `Priority` settings and may include additional context-specific configurations.

## How to Use ai-shell

It's quite simple to use ai-shell. You input an instruction, and the AI spits out a command. You can revise this command, and then either execute it or exit. That's all there is to it.

A neat feature is that you can tweak how the AI understands your instructions using context. You can enable or disable different contexts and prioritize them in accordance with your needs.

## Contributing

Feel free to contribute to ai-shell. I don't have a formal contribution guide yet, but just submit a PR, and I'll take a look.

## License

For details on the license, please check the [LICENSE](https://github.com/sett17/ai-shell/blob/main/LICENSE) file.

