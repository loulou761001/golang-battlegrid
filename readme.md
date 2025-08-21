# BattleGrid

BattleGrid is a turn-based tactical battle simulator written in **Go**, running in the terminal.  
The game features grid-based combat where units move, attack, and interact based on their stats.

---

## Features
- Turn-based system with alternating teams
- Units with customizable stats:
  - Attack, Defense, Health, Range
- Dice-based combat resolution for randomness
- Movement system (Manhattan & diagonal handling)
- Weighted target selection for AI behavior
- Terminal interface with colored battlefield display
- Victory detection between teams

---

## Project Structure
```
battle-sim/
├── go.mod
├── go.sum
├── main.go
├── assets/        # Assets or predefined data
├── gamelogic/     # Core battle logic and AI
├── state/         # Game state storage
├── types/         # Unit and combat types
└── ui/            # Terminal UI
```

---

## Installation
1. Clone this repository:
   ```bash
   git clone https://github.com/your-username/battlegrid.git
   cd battlegrid
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Run the game:
   ```bash
   go run main.go
   ```

---

## Controls
- **Arrow keys** → Move cursor
- **Ctrl+C** → Quit game

---

## Example Gameplay
When the game starts, two teams of units are placed on a grid.  
Each unit will:
1. Move toward the closest enemy.
2. Attack if the enemy is within range.
3. Combat is resolved using attack & defense stats + dice rolls.

Victory is declared when all opposing units are defeated.

---

## Future Improvements
- [ ] Web-based interface (instead of CLI)
- [ ] Better AI decision-making
- [ ] Unit abilities & morale
- [ ] Procedural maps
- [ ] Multiplayer support

---

## License
MIT License. Free to use, modify, and distribute.
