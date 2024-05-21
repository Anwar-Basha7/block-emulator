# Block Emulator

## Overview

Block Emulator is a Go-based simulation tool designed to emulate blockchain behavior and consensus algorithms, particularly PBFT (Practical Byzantine Fault Tolerance). This project aims to provide a robust platform for testing and analyzing blockchain mechanisms in a controlled environment.

## Features

Blockchain Simulation: Mimics the behavior of a blockchain network.
Consensus Algorithm: Implements PBFT for consensus.
Modular Design: Easily extendable for additional features and algorithms.
Logging: Comprehensive logging for monitoring and debugging.

## Installation

To get started with Block Emulator, ensure you have Go installed on your system. Clone the repository and build the project:

  git clone https://github.com/Anwar-Basha7/block-emulator.git
  
  cd block-emulator
  
  go build

## Usage

Run the emulator with the following command:

   go run main.go -n 0 -N 4 -s 0 -S 10 -m 0

## Project Structure

  emulator/: Contains the core emulator logic.
  
  logger/: Manages logging functionalities.
  
  pbft/: Implements the PBFT consensus algorithm.
  
  main.go: Entry point of the application.
