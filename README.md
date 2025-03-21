# Aurora: An Ebitengine Shooter

## Overview

Aurora is a simple 2D shooter game built with Golang and the Ebitengine game library. This project is currently in early development with basic gameplay features implemented.

<img src="assets/preview.jpg" alt="Game Screenshot" width="480"/>

## Current Features

- Basic top-down shooter mechanics
- Player character with movement and shooting
- Simple enemy AI with random movement patterns
- Projectile collision detection
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