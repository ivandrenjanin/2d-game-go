# 2D Game with Go and Raylib

This project is being used to learn Go features in combination with the Raylib game programming library and the Entity-Component-System (ECS) framework.

# Dependencies

Raylib: A simple and easy-to-use library to learn and enjoy video games programming. Raylib is highly inspired by Borland BGI graphics lib and by XNA framework. The library provides a simple and intuitive interface towards a set of powerful tools, allowing for the development of both complex and simple games and applications.

Project structure

- components
- core
- systems
- constants
- main.go

components - Contains the Component definitions. In an ECS paradigm, a Component is a data structure that holds the attributes of entities. It does not contain any behavior. For example, a Position component might contain "x" and "y" attributes.
core - Core is where essential game functions and structures are defined. This could include the game loop, initialization and shutdown routines, and any utility or helper functions that aren't part of a game entity's behavior.
systems - Contains the System implementations. Systems in ECS are where the behavior logic lives. They operate on Components, reading and writing their data. For example, a Movement system might use the data in the Position and Velocity components to update an entity's position over time.
constants - Contains some commonly used attributes like colors
entities - Contains factories and builders for game entities
main.go - Main application entry point.
