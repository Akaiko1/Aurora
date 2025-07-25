# Aurora: An Ebitengine Shooter

## Overview

Aurora is a simple 2D shooter game built with Golang and the Ebitengine game library. This project is currently in early development with basic gameplay features implemented.

<img src="assets/demo.gif" alt="Game Demo" width="480"/>

## Current Features

- Basic top-down shooter mechanics
- Player character with movement and shooting
- **Extensible weapon system** with 4+ weapon types:
  - Normal: 5 projectiles, disappear on hit, balanced stats
  - Piercing: 2 projectiles, pierce through enemies, faster speed
  - Rapid Fire: 8 projectiles, very fast firing, thinner bullets
  - Heavy Cannon: 1 projectile, slow but powerful, large bullets
- Simple enemy AI with random movement patterns
- Optimized collision detection with spatial partitioning
- Grazing system for scoring points by narrowly avoiding enemy bullets
- Simple level progression with phases and scenarios
- Togglable hitbox display (press B)
- Background tiles with grass/no-grass variants

## Getting Started

### Prerequisites

To run this game, you need to have Go installed on your machine. You can download it from the [official Go website](https://golang.org/dl/).

### Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/Akaiko1/aurora
    cd aurora
    ```

2. Install the dependencies:

    ```sh
    go mod tidy
    ```

### Running the Game

To run the game, execute the following command in your terminal:

```sh
go run main.go
```

## Controls

- **Arrow Keys**: Move your character
- **Space**: Shoot projectiles
- **B**: Toggle hitbox display
- **1**: Normal weapon (5 projectiles, balanced)
- **2**: Piercing weapon (2 projectiles, pierce enemies)
- **3**: Rapid Fire weapon (8 projectiles, very fast)
- **4**: Heavy Cannon weapon (1 projectile, slow but powerful)

## Project Structure

- `main.go`: The main entry point of the game
- `internals/game/`: Contains game logic and rendering code
- `internals/entities/`: Player, enemy, and projectile definitions
- `internals/physics/`: Simple collision detection with hitboxes
- `internals/events/`: Event handlers for game objects
- `internals/inputs/`: Image and font loading utilities
- `internals/config/`: Game constants like screen size and speeds
- `assets/`: Game sprites and fonts

## Planned Features

- More enemy types and attack patterns
- Power-ups and special abilities
- Score system improvements
- Sound effects and music
- Menu system and game settings

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests with improvements or bug fixes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.