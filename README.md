# xay

`xay` is a command-line tool that works similarly to macOS's `say` command, but uses OpenAI's text-to-speech API to generate more natural-sounding speech.

## Installation

You can install `xay` in one of two ways:

1. Direct installation (recommended):
```bash
go install github.com/roland-zhou/xay@latest
```

2. Manual installation:
   1. Make sure you have Go 1.21 or later installed
   2. Clone this repository
   3. Run `go install` in the project directory

After installation, make sure the Go binary directory is in your PATH. You can add it by running:

```bash
# For macOS/Linux
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.zshrc  # or ~/.bashrc
source ~/.zshrc  # or ~/.bashrc

# For Windows (PowerShell)
$env:Path += ";$(go env GOPATH)\bin"
```

You can verify the installation by running:
```bash
xay hello
```

## Configuration

Before using `xay`, you need to set up your OpenAI API key:

```bash
export OPENAI_API_KEY=your_api_key_here
```

You can add this line to your shell's configuration file (e.g., `~/.zshrc` or `~/.bashrc`) to make it permanent.

## Usage

Basic usage:
```bash
xay "Hello, world!"
```

The command will use OpenAI's text-to-speech API to speak the provided text.

## Features

- Natural-sounding speech using OpenAI's text-to-speech API
- Simple command-line interface similar to macOS's `say` command
- Environment-based API key configuration

## Requirements

- Go 1.21 or later
- OpenAI API key
- macOS (for audio playback)

## License

MIT 