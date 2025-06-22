package pluginsdk

import (
	"github.com/hashicorp/go-plugin"
)

// ServePlugin はプラグインサーバーを起動するヘルパー関数
func ServePlugin(impl Plugin) {
	// プラグインのRPCマップにプラグインインスタンスを設定
	pluginMap := map[string]plugin.Plugin{
		"gmacs-plugin": &RPCPlugin{Impl: impl},
	}

	// プラグインサーバーを起動
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: Handshake,
		Plugins:         pluginMap,
	})
}

// NewHostStub はテスト用のHostInterfaceスタブを作成
func NewHostStub() HostInterface {
	return &HostStub{}
}

// HostStub はテスト用のHostInterface実装
type HostStub struct {
	messages []string
	status   string
}

func (h *HostStub) GetCurrentBuffer() BufferInterface {
	return &BufferStub{}
}

func (h *HostStub) GetCurrentWindow() WindowInterface {
	return &WindowStub{}
}

func (h *HostStub) SetStatus(message string) {
	h.status = message
}

func (h *HostStub) ShowMessage(message string) {
	h.messages = append(h.messages, message)
}

func (h *HostStub) ExecuteCommand(name string, args ...interface{}) error {
	return nil
}

func (h *HostStub) SetMajorMode(bufferName, modeName string) error {
	return nil
}

func (h *HostStub) ToggleMinorMode(bufferName, modeName string) error {
	return nil
}

func (h *HostStub) AddHook(event string, handler func(...interface{}) error) {
	// スタブ実装
}

func (h *HostStub) TriggerHook(event string, args ...interface{}) {
	// スタブ実装
}

func (h *HostStub) CreateBuffer(name string) BufferInterface {
	return &BufferStub{name: name}
}

func (h *HostStub) FindBuffer(name string) BufferInterface {
	return &BufferStub{name: name}
}

func (h *HostStub) SwitchToBuffer(name string) error {
	return nil
}

func (h *HostStub) OpenFile(path string) error {
	return nil
}

func (h *HostStub) SaveBuffer(bufferName string) error {
	return nil
}

func (h *HostStub) GetOption(name string) (interface{}, error) {
	return nil, nil
}

func (h *HostStub) SetOption(name string, value interface{}) error {
	return nil
}

// GetMessages はテスト用にメッセージ履歴を取得
func (h *HostStub) GetMessages() []string {
	return h.messages
}

// GetStatus はテスト用にステータスを取得
func (h *HostStub) GetStatus() string {
	return h.status
}

// BufferStub はテスト用のBufferInterface実装
type BufferStub struct {
	name    string
	content string
	cursor  int
	dirty   bool
}

func (b *BufferStub) Name() string {
	return b.name
}

func (b *BufferStub) Content() string {
	return b.content
}

func (b *BufferStub) SetContent(content string) {
	b.content = content
	b.dirty = true
}

func (b *BufferStub) InsertAt(pos int, text string) {
	if pos < 0 || pos > len(b.content) {
		return
	}
	b.content = b.content[:pos] + text + b.content[pos:]
	b.dirty = true
}

func (b *BufferStub) DeleteRange(start, end int) {
	if start < 0 || end > len(b.content) || start > end {
		return
	}
	b.content = b.content[:start] + b.content[end:]
	b.dirty = true
}

func (b *BufferStub) CursorPosition() int {
	return b.cursor
}

func (b *BufferStub) SetCursorPosition(pos int) {
	b.cursor = pos
}

func (b *BufferStub) MarkDirty() {
	b.dirty = true
}

func (b *BufferStub) IsDirty() bool {
	return b.dirty
}

func (b *BufferStub) Filename() string {
	return b.name
}

// WindowStub はテスト用のWindowInterface実装
type WindowStub struct {
	buffer       BufferInterface
	width, height int
	scrollOffset int
}

func (w *WindowStub) Buffer() BufferInterface {
	if w.buffer == nil {
		w.buffer = &BufferStub{}
	}
	return w.buffer
}

func (w *WindowStub) SetBuffer(buffer BufferInterface) {
	w.buffer = buffer
}

func (w *WindowStub) Width() int {
	if w.width == 0 {
		return 80
	}
	return w.width
}

func (w *WindowStub) Height() int {
	if w.height == 0 {
		return 24
	}
	return w.height
}

func (w *WindowStub) ScrollOffset() int {
	return w.scrollOffset
}

func (w *WindowStub) SetScrollOffset(offset int) {
	w.scrollOffset = offset
}