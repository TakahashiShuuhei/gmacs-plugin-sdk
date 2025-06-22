# gmacs Plugin SDK

This package provides the Software Development Kit (SDK) for creating gmacs plugins. It includes all the necessary interfaces, RPC communication layer, and helper functions for plugin development.

## Quick Start

### 1. Create a new plugin project

```bash
mkdir my-gmacs-plugin
cd my-gmacs-plugin
go mod init my-gmacs-plugin
```

### 2. Add dependency

```bash
go mod edit -require github.com/TakahashiShuuhei/gmacs/plugin-sdk@latest
```

### 3. Implement your plugin

```go
package main

import (
    "context"
    pluginsdk "github.com/TakahashiShuuhei/gmacs/plugin-sdk"
)

type MyPlugin struct{}

func (p *MyPlugin) Name() string { return "my-plugin" }
func (p *MyPlugin) Version() string { return "1.0.0" }
func (p *MyPlugin) Description() string { return "My awesome plugin" }

func (p *MyPlugin) Initialize(ctx context.Context, host pluginsdk.HostInterface) error {
    host.ShowMessage("My plugin loaded!")
    return nil
}

func (p *MyPlugin) Cleanup() error { return nil }

func (p *MyPlugin) GetCommands() []pluginsdk.CommandSpec {
    return []pluginsdk.CommandSpec{
        {
            Name:        "my-command",
            Description: "My awesome command",
            Interactive: true,
            Handler:     "HandleMyCommand",
        },
    }
}

func (p *MyPlugin) GetMajorModes() []pluginsdk.MajorModeSpec { return nil }
func (p *MyPlugin) GetMinorModes() []pluginsdk.MinorModeSpec { return nil }
func (p *MyPlugin) GetKeyBindings() []pluginsdk.KeyBindingSpec { return nil }

func main() {
    pluginsdk.ServePlugin(&MyPlugin{})
}
```

### 4. Build your plugin

```bash
go build -o my-plugin
```

## API Reference

### Plugin Interface

All plugins must implement the `Plugin` interface:

```go
type Plugin interface {
    // Plugin metadata
    Name() string
    Version() string
    Description() string
    
    // Lifecycle
    Initialize(ctx context.Context, host HostInterface) error
    Cleanup() error
    
    // Feature provision
    GetCommands() []CommandSpec
    GetMajorModes() []MajorModeSpec
    GetMinorModes() []MinorModeSpec
    GetKeyBindings() []KeyBindingSpec
}
```

### Host Interface

The host (gmacs) provides these APIs to plugins:

```go
type HostInterface interface {
    // Editor operations
    GetCurrentBuffer() BufferInterface
    GetCurrentWindow() WindowInterface
    SetStatus(message string)
    ShowMessage(message string)
    
    // Command execution
    ExecuteCommand(name string, args ...interface{}) error
    
    // Mode management
    SetMajorMode(bufferName, modeName string) error
    ToggleMinorMode(bufferName, modeName string) error
    
    // Events and hooks
    AddHook(event string, handler func(...interface{}) error)
    TriggerHook(event string, args ...interface{})
    
    // Buffer operations
    CreateBuffer(name string) BufferInterface
    FindBuffer(name string) BufferInterface
    SwitchToBuffer(name string) error
    
    // File operations
    OpenFile(path string) error
    SaveBuffer(bufferName string) error
    
    // Configuration
    GetOption(name string) (interface{}, error)
    SetOption(name string, value interface{}) error
}
```

## Helper Functions

### ServePlugin

```go
func ServePlugin(impl Plugin)
```

Starts the plugin server with the provided plugin implementation.

### NewHostStub

```go
func NewHostStub() HostInterface
```

Creates a stub implementation of HostInterface for testing.

## Testing

Use the provided stubs for unit testing:

```go
func TestMyPlugin(t *testing.T) {
    plugin := &MyPlugin{}
    host := pluginsdk.NewHostStub()
    
    err := plugin.Initialize(context.Background(), host)
    if err != nil {
        t.Fatalf("Failed to initialize plugin: %v", err)
    }
    
    // Test plugin functionality
    commands := plugin.GetCommands()
    if len(commands) != 1 {
        t.Errorf("Expected 1 command, got %d", len(commands))
    }
}
```

## Examples

See the [gmacs-example-plugin](https://github.com/TakahashiShuuhei/gmacs-example-plugin) repository for a complete example implementation.

## Communication Protocol

This SDK uses HashiCorp's go-plugin library for RPC communication between gmacs and plugins. Plugins run as separate processes and communicate via RPC calls, providing isolation and stability.

## Contributing

This SDK is part of the gmacs project. Please report issues and contribute improvements through the main gmacs repository.