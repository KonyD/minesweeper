# 🧨 Minesweeper

A classic **Minesweeper** game built using **Go** and the **[Raylib](https://www.raylib.com/)** game development library.

![screenshot](./screenshots/screenshot.png)

### 📥 Clone and Run

```bash
git clone https://github.com/KonyD/minesweeper.git
cd minesweeper
go run main.go
```

### 🏗️ Building the Game

##### 🪟 Windows (with MinGW or TDM-GCC)
```bash
go build -ldflags "-H=windowsgui" -o minesweeper.exe
```

##### 🐧 Linux

```bash
go build -o minesweeper
```

##### 🍏 macOS

```bash
go build -o minesweeper
```