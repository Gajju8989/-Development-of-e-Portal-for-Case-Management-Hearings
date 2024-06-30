---

# Swiggy Assignment

Author: Gajendra Singh

## Description
This project implements a simulated arena where players can fight. It includes functionalities for player management (health, strength, attack), dice rolling, and simulating battles between players.


## Dependencies
- Go 1.21.6

## Installation
1. Navigate to the project directory: 
   cd swiggy-Assignment
  Ensure you have Go installed and configured properly.


Directory Structure Overview
arena

service
service.go: Defines the interface for arena services.
service_fight_impl.go: Implements methods for simulating fights between players.
service_fight_impl_test.go: Unit tests for the fight implementation.
di

container.go: Configures dependency injection for the project.
dice

service
service.go: Defines the interface for dice services.
service_impl_rolldice.go: Implements methods for rolling dice.
service_impl_rolldice_test.go: Unit tests for the dice rolling implementation.
player

model
models.go: Defines the data model for players.
service
service.go: Defines the interface for player services.
service_is_alive_impl.go: Implements methods for checking if a player is alive.
service_is_alive_impl_test.go: Unit tests for player alive status implementation.
service_take_damage_impl.go: Implements methods for applying damage to players.
service_take_damage_impl_test.go: Unit tests for applying damage to players.
service_validate_player_impl.go: Implements methods for validating player attributes.
service_validate_player_impl_test.go: Unit tests for player validation.

Running the Project
Navigate to the project directory:
bash
cd swiggy-Assignment

Run the project:

bash
go run main.go

