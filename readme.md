# Airsettle

Welcome to Airsettle, a blockchain platform focused on secure and efficient data management. It features custom execution layers, batch processing, and a democratic validator system.

## Features

- **Blockchain Initialization**: Custom 'air' address prefix, streamlined project structure.
- **Execution Layer Enhancements**: 'Exelayer' for validators and voting power, dynamic management, and robust queries.
- **Verification Key Management**: Efficient 'vkey' type for managing lengthy verification keys.
- **Batch Processing Mechanics**: 'batch_min' and 'batch_max' structures, robust batch handling, and detailed query functions.
- **Validator Management**: 'Poll' system for democratic validator selection, comprehensive management messages, and detailed poll information queries.

## Getting Started

### Prerequisites

- [Ignite Cli](https://github.com/ignite/cli/releases/tag/v0.27.1) v0.27.1
- [Go](https://golang.org/doc/install) v1.20.+

### Installation

#### 1. Clone the repository

    git clone https://github.com/airchains-network/airsettle.git

#### 2. Navigate to the project directory

    cd airsettle

#### 3. Switch to the DevNet Release

    git checkout v0.0.1-beta

#### 4. Set Environment Variable

Add the path `/Users/<your_pc_username>/go/bin` to your environment variables for easy access to 'AirSettle'. Replace `<your_pc_username>` with your actual PC username.

#### 5. Initialize the Chain

    ignite chain init

#### 6. Delete the Existing Configuration Folder

Remove the folder `~/.airsettle` if it exists.

    rm -rf ~/.airsettle

#### 7. Initialize the Node with the Moniker

    airsettled init <moniker> --chain-id airsettle

#### 8. Update Genesis Configuration

Replace the contents of `~/.airsettle/config/genesis.json` with the contents from the `docs/node/genesis.json` file.

    cp docs/node/genesis.json ~/.airsettle/config/genesis.json

#### 9. Update Configuration

1. Edit `~/.airsettle/config/config.toml` to set `persistent_peers`:

```toml
persistent_peers = "067778c3c6c2e098237c867359d6e02b7773d8ea@34.100.195.165:26656"
```

#### 10. Start the Node

    airsettled start
